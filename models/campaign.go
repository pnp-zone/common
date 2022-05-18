package models

import "github.com/myOmikron/echotools/utilitymodels"

type Campaign struct {
	utilitymodels.Common
	Name        string  `json:"name" gorm:"not null"`
	InviteToken *string `json:"invite_token" gorm:"unique"`
	Player      []*User `json:"player" gorm:"many2many:campaign_player;"`
	GameMaster  []*User `json:"game_master" gorm:"many2many:campaign_gamemaster;"`
}
