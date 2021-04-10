package web

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, code int, dto interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if dto == nil {
		return
	}
	b, err := json.Marshal(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

func UnmarshalRequest(r *http.Request, dto interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		return err
	}
	return nil
}
