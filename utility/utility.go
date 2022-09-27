package utility

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// RequestBodyToStruct destination using &struct
func RequestBodyToStruct(w http.ResponseWriter, body io.ReadCloser, destination interface{}) error {
	// Read body
	b, err := ioutil.ReadAll(body)
	defer body.Close()
	if err != nil {
		return err
	}

	// Unmarshal
	err = json.Unmarshal(b, destination)
	if err != nil {
		return err
	}
	return nil
}
