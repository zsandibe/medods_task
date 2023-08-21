package repository

type Jwt struct {
	UserGUID     string `json:"user_guid"`
	AccessToken  string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}
