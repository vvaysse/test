package httpIn

import (
	"net/http"
	"strings"
)

func GetSender(r *http.Request) string {
	return r.RemoteAddr
}

func IsLocalhost(ip string) bool {
	if strings.Contains(ip, "[::1]") || strings.Contains(ip, "127.0.0.1") {
		return true
	}
	return false
}
