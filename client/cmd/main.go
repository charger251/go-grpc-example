package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"
	pb "proto" // Импорт сгенерированного protobuf-кода
	"google.golang.org/grpc"
)

func main() {
	// Установить соединение с сервером gRPC
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
	log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Получить новый экземпляр клиента
	client := pb.NewKeyValueClient(conn)

	var action, key, value string

	// Ожидается что-то вроде "set foo bar"
	if len(os.Args) > 2 {
	action, key = os.Args[1], os.Args[2]
	value = strings.Join(os.Args[3:], " ")
	}

	// Установить 1-секундный тайм-аут с по­мощью context.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	// Вызвать client.Get() или client.Put().
	switch action {
	case "get":
    r, err := client.Get(ctx, &pb.GetRequest{Key: key})
    if err != nil {
      log.Fatalf("could not get value for key %s: %v\n", key, err)
    }
    log.Printf("Get %s returns: %s", key, r.Value)
	case "put":
    _, err := client.Put(ctx, &pb.PutRequest{Key: key, Value: value})
    if err != nil {
      log.Fatalf("could not put key %s: %v\n", key, err)
    }
    log.Printf("Put %s", key)
	default:
	  log.Fatalf("Syntax: go run [get|put] KEY VALUE...")
  }
}