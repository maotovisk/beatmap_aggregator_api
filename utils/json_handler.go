package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JsonHandler struct {
	w http.ResponseWriter
	r *http.Request
}

func NewJsonHandler(w http.ResponseWriter, r *http.Request) *JsonHandler {
	return &JsonHandler{
		w,
		r,
	}
}

func (jw *JsonHandler) ParseBody(obj any) error {
	dec := json.NewDecoder(jw.r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&obj)
	if err != nil {
		switch err {
		case io.EOF, io.ErrUnexpectedEOF:
			return fmt.Errorf("Invalid or missing JSON string")
		default:
			return err
		}
	}
	return nil
}

func (jw *JsonHandler) WriteResponse(obj any) {
	jw.WriteResponseWithStatus(obj, http.StatusOK)
}

func (jw *JsonHandler) WriteResponseWithStatus(obj any, status int) {
	jw.w.Header().Set("Content-Type", "application/json")
	jw.w.WriteHeader(status)

	if obj != nil {
		err := json.NewEncoder(jw.w).Encode(obj)
		if err != nil {
			http.Error(jw.w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (jw *JsonHandler) WriteMessageWithStatus(message string, status int) {
	type Message struct {
		Message string `json:"message"`
	}

	jw.WriteResponseWithStatus(&Message{Message: message}, status)
}
