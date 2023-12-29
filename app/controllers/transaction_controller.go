package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
)

type TransactionController struct {
}

func (controller TransactionController) Deposit(context *gin.Context) {
	name := context.PostForm("name")
	log.Println("name", name)
}

func (controller TransactionController) Withdrawal(context *gin.Context) {
	name := context.PostForm("name")
	log.Println("name", name)
}
