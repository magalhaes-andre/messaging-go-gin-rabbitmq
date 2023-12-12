package main

import "github.com/gin-gonic/gin"

func SimpleMessage(context *gin.Context) {
	message := context.Param("message")
	context.ShouldBindJSON(message)

}

func ComplexMessage() {

}
