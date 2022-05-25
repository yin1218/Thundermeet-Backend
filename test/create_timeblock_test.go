package timeblock_test

import (
	"fmt"
	"testing"
	"thundermeet_backend/app/service"
	"time"

	"github.com/stretchr/testify/assert"

	"thundermeet_backend/app/dao"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// a successful case: timeblock created succesfully
func TestCreate_timeblock_ok(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	dao.InitializeTest(gdb)

	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	start_t := time.Date(2021, 1, 1, 9, 0, 0, 0, beijing)

	mock.ExpectQuery(`SELECT * FROM "timeblocks" WHERE time_block_id = $1`).WithArgs("2021-01-01T05:30:00+08:00A361").WillReturnRows(sqlmock.NewRows([]string{"time_block_id", "event_id", "block_time"}))
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "timeblocks" ("time_block_id","event_id","block_time") VALUES ($1,$2,$3)`).WithArgs("2021-01-01T05:30:00+08:00A361", 361, start_t).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = service.CreateOneTimeblock("2021-01-01T05:30:00+08:00A361", 361, start_t)

	assert.Nil(t, err)
}

// a fail case: timeblock already exist
func TestCreate_timeblock_fail(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	dao.InitializeTest(gdb)

	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	start_t := time.Date(2021, 1, 1, 9, 0, 0, 0, beijing)

	mock.ExpectQuery(`SELECT * FROM "timeblocks" WHERE time_block_id = $1`).WithArgs("2021-01-01T05:30:00+08:00A361").WillReturnError(fmt.Errorf("timeblock exist"))

	err = service.CreateOneTimeblock("2021-01-01T05:30:00+08:00A361", 361, start_t)

	assert.Error(t, err)
}
