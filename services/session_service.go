package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ablancas22/messenger-frontend/models"
	"github.com/go-redis/redis"
	"github.com/gofrs/uuid"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// generateUuid
func generateUuid() string {
	curId, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return curId.String()
}

// SHA runs encrypts an input string
func SHA(str string) string {
	bytes := []byte(str)
	// Converts string to sha2
	h := sha256.New()                   // new sha256 object
	h.Write(bytes)                      // data is now converted to hex
	code := h.Sum(nil)                  // code is now the hex sum
	codestr := hex.EncodeToString(code) // converts hex to string
	return codestr
}

// RedisManager is used for managing the web app's user sessions and state
type RedisManager struct {
	Client *redis.Client
	Ctx    context.Context
}

// InitRedisClient initializes a redis manager struct for session management
func InitRedisClient() *RedisManager {
	var rm RedisManager
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})
	rm.Client = rdb
	rm.Ctx = ctx
	return &rm
}

// Set a RedisManager value
func (rm *RedisManager) Set(key string, value string) error {
	err := rm.Client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get a RedisManager value
func (rm *RedisManager) Get(key string) (string, error) {
	val, err := rm.Client.Get(key).Result()
	if err != nil {
		return val, err
	}
	return val, nil
}

// Del a RedisManager value
func (rm *RedisManager) Del(key string) error {
	_, err := rm.Client.Del(key).Result()
	if err != nil {
		return err
	}
	return nil
}

// SessionService is a struct wrapper around an initialized RedisManager
type SessionService struct {
	RedisManager *RedisManager
}

// NewSessionService initializes a Session Manager
func NewSessionService() *SessionService {
	redisManager := InitRedisClient()
	return &SessionService{RedisManager: redisManager}
}

// lookupSessionID based on a checkId input
func (m *SessionService) lookupSessionID(checkId string) (string, error) {
	val, err := m.RedisManager.Get(checkId)
	if err != nil {
		return "", err
	}
	z := strings.Split(val, ":")
	if len(z) < 2 {
		err = m.RedisManager.Del(checkId)
		if err != nil {
			return "", err
		}
		return "", errors.New("invalid session string")
	}
	if z[0] == "" {
		return "", nil
	}
	return z[1], nil
}

// IsLoggedIn checks if a user making a request is logged in
func (m *SessionService) IsLoggedIn(r *http.Request) bool {
	var expectedSessionID string
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println(err)
		return false
	}
	// TODO - GET IP FROM r AND COMPARE WITH IP FROM  m.lookupSessionID
	expectedSessionID, err = m.lookupSessionID(cookie.Value)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if expectedSessionID != "" {
		return true
	}
	return false
}

// NewSession initializes a new user session based on an Auth instance
func (m *SessionService) NewSession(auth *models.Auth) (*http.Cookie, error) {
	authStr := auth.GetAuthString()
	checkUuid := generateUuid()
	checkUuid = strings.Replace(checkUuid, "-", "", -1)
	cookieValue := checkUuid + ":" + SHA(checkUuid+strconv.Itoa(rand.Intn(100000000)))
	expire := time.Now().AddDate(0, 0, 1)
	err := m.RedisManager.Set(cookieValue, authStr)
	if err != nil {
		return &http.Cookie{}, err
	}
	return &http.Cookie{Name: "SessionID", Value: cookieValue, Expires: expire, HttpOnly: true}, nil
}

// GetSession returns a user Auth session
func (m *SessionService) GetSession(cookie *http.Cookie) (*models.Auth, error) {
	sessionID := cookie.Value
	authStr, err := m.RedisManager.Get(sessionID)
	if err != nil {
		return &models.Auth{}, err
	}
	auth := models.Auth{Authenticated: false}
	auth.LoadAuthString(authStr)
	return &auth, nil
}

// DeleteSession deletes a user session
func (m *SessionService) DeleteSession(cookie *http.Cookie) error {
	sessionID := cookie.Value
	err := m.RedisManager.Del(sessionID)
	if err != nil {
		return err
	}
	return nil
}
