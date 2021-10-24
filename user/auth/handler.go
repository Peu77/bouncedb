package auth

import (
	"bouncedb/config"
	"bouncedb/utils"
	"fmt"
	"github.com/google/uuid"
	"time"
)

var Tokens []Token

func createToken(permissions []string) {
	var secretKey = utils.Decrypt(uuid.NewString(), config.CurrentConfig.SecretKey)
	fmt.Println("generate new token: " + secretKey)
	Tokens = append(Tokens, Token{secretKey, time.Now().UnixMilli(), permissions})
}
