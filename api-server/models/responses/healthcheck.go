package responses

type HealthCheckStatus string

const (
	HealthCheckStatusOk    HealthCheckStatus = "ok"
	HealthCheckStatusError HealthCheckStatus = "error"
)

func (s HealthCheckStatus) IsOk() bool {
	return s == HealthCheckStatusOk
}

func (s HealthCheckStatus) IsError() bool {
	return s == HealthCheckStatusError
}

type HealthCheckRespose struct {
	Status  HealthCheckStatus `json:"status"`
	Details *string           `json:"details,omitempty"`
}

func NewHealthCheckRespose(status HealthCheckStatus) HealthCheckRespose {
	return HealthCheckRespose{
		Status: status,
	}
}

func (r HealthCheckRespose) WithDetails(details string) HealthCheckRespose {
	r.Details = &details
	return r
}
