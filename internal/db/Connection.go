package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DBConfig contiene la configuración para la conexión a la base de datos
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// NewDBConfig crea una nueva configuración desde variables de entorno
func NewDBConfig() DBConfig {
	return DBConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   os.Getenv("MYSQL_DATABASE_NAME"),
	}
}


// Se establece una conexión con la base de datos
func Connect(config DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		config.User, config.Password, config.Host, config.Port, config.DBName)

	var db *sql.DB
	var err error
	maxRetries := 5
	delay := 2 * time.Second

	for i := range maxRetries {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Intento %d: error al abrir la conexión: %v\n", i+1, err)
			time.Sleep(delay)
			continue
		}

		// Configurar pool de conexiones
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(20)
		db.SetConnMaxLifetime(15 * time.Minute)
		db.SetConnMaxIdleTime(10 * time.Minute)

		if err = db.Ping(); err == nil {
			log.Print("Conexión a la base de datos establecida")
			return db, nil
		}
		log.Printf("Intento %d: error al hacer ping: %v\n", i+1, err)
		db.Close()
		time.Sleep(delay)
	}
	return db, fmt.Errorf("no se pudo conectar a la base de datos después de %d intentos: %v", maxRetries, err)
}
