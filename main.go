package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID          string `json:"id"`
	Item        string  `json:"title"`
	Completed   bool     `json:"completed"`
}

var todos = []todo {
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "1", Item: "Watch Footbal", Completed: false},
}

/////////////////////////////////////////////
// Get todos ---GET Method
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}
////////////////////////////////////////////////

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
// Add todos ---POST Method
func addTodo(context *gin.Context) {
	var newTodo todo

	 err := context.BindJSON(&newTodo)
	 if err != nil {
		return
	 }

	 todos = append(todos, newTodo)
}
//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

func main() {
	router := gin.Default()         // Create server
	router.GET("/todos", getTodos)  // Create endpoint
	router.Run("localhost:9090")    // Run server
}
