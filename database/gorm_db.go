package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormConnection struct {
	Database *gorm.DB
}

func NewPostgres(dns string) (*GormConnection, error) {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, err
	}

	return &GormConnection{
		Database: db,
	}, nil
}

func (g *GormConnection) Connect() error {
	db, err := g.Database.DB()
	if err != nil {
		return fmt.Errorf("error al conectar a la base de datos: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error al conectar al hacer ping: %v", err)
	}

	fmt.Println("Conectado a la base de datos")
	return nil
}

func (g GormConnection) Disconnect() {
	if g.Database != nil {
		db, err := g.Database.DB()
		if err != nil {
			fmt.Println("Desconectado de la base de datos")
		}
		db.Close()
		fmt.Println("Desconectado de la base de datos")
	}
}

func (g *GormConnection) GetDB() interface{} {
	return g.Database
}
