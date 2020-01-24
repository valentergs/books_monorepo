package models

import (
	"gopkg.in/guregu/null.v3"
)

//Tradutores is an exportable type
type Tradutores struct {
	ID          int      `json:"tradutor_id"`
	Nome        string   `json:"nome"`
	Sobrenome   string   `json:"sobrenome"`
	Criado      null.String `json:"criado"`
	CriadoPor   int      `json:"criado_por"`
	Alterado    null.String `json:"alterado"`
	AlteradoPor null.Int      `json:"alterado_por"`
}	