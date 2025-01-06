# gRPC Key-Value Service

A simple gRPC-based Key-Value storage service implemented in Go. The service provides basic CRUD (Create, Read, Update, Delete) operations for key-value pairs.

---

## Features

- **Put**: Add or update a key-value pair in the storage.
- **Get**: Retrieve the value of a given key.
- **Delete**: Remove a key-value pair from the storage.
- **In-memory Storage**: Uses a simple in-memory map for demonstration purposes.

---

## Prerequisites

Before running the project, ensure you have the following installed:

- **Go**: Version 1.19 or higher.
- **Protocol Buffers (`protoc`)**: For generating Go code from `.proto` files.
- **`protoc-gen-go`** and **`protoc-gen-go-grpc`**: Plugins for Go code generation.

Install the Go plugins if not already installed:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Ensure the `GOPATH/bin` is in your system's `PATH` so `protoc` can find the plugins.

---

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/username/grpc-keyvalue-service.git
cd grpc-keyvalue-service
```

### 2. Generate gRPC code

Make sure the Protocol Buffers compiler and plugins are properly installed. Generate the Go code from the `.proto` file:

```bash
protoc --proto_path=proto \
    --go_out=proto --go_opt=paths=source_relative \
    --go-grpc_out=proto --go-grpc_opt=paths=source_relative \
    proto/keyvalue.proto
```

This will generate the following files in the `proto/` directory:
- `keyvalue.pb.go` — Code for Protocol Buffers messages.
- `keyvalue_grpc.pb.go` — Code for gRPC service definitions and interfaces.

### 3. Install dependencies

Use `go mod` to install required dependencies:

```bash
go mod tidy
```

---

## Running the Service

### 1. Start the gRPC Server

Run the gRPC server:

```bash
go run server/main.go
```

By default, the server will start on port `50051`.

---

## Testing the Service

### 1. Using `grpcurl`

You can test the service using the `grpcurl` command-line tool.

#### Install `grpcurl` (if not already installed):
Follow the [official grpcurl installation guide](https://github.com/fullstorydev/grpcurl#installation).

#### Test `Put` Method
Add a key-value pair to the in-memory store:

```bash
grpcurl -plaintext -d '{"key": "example_key", "value": "example_value"}' \
    localhost:50051 KeyValue.Put
```

#### Test `Get` Method
Retrieve the value of a given key:

```bash
grpcurl -plaintext -d '{"key": "example_key"}' \
    localhost:50051 KeyValue.Get
```

Expected response:
```json
{
  "value": "example_value"
}
```

#### Test `Delete` Method
Delete a key-value pair:

```bash
grpcurl -plaintext -d '{"key": "example_key"}' \
    localhost:50051 KeyValue.Delete
```

---

## Project Structure

```
grpc-keyvalue-service/
├── proto/
│   ├── keyvalue.proto           # Protocol Buffers definition
│   ├── keyvalue.pb.go           # Generated Protobuf code
│   ├── keyvalue_grpc.pb.go      # Generated gRPC code
├── server/
│   ├── main.go                  # gRPC server implementation
├── client/
│   ├── main.go                  # gRPC client implementation (optional)
├── go.mod                       # Go module file
└── README.md                    # Project documentation
```

---

## Example `.proto` File

Here is the `.proto` definition used in the project:

```proto
syntax = "proto3";

package keyvalue;

service KeyValue {
  rpc Put (PutRequest) returns (PutResponse);
  rpc Get (GetRequest) returns (GetResponse);
  rpc Delete (DeleteRequest) returns (DeleteResponse);
}

message PutRequest {
  string key = 1;
  string value = 2;
}

message PutResponse {}

message GetRequest {
  string key = 1;
}

message GetResponse {
  string value = 1;
}

message DeleteRequest {
  string key = 1;
}

message DeleteResponse {}
```

---

## Improvements and Next Steps

- Add persistent storage (e.g., a database) instead of in-memory storage.
- Implement error handling for invalid inputs and missing keys.
- Add unit tests for server methods.
- Add authentication to secure the gRPC service.

---

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests to improve the project.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
```

---

### **Что включено в README:**
1. **Обзор проекта**: Описаны функции сервиса.
2. **Инструкция по установке**: Указано, как настроить окружение и сгенерировать код.
3. **Запуск сервера**: Пошаговые указания для запуска сервера.
4. **Тестирование**: Примеры команд для тестирования с `grpcurl`.
5. **Структура проекта**: Описание структуры каталогов.
6. **Пример `.proto` файла**: Приведён полный файл `.proto` для наглядности.
7. **Идеи для улучшений**: Указаны дальнейшие шаги для доработки.
8. **Лицензия и вклад**: Информация о лицензии и как вносить изменения.

Вы можете адаптировать этот файл под свои нужды.
