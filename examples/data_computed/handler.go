package data_computed

import (
	"net/http"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	Counter().Render(r.Context(), w)
}
