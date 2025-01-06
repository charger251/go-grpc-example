package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "proto" // Импорт сгенерированного protobuf-кода
	"google.golang.org/grpc"
)

// Реализация gRPC сервера
type server struct {
	pb.UnimplementedKeyValueServer
}

// Реализация метода Get
func (s *server) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Received GET key=%v", r.Key)
	value, err := Get(r.Key)
	if err != nil {
		log.Printf("Error getting key: %v", err)
		return nil, err
	}
	return &pb.GetResponse{Value: value}, nil
}

// Реализация метода Put
func (s *server) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	log.Printf("Received PUT key=%v, value=%v", r.Key, r.Value)
	err := Put(r.Key, r.Value)
	if err != nil {
		log.Printf("Error putting key: %v", err)
		return nil, err
	}
	return &pb.PutResponse{}, nil
}

// Реализация метода Delete
func (s *server) Delete(ctx context.Context, r *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	log.Printf("Received DELETE key=%v", r.Key)
	err := Delete(r.Key)
	if err != nil {
		log.Printf("Error deleting key: %v", err)
		return nil, err
	}
	return &pb.DeleteResponse{}, nil
}

// Хранилище (in-memory key-value store)
var store = make(map[string]string)
var ErrorNoSuchKey = errors.New("no such key")

// Функция для добавления значения
func Put(key string, value string) error {
	store[key] = value
	return nil
}

// Функция для получения значения
func Get(key string) (string, error) {
	value, ok := store[key]
	if !ok {
		return "", ErrorNoSuchKey
	}
	return value, nil
}

// Функция для удаления значения
func Delete(key string) error {
	_, ok := store[key]
	if !ok {
		return ErrorNoSuchKey
	}
	delete(store, key)
	return nil
}

// Главная функция
func main() {
	// Создаём gRPC-сервер
	s := grpc.NewServer()
	pb.RegisterKeyValueServer(s, &server{}) // Регистрируем сервис

	// Настраиваем слушатель на порту 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Server is listening on port 50051...")
	// Запускаем сервер
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}