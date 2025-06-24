package clock_raf

import (
	"net/http"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	ClockRafExample().Render(r.Context(), w)
}
