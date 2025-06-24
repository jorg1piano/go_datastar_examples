package execute_script

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	datastar "github.com/starfederation/datastar/sdk/go"
)

type ExecuteScriptSignal struct {
	Message string `json:"message"`
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	ExecuteScriptExample().Render(r.Context(), w)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	signal := &ExecuteScriptSignal{}
	datastar.ReadSignals(r, signal)

	sse := datastar.NewSSE(w, r)
	var furElise = []string{"E4", "D#4", "E4", "D#4", "E4", "B3", "D4", "C4", "A3"}

	for _, note := range furElise {
		if note != "-" {
			// Trigger the playNote script we defined in excute_script.templ
			sse.ExecuteScript(fmt.Sprintf("playNote('%s')", note))

			// Let the user know that this script is triggered from the server
			signal.Message = fmt.Sprintf("%s triggered on the server", note)
			jsonStr, _ := json.Marshal(signal)
			sse.MergeSignals(jsonStr)
		}
		time.Sleep(200 * time.Millisecond)
	}
	time.Sleep(200 * time.Millisecond)

	jsonStr, _ := json.Marshal(ExecuteScriptSignal{Message: "Fur Elise is on fire!"})
	sse.MergeSignals(jsonStr)
	time.Sleep(2500 * time.Millisecond)

	jsonStr, _ = json.Marshal(ExecuteScriptSignal{Message: "Click me"})
	sse.MergeSignals(jsonStr)
}
