package utl

import (
	"net"
	"net/http"
	"time"
)

// Returns true if given string is a valid IP address. False otherwise.
func ValidIpStr(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return ip != nil
}

// Checks if IP_Address:Port string is reachable
func IsIpPortStrReachable(ipPortStr string) bool {
	conn, err := net.DialTimeout("tcp", ipPortStr, time.Second*3)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// Checks if internet is available. This is of course a very poor check relying
// on DNS resolution working and this particular URL being available.
func InternetIsAvailable() bool {
	_, err := http.Get("http://httpbin.org/ip")
	return err == nil
}

// Inverse of above function.
func InternetNotAvailable() bool {
	return !InternetIsAvailable()
}
