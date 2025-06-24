package index

import "net/http"

func PageHandler(w http.ResponseWriter, r *http.Request) {
	Page().Render(r.Context(), w)
}
