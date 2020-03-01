package model

import (
	"context"

	"github.com/cymon1997/go-backend/entity"
	"github.com/cymon1997/go-backend/internal/database"
	"github.com/cymon1997/go-backend/internal/log"
	"github.com/cymon1997/go-backend/internal/mq"
	"github.com/cymon1997/go-backend/internal/redis"
)

const healthTag = "Article|Health"

type HealthModel struct {
	redisClient redis.Client
	dbClient    database.Client
	publisher   mq.Publisher
}

func (m *HealthModel) Do(ctx context.Context) (entity.Response, error) {
	var response entity.Response
	err := m.Validate(ctx)
	if err != nil {
		log.ErrorDetail(healthTag, "error validation: %v", err)
		response.Message = err.Error()
		return response, err
	}
	err = m.redisClient.Dial()
	if err != nil {
		log.ErrorDetail(healthTag, "error dial redis: %v", err)
		response.Message = err.Error()
		return response, err
	}
	err = m.dbClient.Dial()
	if err != nil {
		log.ErrorDetail(healthTag, "error dial database: %v", err)
		response.Message = err.Error()
		return response, err
	}
	err = m.publisher.Publish("Health", "data")
	if err != nil {
		log.ErrorDetail(healthTag, "error dial message queue: %v", err)
		response.Message = err.Error()
		return response, err
	}
	return response, nil
}

func (m *HealthModel) Validate(ctx context.Context) error {
	return nil
}
