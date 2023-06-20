package redis

import (
	"Goplus/pkg/logger"
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
)

// RedisClient Redis 服务
type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

// once 确保全局的 Redis 对象只实例一次
var once sync.Once

// Redis 全局 Redis，使用 db 1
var Redis *RedisClient

// ConnectRedis 连接 redis 数据库，设置全局的 Redis 对象
func ConnectRedis(address string, usaername string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, usaername, password, db)
	})
}

// NewClient 创建一个新的 redis 连接
func NewClient(address string, usaername string, password string, db int) *RedisClient {
	// 初始化自定的 RedisClient 实例
	rds := &RedisClient{}
	// 使用默认的 context
	rds.Context = context.Background()
	// 使用 redis 库里的 NewClient 初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: usaername,
		Password: password,
		DB:       db,
	})
	// 测试一下连接
	err := rds.Ping()
	logger.LogIf(err)
	return rds
}

// Ping 用以测试 redis 连接是否正常
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}
