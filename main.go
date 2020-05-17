package main

import (
    "ec-by-go/presentation/controller"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("view/*.html")

    router.GET("/healthcheck/", controller.Healthcheck())

    router.GET("/", controller.Top())

    router.Run()
}
