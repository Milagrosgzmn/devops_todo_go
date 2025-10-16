package constants

// Estados válidos para un item TODO
const (
	StatePending    = "pending"
	StateInProgress = "in_progress"
	StateCompleted  = "completed"
)

// Mensajes de error relacionados con estados
const (
	EstadoInvalido = "estado inválido."
)

// ValidStates contiene todos los estados válidos
var ValidStates = []string{StatePending, StateInProgress, StateCompleted}

// IsValidState verifica si un estado es válido
func IsValidState(state string) bool {
	for _, validState := range ValidStates {
		if state == validState {
			return true
		}
	}
	return false
}
