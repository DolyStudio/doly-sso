package auths

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/zkfmapf123/pdf-bot/business"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	CLIENT_ID     = os.Getenv("GOOGLE_CLIENT_ID")
	CLIENT_SECRET = os.Getenv("GOOGLE_CLIENT_SECRET")
	REDIRECT_URL  = os.Getenv("GOOGLE_REDIRECT_URL")
)

type GoogleAuth struct {
	Ctx *oauth2.Config
}

// google auth 구성
func InitGoogleOauth() GoogleAuth {
	googleAuth := &oauth2.Config{
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectURL:  REDIRECT_URL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	return GoogleAuth{
		Ctx: googleAuth,
	}
}

// func (g GoogleAuth) ExpiredToken(aToken, rToken string) {
// }

func (g GoogleAuth) GetUserInfo(accessToken string) (business.UserToken, error) {
	var userToken business.UserToken
	url := "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken
	resp, err := http.Get(url)
	if err != nil {
		return userToken, err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&userToken)
	return userToken, nil
}

// AccessToken Refreseh
func (g GoogleAuth) RefreshAccessToken(refreshToken string) (*oauth2.Token, error) {
	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	newToken, err := g.Ctx.TokenSource(context.Background(), token).Token()
	if err != nil {
		return nil, err
	}

	return newToken, nil
}
