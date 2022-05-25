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

// a successful case: timeblockparticipant created succesfully
func TestCreate_Timeblock_Participant_ok(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	dao.InitializeTest(gdb)

	mock.ExpectQuery(`SELECT * FROM "timeblock_participants" WHERE user_id = $1 AND time_block_id = $2`).WithArgs("sherry", "2021-01-01T09:00:00+08:00A294").WillReturnRows(sqlmock.NewRows([]string{"time_block_id", "priority", "user_id"}))
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "timeblock_participants" ("user_id","time_block_id","priority") VALUES ($1,$2,$3)`).WithArgs("sherry", "2021-01-01T09:00:00+08:00A294", false).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = service.CreateOneTimeblockParticipant("sherry", "2021-01-01T09:00:00+08:00A294", false)

	assert.Nil(t, err)
}

// a successful case: timeblockparticipant created succesfully
func TestCreate_Timeblock_Participant_fail(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	dao.InitializeTest(gdb)

	mock.ExpectQuery(`SELECT * FROM "timeblock_participants" WHERE user_id = $1 AND time_block_id = $2`).WithArgs("sherry", "2021-01-01T09:00:00+08:00A294").WillReturnError(fmt.Errorf("timeblock participant exist"))

	err = service.CreateOneTimeblockParticipant("sherry", "2021-01-01T09:00:00+08:00A294", false)

	assert.Error(t, err)
}
