package models

import (
	"github.com/go-sql-driver/mysql"
	"gopkg.in/guregu/null.v3"
)

// NullTime is an alias for mysql.NullTime data type
type NullTime struct {
	mysql.NullTime
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
