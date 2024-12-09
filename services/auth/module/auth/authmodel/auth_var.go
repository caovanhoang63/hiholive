package authmodel

type AuthEmailPassword struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type AuthRegister struct {
	FirstName string `json:"firstName" form:"firstName" validate:"required,alpha"`
	LastName  string `json:"lastName" form:"lastName" validate:"required,alpha"`
	AuthEmailPassword
}

type Token struct {
	Token string `json:"token"`
	// ExpiredIn in seconds
	ExpiredIn int `json:"expiredIn"`
}

type TokenResponse struct {
	AccessToken Token `json:"accessToken"`
	// RefreshToken will be used when access token expired
	// to issue new pair access token and refresh token.
	RefreshToken *Token `json:"refreshToken,omitempty"`
}
