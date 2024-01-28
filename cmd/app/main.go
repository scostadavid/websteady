package main

import (
	"fmt"
	"net/http"
	"time"
)

type Service struct {
	Name string
	Url  string
}

type MonitorResponse struct {
	Up bool `json:"up"`
}

func Monitor(url string) (*MonitorResponse, error) {

	// format string
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// browser request mimic
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	resp, err := client.Do(req)

	if err != nil {
		return &MonitorResponse{Up: false}, err
	}
	resp.Body.Close()

	return &MonitorResponse{Up: resp.StatusCode == 200}, nil
}

func main() {
	services := []Service{
		{Name: "Google", Url: "https://www.google.com"},
		{Name: "GitHub", Url: "https://www.github.com"},
		{Name: "Netflix", Url: "https://www.netflix.com"},
		{Name: "Httpbin", Url: "https://httpbin.org/status/400"},
	}

	for _, service := range services {
		response, err := Monitor(service.Url)
		fmt.Printf("Service: %s, Status: %v, err: %v\n", service.Name, response.Up, err)
	}

	fmt.Println("Monitoring complete.")
}
