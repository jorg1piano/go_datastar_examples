package deriving_signals_for_server

import (
	"fmt"
	"net/http"

	datastar "github.com/starfederation/datastar/sdk/go"
)

type CounterSignal struct {
	Count     int `json:"count"`
	Double    int `json:"double"`
	Triple    int `json:"triple"`
	Quadruple int `json:"quadruple"` // Added for completeness, not used in the example

	ServerMessageSignal
}

type ServerMessageSignal struct {
	ServerMessage string `json:"serverMessage"`
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	Example().Render(r.Context(), w)
}

func ProcessCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	s := &CounterSignal{}
	datastar.ReadSignals(r, s)

	// Merge in the updated $serverMessage signal
	datastar.NewSSE(w, r).MarshalAndMergeSignals(&ServerMessageSignal{
		ServerMessage: fmt.Sprintf("Received count: %d, double: %d, triple: %d, quadruple: %d", s.Count, s.Double, s.Triple, s.Quadruple),
	})
}
