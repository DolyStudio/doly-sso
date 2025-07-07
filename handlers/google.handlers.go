package handlers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zkfmapf123/pdf-bot/auths"
	"github.com/zkfmapf123/pdf-bot/utils"
)

var oauthStateString = utils.GenerateRandomString()

type GoogleAuthLoginResponseParams struct {
	AuthUrl string `json:"auth_url"`
	State   string `json:"state"`
}

// 사용자 입력 후 -> 구글 로그인 form 으로 이동
func GoogleAuthLogin(c *fiber.Ctx) error {
	gg := auths.InitGoogleOauth()

	url := gg.Ctx.AuthCodeURL(oauthStateString)
	return c.JSON(GoogleAuthLoginResponseParams{
		AuthUrl: url,
		State:   oauthStateString,
	})
}

type GoogleAuthCallbackResponseParams struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

type GoogleTokenResponseParams struct {
	// AccessToken  string `json:"access_token"`
	// RefreshToken string `json:"refresh_token"`
	Message   string    `json:"message"`
	TokenType string    `json:"token_type"`
	Expired   time.Time `json:"expired"`
}

// http://localhost:5555/auth/google/callback?code=4/0ARtbsJp...&state=abc123
// 1. Google 측에서 해당으로 통신
// 2. Code 자체를 다시 전달
func GoogleAuthCallback(c *fiber.Ctx) error {
	state := c.Query("state")

	if state != oauthStateString {
		return c.Status(400).JSON(GoogleAuthCallbackResponseParams{
			Code:  "",
			Error: "state not equals oauth string",
		})
	}

	// 코드 생성
	code := c.Query("code")
	fmt.Println("Google Code : ", code)

	// token 으로 교환
	gg := auths.InitGoogleOauth()
	token, err := gg.Ctx.Exchange(context.Background(), code)
	if err != nil {
		log.Println(err)

		return c.Status(500).JSON(GoogleAuthCallbackResponseParams{
			Code:  "",
			Error: "Token Exchange Failed",
		})
	}

	// Todo
	// AccessToken, RefreshToken 저장 (Database)
	// AccessToken = 실제 API 호출
	// RefreshToken = 액세스 토큰만료 시 새로 발급

	// 사용자 호출
	userInfo, err := gg.GetUserInfo(token.AccessToken)
	if err != nil {
		log.Panicln(err)

		return c.Status(500).JSON(GoogleAuthCallbackResponseParams{
			Code:  "",
			Error: "Invalid User",
		})
	}
	fmt.Println(userInfo)

	return c.JSON(GoogleTokenResponseParams{
		Message:   "Token Success",
		Expired:   token.Expiry,
		TokenType: token.TokenType,
	})
}
