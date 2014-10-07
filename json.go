package em

import (
	"encoding/json"
	"net/http"
)

type RestError struct {
	Errors map[string]interface{} `json:"errors"`
}

func Required(field, value string) *RestError {
	if value == "" {
		return ValidationError(field, "is required.")
	}
	return nil
}
func ValidationResponse(err *RestError) (int, http.Header, interface{}, error) {
	return 422, nil, err, nil
}
func ValidationError(field, msg string) *RestError {
	rErr := RestError{}
	rErr.Errors = make(map[string]interface{})
	rErr.Errors[field] = msg
	return &rErr
}

//ValidationError - error response returns Unprocessable Entity
//For more information on 422 HTTP Error code see 11.2 WebDAV RFC 4918
//https://tools.ietf.org/html/rfc4918#section-11.2
func Validation(w http.ResponseWriter, entity interface{}) error {
	//422 Unprocessable Entity
	//A 422 HTTP response from the server generally implies that the request
	//was well formed but the API was unable to process it because the
	//content was not semantically correct or meaningful per the API.

	//For more information on 422 HTTP Error code see 11.2 WebDAV RFC 4918
	//https://tools.ietf.org/html/rfc4918#section-11.2
	w.WriteHeader(422)
	return Post(w, "errors", entity)
}

//Post response to post request
func Post(w http.ResponseWriter, key string, entity interface{}) error {
	js, err := json.Marshal(map[string]interface{}{key: entity})
	if err != nil {
		return err
	}
	w.Write(js)
	return nil
}
