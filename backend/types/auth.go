package types

type JWTClaim struct {
	Username string `json:"username"`
	Id       uint   `json:"id"`
	Exp      int64  `json:"exp"`
}
