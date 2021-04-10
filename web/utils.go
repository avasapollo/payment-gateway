package web

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, code int, dto interface{}) {
	if dto == nil {
		w.WriteHeader(code)
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
	w.WriteHeader(code)
	return
}

func UnmarshalRequest(r *http.Request, dto interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		return err
	}
	return nil
}
