package models

import (
	"gopkg.in/guregu/null.v3"
)

//Autores is an exportable type
type Autores struct {
	ID          int      `json:"autor_id"`
	Nome        string   `json:"nome"`
	Criado      string `json:"criado"`
	CriadoPor   int      `json:"criado_por"`
	Alterado    null.String `json:"alterado"`
	AlteradoPor null.Int      `json:"alterado_por"`
}
