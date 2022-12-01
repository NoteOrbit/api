package main

import (
	"net/http"
	"fmt"
	"os"
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


func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run(fmt.Sprintf(":%s", port))
}



// package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/gin-gonic/gin"
// )


// type todo struct {
// 	ID        string `json:"id"`
// 	Item      string `json:"item"`
// 	Completed bool   `json:"completed"`
// 	Price     int    `json:"price"`
// }

// var todos = []todo{
// 	{ID: "1", Item: "Nap x Warin", Completed: true, Price: 20},
// 	{ID: "2", Item: "Sangob", Completed: false, Price: 11},
// 	{ID: "3", Item: "SongSarn", Completed: false, Price: 13},
// 	{ID: "4", Item: "LIFE Roasters", Completed: false, Price: 5},
// 	{ID: "5", Item: "BalconyKiss Coffee", Completed: false, Price: 3},
// 	{ID: "6", Item: "Anna Roasters", Completed: false, Price: 2},
// 	{ID: "7", Item: "Amarna", Completed: true, Price: 1},
// }

// func main() {
// 	port := os.Getenv("PORT")

// 	if port == "" {
// 		port = "8000"
// 	}

// 	r := gin.Default()

// 	r.GET("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "Hello, world!",
// 		})
// 	})

// 	r.GET("/:name", func(c *gin.Context) {
// 		name := c.Param("name")

// 		c.JSON(200, gin.H{
// 			"message": fmt.Sprintf("Hello, %s!", name),
// 		})
// 	})

// 	r.Run(fmt.Sprintf(":%s", port))
// }