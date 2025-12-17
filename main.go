package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

func logf(format string, args ...any) {
	ts := time.Now().UTC().Format(time.RFC3339Nano)
	fmt.Printf("%s ", ts)
	fmt.Printf(format, args...)
	fmt.Print("\n")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host, _, _ := net.SplitHostPort(r.RemoteAddr)
		logf("%s %s from %s ua=%q", r.Method, r.URL.Path, host, r.UserAgent())

		hostname, _ := os.Hostname()
		ip := net.ParseIP(host)

		headers := make(map[string]string)
		for k, v := range r.Header {
			if len(v) > 0 {
				headers[k] = v[0]
			}
		}

		env := make(map[string]string)
		for _, e := range os.Environ() {
			for i := 0; i < len(e); i++ {
				if e[i] == '=' {
					env[e[:i]] = e[i+1:]
					break
				}
			}
		}

		response := map[string]interface{}{
			"host": map[string]interface{}{
				"hostname": hostname,
				"ip":       ip.String(),
				"ips":      []string{},
			},
			"http": map[string]interface{}{
				"method":      r.Method,
				"baseUrl":     "",
				"originalUrl": r.RequestURI,
				"protocol":    "http",
			},
			"request": map[string]interface{}{
				"params":  map[string]string{"0": r.URL.Path},
				"query":   r.URL.Query(),
				"cookies": map[string]string{},
				"body":    map[string]interface{}{},
				"headers": headers,
			},
			"environment": env,
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	})

	logf("Listening on :8888")
	_ = http.ListenAndServe(":8888", nil)
}
