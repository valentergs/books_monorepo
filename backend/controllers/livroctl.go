package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/valentergs/books_monorepo/models"
	"github.com/valentergs/books_monorepo/utils"
)

//ControllerLivro será exportado
type ControllerLivro struct{}

//TodosLivros será exportado ==========================================
func (c ControllerLivro) TodosLivros(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error

		if r.Method != "GET" {
			// http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			erro.Message = "Método não permitido"
			utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
			return
		}

		rows, err := db.Query("SELECT * FROM livros ORDER BY livro_id DESC;")
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			fmt.Println(err)
			return
		}

		defer rows.Close()

		clts := make([]models.Livros, 0)
		for rows.Next() {
			clt := models.Livros{}
			err := rows.Scan(&clt.ID, &clt.Isbn, &clt.Criado, &clt.CriadoPor, &clt.Alterado, &clt.AlteradoPor, &clt.Titulo, &clt.TituloOriginal, &clt.Autor, &clt.Tradutor, &clt.Cdd, &clt.Cdu, &clt.Ano, &clt.Tema, &clt.Editora, &clt.Paginas, &clt.Idioma, &clt.Formato, &clt.Photourl)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				fmt.Println(err)
				return
			}
			clts = append(clts, clt)
			
		}
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, clts)

	}

}

//LivroUnico será exportado ==========================================
func (c ControllerLivro) LivroUnico(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error
		var livro models.Livros

		if r.Method != "GET" {
			// http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			erro.Message = "Método não permitido"
			utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			erro.Message = "Numero ID inválido"
		}

		row := db.QueryRow("SELECT * FROM livros WHERE livro_id=$1;", id)

		err = row.Scan(&livro.ID, &livro.Isbn, &livro.Criado, &livro.CriadoPor, &livro.Alterado, &livro.AlteradoPor, &livro.Titulo, &livro.TituloOriginal, &livro.Autor, &livro.Tradutor, &livro.Cdd, &livro.Cdu, &livro.Ano, &livro.Tema, &livro.Editora, &livro.Paginas, &livro.Idioma, &livro.Formato, &livro.Photourl)
		
		if err != nil {
			if err == sql.ErrNoRows {
				erro.Message = "Livro inexistente"
				utils.RespondWithError(w, http.StatusBadRequest, erro)
				return
			} else {
				log.Fatal(err)
			}
		}

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, livro)

	}

}

//LivroInserir será exportado ===========================================
func (c ControllerLivro) LivroInserir(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error
		var livro models.Livros
		// var link string

		if r.Method != "POST" {
			// http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			erro.Message = "Método não permitido"
			utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
			return
		}

		json.NewDecoder(r.Body).Decode(&livro)
		
		//isbn := livro.Isbn
		// link := utils.PhotoLink(isbn)

		expressaoSQL := `INSERT INTO livros (isbn, criado_por, titulo, titulo_original, autor, tradutor, cdd, cdu, ano, tema, editora, paginas, idioma, formato, photourl) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15);`
		_, err := db.Exec(expressaoSQL, livro.Isbn, livro.CriadoPor, livro.Titulo, livro.TituloOriginal, livro.Autor, livro.Tradutor, livro.Cdd, livro.Cdu, livro.Ano, livro.Tema, livro.Editora, livro.Paginas, livro.Idioma, livro.Formato, livro.Photourl)

		// if err != nil {
		// 	panic(err)
		// }

		if pgerr, err := err.(*pq.Error); err {
			if pgerr.Code == "23505" {
				erro.Message = "Erro! Livro em duplicidade"
				utils.RespondWithError(w, http.StatusBadRequest, erro)
				return
			} else {
				erro.Message = "Desculpe. Aconteceu algo de errado com nosso Banco de Dados."
				utils.RespondWithError(w, http.StatusBadRequest, erro)
				return
				panic(err)
			}
		}

		

		// row := db.QueryRow("SELECT * FROM livros WHERE isbn=$1;", livro.Isbn)
		// err = row.Scan(&livro.ID, &livro.Isbn, &livro.Criado, &livro.CriadoPor, &livro.Alterado, &livro.AlteradoPor, &livro.Titulo, &livro.TituloOriginal, &livro.Autor, &livro.Tradutor, &livro.Cdd, &livro.Cdu, &livro.Ano, &livro.Tema, &livro.Editora, &livro.Paginas, &livro.Idioma, &livro.Formato, &livro.Photourl)
		// if err != nil {
		// 	panic(err)
		// }

		SuccessMessage := fmt.Sprintf(`Livro %s foi cadastrado com sucesso!`, livro.Titulo) 

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, SuccessMessage)

	}
}

//LivroApagar será exportado =========================================
func (c ControllerLivro) LivroApagar(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var error models.Error

		if r.Method != "DELETE" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		// Params são os valores informados pelo usuario no URL
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			error.Message = "Numero ID inválido"
		}

		db.QueryRow("DELETE FROM livros where livro_id=$1;", id)

		SuccessMessage := "Livro deletado com sucesso!"

		w.Header().Set("Content-Type", "application/json")
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		utils.ResponseJSON(w, SuccessMessage)

	}
}

//LivroEditar será exportado =========================================
func (c ControllerLivro) LivroEditar(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var livro models.Livros
		var error models.Error

		if r.Method != "PUT" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			error.Message = "Numero ID inválido"
		}

		json.NewDecoder(r.Body).Decode(&livro)

		expressaoSQL := `UPDATE livros SET isbn=$1, alterado=$2, alterado_por=$3, titulo=$4, titulo_original=$5, autor=$6, tradutor=$7, cdd=$8, cdu=$9, ano=$10, tema=$11, editora=$12, paginas=$13, idioma=$14, formato=$15, photourl=$16 WHERE livro_id=$17;`
		_, err = db.Exec(expressaoSQL, livro.Isbn, livro.Alterado, livro.AlteradoPor, livro.Titulo, livro.TituloOriginal, livro.Autor, livro.Tradutor, livro.Cdd, livro.Cdu, livro.Ano, livro.Tema, livro.Editora, livro.Paginas, livro.Idioma, livro.Formato, livro.Photourl, id)
		if err != nil {
			panic(err)
		}

		row := db.QueryRow("SELECT * FROM livros WHERE isbn=$1;", livro.Isbn)
		err = row.Scan(&livro.ID, &livro.Isbn, &livro.Criado, &livro.CriadoPor, &livro.Alterado, &livro.AlteradoPor, &livro.Titulo, &livro.TituloOriginal, &livro.Autor, &livro.Tradutor, &livro.Cdd, &livro.Cdu, &livro.Ano, &livro.Tema, &livro.Editora, &livro.Paginas, &livro.Idioma, &livro.Formato, &livro.Photourl)

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, livro)

	}
}
