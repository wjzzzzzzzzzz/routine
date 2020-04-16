package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"log"
)

//权限框架
func main() {
	authEnforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}
	authEnforcer.LoadPolicy()
	router := gin.Default()
	router.Use(Filter(authEnforcer))
	u := new(UserService)
	router.GET("/login", u.login)
	router.POST("/login", u.login)
	router.POST("/info", u.info)
	router.Run()
}

type UserService struct {
}

func (UserService) login(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": "",
	})
}
func (UserService) info(c *gin.Context) {

}

func Filter(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := "p"
		//判断策略中是否存在
		if e.Enforce(sub, obj, act) {
			fmt.Println("通过权限")
		} else {
			fmt.Println("权限没有通过")
			c.JSON(200, gin.H{
				"status":  "error",
				"message": "权限没有通过",
			})
			c.Abort()
		}
	}
}
