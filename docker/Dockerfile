# Multi-stage Dockerfile para distintos entornos 
# Podria dejar solo el stage prod  y la base
#la imagen 1.23 tenia 6 vulnerabilidades, la 1.24 solo 2 en comparacion
FROM golang:1.24-bookworm AS base
WORKDIR /app
COPY go.* ./
RUN go mod download

# Desarrollo con Air - hot reload
FROM base AS dev
RUN go install github.com/cosmtrek/air@latest
COPY . .
CMD ["air"]

# Staging - ya compilamos la app
FROM base AS staging
COPY . .
RUN go build -o todoApp
CMD ["./todoApp"]

# Prod - usamos la base y compilamos la app desactivando CGO
# y optimizando el binario con -ldflags="-s -w"
# -o guarda el binario
FROM base AS builder
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o todoApp

# Imagen final - distroless
# distroless es una imagen practicamente sin nada, solo lo necesario para correr la app
# como no tiene nada, no tiene gclibc, por eso desactivamos cgo
FROM gcr.io/distroless/static-debian11 AS prod
COPY --from=builder /app/todoApp /
CMD ["/todoApp"]