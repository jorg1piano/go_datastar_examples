package count_until

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	datastar "github.com/starfederation/datastar/sdk/go"
)

type CountUntilSignal struct {
	Count       int    `json:"count"`
	IncrementTo int    `json:"incrementTo"`
	Error       string `json:"error"`
}

func PutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Read the signal
	signal := &CountUntilSignal{}
	if err := datastar.ReadSignals(r, signal); err != nil {
		log.Printf("Failed to read signals: %v", err)
		datastar.NewSSE(w, r).MergeSignals([]byte(`{"error": "Failed to read counter state"}`))
		return
	}

	// Return an error if the increment target is to low
	if signal.Count >= signal.IncrementTo {
		datastar.NewSSE(w, r).MergeSignals([]byte(`{"error": "Please choose a higher increment target"}`))
		return
	}

	// every 500ms, increment the counter until it reaches the target
	for signal.Count <= signal.IncrementTo {
		json, err := json.Marshal(signal)
		if err != nil {
			log.Printf("Failed to marshal signal: %v", err)
			datastar.NewSSE(w, r).MergeSignals([]byte(`{"error": "Failed to update counter"}`))
		} else {
			datastar.NewSSE(w, r).MergeSignals(json)
		}
		time.Sleep(500 * time.Millisecond) // Simulate a slow increment
		signal.Count++
	}
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	Example().Render(r.Context(), w)
}
