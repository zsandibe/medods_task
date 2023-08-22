package repository

type UserToken struct {
	UserGUID     string `bson:"user_guid"`
	RefreshToken string `bson:"refresh_token"`
	BindTokens   string `bson:"bind_tokens"`
}
