package main

import (
	"net/http"

	retry_helpers "github.com/birsanion/netopia/api-server/helpers/retry"
	models_db "github.com/birsanion/netopia/api-server/models/db"
	"github.com/birsanion/netopia/api-server/models/events"
	"github.com/birsanion/netopia/api-server/models/requests"
	"github.com/birsanion/netopia/api-server/models/responses"

	"github.com/almerlucke/go-iban/iban"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AbortWithError(c *gin.Context, err error, status int, msg string) {
	logrus.Error(err.Error())
	c.JSON(status, gin.H{"error": msg})
	c.Abort()
}

func CreatePaymentHandler(db *gorm.DB, queue *RabbitMQConnection) func(*gin.Context) {
	return func(c *gin.Context) {
		var reqData requests.InitPaymentPayload
		if err := c.ShouldBindJSON(&reqData); err != nil {
			AbortWithError(c, err, http.StatusBadRequest, "bad_request")
			return
		}

		if _, err := iban.NewIBAN(reqData.Iban); err != nil {
			AbortWithError(c, err, http.StatusBadRequest, "invalid_iban")
			return
		}

		payment := models_db.Payment{
			Amount:            reqData.Amount,
			Currency:          reqData.Currency,
			Description:       reqData.Description,
			InternalReference: reqData.InternalReference,
			Iban:              reqData.Iban,
			Status:            models_db.PaymentStatusNew,
		}
		if err := retry_helpers.RetryDo(c.Request.Context(), func() (err error) {
			return db.Create(&payment).Error
		}); err != nil {
			AbortWithError(c, err, http.StatusInternalServerError, "internal_error")
			return
		}

		if err := retry_helpers.RetryDo(c.Request.Context(), func() (err error) {
			return queue.Publish(c.Request.Context(), events.CreatePayment{
				TransactionID:     payment.TransactionID,
				Amount:            payment.Amount,
				Currency:          payment.Currency,
				Iban:              payment.Iban,
				Description:       payment.Description,
				InternalReference: payment.InternalReference,
			})
		}); err != nil {
			AbortWithError(c, err, http.StatusInternalServerError, "internal_error")
			return
		}

		c.JSON(http.StatusOK, responses.InitPaymentRespose{
			TransactionID: payment.TransactionID,
			Status:        payment.Status,
		})
	}
}

func HealthCheckHandler(db *gorm.DB, queue *RabbitMQConnection) func(*gin.Context) {
	return func(c *gin.Context) {
		if !IsDBAvailable(db) {
			c.JSON(http.StatusOK, responses.NewHealthCheckRespose(responses.HealthCheckStatusError).
				WithDetails("Database unavailable"))
			return
		}

		if !queue.IsAvailable() {
			c.JSON(http.StatusOK, responses.NewHealthCheckRespose(responses.HealthCheckStatusError).
				WithDetails("RabbitMQ unavailable"))
			return
		}

		c.JSON(http.StatusOK, responses.NewHealthCheckRespose(responses.HealthCheckStatusOk))
	}
}
