package models

import (
	"gopkg.in/guregu/null.v3"
)

//Usuarios is an exportable type
type Usuarios struct {
	ID        int      `json:"usuario_id"`
	Criado    null.String `json:"criado"`
	Alterado  null.String `json:"alterado"`
	Nome      string   `json:"nome"`
	Sobrenome string   `json:"sobrenome"`
	Email     string   `json:"email"`
	Senha     string   `json:"senha"`
}
