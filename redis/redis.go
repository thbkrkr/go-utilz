package redis

import (
	"github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
)

// NewConn creates a Redis connection
func NewConn(redisURL string) redis.Conn {
	conn, err := redis.Dial("tcp", redisURL)
	if err != nil {
		logrus.WithError(err).Fatal("Fail to create redis connection")
	}

	_, err = conn.Do("PING")
	if err != nil {
		logrus.WithError(err).Fatal("Fail to ping redis")
	}

	logrus.WithField("url", redisURL).Info("Redis connection ready")

	return conn
}
