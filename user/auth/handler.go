package auth

import (
	"bouncedb/config"
	"bouncedb/utils"
	"github.com/google/uuid"
	"time"
)

var Tokens []Token

func CreateToken(permissions []string) string {
	var secretKey = utils.Encrypt(config.Token, uuid.NewString())
	Tokens = append(Tokens, Token{secretKey, time.Now().UnixMilli(), permissions})
	return secretKey
}
