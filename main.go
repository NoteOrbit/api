package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
	Price     int    `json:"price"`
}

var todos = []todo{
	{ID: "1", Item: "Nap x Warin", Completed: true, Price: 20},
	{ID: "2", Item: "Sangob", Completed: false, Price: 11},
	{ID: "3", Item: "SongSarn", Completed: false, Price: 13},
	{ID: "4", Item: "LIFE Roasters", Completed: false, Price: 5},
	{ID: "5", Item: "BalconyKiss Coffee", Completed: false, Price: 3},
	{ID: "6", Item: "Anna Roasters", Completed: false, Price: 2},
	{ID: "7", Item: "Amarna", Completed: true, Price: 1},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusOK, newTodo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("todos", addTodo)
	router.Run("localhost:8000")
}
