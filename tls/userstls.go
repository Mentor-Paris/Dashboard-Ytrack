package tls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// UserFinal struct which contains a all informmations of each user
type UserFinal struct {
	ID         int     `json:"id"`
	FirstName  string  `json:"firstName"`
	Xp         XpFinal `json:"xp"`
	Avatar_Url string  `json:"avatar_url"`
	Year       string  `json:"created"`
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

// Ytrack struct which contains a login a avatar_url and a email of each user
type Ytrack struct {
	Login      string `json:"login"`
	Avatar_Url string `json:"avatar_url"`
	Year       string `json:"created"`
}

var (
	Usersytrack []Ytrack
	users       []User

	listuserfinal = []UserFinal{}

	listuser []UserFinal
)

// display the whole JSON of the user
func GetAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, listuserfinal)
}

// display information equal to a single user
func GetUserByID(c *gin.Context) {
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

// display the leaderboard
func Leaderboard(c *gin.Context) {
	listuser = []UserFinal{}
	listusers := Leaderboardapi()
	// Call the HTML method of the Context to render a template
	c.HTML(http.StatusOK, "leaderboardtls.html", gin.H{"listusers": listusers})
}

// create the array of the leaderboard
func Leaderboardapi() []UserFinal {

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
	for v := range listuser {
		dateString := listuser[v].Year
		date, error := time.Parse("2006-01-02T15:04:05Z07:00", dateString)

		if error != nil {
			fmt.Println(error)
		}

		dateformat := date.Format("2006")
		listuser[v].Year = dateformat
	}
	return listuser[:20]
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

// retrieve a random number according to the length of the list of user
func random() UserFinal {
	lenghtusers := len(listuserfinal)
	randomuser := rand.Intn(lenghtusers)
	return listuserfinal[randomuser]
}

// read the JSON file
func ReadJsonUserXp() {
	// Open our jsonFile
	jsonFile, err := os.Open("assets/json/userstls.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened userstls.json")
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
	json.Unmarshal(byteValue, &Usersytrack)
}

// merge the json userxp the json user
func MergeJson() {
	for _, i := range users {
		listuserfinal = append(listuserfinal, UserFinal{ID: i.ID, FirstName: i.FirstName, Xp: XpFinal{FormatString(i.Xp.Amount), i.Xp.Amount}})
	}
	for i := range listuserfinal {
		for _, y := range Usersytrack {
			if listuserfinal[i].FirstName == y.Login {
				listuserfinal[i].Avatar_Url = y.Avatar_Url
				listuserfinal[i].Year = y.Year
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
