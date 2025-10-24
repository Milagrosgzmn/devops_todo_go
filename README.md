# DevOps Todo Go - API REST con CI/CD

![Go Version](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?logo=docker)
![CI/CD](https://img.shields.io/badge/CI%2FCD-GitHub%20Actions-2088FF?logo=github-actions)
![License](https://img.shields.io/badge/License-MIT-green)

API REST desarrollada en Go con Gin framework, implementando un sistema completo de CI/CD, contenerización Docker y monitoreo con New Relic.

**Proyecto desarrollado para el Trabajo Práctico Integrador de DevOps.**

---

## 📋 Tabla de Contenidos

- [Descripción](#-descripción)
- [Arquitectura](#-arquitectura)
- [Tecnologías Utilizadas](#-tecnologías-utilizadas)
- [Características](#-características)
- [Requisitos Previos](#-requisitos-previos)
- [Instalación y Configuración](#-instalación-y-configuración)
- [Uso](#-uso)
- [API Endpoints](#-api-endpoints)
- [Testing](#-testing)
- [Docker](#-docker)
- [CI/CD Pipeline](#-cicd-pipeline)
- [Monitoreo](#-monitoreo)
- [Estructura del Proyecto](#-estructura-del-proyecto)


---

##  Descripción

Este proyecto implementa una **API REST completa** para gestión de tareas (Todo Items) con las siguientes capacidades:

- **CRUD completo** de items (Create, Read, Update, Delete)
- **Health check endpoint** para monitoreo
- **Validaciones de estados** (pending, in_progress, completed)
- **Persistencia en MySQL** con migraciones automáticas
- **Testing unitario** con cobertura >60%
- **Contenerización** con Docker multi-stage
- **CI/CD completo** con GitHub Actions
- **Monitoreo APM** con New Relic

---

## 🏗️ Arquitectura



### Componentes:

1. **API Layer**: Gin framework con handlers RESTful
2. **Business Logic**: Repository pattern para acceso a datos
3. **Data Layer**: MySQL con migraciones Goose
4. **Monitoring**: New Relic APM para métricas en tiempo real

---

## 🛠️ Tecnologías Utilizadas

### Backend
- **Go 1.24** - Lenguaje de programación
- **Gin 1.9.1** - Framework web HTTP
- **MySQL 8.0** - Base de datos relacional
- **Goose** - Manejo de migraciones de BD

### DevOps
- **Docker** - Contenerización (multi-stage build)
- **Docker Compose** - Orquestación de contenedores
- **GitHub Actions** - CI/CD automation
- **New Relic** - Application Performance Monitoring

### Testing
- **Go testing** - Framework nativo de testing

---

## ✨ Características

### Funcionales
- CRUD completo de Todo Items
- Validación de estados (pending, in_progress, completed)
- Migraciones de base de datos automáticas
- Health check endpoint

### DevOps
- Dockerfile multi-stage (optimizado para producción)
- Docker Compose para desarrollo local
- Pipeline CI: Lint + Tests + Build
- Pipeline CD: Push automático a Docker Hub
- Branch Protection en rama main
- Monitoreo APM con New Relic

### Testing
- Tests unitarios de handlers
- Tests de validaciones
- Cobertura >60%

---

## 📦 Requisitos Previos

- **Go 1.24+** - [Descargar](https://golang.org/dl/)
- **Docker 20.10+** - [Descargar](https://docs.docker.com/get-docker/)
- **Docker Compose 2.0+** - [Descargar](https://docs.docker.com/compose/install/)
- **Git** - [Descargar](https://git-scm.com/downloads)

---

## 🚀 Instalación y Configuración

### 1. Clonar el repositorio

```bash
git clone https://github.com/Milagrosgzmn/devops_todo_go.git
cd devops_todo_go
```

### 2. Configurar variables de entorno

Crea un archivo `.env` en el directorio `docker/`:

```bash
cd docker
cp .env.example .env
```

Edita el archivo `.env` con tus valores:

```env
# Base de datos
MYSQL_HOST=
MYSQL_PORT=
MYSQL_USER=
MYSQL_PASSWORD=
MYSQL_DATABASE_NAME=

# Zona horaria
TZ=

# New Relic (opcional para desarrollo local)
NEW_RELIC_APP_NAME=
NEW_RELIC_LICENSE_KEY=tu_license_key_aqui
```

### 3. Levantar los servicios con Docker Compose

```bash
cd docker
docker compose up --build
```

La aplicación estará disponible en: **http://localhost:8080**

### 4. Verificar que funciona

```bash
curl http://localhost:8080/health
# Respuesta esperada: {"status":"OK"}
```

---

## 💻 Uso

### Desarrollo local (sin Docker)

```bash
# Instalar dependencias
go mod download

# Configurar variables de entorno
export MYSQL_HOST=localhost
export MYSQL_PORT=3306
export MYSQL_USER=root
export MYSQL_PASSWORD=password
export MYSQL_DATABASE_NAME=todo_db

# Ejecutar la aplicación
go run main.go
```

### Con Docker Compose (recomendado)

```bash
# Levantar todos los servicios
cd docker
docker compose up -d

# Ver logs
docker compose logs -f app

# Detener servicios
docker compose down

# Reconstruir imagen
docker compose up --build
```

---

## API Endpoints

### Health Check

```http
GET /health
```

**Respuesta:**
```json
{
  "status": "OK"
}
```

### Listar todos los items

```http
GET /items
```

**Respuesta:**
```json
[
  {
    "id": 1,
    "title": "Tarea 1",
    "description": "Descripción de la tarea",
    "completed": false,
    "status": "pending",
    "created_at": "2025-10-24T05:00:00Z",
    "updated_at": "2025-10-24T05:00:00Z"
  }
]
```

### Obtener un item por ID

```http
GET /items/:id
```

**Respuesta:**
```json
{
  "id": 1,
  "title": "Tarea 1",
  "description": "Descripción de la tarea",
  "completed": false,
  "status": "pending",
  "created_at": "2025-10-24T05:00:00Z",
  "updated_at": "2025-10-24T05:00:00Z"
}
```

### Crear un nuevo item

```http
POST /items
Content-Type: application/json

{
  "title": "Nueva tarea",
  "description": "Descripción de la nueva tarea",
  "completed": false,
  "status": "pending"
}
```

**Estados válidos:** `pending`, `in_progress`, `completed`

### Actualizar un item

```http
PUT /items/:id
Content-Type: application/json

{
  "title": "Tarea actualizada",
  "description": "Descripción actualizada",
  "completed": true,
  "status": "completed"
}
```

### Eliminar un item

```http
DELETE /items/:id
```

---

## 🧪 Testing

### Ejecutar todos los tests

```bash
go test ./...
```

### Tests con cobertura

```bash
go test -cover ./...
```

### Tests verbose

```bash
go test -v ./...
```

### Tests con reporte de cobertura HTML

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

---

## 🐳 Docker

### Dockerfile Multi-Stage

El proyecto utiliza un **Dockerfile multi-stage** optimizado:

1. **Stage 1 (base)**: Descarga dependencias
2. **Stage 2 (builder)**: Compila el binario (con CGO_ENABLED=0 y ldflags optimizados)
3. **Stage 3 (prod)**: Imagen final con Distroless (ultra-ligera)

**Resultado:** Imagen final de ~15-30 MB

### Construcción manual

```bash
# Construir imagen
docker build -f docker/Dockerfile -t devops-todo:latest .

# Ejecutar contenedor
docker run -p 8080:8080 --env-file docker/.env devops-todo:latest
```

### Docker Compose

Orquesta dos servicios:

- **app**: Aplicación Go (construida desde Dockerfile)
- **db**: MySQL 8.0 con healthcheck

```bash
# Servicios disponibles
docker compose ps

# Logs de un servicio específico
docker compose logs app
docker compose logs db

# Ejecutar comandos en contenedor
docker compose exec app sh
docker compose exec db mysql -u root -p
```

---

## 🔄 CI/CD Pipeline

### Arquitectura del Pipeline

```
┌─────────────┐
│ Push/PR     │
│ al repo     │
└──────┬──────┘
       │
       ▼
┌─────────────────────────────────┐
│   CI Workflow (Pull Requests)   │
│                                  │
│  1. Lint (go vet)       │
│  2. Test (Go 1.23, 1.24)        │
│  3. Build Docker Image          │
└──────┬──────────────────────────┘
       │
       │ (PR aprobado & mergeado)
       ▼
┌─────────────────────────────────┐
│  Publicacion a  dockerhub       │
│  Workflow (Push a main)        │
│                                  │
│  1. Lint                         │
│  2. Test                         │
│  3. Build & Push a Docker Hub  │
│     - Tag: latest                │
│     - Tag: sha-COMMIT_HASH       │
└──────────────────────────────────┘
```

### Workflows

#### 1. **CI - Continuous Integration** (`.github/workflows/ci.yml`)

**Trigger:** Pull requests a `main` y push a otras ramas

**Jobs:**
- **Lint**: `go vet` check
- **Test**: Matrix con Go 1.23 y 1.24
- **Build**: Construcción de imagen Docker (sin push)

#### 2. **CI - Docker Publish** (`.github/workflows/ci-docker-publish.yml`)

**Trigger:** Push a `main` (después de merge de PR)

**Jobs:**
- **Lint**: Verificación de código
- **Test**: Ejecución de tests
- **Build and Push**: Construcción y publicación en Docker Hub
  - Tag `latest` (siempre actualizado)
  - Tag `sha-XXXXXXX` (histórico inmutable)

### Branch Protection

Configuración en rama `main`:

-  Requiere Pull Request antes de merge
-  Requiere aprobaciones (removido -0 aprobaciones dado que es un proyecto individual)
-  Requiere que pasen los status checks (Lint, Test, Build)
-  No permite bypass de settings
-  No permite force push
-  No permite deletion

---

## 📊 Monitoreo

### New Relic APM

El proyecto está integrado con **New Relic** para monitoreo de:

- **Transactions**: Tiempos de respuesta por endpoint
- **Throughput**: Requests por minuto
- **Error Rate**: Tasa de errores
- **Response Times**: Latencia de la aplicación
- **CPU & Memory**: Uso de recursos del sistema

#### Configuración

Variables de entorno requeridas:

```env
NEW_RELIC_APP_NAME= tu_app_name_aqui
NEW_RELIC_LICENSE_KEY=your_license_key_here
```

#### Dashboards disponibles

- **Overview**: Dashboard principal con métricas generales
- **Transactions**: Análisis detallado por endpoint
- **Errors**: Tracking de errores y excepciones


### Seguridad
-  **Secrets management** con GitHub Secrets
-  **.env files** no commiteados (.gitignore)
-  **No hardcoded credentials**

---

## 🔗 Links

- **Repositorio GitHub**: https://github.com/Milagrosgzmn/devops_todo_go
- **Docker Hub**: https://hub.docker.com/r/milagrosgzmn/milagrosgzmn_work_repository
- **New Relic Dashboard**: https://one.newrelic.com/

### Documentación de referencia

- [Go Documentation](https://go.dev/doc/)
- [Gin Framework](https://gin-gonic.com/docs/)
- [Docker Documentation](https://docs.docker.com/)
- [GitHub Actions](https://docs.github.com/en/actions)
- [New Relic Go Agent](https://docs.newrelic.com/docs/apm/agents/go-agent/)
- [Goose Migrations](https://pressly.github.io/goose/)

## Otras herramientas utilizadas

- **Postman** - Para pruebas de API
- **Copilot** - Asistente de codificación AI para crear los test que imitan un deadlock de la base de datos.
- **Claude code** - Asistente de codificación AI para generar el read me y verificar el código.