package bootstrap

import "github.com/go-redis/redis/v8"

func NewRedis(env *Env) *redis.Client {
	rds := redis.NewClient(&redis.Options{
		Addr:     env.RedisHost + ":" + env.RedisPort,
		Password: env.RedisPassword,
		DB:       0, // use default DB
	})

	return rds
}
