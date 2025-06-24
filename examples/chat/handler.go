package chat

import (
	"encoding/json"
	"learning_datastar/util"
	"log"
	"net/http"

	datastar "github.com/starfederation/datastar/sdk/go"
)

type ChatSignal struct {
	Error      string `json:"error"` // Add error field to chat signal to clear errors
	NewMessage string `json:"newMessage"`
}

type ErrorSignal struct {
	Error string `json:"error"`
}

func MessagesHandler(chatter *util.PubSubChannel, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Subscribe to the chatter channel
	chatterID := util.RandomId()
	ch := chatter.Subscribe(chatterID)
	sse := datastar.NewSSE(w, r)
	for {
		select {

		case <-r.Context().Done():
			chatter.Unsubscribe(chatterID)
			return
		case msg := <-ch:
			log.Println("Received message:", msg.Message, chatterID)

			sse.MergeFragmentTempl(
				ChatMessage(msg.Message),
				datastar.WithMergeAppend(),
				datastar.WithSelector("#messages"),
			)
		}
	}
}

var randomFailureCounter = 0

func SendErrorSignal(w http.ResponseWriter, r *http.Request, msg string) {
	errorSignal := &ErrorSignal{Error: msg}
	jsonStr, _ := json.Marshal(errorSignal)
	datastar.NewSSE(w, r).MergeSignals(jsonStr)
}

func SendHandler(chatter *util.PubSubChannel, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	randomFailureCounter++
	if randomFailureCounter%5 == 0 {
		SendErrorSignal(w, r, "Random failure occurred (we fail every 5th request)")
		return
	}

	signal := &ChatSignal{}
	err := datastar.ReadSignals(r, signal)
	if err != nil {
		SendErrorSignal(w, r, err.Error())
		return
	}

	chatter.PushMessage(signal.NewMessage)
	str, _ := json.Marshal(&ChatSignal{NewMessage: ""})
	datastar.NewSSE(w, r).MergeSignals(str)
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	Page().Render(r.Context(), w)
}
