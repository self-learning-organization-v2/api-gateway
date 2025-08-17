package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Servers []Server   `yaml:"servers"`
}

type Server struct {
	Name      string      `yaml:"name"`
	Path      string      `yaml:"path"`
	Demon     string      `yaml:"demon"`
	Port      string      `yaml:"port"`
	Endpoints []Endpoint  `yaml:"endpoints"`
}

type Endpoint struct {
	Path           string  `yaml:"path"`
	Method         string  `yaml:"method"`
	TimeOut        int     `yaml:"timeOut"`
	Authenticated  bool    `yaml:"authenticated"`
}

func main(){

	data, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	rootMux := http.NewServeMux()
	v1Mux := http.NewServeMux()

	rootMux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1Mux))
	// rootMux.Handle("/api/v2/", http.StripPrefix("/api/v2", v2Mux))
	
	for _, service := range config.Servers {
    for _, endpoint := range service.Endpoints {
			s, e := service, endpoint
			dynamic := fmt.Sprintf("%s%s", s.Path, e.Path)
			if dynamic[0] != '/' {
				dynamic = "/" + dynamic
			}
			v1Mux.HandleFunc(dynamic, func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Service: %s, Endpoint: %s %s", s.Name, e.Method, e.Path)
			})
    }
	}

	fmt.Println("Server running on http://127.0.0.1:8080")
	if err := http.ListenAndServe(":8080", rootMux); err != nil {
		fmt.Println("Error:", err)
	}
}