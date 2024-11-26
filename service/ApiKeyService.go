package service

import (
	"IoT-GPS-Backend/entity"
	"IoT-GPS-Backend/repository"
	"crypto/rand"
	"encoding/base64"
	"github.com/google/uuid"
	"time"
)

type ApiKeyService interface {
	CreateApiKey(deviceId string) (string, error)
	ValidateApiKey(apiKey string, deviceId string) (bool, error)
}

type apiKeyService struct {
	apiKeyRepository repository.ApiKeyRepository
}

func NewApiKeyService(apiKeyRepository repository.ApiKeyRepository) ApiKeyService {
	return &apiKeyService{apiKeyRepository}
}

func (a *apiKeyService) CreateApiKey(deviceId string) (string, error) {
	newApiKey, err := generateAPIKey()
	if err != nil {
		return "", err
	}

	key := entity.ApiKey{
		Id:          uuid.New().String(),
		IoTDeviceId: deviceId,
		Key:         newApiKey,
		CreatedAt:   time.Now(),
	}

	_, err = a.apiKeyRepository.Save(key)
	if err != nil {
		return "", err
	}

	return newApiKey, nil
}

func (a *apiKeyService) ValidateApiKey(apiKey string, deviceId string) (bool, error) {
	key, err := a.apiKeyRepository.GetByKey(apiKey)
	if err != nil {
		return false, err
	}

	if key.IoTDeviceId != deviceId {
		return false, nil
	}

	return apiKey == key.Key, nil
}

func generateAPIKey() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}
