package xurl

import "strings"

// TCP unsures that s url contains TCP protocol identifier.
func TCP(s string) string {
	return "tcp://" + Address(s)
}

// HTTP unsures that s url contains HTTP protocol identifier.
func HTTP(s string) string {
	return "http://" + Address(s)
}

// WS unsures that s url contains WS protocol identifier.
func WS(s string) string {
	return "ws://" + Address(s)
}

// Address unsures that address contains localhost as host if non specified.
func Address(address string) string {
	if strings.HasPrefix(address, ":") {
		return "localhost" + address
	}
	return address
}
