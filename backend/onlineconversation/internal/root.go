package internal

import (
	"backend/common/producer"
	"backend/onlineconversation/internal/client"
	"backend/onlineconversation/internal/consumer"
	"backend/onlineconversation/internal/controller"
	"backend/onlineconversation/internal/repository"
	"backend/onlineconversation/internal/service"
	"log/slog"
	"net/http"
	"os"
)

func NewServer() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	slog.SetDefault(logger)

	kp := producer.NewSyncProducer("online-conversation")

	r := repository.NewRepository()

	s := service.NewService(r, kp, client.NewOllamaClient())

	mux := http.NewServeMux()

	kc := consumer.NewConsumer(s, kp)
	go func() {
		err := kc.GetMessage([]string{"online-conversation"})
		if err != nil {
			panic(err)
		}
	}()

	controller.NewController(s, mux)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.Error("fail to listen and serve http",
			"err", err)
		panic(err)
	}
}
