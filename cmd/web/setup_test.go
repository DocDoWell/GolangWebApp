package main

import (
	"os"
	"testing"
	"web-app/models"
)

var testApp application

func TestMain(m *testing.M) {

	testApp = application{
		Models: *models.New(nil),
	}

	os.Exit(m.Run())
}
