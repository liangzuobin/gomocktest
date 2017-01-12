package service

import (
	"fmt"
	"log"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/liangzuobin/gomocktest/double/dao"
)

var conn *dbr.Connection

func conf(driver, dsn string) (err error) {
	conn, err = dbr.Open(driver, dsn, nil)
	return
}

// NewSession new session with userID and mockDB
var NewSesssion = func(userID int64) *dbr.Session {
	log.Println("NewSession() with userID = ", userID)
	return conn.NewSession(nil)
}

func Service(modelID int64) (string, error) {
	sess := NewSesssion(0)
	tx, err := sess.Begin()
	if err != nil {
		return "", err
	}
	model, err := dao.ModelByID(tx, modelID)
	return fmt.Sprintf("%+v", model), err
}
