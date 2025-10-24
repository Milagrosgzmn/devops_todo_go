package main

import (
	"embed"
	"log"

	cfg "github.com/Milagrosgzmn/devops_todo_go.git/internal/config"
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/db"
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/repository"
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/routes"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/newrelic"
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

	// Inicializamos New Relic
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.GetEnv("NEW_RELIC_APP_NAME", "devops_todo_app_go")),
		newrelic.ConfigLicense(cfg.GetEnv("NEW_RELIC_LICENSE_KEY", "")),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		log.Printf("WARN: Fallo la inicialización de New Relic: %v. Continuando sin monitoreo.\n", err)
	} else {
		log.Println("INFO: New Relic inicializo correctamente.")
	}

	// Configuramos el router con New Relic
	router := routes.SetupRouter(repo, app);
	router.Run(":8080")
}