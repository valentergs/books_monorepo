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

//ControllerUsuario será exportado
type ControllerUsuario struct{}

//TodosUsuarios será exportado ==========================================
func (c ControllerUsuario) TodosUsuarios(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error

		if r.Method != "GET" {
			erro.Message = "Método não permitido"
				utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
				return
		}

		rows, err := db.Query("SELECT * FROM usuarios ORDER BY usuario_id DESC;")

		if err != nil {
			if err == sql.ErrNoRows {
				erro.Message = "Não encontramos nenhum usuario"
				utils.RespondWithError(w, http.StatusBadRequest, erro)
				return
			} else {
				log.Fatal(err)
			}
		}

		defer rows.Close()

		clts := make([]models.Usuarios, 0)
		for rows.Next() {
			clt := models.Usuarios{}
			err := rows.Scan(&clt.ID, &clt.Criado, &clt.Alterado, &clt.Nome, &clt.Sobrenome, &clt.Email, &clt.Senha)
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

//UsuarioUnico será exportado ==========================================
func (c ControllerUsuario) UsuarioUnico(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error
		var usuario models.Usuarios

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

		row := db.QueryRow("SELECT * FROM usuarios WHERE usuario_id=$1;", id)

		err = row.Scan(&usuario.ID, &usuario.Criado, &usuario.Alterado, &usuario.Nome, &usuario.Sobrenome, &usuario.Email, &usuario.Senha)
		
		if err != nil {
			if err == sql.ErrNoRows {
				erro.Message = "Usuário inexistente"
				utils.RespondWithError(w, http.StatusBadRequest, erro)
				return
			} else {
				log.Fatal(err)
			}
		}

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, usuario)

	}

}

//UsuarioInserir será exportado ===========================================
func (c ControllerUsuario) UsuarioInserir(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var erro models.Error
		var usuarios models.Usuarios

		if r.Method != "POST" {
			erro.Message = "Método não permitido"
			utils.RespondWithError(w, http.StatusMethodNotAllowed, erro)
			return
		}

		json.NewDecoder(r.Body).Decode(&usuarios)

		expressaoSQL := `INSERT INTO usuarios (criado, nome, sobrenome, email, senha) VALUES ($1,$2,$3,$4,$5);`
		_, err := db.Exec(expressaoSQL, usuarios.Criado, usuarios.Nome, usuarios.Sobrenome, usuarios.Email, usuarios.Senha)
		if err != nil {
			panic(err)
		}

		SuccessMessage := fmt.Sprintf(`Usuario %s %s foi cadastrado com sucesso!`, usuarios.Nome, usuarios.Sobrenome) 

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, SuccessMessage)

	}
}

//UsuariosApagar será exportado =========================================
func (c ControllerUsuario) UsuariosApagar(db *sql.DB) http.HandlerFunc {

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

		db.QueryRow("DELETE FROM usuarios where usuario_id=$1;", id)

		SuccessMessage := "Usuário deletado com sucesso!"

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, SuccessMessage)

	}
}

//UsuarioEditar será exportado =========================================
func (c ControllerUsuario) UsuarioEditar(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var usuarios models.Usuarios
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

		json.NewDecoder(r.Body).Decode(&usuarios)

		expressaoSQL := `UPDATE usuarios SET alterado=$1, nome=$2, sobrenome=$3, email=$4, senha=$5 WHERE usuario_id=$6;`
		_, err = db.Exec(expressaoSQL, usuarios.Alterado, usuarios.Nome, usuarios.Sobrenome, usuarios.Email, usuarios.Senha, id)
		if err != nil {
			panic(err)
		}

		SuccessMessage := "Usuário alterado com sucesso!"

		w.Header().Set("Content-Type", "application/json")

		utils.ResponseJSON(w, SuccessMessage)

	}
}
