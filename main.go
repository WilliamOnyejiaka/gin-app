package main

import "github.com/gin-gonic/gin"

func base(ctx *gin.Context) {
	ctx.JSON(200,gin.H{
		"message":"Hello World",
	})
}

func main(){
	router:= gin.Default();
	router.GET("/",base)
	router.Run()
}
