package repository

type UserToken struct {
	UserGUID     string `json:"user_guid"`
	RefreshToken string `json:"access_token"`
	BindTokens   string `json:"bind_tokens"`
}
