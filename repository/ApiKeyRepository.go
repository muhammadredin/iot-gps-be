package repository

import (
	"IoT-GPS-Backend/entity"
	"database/sql"
	"log"
)

type ApiKeyRepository interface {
	GetByKey(key string) (entity.ApiKey, error)
	Save(device entity.ApiKey) (entity.ApiKey, error)
}

type apiKeyRepository struct {
	db *sql.DB
}

func NewApiKeyRepository(db *sql.DB) ApiKeyRepository {
	return &apiKeyRepository{db}
}

func (a *apiKeyRepository) GetByKey(key string) (entity.ApiKey, error) {
	var apiKey entity.ApiKey
	err := a.db.QueryRow("SELECT * FROM m_api_keys WHERE api_key = $1", key).
		Scan(&apiKey.Id, &apiKey.IoTDeviceId, &apiKey.Key, &apiKey.CreatedAt)
	if err != nil {
		log.Fatal(err)
		return apiKey, err
	}

	return apiKey, nil
}

func (a *apiKeyRepository) Save(apiKey entity.ApiKey) (entity.ApiKey, error) {
	_, err := a.db.Exec("INSERT INTO m_api_keys (id, iot_device_id, api_key, created_at) VALUES ($1, $2, $3, $4)",
		apiKey.Id, apiKey.IoTDeviceId, apiKey.Key, apiKey.CreatedAt)
	if err != nil {
		return entity.ApiKey{}, err
	}
	return apiKey, nil
}
