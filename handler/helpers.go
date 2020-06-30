package handler

import (
	"encoding/json"
	"net/http"
)

func parseBody(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}
