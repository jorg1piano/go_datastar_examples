package dialog

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	datastar "github.com/starfederation/datastar/sdk/go"
)

type QuoteSignal struct {
	Quote string `json:"quote"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sse := datastar.NewSSE(w, r)

	// Stream a loading animation over SSE just for fun
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	showLoadingAnimation(ctx, sse)

	// Return the actual qoute
	quote := getRandomQuote()
	data, _ := json.Marshal(quote)
	sse.MergeSignals(data)
}

func getRandomQuote() QuoteSignal {
	quotes := []string{
		"Beautiful is better than ugly.",
		"Explicit is better than implicit.",
		"Simple is better than complex.",
		"Complex is better than complicated.",
		"Flat is better than nested.",
		"Sparse is better than dense.",
		"Readability counts.",
		"Special cases aren't special enough to break the rules.",
	}

	// Optionally seed once in init
	randomIndex := rand.Intn(len(quotes))
	return QuoteSignal{Quote: quotes[randomIndex]}
}

func showLoadingAnimation(ctx context.Context, sse *datastar.ServerSentEventGenerator) {
	loaderChars := []string{"|", "/", "-", "\\"}
	ticker := time.NewTicker(80 * time.Millisecond)
	defer ticker.Stop()

	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			char := loaderChars[i%len(loaderChars)]
			data, _ := json.Marshal(QuoteSignal{Quote: char})
			sse.MergeSignals(data)
			i++
		}
	}
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	Example().Render(r.Context(), w)
}
