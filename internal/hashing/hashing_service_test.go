package hashing

import (
	"context"
	"testing"

	pb "final-project-kodzimo/proto"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

/*
Этот тест проверяет, что метод CreateHash не возвращает ошибку и возвращает непустой хеш.
*/
func TestCreateHash(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

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
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	service := NewHashingService(client)
	req := &pb.HashRequest{Payload: "test"}

	// Создаем хеш
	_, err := service.CreateHash(context.Background(), req)
	assert.NoError(t, err)

	// Проверяем хеш
	resp, err := service.CheckHash(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Hash)
}

/*
Этот тест проверяет, что метод GetHash не возвращает ошибку и возвращает тот же хеш, который был создан.
*/
func TestGetHash(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

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
Для запуска unit-тестов вы можете использовать go test -run 'Unit',
а для интеграционных тестов - go test -run 'Integration'.
*/
