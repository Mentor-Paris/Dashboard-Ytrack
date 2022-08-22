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
	"strings"
	"time"

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
}

// UserFinal struct which contains a all informmations of each user pyc
type UserFinal struct {
	ID         int     `json:"id"`
	FirstName  string  `json:"firstName"`
	Xp         XpFinal `json:"xp"`
	Avatar_Url string  `json:"avatar_url"`
}

// UserFinalnational struct which contains a all informmations of each user national
type UserFinalNational struct {
	ID         int     `json:"id"`
	FirstName  string  `json:"firstName"`
	Campus     string  `json:"campus"`
	Xp         XpFinal `json:"xp"`
	Avatar_Url string  `json:"avatar_url"`
}

// Event struct that contains the total event of the plateform
type Event struct {
	Status       string       `json:"status"`
	EndEvent     string       `json:"endAt"`
	Path         string       `json:"path"`
	Registration Registration `json:"registration"`
}

// Event struct that contains the total registration date of the event
type Registration struct {
	StartRegistration string `json:"startAt"`
	EndRegistration   string `json:"endAt"`
	StartEvent        string `json:"eventStartAt"`
}

// User struct which contains a Firtsname a ID and a list of xp of each user
type User struct {
	FirstName string `json:"firstName"`
	ID        int    `json:"id"`
	Xp        Xp     `json:"xp"`
}

// XP struct that contains the total XP of each user
type Xp struct {
	Amount int64 `json:"amount"`
}

// XP struct that contains the total XP of each user
type XpFinal struct {
	Amount    string `json:"amount"`
	AmountInt int64
}

// EventDate struct that contains the total of date event of the plateform
type EvenDate struct {
	Debut_reg   string
	Fin_reg     string
	Debut_event string
	Fin_event   string
	Path        string
}

// Logs struct that contains the total logs of each user
type Logs struct {
	Id     int    `json:"id"`
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
	Log    []Log  `json:"log"`
}

// Log struct that contains the total logs of each user
type Log struct {
	Date    string `json:"date"`
	Mentor  string `json:"mentor"`
	Comment string `json:"comment"`
	Clause  string `json:"clause"`
}

type Complete_Stud struct {
	Id      int    `json:"id"`
	Nom     string `json:"nom"`
	Prenom  string `json:"prenom"`
	Point   int    `json:"point"`
	Credit  int    `json:"credit"`
	Guild   string `json:"guild"`
	Discord string `json:"ID_disc"`
}

// perform a task only once
func init() {
	ReadJsonUserXp()
	ReadJsonUsersYtrack()
	ReadJsonNational()
	ReadJsonEvent()
	ReadJsonLogs()
	ReadJsonStudents()
	MergeJsonPYC()
	MergeJsonNational()
	ListEvent()
}

// we initialize the variables of the map, array of single User and leaderboard
var (
	users               []User
	usersytrack         []Ytrack
	usersytracknational []UserNational

	listuserfinal         = []UserFinal{}
	listuserfinalnational = []UserFinalNational{}

	listuser         []UserFinal
	listusernational = []UserFinalNational{}

	listlogs []Logs

	liststudents []Complete_Stud

	events []Event

	listevents []string

	new_events []EvenDate
	new_event  *EvenDate
)

// principal function
func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/", home)

	r.GET("/userPYC", getAllUsersPYC)

	r.GET("/usernational", getAllUsersNational)

	r.GET("/user/:id", getUserByID)

	r.GET("leaderboard", leaderboard)

	r.GET("leaderboardnational", leaderboardnational)

	r.GET("graphique", graphique)

	r.GET("progress", progress)

	r.GET("students", studentslog)

	r.GET("/students/:id", getstudentsByID)

	r.GET("/go-bot", getGobot)

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
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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
		return listuser[i].Xp.AmountInt > listuser[j].Xp.AmountInt
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
		return listusernational[i].Xp.AmountInt > listusernational[j].Xp.AmountInt
	})
	return listusernational
}

// display the leaderboard of ytrack
func graphique(c *gin.Context) {
	var liststructuser []UserFinalNational
	var liststringuser []string
	for _, i := range usersytracknational {
		liststructuser = append(liststructuser, UserFinalNational{Campus: i.Campus})
	}
	for j := range usersytracknational {
		liststringuser = append(liststringuser, liststructuser[j].Campus)
	}
	countcampus := printUniqueValue(liststringuser)
	c.HTML(http.StatusOK, "graphique.html", gin.H{"countcampus": countcampus})
}

// display the date of event of ytrack
func progress(c *gin.Context) {
	c.HTML(http.StatusOK, "progress.html", gin.H{"new_events": new_events})
}

// display the logs of students of ytrack
func studentslog(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, listlogs)
}

func getstudentsByID(c *gin.Context) {
	id := c.Param("id")
	idconvert, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err)
	}

	for _, a := range listlogs {
		if a.Id == idconvert {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

func getGobot(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, liststudents)
}

// external functions

func printUniqueValue(arr []string) map[string]int {
	//Create a   dictionary of values for each element
	dict := make(map[string]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	return dict
}

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
	jsonFile, err := os.Open("assets/json/usersxp.json")
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
	jsonFile, err := os.Open("assets/json/usersnational.json")
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
	jsonFile, err := os.Open("assets/json/usersnationalxp.json")
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

// read the JSON file
func ReadJsonEvent() {
	// Open our jsonFile
	jsonFile, err := os.Open("assets/json/events.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened events.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &events)
}

// read the JSON file
func ReadJsonLogs() {
	// Open our jsonFile
	jsonFile, err := os.Open("assets/json/logsGeneral.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened logsGeneral.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &liststudents)
}

func ReadJsonStudents() {
	// Open our jsonFile
	jsonFile, err := os.Open("assets/json/api.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened api.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &listlogs)
}

// merge the json userxp the json userpyc
func MergeJsonPYC() {
	for _, i := range users {
		listuserfinal = append(listuserfinal, UserFinal{ID: i.ID, FirstName: i.FirstName, Xp: XpFinal{FormatString(i.Xp.Amount), i.Xp.Amount}})
	}
	for i := range listuserfinal {
		for _, y := range usersytrack {
			if listuserfinal[i].FirstName == y.Login {
				listuserfinal[i].Avatar_Url = y.Avatar_Url
				break
			}
		}
	}
}

// merge the json userxp the json usernational
func MergeJsonNational() {
	for _, i := range usersytracknational {
		listuserfinalnational = append(listuserfinalnational, UserFinalNational{ID: i.ID, FirstName: i.FirstName, Xp: XpFinal{FormatString(i.Xp.Amount), i.Xp.Amount}, Campus: i.Campus})
	}
	for i := range listuserfinalnational {
		for _, y := range usersytrack {
			if listuserfinalnational[i].FirstName == y.Login {
				listuserfinalnational[i].Avatar_Url = y.Avatar_Url
				break
			}
		}
	}
}

func FormatString(n int64) string {
	in := strconv.FormatInt(n, 10)
	out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)
	if in[0] == '-' {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ' '
		}
	}
}

// create list, split the path of event, format the date of event
func ListEvent() {
	new_event = new(EvenDate)
	last_str := ""

	for _, i := range events {
		listevents = append(listevents, i.Registration.StartRegistration, i.Registration.EndRegistration, i.Registration.StartEvent, i.EndEvent, i.Path)
	}

	for i := range listevents {

		if i%5 != 4 {
			dateString := listevents[i]
			date, error := time.Parse("2006-01-02T15:04:05Z07:00", dateString)

			dateformat := date.Format("2006-01-02 15:04:05")

			if error != nil {
				fmt.Println(error)
			}

			switch i % 5 {
			case 0:
				new_event.Debut_reg = dateformat
			case 1:
				new_event.Fin_reg = dateformat
			case 2:
				new_event.Debut_event = dateformat
			case 3:
				new_event.Fin_event = dateformat
			}

		} else {
			s := strings.Split(events[i/5].Path, "/")
			if 3 >= len(s) {
				last_str = s[2]
			} else {
				last_str = s[2] + "/" + s[3]
			}
			new_event.Path = last_str
			new_events = append(new_events, *new_event)
			sort.SliceStable(new_events, func(i, j int) bool {
				return new_events[i].Debut_event > new_events[j].Debut_event
			})
			new_event = new(EvenDate)
		}
	}
}
