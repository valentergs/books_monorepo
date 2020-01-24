
package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/valentergs/books_monorepo/models"
	"github.com/valentergs/books_monorepo/utils"
)

//ControllerAutor será exportado
type ControllerAutor struct{}

//TodosAutores será exportado ==========================================
func (c ControllerAutor) TodosAutores(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error

		if r.Method != "GET" {
			erro.Message = "Método não permitido"
				utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
				return
		}

		rows, err := db.Query("SELECT * FROM autores ORDER BY autor_id DESC;")

		if err != nil {
			if err == sql.ErrNoRows {
				erro.Message = "Não encontramos nenhum autor"
				utils.RespondWithError(w, http.StatusBadRequest, erro)
				return
			} else {
				log.Fatal(err)
			}
		}

		defer rows.Close()

		clts := make([]models.Autores, 0)
		for rows.Next() {
			clt := models.Autores{}
			err := rows.Scan(&clt.ID, &clt.Nome, &clt.Sobrenome, &clt.Criado, &clt.CriadoPor, &clt.Alterado, &clt.AlteradoPor)
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

//AutorUnico será exportado ==========================================
func (c ControllerAutor) AutorUnico(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error
		var autores models.Autores

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

		row := db.QueryRow("SELECT * FROM autores WHERE autor_id=$1;", id)

		err = row.Scan(&autores.ID, &autores.Nome, &autores.Sobrenome, &autores.Criado, &autores.CriadoPor, &autores.Alterado, &autores.AlteradoPor)
		
		if err != nil {
			if err == sql.ErrNoRows {
				erro.Message = "Não há autores cadastrados"
				utils.RespondWithError(w, http.StatusBadRequest, erro)
				return
			} else {
				log.Fatal(err)
			}
		}

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, autores)

	}

}

//AutorInserir será exportado ===========================================
func (c ControllerAutor) AutorInserir(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error
		var autores models.Autores

		if r.Method != "POST" {
			erro.Message = "Método não permitido"
			utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
			return
		}

		json.NewDecoder(r.Body).Decode(&autores)

		expressaoSQL := `INSERT INTO autores (nome, sobrenome, criado, criado_por) VALUES ($1,$2,$3,$4);`
		_, err := db.Exec(expressaoSQL, autores.Nome, autores.Sobrenome, autores.Criado, autores.CriadoPor)
		if err != nil {
			panic(err)
		}

		SuccessMessage := `Autor cadastrado com sucesso!`

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, SuccessMessage)

	}
}

//AutoresApagar será exportado =========================================
func (c ControllerAutor) AutoresApagar(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var error models.Error

		if r.Method != "DELETE" {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			error.Message = "Numero ID inválido"
		}

		db.QueryRow("DELETE FROM autores where autor_id=$1;", id)

		SuccessMessage := "Autor deletado com sucesso!"

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, SuccessMessage)

	}
}

//AutorEditar será exportado =========================================
func (c ControllerAutor) AutorEditar(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var autores models.Autores
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

		json.NewDecoder(r.Body).Decode(&autores)

		expressaoSQL := `UPDATE autores SET nome=$1, sobrenome=$2, alterado=$3, alterado_por=$4 WHERE autor_id=$5;`
		_, err = db.Exec(expressaoSQL, autores.Nome, autores.Sobrenome, autores.Alterado, autores.AlteradoPor, id)
		if err != nil {
			panic(err)
		}

		SuccessMessage := "Autor alterado com sucesso!"

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, SuccessMessage)

	}
}
