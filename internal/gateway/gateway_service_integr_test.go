package gateway

import (
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"final-project-kodzimo/internal/hashing"
	"final-project-kodzimo/internal/storage"
	pb "final-project-kodzimo/proto"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestCreateHashHandlerIntegration(t *testing.T) {
	// Подключаемся к реальной базе данных Redis
	redisClient, err := storage.ConnectToRedis()
	if err != nil {
		t.Fatalf("failed to connect to Redis: %v", err)
	}

	// Создаем HashingService и gRPC-сервер
	hashingService := hashing.NewHashingService(redisClient)
	server := grpc.NewServer()
	pb.RegisterHashingServer(server, &hashing.Server{HashingService: hashingService})

	// Запускаем gRPC-сервер в отдельной горутине
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			t.Fatalf("failed to listen: %v", err)
		}
		if err := server.Serve(lis); err != nil {
			t.Fatalf("failed to serve: %v", err)
		}
	}()

	// Создаем GatewayService с реальным gRPC-клиентом
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	gw := &GatewayService{
		HashingClient: pb.NewHashingClient(conn),
	}

	// Отправляем HTTP-запрос к CreateHashHandler
	req, err := http.NewRequest("POST", "/createhash", strings.NewReader("test"))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(gw.CreateHashHandler)
	handler.ServeHTTP(rr, req)

	// Проверяем, что ответ содержит ожидаемый HTTP-статус и тело
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.NotEmpty(t, rr.Body.String())
}
