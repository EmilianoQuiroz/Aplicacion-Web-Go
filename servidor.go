// Servidor de la aplicacion ç
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Mi pagina principal")
	})
	http.HandleFunc("/Productos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pagina productos")
	})
	// Redireccionamiento
	http.HandleFunc("/Categorias", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", 301)
	})
	// Not found error 404
	http.HandleFunc("/Noencontro", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	// Error
	http.HandleFunc("/Error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Error en el servidor", 500)
	})
	http.ListenAndServe(":8000", nil)
}
