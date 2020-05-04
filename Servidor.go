/*
	Ejemplo 1: hola mundo en servidor HTTP con Go
	@author parzibyte
*/
 
package main
 
import (
	"fmt"// Imprimir en consola
	"io"// Ayuda a escribir en la respuesta
	"log"//Loguear si algo sale mal
	"net/http"// El paquete HTTP
)
 
func main() {
 
	http.HandleFunc("/hola", func(w http.ResponseWriter, peticion *http.Request) {
		io.WriteString(w, "Solicitaste hola")
	})
	direccion := ":3030" // Como cadena, no como entero; porque representa una dirección
	fmt.Println("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))
}