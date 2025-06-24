package data_on_and_data_attr

import (
	"net/http"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	DataOnExample().Render(r.Context(), w)
}
