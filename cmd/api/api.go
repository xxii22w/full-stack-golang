package main

import (
	"time"

	"github.com/xxii22w/fullstackgo/internal/auth"
	"github.com/xxii22w/fullstackgo/internal/store"
	"go.uber.org/zap"
)

type application struct {
	config        Config
	store         store.Storage
	authenticator auth.Authenticator
	logger        *zap.SugaredLogger
}

type Config struct {
	addr        string
	db          dbConfig
	env         string
	apiURL      string
	frontendURL string
	auth        authConfig
	redisCfg    redisConfig
	// rateLimiter ratelimiter.Config

}

type authConfig struct {
	basic basicConfig
	token tokenConfig
}

type tokenConfig struct {
	secret string
	exp    time.Duration
	iss    string
}

type basicConfig struct {
	user string
	pass string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type redisConfig struct {
	addr    string
	pw      string
	db      int
	enabled bool
}
