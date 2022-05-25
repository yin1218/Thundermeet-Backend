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

// a successful case: user status for event timeblocks fetched succesfully
func TestDelete_timeblocks_from_event_ok(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	dao.InitializeTest(gdb)

	mock.ExpectExec(`DELETE from emails where user_id = $1 AND time_block_id = $2`).WithArgs("sherry", "2021-01-01T03:30:00+08:00A294").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = service.DeleteTimeblocksFromEvent(294, []string{"2021-01-01T03:30:00+08:00"}, "sherry")
	assert.Nil(t, err)
}
