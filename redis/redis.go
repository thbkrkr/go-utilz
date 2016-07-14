package redis

import (
	"github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
)

var defaultRedisURL = ":6379"

// Conn creates a Redis connection using the env var REDIS_URL
func Conn(redisURL string) redis.Conn {
	if redisURL == "" {
		redisURL = defaultRedisURL
	}

	conn, err := redis.Dial("tcp", redisURL)
	if err != nil {
		logrus.Fatal(err)
	}

	_, err = conn.Do("PING")
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.WithField("url", redisURL).Info("Redis connection ready")

	return conn
}
