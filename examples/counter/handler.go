package counter

import (
	"encoding/json"
	"log"
	"net/http"

	datastar "github.com/starfederation/datastar/sdk/go"
)

type CounterSignal struct {
	Count       int `json:"count"`
	IncrementBy int `json:"incrementBy"`
}

func CounterIncrementPutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Read the signal
	signal := &CounterSignal{}
	if err := datastar.ReadSignals(r, signal); err != nil {
		log.Printf("Failed to read signals: %v", err)
		datastar.NewSSE(w, r).MergeSignals([]byte(`{"error": "Failed to read counter state"}`))
		return
	}

	// Increment the counter
	signal.Count += signal.IncrementBy

	json, err := json.Marshal(signal)
	if err != nil {
		log.Printf("Failed to marshal signal: %v", err)
		datastar.NewSSE(w, r).MergeSignals([]byte(`{"error": "Failed to update counter"}`))
		return
	}

	// Send the updated signal back to the client
	datastar.NewSSE(w, r).MergeSignals(json)
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	Counter().Render(r.Context(), w)
}
