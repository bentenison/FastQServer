package models

type RefreshToken struct {
	ID  string `json:"-"`
	SS  string `json:"refreshtoken"`
	UID string `json:"-"`
}

type IdToken struct {
	SS string `json:"idtoken"`
}

type TokenPair struct {
	IdToken
	RefreshToken
}
