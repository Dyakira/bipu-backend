package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-gonic/contrib/sessions"
	"log"
	"golang.org/x/oauth2"
	"time"
	"github.com/google/go-github/github"
	"bipu-backend/src/main/model"
	"encoding/json"
	"encoding/base64"
	"crypto/rand"
	oauth2gh "golang.org/x/oauth2/github"
	"bipu-backend/src/main/middleware"
)



var (
	conf  *oauth2.Config
	state string
)

var redirectURL, credFile string

var scopes = []string{
	"repo",
	// You have to select your own scope from here -> https://developer.github.com/v3/oauth/#scopes
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func Setup(scopes []string) {

	conf = &oauth2.Config{
		ClientID:     middleware.Config.Credentials.ClientID,
		ClientSecret: middleware.Config.Credentials.ClientSecret,
		RedirectURL:  middleware.Config.Credentials.RedirectURL,
		Scopes:       scopes,
		Endpoint:     oauth2gh.Endpoint,
	}
}

func LoginHandler(ctx *gin.Context) {
	state = randToken()
	session := sessions.Default(ctx)
	session.Set("state", state)
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": GetLoginURL(state),
	})
}

func GetLoginURL(state string) string {
	Setup(scopes)
	return conf.AuthCodeURL(state)
}

func GithubCallback(ctx *gin.Context) {
	session := sessions.Default(ctx)
	//retrievedState := session.Get("state")
	//if retrievedState != ctx.Query("state") {
	//	ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
	//	return
	//}

	tok, err := conf.Exchange(oauth2.NoContext, ctx.Query("code"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	client := github.NewClient(conf.Client(oauth2.NoContext, tok))
	user, _, err := client.Users.Get(oauth2.NoContext, "")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tokenJson, err := json.Marshal(tok)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	modelUser := model.User{Openid: *user.ID, Name: *user.Name,
		AvatarUrl: *user.AvatarURL, OpenToken: string(tokenJson), Bio: *user.Bio, Source: "github", CreateAt: time.Now(), UpdateAt: time.Now()}

	model.InsertUserInfo(&modelUser)
	session.Set("uid", modelUser.Uid)
	session.Save()
}
