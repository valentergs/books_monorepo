package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"github.com/valentergs/books_monorepo/controllers"
	"github.com/valentergs/books_monorepo/driver"
)

var db *sql.DB

func main() {
	db := driver.ConnectDB()
	usuarioctl := controllers.ControllerUsuario{}
	livroctl := controllers.ControllerLivro{}
	autorctl := controllers.ControllerAutor{}
	tradutorctl := controllers.ControllerTradutor{}
	comentctl := controllers.Controllercomentarios{}

	// gorilla.mux
	router := mux.NewRouter()

	// LIVRO URL ====================================
	router.HandleFunc("/livros", livroctl.TodosLivros(db)).Methods("GET")
	router.HandleFunc("/livros/{id}", livroctl.LivroUnico(db)).Methods("GET")
	router.HandleFunc("/livros/inserir", livroctl.LivroInserir(db)).Methods("POST")
	router.HandleFunc("/livros/apagar/{id}", livroctl.LivroApagar(db)).Methods("DELETE")
	router.HandleFunc("/livros/editar/{id}", livroctl.LivroEditar(db)).Methods("PUT")

	// USUARIO URL ====================================
	router.HandleFunc("/usuarios", usuarioctl.TodosUsuarios(db)).Methods("GET")
	router.HandleFunc("/usuarios/{id}", usuarioctl.UsuarioUnico(db)).Methods("GET")
	router.HandleFunc("/usuarios/inserir", usuarioctl.UsuarioInserir(db)).Methods("POST")
	router.HandleFunc("/usuarios/apagar/{id}", usuarioctl.UsuariosApagar(db)).Methods("DELETE")
	router.HandleFunc("/usuarios/editar/{id}", usuarioctl.UsuarioEditar(db)).Methods("PUT")

	// AUTOR URL ====================================
	router.HandleFunc("/autores", autorctl.TodosAutores(db)).Methods("GET")
	router.HandleFunc("/autores/{id}", autorctl.AutorUnico(db)).Methods("GET")
	router.HandleFunc("/autores/inserir", autorctl.AutorInserir(db)).Methods("POST")
	router.HandleFunc("/autores/apagar/{id}", autorctl.AutoresApagar(db)).Methods("DELETE")
	router.HandleFunc("/autores/editar/{id}", autorctl.AutorEditar(db)).Methods("PUT")

	// TRADUTOR URL ====================================
	router.HandleFunc("/tradutores", tradutorctl.TodosTradutor(db)).Methods("GET")
	router.HandleFunc("/tradutores/{id}", tradutorctl.TradutorUnico(db)).Methods("GET")
	router.HandleFunc("/tradutores/inserir", tradutorctl.TradutorInserir(db)).Methods("POST")
	router.HandleFunc("/tradutores/apagar/{id}", tradutorctl.TradutoresApagar(db)).Methods("DELETE")
	router.HandleFunc("/tradutores/editar/{id}", tradutorctl.TradutorEditar(db)).Methods("PUT")

	// COMENTARIO URL ====================================
	router.HandleFunc("/comentarios", comentctl.Todoscomentarios(db)).Methods("GET")
	router.HandleFunc("/comentarios/inserir", comentctl.ComentariosInserir(db)).Methods("POST")
	router.HandleFunc("/comentarios/apagar/{id}", comentctl.ComentariosApagar(db)).Methods("DELETE")
	router.HandleFunc("/comentarios/editar/{id}", comentctl.ComentariosEditar(db)).Methods("PUT")

	// CORS ==========================================================
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		Debug:          true,
	})

	// Insert the middleware
	handler := c.Handler(router)

	log.Println("Listen on port 8080...")
	// Quando  usar o CORS colocar "handler" no lugar do "router"
	// Quando usar ROUTER usar "router"
	log.Fatal(http.ListenAndServe(":8080", handler))
}
