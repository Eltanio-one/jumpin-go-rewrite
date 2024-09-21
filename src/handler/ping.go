package handler

import (
	"fmt"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	// eventually we will make this ping request return config info
	// currently just used as a healthcheck
	_, err := w.Write([]byte("ping"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
