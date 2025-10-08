package handler

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type ErrorResponse struct {
    Message   string `json:"message"`
    ErrorCode int    `json:"errorCode"`
}

type SuccessResponse struct {
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

func sendError(w http.ResponseWriter, code int, msg string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)

    response := ErrorResponse{
        Message:   msg,
        ErrorCode: code,
    }

    json.NewEncoder(w).Encode(response)
}

func sendSuccess(w http.ResponseWriter, op string, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    response := SuccessResponse{
        Message: fmt.Sprintf("operation from handler: %s successfully", op),
        Data:    data,
    }

    json.NewEncoder(w).Encode(response)
}
