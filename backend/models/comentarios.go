package models

import (
	"gopkg.in/guregu/null.v3"
)

//Comentarios is an exportable type
type Comentarios struct {
	ID          int      `json:"comentario_id"`
	Criado      null.String `json:"criado"`
	CriadoPor   int   `json:"criado_por"`
	Alterado    null.String `json:"alterado"`
	// AlteradoPor null.Int      `json:"alterado_por"`
	Livro       int      `json:"livro"`
	Text        string   `json:"texto"`
}