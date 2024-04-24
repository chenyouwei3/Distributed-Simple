package main

import (
	"Distributed-Simple/log"
	"Distributed-Simple/service"
	"context"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	ctx, err := service.Start(context.Background(), "Log Service", "4000", log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
