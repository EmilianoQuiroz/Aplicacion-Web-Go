// Servidor de la aplicacion ç
package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Mi pagina principal")
	})
	http.HandleFunc("/Productos", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Pagina productos")
	})

	http.ListenAndServe(":8000", nil)
}