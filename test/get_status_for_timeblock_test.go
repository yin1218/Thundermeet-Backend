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

// a successful case: user status for event timeblocks fetched succesfully
func TestGet_status_for_timeblock_ok(t *testing.T) {
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

	mock.ExpectQuery(`SELECT * FROM "timeblock_participants" WHERE time_block_id = $1 AND user_id = $2`).WithArgs("2021-01-01T03:00:00+08:00A294", "sherry").
		WillReturnRows(sqlmock.NewRows([]string{"time_block_id", "priority", "user_id"}).
			AddRow("2021-01-01T03:00:00+08:00A294", false, "sherry"))

	mock.ExpectQuery(`SELECT * FROM "timeblock_participants" WHERE time_block_id = $1 AND user_id = $2`).WithArgs("2021-01-01T03:30:00+08:00A294", "sherry").
		WillReturnRows(sqlmock.NewRows([]string{"time_block_id", "priority", "user_id"}).
			AddRow("2021-01-01T03:30:00+08:00A294", true, "sherry"))

	normal, priority, err := service.GetStatusForTimeblock("sherry", 294)

	fmt.Print(priority)

	assert.Nil(t, err)
	assert.NotNil(t, normal)
	assert.NotNil(t, priority)
	assert.Equal(t, 1, len(normal))
	assert.Equal(t, 1, len(priority))
}
