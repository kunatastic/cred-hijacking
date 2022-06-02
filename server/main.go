package main

import (
	"fmt"
	"net/http"
	"net"
	"strings"
)

func Credentials(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: credentials", r.Method)
	switch r.Method {
		case "POST":
			if err := r.ParseForm(); err != nil {
				fmt.Println(w, "ParseForm() err: %v", err)
				return
			}
			credentials := r.FormValue("credentials")
			ip, _:= getIP(r)
			fmt.Println("Credentials =", credentials, "IPAddress =", ip)
		default:
			fmt.Println(w, "Sorry, only  POST method is supported.")
	}
}

func getIP(r *http.Request) (string, error) {
    //Get IP from the X-REAL-IP header
    ip := r.Header.Get("X-REAL-IP")
    netIP := net.ParseIP(ip)
    if netIP != nil {
        return ip, nil
    }

    //Get IP from X-FORWARDED-FOR header
    ips := r.Header.Get("X-FORWARDED-FOR")
    splitIps := strings.Split(ips, ",")
    for _, ip := range splitIps {
        netIP := net.ParseIP(ip)
        if netIP != nil {
            return ip, nil
        }
    }

    //Get IP from RemoteAddr
    ip, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        return "", err
    }
    netIP = net.ParseIP(ip)
    if netIP != nil {
        return ip, nil
    }
    return "", fmt.Errorf("No valid ip found")
}


func main() {
	http.HandleFunc("/credentials", Credentials)
	fmt.Printf("Server Started...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}