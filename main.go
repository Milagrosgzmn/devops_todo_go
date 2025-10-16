package main

import (
	"embed"
	"log"

	cfg "github.com/Milagrosgzmn/devops_todo_go.git/internal/config"
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/db"
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/repository"
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/routes"
	"github.com/joho/godotenv"
)

//go:embed internal/migrations/*.sql
var embedFS embed.FS
func main() {
	// cargamos el archivo .env (opcional en Docker, donde las vars vienen del compose)
	err := godotenv.Load()
	if err != nil {
		log.Println("WARN: .env no encontrado, se usaran variables de entorno del sistema");
	}
	// Configuramos la conexión a la base de datos
	config := db.NewDBConfig()
	// Nos conectamos
	dbInstance, err := db.Connect(config)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
	// Ejecutamos las migraciones antes de iniciar el servidor
	if err := cfg.RunMigrations(dbInstance, embedFS); err != nil {
		log.Fatalf("Error al ejecutar migraciones: %v", err)
	}

	// Inicializamos el repositorio
	var repo repository.IRepository = repository.NewItemMySqlRepository(dbInstance)

	// Configuramos el router
	router := routes.SetupRouter(repo);
	router.Run(":8080")
}