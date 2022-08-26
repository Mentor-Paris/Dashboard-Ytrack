package logs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Logs struct that contains the total logs of each user
type Logs struct {
	Id     int    `json:"id"`
	Nom    string `json:"nom"`
	Prenom string `json:"prenom"`
	Log    []Log  `json:"log"`
}

// Log struct that contains the total logs of each user
type Log struct {
	Id_L    int    `json:"id_l"`
	Date    string `json:"date"`
	Mentor  string `json:"mentor"`
	Comment string `json:"comment"`
	Clause  string `json:"clause"`
}

var (
	listlogs []Logs
	newLogs  Logs
	new_log  Log
)

// display the logs of students of ytrack
func Studentslog(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, listlogs)
}

func GetstudentsByID(c *gin.Context) {
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

func Createstudentslog(c *gin.Context) {

	Id, ok := c.GetQuery("Id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	Idconvert, err := strconv.Atoi(Id)

	if err != nil {
		fmt.Println(err)
	}

	Id_L, ok := c.GetQuery("Id_L")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id_l query parameter"})
		return
	}

	Id_Lconvert, err := strconv.Atoi(Id_L)

	if err != nil {
		fmt.Println(err)
	}

	Nom, ok := c.GetQuery("Nom")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing discord query parameter"})
		return
	}

	Prenom, ok := c.GetQuery("Prenom")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing discord query parameter"})
		return
	}

	Date, ok := c.GetQuery("Date")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing discord query parameter"})
		return
	}

	Mentor, ok := c.GetQuery("Mentor")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing discord query parameter"})
		return
	}

	Comment, ok := c.GetQuery("Comment")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing discord query parameter"})
		return
	}

	Clause, ok := c.GetQuery("Clause")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing discord query parameter"})
		return
	}

	logs, err := getLogsById(Idconvert)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Students not found."})
		return
	}

	newLogs = *new(Logs)

	for _, v := range listlogs {
		if v.Id == Idconvert {

			newLogs = v
			newLogs.Id = Idconvert
			newLogs.Nom = Nom
			newLogs.Prenom = Prenom

			new_log.Id_L = Id_Lconvert
			new_log.Date = Date
			new_log.Mentor = Mentor
			new_log.Comment = Comment
			new_log.Clause = Clause
			newLogs.Log = append(newLogs.Log, new_log)
		}
	}

	if logs.Id == newLogs.Id {
		logs.Log = newLogs.Log
	}

	c.IndentedJSON(http.StatusCreated, listlogs)
}

func Deletestudentslog(c *gin.Context) {

	Id, ok := c.GetQuery("Id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	Idconvert, err := strconv.Atoi(Id)

	if err != nil {
		fmt.Println(err)
	}

	Id_L, ok := c.GetQuery("Id_L")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id_l query parameter"})
		return
	}

	Id_Lconvert, err := strconv.Atoi(Id_L)

	if err != nil {
		fmt.Println(err)
	}

	for j, i := range listlogs {
		if i.Id == Idconvert {
			if j == len(listlogs) {
				continue
			} else {
				for k, v := range i.Log {
					if v.Id_L == Id_Lconvert {
						if k == len(listlogs) {
							continue
						} else {
							listlogs[j].Log = append(listlogs[j].Log[:k], listlogs[j].Log[k+1:]...)
							if Idconvert > 0 {
								Idconvert = Idconvert - 1
							}
							continue
						}
					}
				}

			}
		}
	}

	c.IndentedJSON(http.StatusOK, listlogs)
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
	json.Unmarshal(byteValue, &listlogs)
}

func getLogsById(id int) (*Logs, error) {
	for i, b := range listlogs {
		if b.Id == id {
			return &listlogs[i], nil
		}
	}
	return nil, errors.New("logs not found")
}
