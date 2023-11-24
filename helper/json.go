package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestToBody(r *http.Request, result any) {
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(result)
	PanicError(err)
}

func WriteRequestToBody(w http.ResponseWriter, response any) {
	w.Header().Add("content-type", "application/json")
	encode := json.NewEncoder(w)
	err := encode.Encode(response)
	PanicError(err)
}
