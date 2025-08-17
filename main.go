package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sajal-jayanto/api-gateway/api"
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
			
			path := fmt.Sprintf("%s%s", s.Path, e.Path)
			
			/// if you want to custom load balance add your logic here to select the s.Dome ans you want and put you owns algorithm  
			servicePath := fmt.Sprintf("%s:%s%s",s.Demon, s.Port, path)

			switch e.Method{
				case "GET":{
					api.GET(v1Mux, path, servicePath)
				}
				case "POST":{
					api.POST(v1Mux, path, servicePath)
				}
				case "PUT":{
					api.PUT(v1Mux, path, servicePath)
				}
				case "PATCH":{
					api.PATCH(v1Mux, path, servicePath)
				}
				case "HEAD":{
					api.HEAD(v1Mux, path, servicePath)
				}
				case "OPTIONS":{
					api.OPTIONS(v1Mux, path, servicePath)
				}
				case "DELETE":{
					api.DELETE(v1Mux, path, servicePath)
				}
			}
    }
	}

	fmt.Println("Server running on http://127.0.0.1:8080")
	if err := http.ListenAndServe(":8080", rootMux); err != nil {
		fmt.Println("Error:", err)
	}
}
