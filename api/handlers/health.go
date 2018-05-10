package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthCheck handler da rota GET /health
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Setando um header http de resposta
	w.Header().Set("content-type", "application/json")

	// Gerando um objeto customizado à partir de um map, e o convertendo em json
	response, _ := json.Marshal(map[string]interface{}{
		"status": "up",
	})

	// Write escreve o conteúdo do slice de bytes no corpo da resposta
	w.Write(response)
	// WriteHeader seta o status code da resposta. É importante frisar que ele só pode ser chamado
	// uma única vez no contexto da resposta. Chamadas subsequentes são ignoradas, portanto convém
	// chamar essa função quando você estiver prestes a retornar do handler
	w.WriteHeader(http.StatusOK)
	return
}
