package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func IniciarConexao() {
	dsn := fmt.Sprintf("host=%s user=root password=root dbname=root port=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"))
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panicf("Erro a conectar no banco de dados: %v\n", err)
	}
}
