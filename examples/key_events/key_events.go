package key_events

import (
	"net/http"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	keyEventsExample().Render(r.Context(), w)
}
