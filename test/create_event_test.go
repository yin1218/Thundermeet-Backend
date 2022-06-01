package timeblock_test

import (
	"fmt"
	"testing"
	"thundermeet_backend/app/service"

	"github.com/stretchr/testify/assert"

	"thundermeet_backend/app/dao"

	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Create_Event_test_ok(t *testing.T) {
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
	mock.ExpectQuery(`INSERT INTO "events" ("eventName","isPriorityEnabled", "isConfirmed", "startTime", "endTime", "dateOrDays", "startDate", "endDate", "adminId") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`).WithArgs("test", false, false, "1100", "2000", "2022-05-29 08:00:00", "2022-05-31 08:00:00", false, "sherry").WillReturnError(fmt.Errorf("event exist"))
	mock.ExpectRollback()

	startDate, _ := time.Parse("2006-01-02", "2022-05-29")
	endDate, _ := time.Parse("2006-01-02", "2022-05-31")
	isPriorityEnabled := new(bool)
	*isPriorityEnabled = true
	dateOrDays := new(bool)
	*dateOrDays = false

	event, err := service.CreateEvent("test", isPriorityEnabled, "1100", "2000", dateOrDays, "", "", startDate, endDate, "sherry", "")

	assert.Nil(t, err)
	assert.NotNil(t, event)
}
