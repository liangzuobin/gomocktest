package service

import (
	"database/sql"
	"log"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	"github.com/stretchr/testify/assert"
)

func MockNewSession(db *sql.DB) func(int64) *dbr.Session {
	return func(int64) *dbr.Session {
		conn := &dbr.Connection{
			Dialect:       dialect.MySQL,
			DB:            db,
			EventReceiver: &dbr.NullEventReceiver{},
		}
		return &dbr.Session{
			Connection:    conn,
			EventReceiver: conn.EventReceiver,
		}
	}
}

func TestService(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// important
	NewSesssion = MockNewSession(db)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT (.+) FROM model").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("1", "foo"))
	mock.ExpectCommit()

	str, err := Service(100)
	log.Println("model = ", str)
	assert.NoError(t, err)
}
