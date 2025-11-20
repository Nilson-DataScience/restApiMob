package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.GET("/", func(context *Context) {
		context.JSON(200, H{
			"message": "hello go x from vercel !!!!",
		})
	})

	server.GET("/mob", func(context *Context) {
		data, err := os.ReadFile("db.json")
		if err != nil {
			context.JSON(500, H{
				"error": "MENSAGEM DE ERRO ATUALIZADA: Arquivo não encontrado!",
			})
			return
		}

		var jsonData map[string]interface{}
		if err := json.Unmarshal(data, &jsonData); err != nil {
			context.JSON(500, H{
				"error": "Erro ao decodificar JSON!",
			})
			return
		}

		context.JSON(200, jsonData)
	})

	server.GET("/hello", func(context *Context) {
		name := context.Query("name")
		if name == "" {
			context.JSON(400, H{
				"message": "name not found",
			})
		} else {
			context.JSON(200, H{
				"data": fmt.Sprintf("Hello %s!", name),
			})
		}
	})
	server.GET("/user/:id", func(context *Context) {
		context.JSON(400, H{
			"data": H{
				"id": context.Param("id"),
			},
		})
	})
	server.GET("/long/long/long/path/*test", func(context *Context) {
		context.JSON(200, H{
			"data": H{
				"url": context.Path,
			},
		})
	})
	server.Handle(w, r)
}
