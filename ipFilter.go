package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func ipFilter(r *http.Request) error {
	var err error
	fmt.Println(r.RemoteAddr)
	ip := strings.Split(r.RemoteAddr, ":")[0]
	if ip != "127.0.0.1" {
		err = errors.New("ip not allowed to post images")
	}

	return err
}
