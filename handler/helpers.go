package handler

import (
	"encoding/json"
	"net/http"

	"github.com/armedi/learn-go/lib/api"
)

func parseBody(r *http.Request, dst interface{}) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		return api.ErrBadRequest(err.Error(), "")
	}
	return nil
}
