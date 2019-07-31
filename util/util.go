package util

import (
	"encoding/json"
	"math/rand"
	"mime"
	"net/http"
	"strings"
	"time"
)

func RespondWithObject(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func RespondWithErrorObject(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(data)
}

func GenerateRandomInRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(max-min) + min
	return result
}

// Determine whether the request `content-type` includes a
// server-acceptable mime-type
//
// Failure should yield an Eror
func HasContentType(r *http.Request, mimetype string) bool {
	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		return mimetype == "application/octet-stream"
	}

	for _, v := range strings.Split(contentType, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimetype {
			return true
		}
	}
	return false
}
