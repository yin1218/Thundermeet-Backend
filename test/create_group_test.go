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

func Create_Group_test_ok(t *testing.T) {
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
	mock.ExpectQuery(`INSERT INTO "groups" ("groupName ","user_id") VALUES ($1,$2)`).WithArgs("test", "sherry").WillReturnError(fmt.Errorf("event exist"))
	mock.ExpectRollback()

	group, err := service.CreateGroup("test", "sherry")

	assert.Nil(t, err)
	assert.NotNil(t, group)
}
