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
	reqIP := strings.Split(r.RemoteAddr, ":")[0]
	for _, ip := range config.BlockedIPs {
		if reqIP == ip {
			err = errors.New("ip not allowed to post images")
		}
	}

	for _, ip := range config.AllowedIPs {
		if reqIP != ip {
			err = errors.New("ip not allowed to post images")
		}
	}
	return err
}
