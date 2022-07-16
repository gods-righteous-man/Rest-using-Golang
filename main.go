package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{"1", "Clean Room", false},
	{"2", "Read book", true},
	{"3", "Do Laundry", true},
}

func addTodo(context *gin.Context) {

	var newTodo Todo

	if err := context.BindJSON(&newTodo); err != nil {

		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)

}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodo(context *gin.Context) {

	id := context.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)

}

func getTodoById(id string) (*Todo, error) {

	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}

	}
	return nil, errors.New("Todo not found")
}

func toggleTodoStatus(context *gin.Context) {

	id := context.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})

	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)

}

func main() {

	router := gin.Default() //for creating server

	router.GET("/todos", getTodos) //endpoint

	router.POST("/todos", addTodo)

	router.GET("/todos/:id", getTodo) //read specific

	router.PATCH("/todos/:id", toggleTodoStatus) //update

	router.Run("localhost:9090") //to run server
}
