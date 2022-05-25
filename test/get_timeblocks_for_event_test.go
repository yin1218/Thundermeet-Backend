package timeblock_test

import (
	"testing"
	"thundermeet_backend/app/service"
	"time"

	"github.com/stretchr/testify/assert"

	"thundermeet_backend/app/dao"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// a successful case: timeblock fetched succesfully
func TestGet_timeblocks_for_event_ok(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	dao.InitializeTest(gdb)

	mock.ExpectQuery(`SELECT * FROM "timeblocks" WHERE event_id = $1`).WithArgs(294).
		WillReturnRows(sqlmock.NewRows([]string{"time_block_id", "event_id", "block_time"}).
			AddRow("2021-01-01T03:00:00+08:00A294", 294, time.Now()).
			AddRow("2021-01-01T03:30:00+08:00A294", 294, time.Now()))

	timeblocks, err := service.GetTimeblocksForEvent(294)

	assert.Nil(t, err)
	assert.NotNil(t, timeblocks)
	assert.Equal(t, 2, len(timeblocks))
}
