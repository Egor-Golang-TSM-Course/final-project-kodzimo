package hashing

import (
	"context"
	"testing"

	"final-project-kodzimo/internal/storage"
	pb "final-project-kodzimo/proto"

	"github.com/stretchr/testify/assert"
)

/*
Этот тест проверяет, что метод CreateHash не возвращает ошибку и возвращает непустой хеш.
*/
func TestCreateHash(t *testing.T) {
	client, err := storage.ConnectToRedis()
	if err != nil {
		t.Fatalf("failed to connect to Redis: %v", err)
	}

	service := NewHashingService(client)
	req := &pb.HashRequest{Payload: "test"}

	resp, err := service.CreateHash(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Hash)
}

/*
Этот тест проверяет, что метод CheckHash не возвращает ошибку и возвращает непустой хеш.
*/
func TestCheckHash(t *testing.T) {
	client, err := storage.ConnectToRedis()
	if err != nil {
		t.Fatalf("failed to connect to Redis: %v", err)
	}

	service := NewHashingService(client)
	req := &pb.HashRequest{Payload: "test"}

	// Создаем хеш
	createResp, err := service.CreateHash(context.Background(), req)
	assert.NoError(t, err)

	// Проверяем хеш
	checkReq := &pb.HashRequest{Payload: createResp.GetHash()}
	checkResp, err := service.CheckHash(context.Background(), checkReq)

	assert.NoError(t, err)
	if checkResp != nil {
		assert.NotEmpty(t, checkResp.Hash)
	}
}

/*
Этот тест проверяет, что метод GetHash не возвращает ошибку и возвращает тот же хеш, который был создан.
*/
func TestGetHash(t *testing.T) {
	client, err := storage.ConnectToRedis()
	if err != nil {
		t.Fatalf("failed to connect to Redis: %v", err)
	}

	service := NewHashingService(client)
	req := &pb.HashRequest{Payload: "test"}

	// Создаем хеш
	createResp, err := service.CreateHash(context.Background(), req)
	assert.NoError(t, err)

	// Получаем хеш
	getReq := &pb.HashRequest{Payload: createResp.GetHash()}
	getResp, err := service.GetHash(context.Background(), getReq)

	assert.NoError(t, err)
	assert.NotNil(t, getResp)
	assert.Equal(t, req.GetPayload(), getResp.GetHash())
}

/*
Для запуска тестов вы можете использовать go test -run 'Имя_теста'.
*/
