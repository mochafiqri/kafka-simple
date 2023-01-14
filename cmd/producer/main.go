package main

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/segmentio/kafka-go"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	var serverConf = &http.Server{
		Addr: ":8080",
	}

	var p = initProducer()

	e.POST("/send-message", func(c echo.Context) error {
		var ctx = context.Background()
		var tmp = c.FormValue("message")
		var topic = c.FormValue("topic")
		var message = map[string]interface{}{
			"message": tmp,
		}

		val, err := json.Marshal(message)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}

		kafkaMsg := kafka.Message{
			Key:   []byte("Simple Message"),
			Topic: topic,
			Value: val,
		}

		err = p.WriteMessages(ctx, kafkaMsg)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, message)
	})

	e.StartServer(serverConf)
}

func initProducer() *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP("localhost:29092"),
		Balancer:               &kafka.CRC32Balancer{Consistent: true},
		RequiredAcks:           kafka.RequireAll,
		AllowAutoTopicCreation: true,
	}
}
