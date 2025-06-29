package signals_and_the_dom

import (
	"learning_datastar/util"
	"net/http"

	datastar "github.com/starfederation/datastar/sdk/go"
)

type PersonaliaSignal struct {
	Name     string `json:"name"`
	JobTitle string `json:"jobTitle"`
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	Example(PersonaliaSignal{}).Render(r.Context(), w)
}

func MorphHandlerTargetUsingIfMissing(w http.ResponseWriter, r *http.Request) {
	// This handler replaces the DOM element with a new one
	// containing the updated signals.
	signal := &PersonaliaSignal{}
	datastar.ReadSignals(r, signal)

	datastar.NewSSE(w, r).MergeFragmentTempl(
		DataSignalsIfMissing(PersonaliaSignal{}, util.RandomColor()), // This will not include the updated signal, since we use __ifmissing
		datastar.WithMergeMorph(),
		datastar.WithSelectorID(morphTargetIDWithIfMissing),
	)
}

func MorphHandlerReplaceWith(w http.ResponseWriter, r *http.Request) {
	// This handler replaces the DOM element with a new one
	// containing the updated signals.
	signal := &PersonaliaSignal{}
	datastar.ReadSignals(r, signal)

	datastar.NewSSE(w, r).MergeFragmentTempl(
		DataSignals(PersonaliaSignal{}, util.RandomColor()), // This will use the passed in signal and effectively replace it on the client
		datastar.WithMergeMorph(),
		datastar.WithSelectorID(morphTargetIDWithoutIfMissing),
	)
}
