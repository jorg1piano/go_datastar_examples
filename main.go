package main

import (
	"fmt"
	"learning_datastar/examples/binding_lists"
	"learning_datastar/examples/chat"
	"learning_datastar/examples/clock_raf"
	"learning_datastar/examples/count_until"
	"learning_datastar/examples/counter"
	"learning_datastar/examples/data_computed"
	"learning_datastar/examples/data_on_and_data_attr"
	"learning_datastar/examples/deriving_signals_for_server"
	"learning_datastar/examples/dialog"
	"learning_datastar/examples/execute_script"
	"learning_datastar/examples/form_validation"
	"learning_datastar/examples/index"
	"learning_datastar/examples/key_events"
	"learning_datastar/examples/signals_and_the_dom"
	"learning_datastar/util"
	"net/http"
)

var pubsub = util.NewPubSub()

func main() {
	// Handle the routes
	http.HandleFunc("/", index.PageHandler)

	http.HandleFunc("/counter", counter.PageHandler)
	http.HandleFunc("/counter/increment", counter.CounterIncrementPutHandler)

	http.HandleFunc("/count-until", count_until.PageHandler)
	http.HandleFunc("/count-until/increment-slowly", count_until.PutHandler)

	http.HandleFunc("/dialog", dialog.PageHandler)
	http.HandleFunc("/dialog/quote", dialog.GetHandler)

	http.HandleFunc("/chat", chat.PageHandler)
	http.HandleFunc("/chat/messages", func(w http.ResponseWriter, r *http.Request) { chat.MessagesHandler(pubsub, w, r) })
	http.HandleFunc("/chat/send", func(w http.ResponseWriter, r *http.Request) { chat.SendHandler(pubsub, w, r) })

	http.HandleFunc("/clock-raf", clock_raf.PageHandler)

	http.HandleFunc("/data-on-and-data-attr", data_on_and_data_attr.PageHandler)

	http.HandleFunc("/execute-script", execute_script.PageHandler)
	http.HandleFunc("/execute-script/start", execute_script.PostHandler)

	http.HandleFunc("/form-validation", form_validation.PageHandler)
	http.HandleFunc("/form-validation/validate", form_validation.PostHandler)

	http.HandleFunc("/key-events", key_events.PageHandler)

	http.HandleFunc("/data-computed", data_computed.PageHandler)

	http.HandleFunc("/binding-lists", binding_lists.PageHandler)

	http.HandleFunc("/computed-signals-to-server", deriving_signals_for_server.PageHandler)
	http.HandleFunc("/computed-signals-to-server/process", deriving_signals_for_server.ProcessCounter)

	http.HandleFunc("/signals-if-missing", signals_and_the_dom.PageHandler)
	http.HandleFunc("/signals-if-missing/replace-with-data-signals-if-missing", signals_and_the_dom.MorphHandlerTargetUsingIfMissing)
	http.HandleFunc("/signals-if-missing/replace-with", signals_and_the_dom.MorphHandlerReplaceWith)

	// Start the server
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
