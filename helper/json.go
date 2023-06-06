package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody (r *http.Request, result interface{}){
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}



func WriteResponseBody (write http.ResponseWriter, response interface{}){
	write.Header().Add("Content-Type", "application/json")
	enconder := json.NewEncoder(write)
	err := enconder.Encode(response)

	PanicIfError(err)
}
