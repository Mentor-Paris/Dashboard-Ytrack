package main

import (
	"Restful-API/api"
	"Restful-API/events"
	"Restful-API/graph"
	"Restful-API/logs"
	"Restful-API/national"
	"Restful-API/pyc"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// perform a task only once
func init() {
	pyc.ReadJsonUserXp()
	pyc.ReadJsonUsersYtrack()
	national.ReadJsonNational()
	events.ReadJsonEvent()
	logs.ReadJsonLogs()
	api.ReadJsonStudents()
	pyc.MergeJsonPYC()
	national.MergeJsonNational()
	events.ListEvent()
}

// principal function
func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/", home)

	r.GET("/userPYC", pyc.GetAllUsersPYC)

	r.GET("/usernational", national.GetAllUsersNational)

	r.GET("/user/:id", pyc.GetUserByID)

	r.GET("leaderboard", pyc.Leaderboard)

	r.GET("leaderboardnational", national.Leaderboardnational)

	r.GET("graphique", graph.Graphique)

	r.GET("progress", events.Progress)

	r.GET("students", logs.Studentslog)

	r.POST("students", logs.Createstudentslog)

	r.DELETE("students", logs.Deletestudentslog)

	r.GET("/students/:id", logs.GetstudentsByID)

	r.GET("/go-bot", api.GetGobot)

	r.POST("/go-bot", api.CreateGobot)

	r.PATCH("/go-bot", api.PatchGobot)

	r.DELETE("/go-bot", api.DeleteGobot)

	r.SetFuncMap(template.FuncMap{"add": add})

	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "assets/css")
	r.Static("/img", "assets/img")
	r.Static("/js", "assets/js")
	r.Static("/files", "assets/files")
	r.Static("/json", "assets/json")

	fmt.Println("\n" + "Voici le lien du serveur :" + " http://localhost:8080/")
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// display the home page
func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home"})
}

// add +1 to the index of the top leaderboard
func add(x, y int) int {
	return x + y
}
