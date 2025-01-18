package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/Josh-LeBlanc/gurl/store"
    "github.com/Josh-LeBlanc/gurl/handler"
)

func main() {
    // i believe this is called the router
    r := gin.Default()

    // we route the path to a handler function
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "welcome": "to gurl, the go url shortener",
        })
    })

    r.POST("/create-short-url", func(c *gin.Context) {
        handler.CreateShortUrl(c)
    })

    r.GET("/:shortUrl", func(c *gin.Context) {
        handler.HandleShortUrlRedirect(c)
    })

    // initialize store
    store.InitializeStore()

    // run it and check for errors
    err := r.Run(":9808")
    if err != nil {
        panic(fmt.Sprintf("failed to start web server. Error: %v", err))
    }
}
