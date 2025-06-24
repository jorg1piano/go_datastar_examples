package form_validation

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	datastar "github.com/starfederation/datastar/sdk/go"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	formValidationExample(FormSignal{}).Render(r.Context(), w)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	formSignal := &FormSignal{}
	err := datastar.ReadSignals(r, formSignal)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read signals: %v", err), http.StatusBadRequest)
		log.Println(err)
		return
	}
	// Update the dirty flags for each field
	if formSignal.FirstName != "" {
		formSignal.FirstNameDirty = true
	}
	if formSignal.LastName != "" {
		formSignal.LastNameDirty = true
	}
	if formSignal.Email != "" {
		formSignal.EmailDirty = true
	}

	// If a field is dirty, validate it and show an error message if it's invalid
	if formSignal.FirstNameDirty && formSignal.FirstName != "John" {
		formSignal.FirstNameInValidMessage = fmt.Sprintf("%s is not a valid first name. First name must be 'John'.", formSignal.FirstName)
	} else {
		formSignal.FirstNameInValidMessage = ""
	}

	if formSignal.LastNameDirty && formSignal.LastName != "Doe" {
		formSignal.LastNameInValidMessage = fmt.Sprintf("%s is not a valid last name. Last name must be 'Doe'.", formSignal.LastName)
	} else {
		formSignal.LastNameInValidMessage = ""
	}

	if formSignal.EmailDirty && !validEmail(formSignal.Email) {
		formSignal.EmailInValidMessage = "Please provide a valid email"
	} else {
		formSignal.EmailInValidMessage = ""
	}

	// Check that all the fields are valid
	formSignal.IsFormValid = isFormValid(formSignal)

	// Merge the updated signal back to the client
	sse := datastar.NewSSE(w, r)
	jsonStr, _ := json.Marshal(formSignal)
	sse.MergeSignals(jsonStr)
}

func isFormValid(s *FormSignal) bool {
	return s.FirstName == "John" &&
		s.LastName == "Doe" &&
		validEmail(s.Email) &&
		s.FirstNameDirty &&
		s.LastNameDirty &&
		s.EmailDirty
}

func validEmail(email string) bool {
	emailRegex := `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$` // probably not correct, butt good enough for this example
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
