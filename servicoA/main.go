package main

import (
	"fmt"
	"io"
	"net/http"
	"servicoa/config"
)

func zipCodeHandler(w http.ResponseWriter, r *http.Request) {
	// Parâmetro
	zipCodeParam := r.URL.Query().Get("cep")
	if zipCodeParam == "" || len(zipCodeParam) != 8 {
		w.WriteHeader(422)
		w.Write([]byte(config.INVALID_ZIP_CODE))
		return
	}

	resp, err := http.Get("http://localhost:8081/?cep=" + zipCodeParam)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err.Error())
		w.WriteHeader(404)
		w.Write([]byte(config.ZIP_CODE_NOT_FOUND))
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(config.ZIP_CODE_NOT_FOUND))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))

}

func main() {
	http.HandleFunc("/", zipCodeHandler)
	http.ListenAndServe(":8080", nil)
}
