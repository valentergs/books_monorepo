package models

import (
	"gopkg.in/guregu/null.v3"
)

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


