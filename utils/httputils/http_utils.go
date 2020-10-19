package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/emadghaffari/res_errors/errors"
)

// ResponseJSON func
func ResponseJSON(w http.ResponseWriter, status int, text interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(text)
}

// ResponseError func
func ResponseError(w http.ResponseWriter, err *errors.ResError) {
	ResponseJSON(w, err.Status, err)
}
