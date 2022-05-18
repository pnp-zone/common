package models

import (
	"github.com/myOmikron/echotools/utilitymodels"
)

type User struct {
	utilitymodels.Common
	UserID             uint
	User               utilitymodels.User
	ForcePasswordReset bool
}
