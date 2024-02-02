package helper

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type response struct {
	Message string `json:"message"`
}

func FromJSON(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(data)
}

func ToJSON(data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func GetIDFromURL(r *http.Request) (int, error) {
	id := chi.URLParam(r, "id")
	return strconv.Atoi(id)
}

func WriteJSONData(w http.ResponseWriter, code int, data []byte) {
	w.WriteHeader(code)
	w.Write(data)
}

func WriteJSONMessage(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response{Message: message})
}

