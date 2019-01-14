package main

import (
	"net/http"
	//"strings"
	"net"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
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
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
			allIPs += ip.String() + ", "
		}
	}
	message := "<h1>Hello from " + allIPs + "</h1>"
	w.Write([]byte(message))
}
func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
