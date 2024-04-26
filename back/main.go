package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.HTML(
		// Установка статуса HTTP на 200
		http.StatusOK,
		// Использование index.html шаблона
		"index.html",
		// Установка title на "Home Page"
		gin.H{
			"title": "Home Page",
		},
	)
}


//кеш вакансий, чтобы не грузить нейронку
var Cash map[string] []string

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./front/*")
	r.StaticFile("/script.js", "./front/script.js")
	r.StaticFile("/style.css", "./front/style.css")
	r.StaticFile("/favicon.ico", "./front/icon.ico")

	r.GET("/", home)

	r.Run(":8080")
}