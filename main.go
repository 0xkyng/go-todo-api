package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	// json tag to de-serialize json todo
	ID          string `json:"id"`
	Item        string  `json:"item"`
	Completed   bool     `json:"completed"`
}

var todos = []todo {
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Watch Footbal", Completed: false},
}

/////////////////////////////////////////////
// Get todos ---GET Method
func getTodos(context *gin.Context) {
	// Context contains info about the incoming
	// http request
	context.IndentedJSON(http.StatusOK, todos) // Converting []todo adta into json
}
////////////////////////////////////////////////

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
// Add todos ---POST Method
func addTodo(context *gin.Context) {
	var newTodo todo
 // using BindJson method to serialize todo or extract data
 // From []todo into newTodo
	 err := context.BindJSON(&newTodo)
	 if err != nil {
		return
	 }

	 todos = append(todos, newTodo)

	 context.IndentedJSON(http.StatusCreated, newTodo)
}
//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

//---------------------------------------------
// Update todos
func updateTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)


}
//---------------------------------------------

func getTodoById(id string) (*todo, error){
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()         // Create server
	router.GET("/todos", getTodos)  // Create endpoint for GET Method
	router.GET("/todos/:id", getTodo)  // Create endpoint a specific id
	router.PATCH("/todos/:id", updateTodoStatus)  // Create endpoint for PATCH Method
	router.POST("/todos", addTodo)  // endpoint for POST method
	router.Run("localhost:9090")    // Run server
}
