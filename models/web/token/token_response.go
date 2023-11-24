package token

type TokenResponse struct {
	Token        string `json:"token"`
	TokenRefresh string `json:"token_refresh"`
}
