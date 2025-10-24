package constants

import "net/http"

// Códigos de estado HTTP
const (
	StatusOK                  = http.StatusOK                  // 200
	StatusCreated             = http.StatusCreated             // 201
	StatusNoContent           = http.StatusNoContent           // 204
	StatusBadRequest          = http.StatusBadRequest          // 400
	StatusNotFound            = http.StatusNotFound            // 404
	StatusInternalServerError = http.StatusInternalServerError // 500
)

// Mensajes de respuesta
const (
	// Mensajes de éxito
	ItemsObtenidos     = "Items obtenidos exitosamente"
	ItemObtenido       = "Item obtenido exitosamente"
	ItemCreado         = "Item creado exitosamente"
	ItemActualizado    = "Item actualizado exitosamente"
	ItemEliminado      = "Item eliminado exitosamente"

	// Mensajes de error
	IDInvalido         = "ID de item inválido"
	CuerpoInvalido     = "Cuerpo de la petición inválido"
	ItemNoEncontrado   = "Item no encontrado"
	ErrorInterno       = "Error interno del servidor"
	ErrorBaseDatos     = "Error en la base de datos"
	CamposRequeridos   = "Faltan campos requeridos"
)
