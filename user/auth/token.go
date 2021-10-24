package auth

type Token struct {
	key         string
	created     int64
	permissions []string
}
