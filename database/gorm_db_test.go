package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGormDbConnection(t *testing.T) {
	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	connection, err := NewPostgres(db)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = connection.Connect()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	instance := connection.GetDB()
	if instance == nil {
		t.Errorf("Expected no error, got %v", err)
	}

	connection.Disconnect()

}
