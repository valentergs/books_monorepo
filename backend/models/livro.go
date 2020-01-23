package models

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"gopkg.in/guregu/null.v3"
)

// NullTime is an alias for mysql.NullTime data type
type NullTime struct {
	mysql.NullTime
}

//Usuarios is an exportable type
type Usuarios struct {
	ID        int      `json:"usuario_id"`
	Nome      string   `json:"nome"`
	Sobrenome string   `json:"sobrenome"`
	Criado    null.String `json:"criado"`
	Alterado  null.String `json:"alterado"`
	Email     string   `json:"email"`
	Senha     string   `json:"senha"`
}

//Autores is an exportable type
type Autores struct {
	ID          int      `json:"autor_id"`
	Nome        string   `json:"nome"`
	Sobrenome   string   `json:"sobrenome"`
	Criado      null.String `json:"criado"`
	CriadoPor   int      `json:"criado_por"`
	Alterado    null.String `json:"alterado"`
	AlteradoPor int      `json:"alterado_por"`
}

//Tradutores is an exportable type
type Tradutores struct {
	ID          int      `json:"tradutor_id"`
	Nome        string   `json:"nome"`
	Sobrenome   string   `json:"sobrenome"`
	Criado      null.String `json:"criado"`
	CriadoPor   int      `json:"criado_por"`
	Alterado    null.String `json:"alterado"`
	AlteradoPor int      `json:"alterado_por"`
}

//Livros is an exportable type
type Livros struct {
	ID             int      `json:"livro_id"`
	Isbn           string   `json:"isbn"`
	Criado         null.String `json:"criado"`
	CriadoPor      int   `json:"criado_por"`
	Alterado       null.String `json:"alterado"`
	AlteradoPor    null.Int      `json:"alterado_por"`
	Titulo         string   `json:"titulo"`
	TituloOriginal string   `json:"titulo_original"`
	Autor          int      `json:"autor"`
	Tradutor       int      `json:"tradutor"`
	Cdd            string   `json:"cdd"`
	Cdu            null.Int   `json:"cdu"`
	Ano            string   `json:"ano"`
	Tema           string   `json:"tema"`
	Editora        string   `json:"editora"`
	Paginas        string   `json:"paginas"`
	Idioma         string   `json:"idioma"`
	Formato        string   `json:"formato"`
	Photourl       null.String   `json:"photourl"`
}

//Biblioteca is an exportable type
type Biblioteca struct {
	Livro   int `json:"livro_id"`
	Usuario int `json:"usuario_id"`
}

//Cdd is an exportable type
type Cdd struct {
	Cdd  int `json:"cdd"`
	Tema int `json:"Tema"`
}

//CddClasse is an exportable type
type CddClasse struct {
	Cdd  int `json:"cdd"`
	Tema int `json:"Tema"`
}

//Comentarios is an exportable type
type Comentarios struct {
	ID          int      `json:"comentario_id"`
	Criado      null.String `json:"criado"`
	CriadoPor   string   `json:"criado_por"`
	Alterado    null.String `json:"alterado"`
	AlteradoPor int      `json:"alterado_por"`
	Livro       int      `json:"livro"`
	Text        string   `json:"texto"`
}

// MarshalJSON for NullTime
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}
