package provider

import (
	"sync"

	"github.com/cymon1997/go-backend/internal/database"
	"github.com/cymon1997/go-backend/internal/elastic"
	"github.com/cymon1997/go-backend/internal/mq"
	"github.com/cymon1997/go-backend/internal/redis"
)

var (
	dbClient     database.Client
	syncDbClient sync.Once

	redisClient     redis.Client
	syncRedisClient sync.Once

	mqServer     mq.Server
	syncMqClient sync.Once

	esClient     elastic.Client
	syncEsClient sync.Once
)

func GetDBClient() database.Client {
	syncDbClient.Do(func() {
		cfg := GetAppConfig().DBConfig
		dbClient = database.New(cfg)
	})
	return dbClient
}

func GetRedisClient() redis.Client {
	syncRedisClient.Do(func() {
		cfg := GetAppConfig().RedisConfig
		redisClient = redis.New(&cfg)
	})
	return redisClient
}

func GetMQClient() mq.Server {
	syncMqClient.Do(func() {
		mqServer = mq.New()
	})
	return mqServer
}

func GetESClient() elastic.Client {
	syncEsClient.Do(func() {
		cfg := GetAppConfig().ESConfig
		esClient = elastic.New(cfg)
	})
	return esClient
}
