package data_persist

import (
	"net/http"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	ExampleDataPersist().Render(r.Context(), w)
}
func PageHandlerVanillaJS(w http.ResponseWriter, r *http.Request) {
	ExampleVanillaJS().Render(r.Context(), w)
}
