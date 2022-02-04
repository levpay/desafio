package controller

import (
	"desafio/data"
	"desafio/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsertUUIDToSuper(t *testing.T) {
	heroWithoutUUID := models.HeroAPIResponse{
		Results: []models.HeroAPIResults{
			{
				Name:        "Joker Mock",
				Powerstats:  models.Powerstats{Intelligence: "10", Power: "10"},
				Biography:   models.Biography{FullName: "Joker Mock", Alignment: "bad"},
				Work:        models.Work{Occupation: "witcher"},
				Connections: models.Connections{GroupAffiliations: "-", Relatives: "-"},
				Image:       models.Image{URL: "jpg"},
			},
		},
	}

	got := insertUUIDtoSuper(heroWithoutUUID)

	if got[0].UUID == "" {
		t.Errorf("Test failed: expected something but got: %v", got[0].UUID)
	}
}

func TestCreateSuper(t *testing.T) {
	t.Run("Test creation controller", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		data.DB = db
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectExec("INSERT INTO supers").WithArgs().WillReturnResult(sqlmock.NewResult(1, 2))
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "localhost:3000/api/create?name=joker", nil)
		if err != nil {
			t.Errorf("Could not make new request: %s", err)
		}

		CreateSuper(recorder, request)
		got, _ := ioutil.ReadAll(recorder.Body)
		t.Log(string(got))

	})
}

func TestSearchSuper(t *testing.T) {
	t.Run("Test search controller", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		data.DB = db
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		rows := sqlmock.NewRows([]string{"uuid", "hero_name", "full_name", "alignment", "intelligence", "power", "occupation", "image", "group_connections", "relatives"}).
			AddRow("uuid", "joker mock", "joker mock", "bad", "1", "1", "mock", "jpg", "avengers", "geralt of rivia, mary jane")

		mock.ExpectQuery("SELECT").WithArgs("joker").WillReturnRows(rows)

		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "localhost:3000/api/search?name=joker", nil)
		if err != nil {
			t.Errorf("Could not make new request: %s", err)
		}

		SearchForSuper(recorder, request)
		got, _ := ioutil.ReadAll(recorder.Body)
		expected := []byte(`{"status":"success","results":[{"uuid":"uuid","name":"joker mock","powerstats":{"intelligence":"1","power":"1"},"biography":{"full-name":"joker mock","alignment":"bad"},"occupation":"mock","relatives-count":1,"group-affiliations":["avengers"],"image":"jpg"}]}`)

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Test failed: got %v but expected %v", string(got), string(expected))
		}
	})
}
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
