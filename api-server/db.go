package main

import (
	"regexp"
	"time"

	models_db "github.com/birsanion/netopia/api-server/models/db"
	"github.com/go-sql-driver/mysql"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	addrWithProto = regexp.MustCompile(`^(\w+)\((.*)\)$`)
)

func NewDbConnection(cfg Config) (*gorm.DB, error) {
	network, address, err := parseAddress(cfg.DBHost)
	if err != nil {
		return nil, err
	}

	config := mysql.NewConfig()
	config.User = cfg.DBUser
	config.Net = network
	config.Addr = address
	config.Passwd = cfg.DBPassword
	config.DBName = cfg.DBName
	config.Collation = "utf8mb4_unicode_ci"
	config.ParseTime = true
	config.Loc = time.UTC

	db, err := gorm.Open(gmysql.Open(config.FormatDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models_db.Payment{}); err != nil {
		return nil, err
	}

	sdb, err := db.DB()
	if err != nil {
		return db, err
	}

	err = sdb.Ping()
	if err != nil {
		sdb.Close()
		return nil, err
	}

	return db, nil
}

func IsDBAvailable(db *gorm.DB) bool {
	if db == nil {
		return false
	}

	sdb, err := db.DB()
	if err != nil {
		return false
	}

	if err = sdb.Ping(); err != nil {
		return false
	}

	return true
}

func parseAddress(str string) (network, address string, err error) {
	m := addrWithProto.FindStringSubmatch(str)
	if m == nil {
		network = "tcp"
		address = str
		return
	}

	network = m[1]
	address = m[2]
	return
}
