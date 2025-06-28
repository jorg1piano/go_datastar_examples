package binding_lists

import (
	"net/http"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	openingHours := OpeningHoursSignals{
		[]OpeningHours{
			{Day: "Monday", Open: "09:00", Close: "17:00"},
			{Day: "Tuesday", Open: "10:00", Close: "16:00"},
			{Day: "Wednesday", Open: "11:00", Close: "15:00"},
			{Day: "Thursday", Open: "12:00", Close: "14:00"},
			{Day: "Friday", Open: "13:00", Close: "13:30"},
			{Day: "Saturday", Open: "14:00", Close: "15:00"},
			{Day: "Sunday", Open: "15:00", Close: "16:00"},
		},
	}
	OpeningHoursExample(openingHours).Render(r.Context(), w)
}
