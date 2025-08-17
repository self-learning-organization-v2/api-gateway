package api

import (
	"encoding/json"
	"fmt"
	"io"
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

		res, err := http.Get(servicePath)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(map[string]any{
				"status": http.StatusBadGateway,
				"error":  err.Error(),
			})
			return
		}
    defer res.Body.Close()

		for key, values := range res.Header {
			for _, v := range values {
				w.Header().Add(key, v)
			}
		}

		w.WriteHeader(res.StatusCode)
		if _, err := io.Copy(w, res.Body); err != nil {
			fmt.Println("error copying response:", err)
		}
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

		defer r.Body.Close()
		req, err := http.NewRequest(http.MethodPost, servicePath, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(map[string]any{
				"status": http.StatusBadGateway,
				"error":  err.Error(),
			})
			return
		}

		req.Header = r.Header.Clone()

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(map[string]any{
				"status": http.StatusBadGateway,
				"error":  err.Error(),
			})
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)
		for k, v := range resp.Header {
			for _, val := range v {
				w.Header().Add(k, val)
			}
		}

		if _, err := io.Copy(w, resp.Body); err != nil {
			fmt.Println("error copying response:", err)
		}

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

		defer r.Body.Close()
		req, err := http.NewRequest(http.MethodPost, servicePath, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(map[string]any{
				"status": http.StatusBadGateway,
				"error":  err.Error(),
			})
			return
		}

		req.Header = r.Header.Clone()

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(map[string]any{
				"status": http.StatusBadGateway,
				"error":  err.Error(),
			})
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)
		for k, v := range resp.Header {
			for _, val := range v {
				w.Header().Add(k, val)
			}
		}

		if _, err := io.Copy(w, resp.Body); err != nil {
			fmt.Println("error copying response:", err)
		}
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

		defer r.Body.Close()
		req, err := http.NewRequest(http.MethodPost, servicePath, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(map[string]any{
				"status": http.StatusBadGateway,
				"error":  err.Error(),
			})
			return
		}

		req.Header = r.Header.Clone()

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(map[string]any{
				"status": http.StatusBadGateway,
				"error":  err.Error(),
			})
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(resp.StatusCode)
		for k, v := range resp.Header {
			for _, val := range v {
				w.Header().Add(k, val)
			}
		}

		if _, err := io.Copy(w, resp.Body); err != nil {
			fmt.Println("error copying response:", err)
		}

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

		defer r.Body.Close()
		req, err := http.NewRequest(http.MethodDelete, servicePath, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(map[string]any{
				"status": http.StatusBadGateway,
				"error":  err.Error(),
			})
			return
		}

		req.Header = r.Header.Clone()

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(map[string]any{
				"status": http.StatusBadGateway,
				"error":  err.Error(),
			})
			return
		}
		defer resp.Body.Close()

		for k, v := range resp.Header {
			for _, val := range v {
				w.Header().Add(k, val)
			}
		}

		w.WriteHeader(resp.StatusCode)

		if _, err := io.Copy(w, resp.Body); err != nil {
			fmt.Println("error copying response:", err)
		}
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

		defer r.Body.Close()
		req, err := http.NewRequest(http.MethodOptions, servicePath, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(map[string]any{
				"status": http.StatusBadGateway,
				"error":  err.Error(),
			})
			return
		}

		req.Header = r.Header.Clone()

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(map[string]any{
				"status": http.StatusBadGateway,
				"error":  err.Error(),
			})
			return
		}
		defer resp.Body.Close()

		for k, v := range resp.Header {
			for _, val := range v {
				w.Header().Add(k, val)
			}
		}

		w.WriteHeader(resp.StatusCode)

		if _, err := io.Copy(w, resp.Body); err != nil {
			fmt.Println("error copying response:", err)
		}
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

		defer r.Body.Close()
		req, err := http.NewRequest(http.MethodHead, servicePath, nil)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(map[string]any{
				"status": http.StatusBadGateway,
				"error":  err.Error(),
			})
			return
		}

		req.Header = r.Header.Clone()

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(map[string]any{
				"status": http.StatusBadGateway,
				"error":  err.Error(),
			})
			return
		}
		defer resp.Body.Close()

		for k, v := range resp.Header {
			for _, val := range v {
				w.Header().Add(k, val)
			}
		}

		w.WriteHeader(resp.StatusCode)
	})
}