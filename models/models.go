package models

import (
	"github.com/go-redis/redis"
	"time"
)

type clip struct {
	ID           int           `json:"id"`
	Message      string        `json:"clip"`
	Username     string        `json:"username"`
	AuthCode     string        `json:"auth_code"`
	AuthDuration time.Duration `json:"auth_duration"`
}

var CLIENT = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "redisXgolang",
	DB:       0,
})
