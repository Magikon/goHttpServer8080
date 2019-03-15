package main

import (
	"net"
	"net/http"
	"os"
	"strconv"
)

func sayIP(w http.ResponseWriter, r *http.Request) {
	var allIPs string
	ifaces, err := net.Interfaces()
	if err != nil {
		os.Exit(1)
	}
	for _, i := range ifaces {

		addrs, err := i.Addrs()
		if err != nil {
			os.Exit(1)
		}
		for _, addr := range addrs {
			if IPNett, ok := addr.(*net.IPNet); ok && !IPNett.IP.IsLoopback() {
				allIPs += strconv.Itoa(i.Index) + ": " + addr.String() + "\r\n"
			}
		}
	}
	message := "Container ips \r\n" + allIPs
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayIP)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
