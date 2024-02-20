package handler

import (
	"net/http"
)

type indexData struct {
	SourceIP string
}

func Index(t TemplateBuilder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := indexData{
			SourceIP: r.RemoteAddr,
		}

		if data.SourceIP == `` {
			data.SourceIP = `World`
		}

		bodyHTML, err := t.Build(`index.html`, data)
		writeResponse(w, bodyHTML, err)
	}
}
