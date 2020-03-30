package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/golang/glog"
)

func OK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func InternalError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func BadRequest(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func Unauthorized(w http.ResponseWriter) {
	http.Error(w, "", http.StatusUnauthorized)
}

func RespondJSON(w http.ResponseWriter, payload interface{}) {
	js, err := json.Marshal(payload)
	if err != nil {
		glog.Error(err)
		InternalError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func RespondText(w http.ResponseWriter, text string) {
	w.Header().Set("Content-Type", "plain/text")
	w.Write([]byte(text))
}

func ReadJSONBody(body io.ReadCloser, data interface{}) error {
	decoder := json.NewDecoder(body)
	defer body.Close()
	if e := decoder.Decode(data); e != nil {
		return e
	}
	return nil
}
