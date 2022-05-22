package service

// import (
// 	"database/sql"
// 	"log"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestSelectOneEventOk(t *testing.T) {
// 	db, mock := NewMock()
// 	mock.ExpectQuery("SELECT firstname, lastname, email, age FROM Users;").
// 		WillReturnRows(sqlmock.NewRows([]string{"firstname", "lastname", "email", "age"}).
// 			AddRow("pepe", "guerra", "pepe@gmail.com", 34))
// 	subject := UserProvider{
// 		DatabaseProvider: repositories.NewMockDBProvider(db, nil),
// 	}
// 	resp, err := subject.GetUsers()

// 	assert.Nil(t, err)
// 	assert.NotNil(t, resp)
// 	assert.Equal(t, 1, len(resp))
// }

// func NewMock() (*sql.DB, sqlmock.Sqlmock) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}

// 	return db, mock
// }
