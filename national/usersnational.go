package national

import (
	"Restful-API/pyc"
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

// UserNational struct which contains a Firtsname a ID and a list of xp a Campus of each user national
type UserNational struct {
	FirstName string `json:"firstName"`
	ID        int    `json:"id"`
	Campus    string `json:"campus"`
	Xp        Xp     `json:"xp"`
}

// XP struct that contains the total XP of each user
type Xp struct {
	Amount int64 `json:"amount"`
}

// UserFinalnational struct which contains a all informmations of each user national
type UserFinalNational struct {
	ID         int     `json:"id"`
	FirstName  string  `json:"firstName"`
	Campus     string  `json:"campus"`
	Xp         XpFinal `json:"xp"`
	Avatar_Url string  `json:"avatar_url"`
	Year       string  `json:"created"`
}

// XP struct that contains the total XP of each user
type XpFinal struct {
	Amount    string `json:"amount"`
	AmountInt int64
}

var (
	listuserfinalnational = []UserFinalNational{}
	listusernational      = []UserFinalNational{}
	Usersytracknational   []UserNational
)

// display the whole JSON of the user of ytrack
func GetAllUsersNational(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, listuserfinalnational)
}

// display the leaderboard of ytrack
func Leaderboardnational(c *gin.Context) {
	listusernational = []UserFinalNational{}
	listusersnational := Leaderboardapinational()
	c.HTML(http.StatusOK, "leaderboardnational.html", gin.H{"listusersnational": listusersnational, "title": "Leaderboard National"})
}

// create the array of the leaderboard of ytrack
func Leaderboardapinational() []UserFinalNational {
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
	for v := range listusernational {
		dateString := listusernational[v].Year

		if dateString == "" {
			listusernational[v].Year = "Pas de Promo"
			listusernational[v].Avatar_Url = "https://via.placeholder.com/360x360"
		} else {
			date, error := time.Parse("2006-01-02T15:04:05Z07:00", dateString)

			if error != nil {
				fmt.Println(error)
			}

			dateformat := date.Format("2006")
			listusernational[v].Year = dateformat
		}
	}

	return listusernational
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
func randomnational() UserFinalNational {
	lenghtusers := len(listuserfinalnational)
	randomuser := rand.Intn(lenghtusers)
	return listuserfinalnational[randomuser]
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
	json.Unmarshal(byteValue, &Usersytracknational)
}

// merge the json userxp the json usernational
func MergeJsonNational() {
	for _, i := range Usersytracknational {
		listuserfinalnational = append(listuserfinalnational, UserFinalNational{ID: i.ID, FirstName: i.FirstName, Xp: XpFinal{FormatString(i.Xp.Amount), i.Xp.Amount}, Campus: i.Campus})
	}
	for i := range listuserfinalnational {
		for _, y := range pyc.Usersytrack {
			if listuserfinalnational[i].FirstName == y.Login {
				listuserfinalnational[i].Avatar_Url = y.Avatar_Url
				listuserfinalnational[i].Year = y.Year
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
