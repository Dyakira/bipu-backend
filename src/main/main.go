package main

import (
	"bipu-backend/src/main/controller"
	"bipu-backend/src/main/middleware"
	"flag"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"bipu-backend/src/main/model"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func init() {
	fmt.Print(GetCurrentDirectory())
	err := initGlobalConfig()
	//出现初始化问题，终止服务
	if err != nil {
		fmt.Print(err)
		syscall.Exit(1)
	}
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

/*从配置文件conf/config.json加载全局配置文件*/
func initGlobalConfig() error {
	if conf, err := middleware.GetConfig(); err == nil {
		middleware.Config = conf
		model.CommomInitDB()
		model.InitDB()
	} else {
		return err
	}
	return nil
}

func main() {
	//gin.SetMode(gin.ReleaseMode)

	flag.Parse()

	router := gin.Default()
	// init settings for github auth
	store := sessions.NewCookieStore([]byte(middleware.Config.Session.SessionSecret))
	router.Use(sessions.Sessions(middleware.Config.Session.SessionName, store))

	router.Use(gin.Logger())

	v1 := router.Group("/v1")
	{
		v1.GET("/login", controller.LoginHandler)
		v1.GET("auth/github/callback", controller.GithubCallback)
	}

	// user controller
	{
		v1.POST("/user/register", controller.InsertUser)
		v1.POST("/user/loginnew", controller.UserLogin)
	}
	// volume controller
	{
		v1.GET("/volume", controller.QueryVolumesAll)
		v1.POST("/volume", controller.InsertVolume)
		v1.GET("/volume/uid/:uid", controller.QueryVolumesByUid)
		v1.GET("/volume/vid/:vid", controller.QueryVolumesByVid)
	}

	authorized := v1.Group("/")
	authorized.Use(middleware.Auth())
	{
		authorized.GET("/logout", controller.UserLogoutHandler)
		authorized.GET("/userInfo", controller.UserInfo)
	}

	router.Run("127.0.0.1:8081")

	if pid := syscall.Getpid(); pid != 1 {
		ioutil.WriteFile("bipu_server.pid", []byte(strconv.Itoa(pid)), 0777)
		defer os.Remove("bipu_server.pid")
	}
}
