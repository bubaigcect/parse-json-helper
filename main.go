package parsejson

import (
	"encoding/json"
	"net/http"

	response "github.com/bubaigcect/parse-json-helper/common"
)

func Response(w http.ResponseWriter, message string, data interface{}, statusCodes ...int) {
	successStatusCodes := []int{200, 201, 202, 203, 204, 205, 206, 207, 208, 226}
	status := true
	code := 200
	if len(statusCodes) > 0 {
		code = statusCodes[0]
	}

	status = statusCodePresent(code, successStatusCodes)

	response := response.Response{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data,
	}

	if len(statusCodes) > 1 {
		responseCode := statusCodes[1]
		response.ResponseCode = responseCode
	}

	if data != nil {
		response.Data = data
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Function to check if an code is present in a slice of status codes
func statusCodePresent(statusCode int, statusCodes []int) bool {
	for _, code := range statusCodes {
		if code == statusCode {
			return true
		}
	}

	return false
}
