package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserNational struct which contains a Firtsname a ID and a list of xp a Campus of each user national
type UserNational struct {
	FirstName string `json:"firstName"`
	ID        int    `json:"id"`
	Campus    string `json:"campus"`
	Xp        Xp     `json:"xp"`
}

// Ytrack struct which contains a login a avatar_url and a email of each user
type Ytrack struct {
	Login      string `json:"login"`
	Avatar_Url string `json:"avatar_url"`
	Email      string `json:"email"`
}

// UserFinal struct which contains a all informmations of each user pyc
type UserFinal struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	Email      string `json:"email"`
	Xp         Xp     `json:"xp"`
	Avatar_Url string `json:"avatar_url"`
}

// UserFinalnational struct which contains a all informmations of each user national
type UserFinalNational struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	Campus     string `json:"campus"`
	Email      string `json:"email"`
	Xp         Xp     `json:"xp"`
	Avatar_Url string `json:"avatar_url"`
}

// User struct which contains a Firtsname a ID and a list of xp of each user
type User struct {
	FirstName string `json:"firstName"`
	ID        int    `json:"id"`
	Xp        Xp     `json:"xp"`
}

// XP struct that contains the total XP of each user
type Xp struct {
	Amount int `json:"amount"`
}

// perform a task only once
func init() {
	ReadJsonUserXp()
	ReadJsonUsersYtrack()
	ReadJsonNational()
	MergeJsonPYC()
	MergeJsonNational()
}

// we initialize the variables of the map, array of single User and leaderboard
var users map[string]User
var usersytrack []Ytrack
var usersytracknational []UserNational

var listuserfinal = []UserFinal{}
var listuserfinalnational = []UserFinalNational{}

var listuser []UserFinal
var listusernational = []UserFinalNational{}

// principal function
func main() {

	r := gin.Default()

	r.GET("/", home)

	r.GET("/userPYC", getAllUsersPYC)

	r.GET("/usernational", getAllUsersNational)

	r.GET("/user/:id", getUserByID)

	r.GET("leaderboard", leaderboard)

	r.GET("leaderboardnational", leaderboardnational)

	r.SetFuncMap(template.FuncMap{"add": add})

	r.LoadHTMLGlob("templates/*")

	fmt.Println("\n" + "Voici le lien du serveur :" + " http://localhost:8080/")
	r.Run()
}

// route functions

// display the home page
func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home"})
}

// display the whole JSON of the user of PYC
func getAllUsersPYC(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, listuserfinal)
}

// display the whole JSON of the user of ytrack
func getAllUsersNational(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, listuserfinalnational)
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

// display the leaderboard of PYC
func leaderboard(c *gin.Context) {
	listuser = []UserFinal{}
	listusers := leaderboardapi()
	// Call the HTML method of the Context to render a template
	c.HTML(http.StatusOK, "leaderboardPyc.html", gin.H{"listusers": listusers})
}

// create the array of the leaderboard of PYC
func leaderboardapi() []UserFinal {

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
}

// display the leaderboard of ytrack
func leaderboardnational(c *gin.Context) {
	listusernational = []UserFinalNational{}
	listusersnational := leaderboardapinational()
	c.HTML(http.StatusOK, "leaderboardnational.html", gin.H{"listusersnational": listusersnational})
}

// create the array of the leaderboard of ytrack
func leaderboardapinational() []UserFinalNational {
	for comp := 0; comp < len(listuserfinalnational); comp++ {
		id := randomnational()
		for verifnational(id, listusernational) {
			id = randomnational()
		}
		listusernational = append(listusernational, id)
	}
	sort.SliceStable(listusernational, func(i, j int) bool {
		return listusernational[i].Xp.Amount > listusernational[j].Xp.Amount
	})
	return listusernational
}

// external functions

// add +1 to the index of the top leaderboard
func add(x, y int) int {
	return x + y
}

// check if the randomly chosen id is not already in the list
func verif(a UserFinal, b []UserFinal) bool {
	for _, comp := range b {
		if a.ID == comp.ID {
			return true
		}
	}
	return false
}

// check if the randomly chosen id is not already in the list
func verifnational(a UserFinalNational, b []UserFinalNational) bool {
	for _, comp := range b {
		if a.ID == comp.ID {
			return true
		}
	}
	return false
}

// retrieve a random number according to the length of the list of user
func random() UserFinal {
	lenghtusers := len(listuserfinal)
	randomuser := rand.Intn(lenghtusers)
	return listuserfinal[randomuser]
}

// retrieve a random number according to the length of the list of user
func randomnational() UserFinalNational {
	lenghtusers := len(listuserfinalnational)
	randomuser := rand.Intn(lenghtusers)
	return listuserfinalnational[randomuser]
}

// read the JSON file
func ReadJsonUserXp() {
	// Open our jsonFile
	jsonFile, err := os.Open("usersxp.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened usersxp.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)
}

// read the JSON file
func ReadJsonUsersYtrack() {
	// Open our jsonFile
	jsonFile, err := os.Open("usersnational.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened usersnational.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &usersytrack)
}

// read the JSON file
func ReadJsonNational() {
	// Open our jsonFile
	jsonFile, err := os.Open("usersnationalxp.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened usersnationalxp.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &usersytracknational)
}

// merge the json userxp the json userpyc
func MergeJsonPYC() {
	for _, i := range users {
		listuserfinal = append(listuserfinal, UserFinal{ID: i.ID, FirstName: i.FirstName, Xp: i.Xp})
	}
	for i := range listuserfinal {
		for _, y := range usersytrack {
			if listuserfinal[i].FirstName == y.Login {
				listuserfinal[i].Email = y.Email
				listuserfinal[i].Avatar_Url = y.Avatar_Url
				break
			}
		}
	}
}

// merge the json userxp the json usernational
func MergeJsonNational() {
	for _, i := range usersytracknational {
		listuserfinalnational = append(listuserfinalnational, UserFinalNational{ID: i.ID, FirstName: i.FirstName, Xp: i.Xp, Campus: i.Campus})
	}
	for i := range listuserfinalnational {
		for _, y := range usersytrack {
			if listuserfinalnational[i].FirstName == y.Login {
				listuserfinalnational[i].Email = y.Email
				listuserfinalnational[i].Avatar_Url = y.Avatar_Url
				break
			}
		}
	}
}
