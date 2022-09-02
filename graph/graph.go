package graph

import (
	"Restful-API/national"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ()

// display the leaderboard of ytrack
func Graphique(c *gin.Context) {
	var liststructuser []national.UserFinalNational
	var liststringuser []string
	for _, i := range national.Usersytracknational {
		liststructuser = append(liststructuser, national.UserFinalNational{Campus: i.Campus})
	}
	for j := range national.Usersytracknational {
		liststringuser = append(liststringuser, liststructuser[j].Campus)
	}
	countcampus := printUniqueValue(liststringuser)
	c.HTML(http.StatusOK, "graphique.html", gin.H{"countcampus": countcampus, "title": "Statistiques"})
}

func printUniqueValue(arr []string) map[string]int {
	//Create a   dictionary of values for each element
	dict := make(map[string]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	return dict
}
