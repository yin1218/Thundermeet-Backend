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

func TestSelect_one_event_ok(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	dao.InitializeTest(gdb)

	mock.ExpectQuery(`SELECT * FROM "events" WHERE event_id = $1 ORDER BY "events"."event_id" LIMIT 1`).WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"event_id", "event_name"}).
			AddRow(1, "test event"))

	event, err := service.SelectOneEvent(1)

	// fmt.Println(err)
	// fmt.Println(event)

	assert.Nil(t, err)
	assert.NotNil(t, event)
}
