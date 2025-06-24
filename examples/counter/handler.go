package counter

import (
	"net/http"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	Counter().Render(r.Context(), w)
}

func CounterIncrementPutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Counter incremented successfully"))
}
