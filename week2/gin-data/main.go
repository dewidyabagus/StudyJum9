package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const defAddress = ":8080"

type Person struct {
	ID   string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"min=2"`
}

func main() {
	// Membuat Gin Engine dengan default configuration
	router := gin.Default()
	router.LoadHTMLGlob("templates/*") // Set templates location

	// Gin mendukung banyak format rending seperi html, json, xml, toml, dan yaml
	router.GET("/biodata/:id", func(ctx *gin.Context) {
		var (
			id         = ctx.Param("id")
			renderType = ctx.Query("render")
			biodata    = gin.H{
				"id":         id,
				"email":      "john.wick@example.com",
				"first_name": "John",
				"last_name":  "wick",
				"hobbies":    []string{"coding", "bersepeda", "berenang"},
				"age":        30,
			}
		)

		switch renderType {
		default:
			ctx.String(http.StatusOK, fmt.Sprintf("%+v", biodata))

		case "html":
			ctx.HTML(http.StatusOK, "biodata.html", biodata)

		case "json":
			ctx.JSON(http.StatusOK, biodata)

		case "xml":
			ctx.Negotiate(http.StatusOK, gin.Negotiate{Offered: []string{gin.MIMEXML, gin.MIMEXML2}, Data: biodata})

		case "toml":
			ctx.TOML(http.StatusOK, biodata)

		case "yaml":
			ctx.YAML(http.StatusOK, biodata)
		}
	})

	router.GET("/person/:id/:name", func(ctx *gin.Context) {
		var params Person
		if err := ctx.ShouldBindUri(&params); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		// Binding data query menggunakan tag form, terus untuk method binding menggunakan ShouldBind()
		ctx.JSON(http.StatusOK, params)
	})

	// Memulai HTTP Server dan setup graceful shutdown
	server := &http.Server{
		Addr:    defAddress,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("listen and serve failed:", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("shutdown server failed:", err.Error())
	}
}
