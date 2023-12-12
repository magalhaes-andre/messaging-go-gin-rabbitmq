package main

import "github.com/gin-gonic/gin"

func HandleRequests() {
	router := gin.Default()
	router.POST("/simple/:message", SimpleMessage)
	router.POST("/complex/", ComplexMessage)
}
