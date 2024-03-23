package gateway

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	pb "final-project-kodzimo/proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

// HashingClientMock является мок-объектом для pb.HashingClient
type HashingClientMock struct {
	mock.Mock
}

// CheckHash является фиктивной реализацией метода CheckHash
func (m *HashingClientMock) CheckHash(ctx context.Context, in *pb.HashRequest, opts ...grpc.CallOption) (*pb.HashResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.HashResponse), args.Error(1)
}

// GetHash является фиктивной реализацией метода CheckHash
func (m *HashingClientMock) GetHash(ctx context.Context, in *pb.HashRequest, opts ...grpc.CallOption) (*pb.HashResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.HashResponse), args.Error(1)
}

// CreateHash является фиктивной реализацией метода CheckHash
func (m *HashingClientMock) CreateHash(ctx context.Context, in *pb.HashRequest, opts ...grpc.CallOption) (*pb.HashResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.HashResponse), args.Error(1)
}

/*
Этот тест проверяет, что CheckHashHandler возвращает статус 200 OK при получении POST-запроса.
В этом примере мы создаем мок-объект HashingClientMock, который возвращает фиктивный хеш и nil-ошибку
при вызове CheckHash. Затем мы используем этот мок-объект при создании GatewayService в нашем тесте.
*/

func TestCheckHashHandler(t *testing.T) {
	hashingClientMock := new(HashingClientMock)
	hashingClientMock.On("CheckHash", mock.Anything, mock.Anything).Return(&pb.HashResponse{Hash: "testhash"}, nil)

	gw := &GatewayService{
		HashingClient: hashingClientMock,
	}

	req, err := http.NewRequest("POST", "/checkhash", strings.NewReader("test"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(gw.CheckHashHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// Здесь вы можете добавить дополнительные проверки для тела ответа
}

/*
В этом тесте мы будем проверять, что обработчик возвращает http.StatusMethodNotAllowed при получении
запроса с неправильным методом HTTP. Этот тест отправляет GET-запрос к CheckHashHandler и проверяет,
что возвращается статус http.StatusMethodNotAllowed.
*/

func TestCheckHashHandlerInvalidMethod(t *testing.T) {
	hashingClientMock := new(HashingClientMock)
	hashingClientMock.On("CheckHash", mock.Anything, mock.Anything).Return(&pb.HashResponse{Hash: "testhash"}, nil)

	gw := &GatewayService{
		HashingClient: hashingClientMock,
	}

	req, err := http.NewRequest("GET", "/checkhash", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(gw.CheckHashHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

/*
Unit-тесты могут быть написаны для каждого из ваших обработчиков HTTP (CheckHashHandler,
GetHashHandler, CreateHashHandler). Эти тесты могут проверять, что обработчики правильно
обрабатывают запросы и возвращают ожидаемые HTTP-статусы и тела ответов. В Go вы можете
использовать пакет net/http/httptest для создания фиктивных HTTP-запросов и записи ответов.

Интеграционные тесты могут быть написаны для проверки взаимодействия между вашими обработчиками HTTP
и gRPC-сервером. Эти тесты могут проверять, что обработчики правильно вызывают соответствующие методы
на gRPC-клиенте и корректно обрабатывают ответы. Для этих тестов вам может потребоваться фиктивный
gRPC-сервер или мок-объекты.
*/
