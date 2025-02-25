package config

import (
	"errors"
	"os"
	"time"
)

type AuthenticationConfig struct {
	accessSecretKey  string
	refreshSecretKet string
	refreshTime      time.Duration
	accessTime       time.Duration
	authPrefix       string
}

func NewAuthConfig() (*AuthenticationConfig, error) {
	accessKey := os.Getenv("ACCESS_SECRET_KEY")
	if accessKey == "" {
		return nil, errors.New("don't parse accessSecretKey")
	}
	refreshKey := os.Getenv("REFRESH_SECRET_KEY")
	if refreshKey == "" {
		return nil, errors.New("don't parse refreshSecretKey")
	}

	authPrefix := os.Getenv("AUTH_PREFIX")
	if authPrefix == "" {
		return nil, errors.New("don't parse prefix Authentication")
	}

	refreshTime := 60 * time.Minute
	accessTime := 5 * time.Minute

	return &AuthenticationConfig{
		accessSecretKey:  accessKey,
		refreshSecretKet: refreshKey,
		refreshTime:      refreshTime,
		accessTime:       accessTime,
		authPrefix:       authPrefix,
	}, nil
}

func (a *AuthenticationConfig) AccessSecretKey() string {
	return a.accessSecretKey
}

func (a *AuthenticationConfig) RefreshSecretKey() string {
	return a.refreshSecretKet
}

func (a *AuthenticationConfig) RefreshTime() time.Duration {
	return a.refreshTime
}

func (a *AuthenticationConfig) AccessTime() time.Duration {
	return a.accessTime
}

func (a *AuthenticationConfig) AuthenticationPrefix() string {
	return a.authPrefix
}
