package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GET(mux *http.ServeMux, path string, servicePath string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]any{ 
				"status": http.StatusMethodNotAllowed,
				"error": "Method Not Allowed",
			})
			return
		}

		fmt.Println("GET:", path, servicePath)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  http.StatusOK,
			"message": "GET request handled successfully",
		})
		// http.Get('/')
		
	})
}

func POST(mux *http.ServeMux, path string, servicePath string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]any{ 
				"status": http.StatusMethodNotAllowed,
				"error": "Method Not Allowed",
			})
			return
		}

		fmt.Println("POST:", path, servicePath)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  http.StatusOK,
			"message": "POST request handled successfully",
		})
		// http.Post()
	})
}

func PUT(mux *http.ServeMux, path string, servicePath string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]any{ 
				"status": http.StatusMethodNotAllowed,
				"error": "Method Not Allowed",
			})
			return
		}

		fmt.Println("Put:", path, servicePath)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  http.StatusOK,
			"message": "Put request handled successfully",
		})
		// http.Put()
	})
}

func PATCH(mux *http.ServeMux, path string, servicePath string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPatch {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]any{ 
				"status": http.StatusMethodNotAllowed,
				"error": "Method Not Allowed",
			})
			return
		}

		fmt.Println("Patch:", path, servicePath)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  http.StatusOK,
			"message": "Patch request handled successfully",
		})
		// http.Patch()
	})
}

func DELETE(mux *http.ServeMux, path string, servicePath string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]any{ 
				"status": http.StatusMethodNotAllowed,
				"error": "Method Not Allowed",
			})
			return
		}

		fmt.Println("Delete:", path, servicePath)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  http.StatusOK,
			"message": "Delete request handled successfully",
		})
		
		// http.Delete()
	})
}

func OPTIONS(mux *http.ServeMux, path string, servicePath string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodOptions {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]any{ 
				"status": http.StatusMethodNotAllowed,
				"error": "Method Not Allowed",
			})
			return
		}

		fmt.Println("Options:", path, servicePath)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  http.StatusOK,
			"message": "Options request handled successfully",
		})
		// http.Options()
	})
}

func HEAD(mux *http.ServeMux, path string, servicePath string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodHead {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]any{ 
				"status": http.StatusMethodNotAllowed,
				"error": "Method Not Allowed",
			})
			return
		}

		fmt.Println("Head:", path, servicePath)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  http.StatusOK,
			"message": "Head request handled successfully",
		})
		// http.Head()
	})
}