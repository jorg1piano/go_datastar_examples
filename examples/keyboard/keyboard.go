package keyboard

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	KeyboardLayout().Render(r.Context(), w)
}
