package request

// OAuth2PasswordCredentialRequest POST oauth2/token validation form
type OAuth2PasswordCredentialRequest struct {
	GrantType    string `form:"grant_type" validate:"required,oneof=password"`
	Scope        string `form:"scope" validate:"required"`
	Username     string `form:"username" validate:"required"`
	Password     string `form:"password" validate:"required"`
	ClientID     string `form:"client_id" validate:"required"`
	ClientSecret string `form:"client_secret" validate:"required"`
}

// OAuth2RefreshTokenRequest POST oauth2/token validation form
type OAuth2RefreshTokenRequest struct {
	GrantType    string `form:"grant_type" validate:"required,oneof=refresh_token"`
	Scope        string `form:"scope" validate:"required"`
	ClientID     string `form:"client_id" validate:"required"`
	ClientSecret string `form:"client_secret" validate:"required"`
	RefreshToken string `form:"refresh_token" validate:"required"`
}

// OAuth2ClientCredentialRequest POST oauth2/token validation form
type OAuth2ClientCredentialRequest struct {
	GrantType    string `form:"grant_type" validate:"required,oneof=client_credentials"`
	Scope        string `form:"scope" validate:"required"`
	ClientID     string `form:"client_id" validate:"required"`
	ClientSecret string `form:"client_secret" validate:"required"`
}
