package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Making headers variadic allows us to make it optional
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for _, header := range headers {
		for key, value := range header {
			w.Header()[key] = value
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	return err
}

// Data will be passed as a pointer
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024 // json over a megabyte is suspicious
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	err := dec.Decode(data)
	if err != nil {
		return err
	}

	// decoding information into an empty struct
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain single JSON value")
	}

	return nil
}

// using variadics for optionals here is is pretty awkward because ideally we'd like to
// receive either 0 or 1 arguments for status but not more. Since we can't add this into
// the typing of the function, we can at least mention it in the sent json error.
func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()
	if len(status) > 1 {
		payload.Message = fmt.Sprintf(`%v\nNote: recieved more than one status code. If this message is being sent there is a bug in the backend error handling code.`, payload.Message)
	}
	return app.writeJSON(w, statusCode, payload)
}
