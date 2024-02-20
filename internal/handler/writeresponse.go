package handler

import (
	"net/http"
)

func writeResponse(w http.ResponseWriter, handlerBody []byte, handlerErr error) {
	if handlerErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(handlerBody)
}
