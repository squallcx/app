package app

import (
	"fmt"
	"net/http"
	"strings"
)

func ipHandler(w http.ResponseWriter, r *http.Request) {
	ip := strings.Split(r.RemoteAddr, ":")[0]
	fmt.Fprintf(w, "%s", ip)
}
