package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Config struct {
	Servers []Server   `json:"servers"`
}

type Server struct {
	Name      string      `json:"name"`
	Path      string      `json:"path"`
	Demon     string      `json:"demon"`
	Port      string      `json:"port"`
	Endpoints []Endpoint  `json:"endpoints"`
}

type Endpoint struct {
	Path           string  `json:"path"`
	Method         string  `json:"method"`
	TimeOut        int     `json:"timeOut"`
	Authenticated  bool    `json:"authenticated"`
}

func main(){

	data, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
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
			switch e.Method{
				case "GET":{
					GET(v1Mux, dynamic)
				}
				case "POST":{
					POST(v1Mux, dynamic)
				}
				case "PUT":{
					PUT(v1Mux, dynamic)
				}
				case "PATCH":{
					PATCH(v1Mux, dynamic)
				}
				case "HEAD":{
					HEAD(v1Mux, dynamic)
				}
				case "OPTIONS":{
					OPTIONS(v1Mux, dynamic)
				}
				case "DELETE":{
					DELETE(v1Mux, dynamic)
				}
			}
    }
	}

	fmt.Println("Server running on http://127.0.0.1:8080")
	if err := http.ListenAndServe(":8080", rootMux); err != nil {
		fmt.Println("Error:", err)
	}
}


func GET(mux *http.ServeMux, path string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		// http.Get('/')
		
	})
}

func POST(mux *http.ServeMux, path string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		// http.Post()
	})
}

func PUT(mux *http.ServeMux, path string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		// http.Put()
	})
}

func PATCH(mux *http.ServeMux, path string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		// http.Patch()
	})
}

func DELETE(mux *http.ServeMux, path string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		// http.Delete()
	})
}

func OPTIONS(mux *http.ServeMux, path string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodOptions {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		// http.Options()
	})
}

func HEAD(mux *http.ServeMux, path string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodHead {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		// http.Head()
	})
}