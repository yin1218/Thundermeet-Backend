package timeblock_test

import (
	"fmt"
	"testing"
	"thundermeet_backend/app/service"

	"github.com/stretchr/testify/assert"

	"thundermeet_backend/app/dao"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func register_user_ok(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{

		Conn: db,
	}), &gorm.Config{})

	dao.InitializeTest(gdb)

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users" ("user_Id", "user_Name", "password_Hash", "password_Answer") VALUES ($1,$2,$3,$4)`).WithArgs("viola20", "viola", "1234", "NTU")
	mock.ExpectRollback()
	err = service.RegisterOneUser("viola20", "viola", "1234", "NTU")

	assert.Nil(t, err)

}
func register_user_fail(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{

		Conn: db,
	}), &gorm.Config{})

	dao.InitializeTest(gdb)

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users" ("user_Id", "user_Name", "password_Hash", "password_Answer") VALUES ($1,$2,$3,$4)`).WithArgs("viola20", "viola", "1234", "NTU").
		WillReturnError(fmt.Errorf("event exist"))
	mock.ExpectRollback()
	err = service.RegisterOneUser("viola20", "viola", "1234", "NTU")

	assert.Nil(t, err)

}
