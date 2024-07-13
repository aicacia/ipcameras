package model

type CredentialsST struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
} // @name Credentials

type TokenST struct {
	AccessToken           string `json:"accessToken" validate:"required"`
	TokenType             string `json:"tokenType" validate:"required"`
	IssuedTokenType       string `json:"issuedTokenType" validate:"required"`
	ExpiresIn             int64  `json:"expiresIn" validate:"required"`
	RefreshToken          string `json:"refreshToken" validate:"required"`
	RefreshTokenExpiresIn int64  `json:"refreshTokenExpiresIn" validate:"required"`
} // @name Token
