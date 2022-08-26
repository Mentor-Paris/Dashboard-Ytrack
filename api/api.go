package api

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

type Complete_Stud struct {
	Id      int    `json:"id"`
	Nom     string `json:"nom"`
	Prenom  string `json:"prenom"`
	Point   int    `json:"point"`
	Credit  int    `json:"credit"`
	Guild   string `json:"guild"`
	Discord string `json:"ID_disc"`
}

var (
	liststudents []Complete_Stud
	newStud      Complete_Stud
)

func GetGobot(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, liststudents)
}

func CreateGobot(c *gin.Context) {

	Id, ok := c.GetQuery("Id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	Idconvert, err := strconv.Atoi(Id)

	if err != nil {
		fmt.Println(err)
	}

	Nom, ok := c.GetQuery("Nom")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing nom query parameter"})
		return
	}

	Prenom, ok := c.GetQuery("Prenom")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing prenom query parameter"})
		return
	}

	Point, ok := c.GetQuery("Point")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing point query parameter"})
		return
	}

	Pointconvert, err := strconv.Atoi(Point)

	if err != nil {
		fmt.Println(err)
	}

	Credit, ok := c.GetQuery("Credit")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing credit query parameter"})
		return
	}

	Creditconvert, err := strconv.Atoi(Credit)

	if err != nil {
		fmt.Println(err)
	}

	Guild, ok := c.GetQuery("Guild")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing guild query parameter"})
		return
	}

	Discord, ok := c.GetQuery("Discord")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing discord query parameter"})
		return
	}

	newStud.Id = Idconvert
	newStud.Nom = Nom
	newStud.Prenom = Prenom
	newStud.Point = Pointconvert
	newStud.Credit = Creditconvert
	newStud.Guild = Guild
	newStud.Discord = Discord

	liststudents = append(liststudents, newStud)

	c.IndentedJSON(http.StatusCreated, liststudents)
}

func PatchGobot(c *gin.Context) {

	Id, ok := c.GetQuery("Id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	Idconvert, err := strconv.Atoi(Id)

	if err != nil {
		fmt.Println(err)
	}

	NewId, ok := c.GetQuery("NewId")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing NewId query parameter"})
		return
	}

	NewIdconvert, err := strconv.Atoi(NewId)

	if err != nil {
		fmt.Println(err)
	}

	Nom, ok := c.GetQuery("Nom")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing nom query parameter"})
		return
	}

	Prenom, ok := c.GetQuery("Prenom")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing prenom query parameter"})
		return
	}

	Point, ok := c.GetQuery("Point")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing point query parameter"})
		return
	}

	Pointconvert, err := strconv.Atoi(Point)

	if err != nil {
		fmt.Println(err)
	}

	Credit, ok := c.GetQuery("Credit")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing credit query parameter"})
		return
	}

	Creditconvert, err := strconv.Atoi(Credit)

	if err != nil {
		fmt.Println(err)
	}

	Guild, ok := c.GetQuery("Guild")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing guild query parameter"})
		return
	}

	Discord, ok := c.GetQuery("Discord")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing discord query parameter"})
		return
	}

	students, err := getStudById(Idconvert)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Students not found."}) //return custom request for bad request or book not found
		return
	}

	students.Id = NewIdconvert
	students.Nom = Nom
	students.Prenom = Prenom
	students.Point = Pointconvert
	students.Credit = Creditconvert
	students.Guild = Guild
	students.Discord = Discord

	c.IndentedJSON(http.StatusOK, students)
}

func DeleteGobot(c *gin.Context) {

	Id, ok := c.GetQuery("Id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
		return
	}

	Idconvert, err := strconv.Atoi(Id)

	if err != nil {
		fmt.Println(err)
	}

	for j, i := range liststudents {
		if i.Id == Idconvert {

			if j == len(liststudents) {
				continue
			} else {
				liststudents = append(liststudents[:j], liststudents[j+1:]...)
				if Idconvert > 0 {
					Idconvert = Idconvert - 1
				}
				continue
			}
		}
	}

	c.IndentedJSON(http.StatusOK, liststudents)
}

func getStudById(id int) (*Complete_Stud, error) {
	for i, b := range liststudents {
		if b.Id == id {
			return &liststudents[i], nil
		}
	}
	return nil, errors.New("students not found")
}

// read the JSON file
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
	json.Unmarshal(byteValue, &liststudents)
}
