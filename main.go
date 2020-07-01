package main

import (
	"context"
	"github.com/coding-socks/glue"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	_ = godotenv.Load(glue.EnvironmentFilePath())

	go func() {
		<-s
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		_ = glue.Shutdown(ctx)
	}()

	if err := glue.Handle(glue.NewConfig()); err != nil {
		log.Fatal(err)
	}
}
