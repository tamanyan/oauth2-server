package form

// OAuth2TokenForm POST oauth2/token validation form
type OAuth2TokenForm struct {
	GrantType    string `form:"grant_type" validate:"required,oneof=password"`
	Scope        string `form:"scope" validate:"required"`
	Username     string `form:"username" validate:"required"`
	Password     string `form:"password" validate:"required"`
	ClientID     string `form:"client_id" validate:"required"`
	ClientSecret string `form:"client_secret" validate:"required"`
}
