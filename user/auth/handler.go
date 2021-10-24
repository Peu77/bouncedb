package auth

import (
	"bouncedb/config"
	"bouncedb/utils"
	"fmt"
	"github.com/google/uuid"
	"time"
)

var Tokens []Token

func CreateToken(permissions []string) {
	var secretKey = utils.Encrypt(config.Token, uuid.NewString())
	fmt.Println("generate new token: " + secretKey)
	Tokens = append(Tokens, Token{secretKey, time.Now().UnixMilli(), permissions})

}
