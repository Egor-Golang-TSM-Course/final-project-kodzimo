package main

import (
	"final-project-kodzimo/internal/gateway"
	"final-project-kodzimo/internal/hashing"
	"final-project-kodzimo/internal/storage"
	pb "final-project-kodzimo/proto"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	redisClient, err := storage.ConnectToRedis()
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	hashingService := hashing.NewHashingService(redisClient)

	/*
		Этот код (ниже) создает gRPC сервер и регистрирует ваш Hashing Service на этом сервере.
		Затем он начинает слушать входящие запросы на порту 50051.
	*/

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHashingServer(s, &hashing.Server{HashingService: hashingService})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	/*
		В этом коде мы создаем новый экземпляр GatewayService, регистрируем обработчики HTTP для каждого
		из методов, а затем запускаем HTTP-сервер, который слушает на порту 8080. Обратите внимание, что мы
		запускаем gRPC сервер в отдельной горутине, чтобы основной поток мог продолжить и запустить HTTP-сервер.
	*/

	// Создаем соединение с gRPC сервером
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	// Создаем новый Gateway Service
	gw := &gateway.GatewayService{
		HashingClient: pb.NewHashingClient(conn),
	}

	// Регистрируем обработчики HTTP
	http.HandleFunc("/checkhash", gw.CheckHashHandler)
	http.HandleFunc("/gethash", gw.GetHashHandler)
	http.HandleFunc("/createhash", gw.CreateHashHandler)

	// Запускаем HTTP-сервер
	log.Fatal(http.ListenAndServe(":8080", nil))
}
