package repository

type Jwt struct {
	UserGUID     string `json:"user_guid"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
