package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	*http.Response
}

// ReadJSON binds json to Go type and optionally runs validation
func (r *Response) ReadJSON(obj interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		return errors.New("unable to read input stream")
	}

	err = json.Unmarshal(body, obj)
	if err != nil {
		log.Error(err)
		return errors.New("invalid json data")
	}

	return nil
}
