package httpUtils

import (
	"log"
	"net/http"
)

func HandleError(w *http.ResponseWriter, code int, responseText string, logMessage string, err error) {
	errorMessage := ""
	writer := *w

	if err != nil {
		errorMessage = err.Error()
	}

	if logMessage != "" {
		log.Println(logMessage)
	}
	if errorMessage != "" {
		log.Println(errorMessage)
	}
	writer.WriteHeader(code)
	writer.Write([]byte(responseText))
}
