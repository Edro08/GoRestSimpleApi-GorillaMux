package test

import (
	"GoRestSimpleApi/kit/platform/storage"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	mysqlURI := storage.GetmysqlURI()
	_, err := storage.SetupStorage(mysqlURI)

	if err != nil {
		t.Errorf(err.Error())
	}
}
