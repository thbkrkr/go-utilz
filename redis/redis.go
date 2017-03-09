package redis

import (
	"os"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
)

var maxConnections = 1000

func init() {
	maxConn := os.Getenv("REDIS_MAX_CONNECTIONS")
	if maxConn != "" {
		maxConnections, _ = strconv.Atoi(maxConn)
	}
}

// NewConn creates a redis connection
func NewConn(redisURL string, redisPwd string) redis.Conn {
	conn, err := redis.Dial("tcp", redisURL, redis.DialPassword(redisPwd))
	if err != nil {
		logrus.WithError(err).WithField("url", redisURL).
			Fatal("Fail to create redis connection")
	}

	_, err = conn.Do("PING")
	if err != nil {
		logrus.WithError(err).WithField("url", redisURL).
			Fatal("Fail to ping redis")
	}

	logrus.WithField("url", redisURL).Info("Redis connection ready")

	return conn
}

// NewPool creates a redis connection pool
func NewPool(redisURL string, redisPwd string) *redis.Pool {
	pool := redis.NewPool(func() (redis.Conn, error) {
		conn, err := redis.Dial("tcp", redisURL, redis.DialPassword(redisPwd))
		if err != nil {
			logrus.WithError(err).Fatal("Fail to create redis connection pool")
			return nil, err
		}

		_, err = conn.Do("PING")
		if err != nil {
			logrus.WithError(err).Fatal("Fail to ping redis")
			return nil, err
		}

		logrus.WithField("url", redisURL).Info("Redis connection ready")
		return conn, nil
	}, maxConnections)

	return pool
}
