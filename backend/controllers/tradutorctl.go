
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

//ControllerTradutor será exportado
type ControllerTradutor struct{}

//TodosTradutor será exportado ==========================================
func (c ControllerTradutor) TodosTradutor(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error

		if r.Method != "GET" {
			erro.Message = "Método não permitido"
				utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
				return
		}

		rows, err := db.Query("SELECT * FROM tradutores ORDER BY tradutor_id DESC;")

		if err != nil {
			if err == sql.ErrNoRows {
				erro.Message = "Não encontramos nenhum tradutor"
				utils.RespondWithError(w, http.StatusBadRequest, erro)
				return
			} else {
				log.Fatal(err)
			}
		}

		defer rows.Close()

		clts := make([]models.Tradutores, 0)
		for rows.Next() {
			clt := models.Tradutores{}
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

//TradutorUnico será exportado ==========================================
func (c ControllerTradutor) TradutorUnico(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error
		var tradutores models.Tradutores

		if r.Method != "GET" {
			erro.Message = "Método não permitido"
			utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
			return
		}

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			erro.Message = "Numero ID inválido"
		}

		row := db.QueryRow("SELECT * FROM tradutores WHERE tradutor_id=$1;", id)

		err = row.Scan(&tradutores.ID, &tradutores.Nome, &tradutores.Sobrenome, &tradutores.Criado, &tradutores.CriadoPor, &tradutores.Alterado, &tradutores.AlteradoPor)
		
		if err != nil {
			if err == sql.ErrNoRows {
				erro.Message = "Não há tradutores cadastrados"
				utils.RespondWithError(w, http.StatusBadRequest, erro)
				return
			} else {
				log.Fatal(err)
			}
		}

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, tradutores)

	}

}

//TradutorInserir será exportado ===========================================
func (c ControllerTradutor) TradutorInserir(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error
		var tradutores models.Tradutores

		if r.Method != "POST" {
			erro.Message = "Método não permitido"
			utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
			return
		}

		json.NewDecoder(r.Body).Decode(&tradutores)

		expressaoSQL := `INSERT INTO tradutores (nome, sobrenome, criado, criado_por) VALUES ($1,$2,$3,$4);`
		_, err := db.Exec(expressaoSQL, tradutores.Nome, tradutores.Sobrenome, tradutores.Criado, tradutores.CriadoPor)
		if err != nil {
			panic(err)
		}

		SuccessMessage := fmt.Sprintf(`Tradutor %s %s foi cadastrado com sucesso!`, tradutores.Nome, tradutores.Sobrenome) 

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, SuccessMessage)

	}
}

//TradutoresApagar será exportado =========================================
func (c ControllerTradutor) TradutoresApagar(db *sql.DB) http.HandlerFunc {

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

		db.QueryRow("DELETE FROM tradutores where tradutor_id=$1;", id)

		SuccessMessage := "Tradutor deletado com sucesso!"

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, SuccessMessage)

	}
}

//TradutorEditar será exportado =========================================
func (c ControllerTradutor) TradutorEditar(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var tradutores models.Tradutores
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

		json.NewDecoder(r.Body).Decode(&tradutores)

		expressaoSQL := `UPDATE tradutores SET nome=$1, sobrenome=$2, alterado=$3, alterado_por=$4 WHERE tradutor_id=$5;`
		_, err = db.Exec(expressaoSQL, tradutores.Nome, tradutores.Sobrenome, tradutores.Alterado, tradutores.AlteradoPor, id)
		if err != nil {
			panic(err)
		}

		SuccessMessage := "Tradutor alterado com sucesso!"

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, SuccessMessage)

	}
}
