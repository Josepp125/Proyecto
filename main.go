package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//Estructuras
func Saludar(nombre string) string {
	return "Hola " + nombre + " desde la función"
}

type Usuario struct {
	Nombre string
	Edad   int
}

// var templates = template.Must(template.New("T").ParseGlob("templates/*.html"))
var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))

//función para renderizar los templates desde cada Handler
/**si se renderiza un archivo inexistente da "La conexión ha sido reiniciada"
* por lo que hay que manejar el error con http.Error() **/
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

//HandlerError
func manejaError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status) //incluye el StatusError en el mensaje de eerror
	errorTemplate.Execute(rw, nil)
}

func renderTemplate(rw http.ResponseWriter, archivo string, data interface{}) {
	err := templates.ExecuteTemplate(rw, archivo, data)
	if err != nil {
		//http.Error(rw, "No es posible retornar template", http.StatusInternalServerError)
		manejaError(rw, http.StatusInternalServerError)
	}
}

//Handler
func Index(rw http.ResponseWriter, r *http.Request) {
	usuario := Usuario{"Jose", 0}
	//renderTemplate(rw, "inde.html", usuario) //produce el error
	renderTemplate(rw, "index.html", usuario)
}

func Acercade(rw http.ResponseWriter, r *http.Request) {
    usuario := Usuario{"Jose", 50}
	renderTemplate(rw, "acercade.html", usuario)
}
func Servicios(rw http.ResponseWriter, r *http.Request) {
    usuario := Usuario{"Jose", 50}
	renderTemplate(rw, "servicios.html", usuario)
}
func Iniciarsesion(rw http.ResponseWriter, r *http.Request) {
    usuario := Usuario{"Jose", 50}
	renderTemplate(rw, "login.html", usuario)
}
func Preguntas(rw http.ResponseWriter, r *http.Request) {
    usuario := Usuario{"Jose", 50}
	renderTemplate(rw, "preguntas.html", usuario)
}
func Validar(rw http.ResponseWriter, r *http.Request) {
	nombre := r.URL.Query().Get("login")
    password := r.URL.Query().Get("pass")
    usuario := Usuario{"Jose",24}
	if nombre == "jose" && password == "123456"{
		renderTemplate(rw, "perfil.html", usuario)
	} else {
		renderTemplate(rw, "error1.html", usuario)
	}
}

func main() {
	//Archivos estáticos
	archEstaticos := http.FileServer(http.Dir("estaticos"))

	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/acercade", Acercade)
	mux.HandleFunc("/preguntas", Preguntas)
	mux.HandleFunc("/servicios", Servicios)
	mux.HandleFunc("/login", Iniciarsesion)
	mux.HandleFunc("/validar", Validar)


	//Mux de archivos estáticos
	mux.Handle("/estaticos/", http.StripPrefix("/estaticos/", archEstaticos))

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	fmt.Println("Servidor corriendo en http://localhost:8080/")
	log.Fatal(server.ListenAndServe())
}
