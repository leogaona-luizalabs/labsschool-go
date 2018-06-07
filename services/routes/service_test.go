package routes_test

import (
	"errors"
	"testing"

	"github.com/luizalabs/go-training/services/routes"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestListOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		t.Fail()
	}

	defer db.Close()

	columns := []string{"route_key", "modality_id", "threshold", "create_date", "update_date"}
	mock.ExpectQuery("select ").WithArgs(50, 0).WillReturnRows(sqlmock.NewRows(columns).AddRow("400-20000000-23799999", 1, 35, "2018-05-21 13:58:19", "2018-05-21 13:58:19"))

	service := routes.New(db)
	routes, err := service.List(50, 0)

	//assertions
	if err != nil {
		t.Errorf("error returned from service: %s", err.Error())
		t.Fail()
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("error returned from mock: %s", err.Error())
		t.Fail()
	}

	if len(routes) != 1 {
		t.Errorf("wrong len. Expected %d but got %d", 1, len(routes))
		t.Fail()
	}
}

func TestListError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		t.Fail()
	}

	defer db.Close()

	mock.ExpectQuery("select ").WithArgs(50, 0).WillReturnError(errors.New("table not found"))

	service := routes.New(db)
	routes, err := service.List(50, 0)

	//assertions
	if err == nil {
		t.Error("no error returned from service")
		t.Fail()
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("error returned from mock: %s", err.Error())
		t.Fail()
	}

	if len(routes) != 0 {
		t.Errorf("wrong len. Expected %d but got %d", 0, len(routes))
		t.Fail()
	}
}
