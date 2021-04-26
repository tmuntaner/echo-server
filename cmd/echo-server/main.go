package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"sort"
	"strings"
)

type Server struct {
	logger *zap.SugaredLogger
}

func (s *Server) echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/echo" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	s.logger.Infow("Request received",
		"headers", r.Header,
		"url", r.URL,
		"method", r.Method,
		"protocol", r.Proto)

	var body string
	body += fmt.Sprintf("URL: %s\n", r.URL)
	body += fmt.Sprintf("Method: %s\n", r.Method)
	body += fmt.Sprintf("Protocol: %s\n", r.Proto)

	body += fmt.Sprint("\nHeaders:\n")
	keys := make([]string, len(r.Header))
	i := 0
	for key, _ := range r.Header {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := r.Header[key]
		body += fmt.Sprintf("%s: %s\n", key, strings.Join(value, ","))
	}

	_, err := fmt.Fprintf(w, body)
	if err != nil {
		fmt.Printf("echo handler experienced error: %s", err)
	}
}

func main() {
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync() // flushes buffer, if any
	}()

	sugar := logger.Sugar()
	s := Server{
		logger: sugar,
	}

	http.HandleFunc("/echo", s.echoHandler)

	sugar.Info("Starting webserver")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
