package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    // i believe this is called the router
    r := gin.Default()

    // we route the path to a handler function
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "hey! ho!": "let's go!",
        })
    })

    // run it and check for errors
    err := r.Run(":9808")
    if err != nil {
        panic(fmt.Sprintf("failed to start web server. Error: %v", err))
    }
}
