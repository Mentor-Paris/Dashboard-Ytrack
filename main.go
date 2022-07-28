package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
    "os"
	"strconv"
	"encoding/json"
	"sort"
	"math/rand"
)

type Ytrack struct {
	Login string `json:"login"`
	Avatar_Url string `json:"avatar_url"`
	Email string `json:"email"`
}

type UserFinal struct {
	ID     int  `json:"id"`
	FirstName   string `json:"firstName"`
	Email string `json:"email"`
	Xp Xp `json:"xp"`
	Avatar_Url string `json:"avatar_url"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	FirstName   string `json:"firstName"`
	ID     int  `json:"id"`
    Xp Xp `json:"xp"`
}

// Social struct which contains a
// list of links
type Xp struct {
    Amount int `json:"amount"`
}

// perform a task only once
func init() {
	ReadJson()
	ReadJson2()
	MergeJson()
}

// we initialize our Users map[string] of User
// var users map[string]User
var users map[string]User
var usersytrack []Ytrack

var listuserfinal = []UserFinal{}

var listuser []UserFinal

// principal function
func main() {

	r := gin.Default()

	r.GET("/", getAllUsers)

	r.GET("/user/:id", getUserByID)

	r.GET("leaderboard", leaderboard)

	// r.GET("leaderboardnational", leaderboardnational)

	r.LoadHTMLGlob("templates/*")

	fmt.Println("\n"+"Voici le lien du serveur :"+" http://localhost:8080/"+"\n")
	r.Run()
}

// route functions 

// display the whole JSON
func getAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, listuserfinal)
}

// display information equal to a single user
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	idconvert, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err)
	}

	for _, a := range listuserfinal {
        if a.ID == idconvert {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

func leaderboard(c *gin.Context) {
	listuser = []UserFinal{}
	listusers := leaderboardapi()
	  // Call the HTML method of the Context to render a template
	c.HTML(http.StatusOK,"leaderboardPyc.html",gin.H{"listusers": listusers})
		// Set the HTTP status to 200 (OK)
		// Use the index.html template
		// Pass the data that the page uses (in this case, 'title')
	
}

// 
func leaderboardapi() []UserFinal{

	for comp := 0; comp < len(listuserfinal); comp++ {
		id := random()
		for verif(id, listuser) {
			id = random()
		}
		listuser = append(listuser, id)	
	}
	sort.SliceStable(listuser, func(i, j int) bool {
		return listuser[i].Xp.Amount > listuser[j].Xp.Amount
	})
	return listuser[:20]
	// c.IndentedJSON(http.StatusOK, listuser[:20])
}		

func verif(a UserFinal, b []UserFinal) bool {
	for _, comp := range b {
		if a.ID == comp.ID {
			return true
		}
	}
	return false
}

// retrieve a random number according to the length of the map
func random() UserFinal {
	lenghtusers := len(listuserfinal)
	randomuser := rand.Intn(lenghtusers)
	return listuserfinal[randomuser]
}


// read the JSON file
func ReadJson() {
	// Open our jsonFile
    jsonFile, err := os.Open("usersxp.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened usersxp.json"+"\n")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
    json.Unmarshal(byteValue, &users)
}

func ReadJson2() {
	// Open our jsonFile
    jsonFile, err := os.Open("users-ytrack.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened users-ytrack.json"+"\n")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
    json.Unmarshal(byteValue, &usersytrack)
}


func MergeJson() {
	for _ , i := range users {
		listuserfinal = append(listuserfinal, UserFinal{ID: i.ID , FirstName : i.FirstName , Xp : i.Xp})
	}
	for i := range listuserfinal {
		for _ , y := range usersytrack{
			if listuserfinal[i].FirstName == y.Login {
				listuserfinal[i].Email = y.Email
				listuserfinal[i].Avatar_Url = y.Avatar_Url
				break
			}
		}
	}
}
