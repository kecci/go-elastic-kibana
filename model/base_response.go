package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type (
	// BaseResponse is the base response
	BaseResponse struct {
		Code       string      `json:"code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
		Errors     []string    `json:"errors"`
		ServerTime int64       `json:"server_time"`
	}
)

// MapBaseResponse map response
func MapBaseResponse(w http.ResponseWriter, r *http.Request, message string, data interface{}, err error) {
	// Check Request ID
	requestID := r.Header.Get("X-Request-ID")
	if requestID != "" {
		bodyByte, _ := json.Marshal(data)
		fmt.Println("[RESPONSE: ", r.URL.String(), "] REQUEST_ID: ", requestID, " BODY:", string(bodyByte))
	}

	httpCode := http.StatusOK
	code := "SUCCESS"
	var errors []string
	if err != nil {
		httpCode = http.StatusInternalServerError
		code = "INTERNAL_SERVER_ERROR"
		errors = []string{err.Error()}
	}

	// Payload Response
	payload := BaseResponse{
		Code:       code,
		Message:    message,
		Data:       data,
		Errors:     errors,
		ServerTime: time.Now().Unix(),
	}

	// Marshal json response
	jsonResponse, _ := json.MarshalIndent(payload, "", "    ")

	// Write Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(jsonResponse)
}
