package controller

import (
	"desafio/data"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestDeleteSuper(t *testing.T) {
	db, mock, err := sqlmock.New()
	data.DB = db

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("DELETE")).WillReturnResult(sqlmock.NewResult(1, 1))

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "localhost:3000/api/delete?uuid=1234asdf", nil)
	if err != nil {
		t.Errorf("Could not make new request: %s", err)

	}
	DeleteSuper(recorder, request)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
