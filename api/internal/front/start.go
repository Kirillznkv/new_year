package front

import (
	"net/http"
)

func Start() error {
	srv := NewServer()

	return http.ListenAndServe(":80", srv)
}
