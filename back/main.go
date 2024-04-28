package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Form struct{
	Text string `json:"text"`
}

type Rewiew struct{
	Id int `json:"id_request"`
	Grade int `json:"grade"`

}

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

func send(c *gin.Context) {
	
	resp, err := http.Post("http://model:5000/send", "application/json", c.Request.Body)
	if err!=nil{
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.Writer.Write([]byte(err.Error()))
	}
	defer resp.Body.Close()
	b,err := io.ReadAll(resp.Body)
	if err!=nil{
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.Writer.Write([]byte(err.Error()))
        return
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(b)
}

func rewiew(c *gin.Context){
	fmt.Println("rewiew")
	c.Writer.Header().Set("Content-Type", "application/json")

	b,err := io.ReadAll(c.Request.Body)
	if err!=nil{
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.Writer.Write([]byte(err.Error()))
        return
	}
	type R struct{
		Grade int `json:"grade"`
		Text string `json:"text"`
	}

	var r R
	if err = json.Unmarshal(b,&r);err !=nil{
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error()))
        return
	}

	fmt.Println(r)

	connStr := "host=db user=tilt password=tilt_pass dbname=tilt_db sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error()))
    } 
    defer db.Close()
     
    _, err = db.Exec("insert into reviews (points, description) values($1,$2)", 
        r.Grade, r.Text)
    if err != nil{
        c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error()))
    }


	c.Writer.Write([]byte("ok"))
}


//кеш вакансий, чтобы не грузить нейронку
var Cash map[string] []string

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("./front/*.html")
	r.StaticFile("/script.js", "./front/script.js")
	r.StaticFile("/favicon.ico","./front/GBicon.ico")
	r.StaticFS("/css",http.Dir("./front/css"))

	r.GET("/", home)
	r.POST("/send",send)
	r.POST("/rewiew",rewiew)

	r.Run(":8080")
}