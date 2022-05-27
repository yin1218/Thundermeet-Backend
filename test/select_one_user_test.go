package timeblock_test

import (
	"testing"
	"thundermeet_backend/app/service"

	"github.com/stretchr/testify/assert"

	"thundermeet_backend/app/dao"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestSelect_one_user_ok(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{

		Conn: db,
	}), &gorm.Config{})

	dao.InitializeTest(gdb)

	mock.ExpectQuery(`SELECT user_Id,user_Name,password_Hash,password_Answer FROM "users" WHERE User_Id=$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`).WithArgs("sherry").
		WillReturnRows(sqlmock.NewRows([]string{"user_Id", "user_Name", "password_Hash", "password_Answer"}).
			AddRow("sherry", "ChristineWang", "password", "NCCU"))

	user, err := service.SelectOneUser("sherry")

	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestSelect_one_user_fail(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{

		Conn: db,
	}), &gorm.Config{})

	dao.InitializeTest(gdb)

	mock.ExpectQuery(`SELECT user_Id,user_Name,password_Hash,password_Answer FROM "users" WHERE User_Id=$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`).WithArgs("sherry").
		WillReturnRows(sqlmock.NewRows([]string{"user_Id", "user_Name", "password_Hash", "password_Answer"}).
			AddRow("sherry", "ChristineWang", "password", "NCCU"))

	user, err := service.SelectOneUser("sherry")

	assert.Nil(t, err)
	assert.NotNil(t, user)
}
