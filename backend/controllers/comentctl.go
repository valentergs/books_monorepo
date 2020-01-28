package controllers

import (
	"database/sql"
	"encoding/json"
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	// "strconv"

	// "github.com/gorilla/mux"
	"github.com/valentergs/books_monorepo/models"
	"github.com/valentergs/books_monorepo/utils"
)

//Controllercomentarios será exportado
type Controllercomentarios struct{}

//Todoscomentarios será exportado ==========================================
func (c Controllercomentarios) Todoscomentarios(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error

		if r.Method != "GET" {
			erro.Message = "Método não permitido"
			utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
			return
		}

		rows, err := db.Query("SELECT * FROM comentarios ORDER BY comentario_id DESC;")

		if err != nil {
			if err == sql.ErrNoRows {
				erro.Message = "Não encontramos nenhum comentario"
				utils.RespondWithError(w, http.StatusBadRequest, erro)
				return
			} else {
				log.Fatal(err)
			}
		}

		defer rows.Close()

		clts := make([]models.Comentarios, 0)
		for rows.Next() {
			clt := models.Comentarios{}
			err := rows.Scan(&clt.ID, &clt.Criado, &clt.CriadoPor, &clt.Alterado, &clt.Livro, &clt.Text)
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

// //comentariosUnico será exportado ==========================================
// func (c Controllercomentarios) ComentariosUnico(db *sql.DB) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var erro models.Error
// 		var comentarios models.comentariosarios

// 		if r.Method != "GET" {
// 			// http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
// 			erro.Message = "Método não permitido"
// 			utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
// 			return
// 		}

// 		params := mux.Vars(r)
// 		id, err := strconv.Atoi(params["id"])
// 		if err != nil {
// 			erro.Message = "Numero ID inválido"
// 		}

// 		row := db.QueryRow("SELECT * FROM comentarios WHERE comentarios_id=$1;", id)

// 		err = row.Scan(&comentarios.ID, &comentarios.Criado, &comentarios.Alterado, &comentarios.Nome, &comentarios.Sobrenome, &comentarios.Email, &comentarios.Senha)

// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				erro.Message = "Comentário inexistente"
// 				utils.RespondWithError(w, http.StatusBadRequest, erro)
// 				return
// 			} else {
// 				log.Fatal(err)
// 			}
// 		}

// 		w.Header().Set("Content-Type", "application/json")

// 		utils.ResponseJSON(w, comentarios)

// 	}

// }

//ComentariosInserir será exportado ===========================================
func (c Controllercomentarios) ComentariosInserir(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error
		var comentarios models.Comentarios

		if r.Method != "POST" {
			erro.Message = "Método não permitido"
			utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
			return
		}

		json.NewDecoder(r.Body).Decode(&comentarios)

		expressaoSQL := `INSERT INTO comentarios (criado_por, livro, texto) VALUES ($1,$2,$3);`
		_, err := db.Exec(expressaoSQL, comentarios.CriadoPor, comentarios.Livro, comentarios.Text)
		if err != nil {
			panic(err)
		}

		SuccessMessage := `Comentário cadastrado com sucesso!`

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, SuccessMessage)

	}
}

//ComentariosApagar será exportado =========================================
func (c Controllercomentarios) ComentariosApagar(db *sql.DB) http.HandlerFunc {

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

		db.QueryRow("DELETE FROM comentarios where comentarios_id=$1;", id)

		SuccessMessage := "Usuário deletado com sucesso!"

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, SuccessMessage)

	}
}

// //comentariosEditar será exportado =========================================
// func (c Controllercomentarios) ComentariosEditar(db *sql.DB) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var comentarios models.comentarios
// 		var error models.Error

// 		if r.Method != "PUT" {
// 			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
// 			return
// 		}

// 		params := mux.Vars(r)
// 		id, err := strconv.Atoi(params["id"])
// 		if err != nil {
// 			error.Message = "Numero ID inválido"
// 		}

// 		json.NewDecoder(r.Body).Decode(&comentarios)

// 		expressaoSQL := `UPDATE comentarios SET alterado=$1, nome=$2, sobrenome=$3, email=$4, senha=$5 WHERE comentarios_id=$6;`
// 		_, err = db.Exec(expressaoSQL, comentarios.Alterado, comentarios.Nome, comentarios.Sobrenome, comentarios.Email, comentarios.Senha, id)
// 		if err != nil {
// 			panic(err)
// 		}

// 		SuccessMessage := "Usuário alterado com sucesso!"

// 		w.Header().Set("Content-Type", "application/json")

// 		utils.ResponseJSON(w, SuccessMessage)

// 	}
// }
