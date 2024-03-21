package httpOut

import (
	"net/http"
)

func TextResp(w http.ResponseWriter, resp string) {
	// Set the content type as text/plain
	w.Header().Set("Content-Type", "text/plain")
	// Write the response
	w.Write([]byte(resp))
}
