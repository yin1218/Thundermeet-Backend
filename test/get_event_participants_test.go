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

func get_event_participants_ok(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	dao.InitializeTest(gdb)

	mock.ExpectQuery(`SELECT * FROM "events" WHERE event_id = $1`).WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"participants"}).
			AddRow("sherry").
			AddRow("wpbag"))

	event, err := service.GetEventParticipants(1)

	// fmt.Print(event)

	assert.Nil(t, err)
	assert.NotNil(t, event)
}
