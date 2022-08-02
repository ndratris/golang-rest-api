package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)


type todo struct {
	ID		string		`json:"id"`
	Name	string		`json:"name"`
	Status	bool		`json:"status"`
}

var data = [] todo {
	{ID:"1", Name:"Hendra", Status:false },
	{ID:"2", Name:"Dzaki", Status:true },
	{ID:"3", Name:"Hendra", Status:false },
	{ID:"4", Name:"Mita test", Status:false },
} 

func getData(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, data)
} 
func addData (context *gin.Context) {
var newData todo
if err := context.BindJSON(&newData);err != nil {
	return
}
data = append(data, newData)
context.IndentedJSON(http.StatusCreated, newData)

}
func getDataId(id string)(*todo , error){
	for i, t:= range data {
		if t.ID == id {
			return &data[i] , nil
		}
	}
	return nil, errors.New("data not found")
}

func getDataDetails(context *gin.Context) {
	id := context.Param("id")
	todo , err := getDataId(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message" :"data not found"})
		return 
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleDataStatus (context *gin.Context) {
	id := context.Param("id")
	todo , err := getDataId(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message" :"data not update, id not found"})
		return 
	}
	todo.Status = !todo.Status
	context.IndentedJSON(http.StatusOK, todo)
}
func main() {
	router := gin.Default()
	router.GET("/ndra", getData)
	router.GET("/ndra/:id", getDataDetails)
	router.PATCH("/ndra/:id", toggleDataStatus)
	router.POST("/ndra", addData)
	router.Run("localhost:9090")
}