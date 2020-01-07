package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"

	logex "github.com/chouandy/go-sdk/log"
	redisex "github.com/chouandy/go-sdk/redis"
)

const accessTokenBytesLength = 32

// AccessTokenTTL 30 mins = 1800 sec
const AccessTokenTTL = 30 * time.Minute

const (
	userAccessTokenKeyFormat  = "%d:%s"
	userAccessTokensKeyFormat = "%d:*"
)

// NewAccessToken new access token
func NewAccessToken() string {
	// New rand bytes
	b := make([]byte, accessTokenBytesLength)
	rand.Read(b)

	// Encode to hex
	enc := make([]byte, len(b)*2)
	hex.Encode(enc, b)

	return string(enc)
}

// NewAccessTokenByUserID new access token by user id
func NewAccessTokenByUserID(userID uint64) (string, error) {
	// New redis client
	client, err := redisex.NewClient()
	defer client.Close()
	if err != nil {
		return "", err
	}

	// New access token
	accessToken := NewAccessToken()

	// Access key as redis key
	if err := client.Set(accessToken, userID, AccessTokenTTL); err != nil {
		logex.Log.Error(err)
		return "", err
	}

	// User id and access key as redis key
	key := fmt.Sprintf(userAccessTokenKeyFormat, userID, accessToken)
	if err := client.Set(key, userID, AccessTokenTTL); err != nil {
		logex.Log.Error(err)
		return "", err
	}

	return accessToken, nil
}

// GetUserIDByAccessToken get user id by access token
func GetUserIDByAccessToken(accessToken string) (uint64, error) {
	// New redis client
	client, err := redisex.NewClient()
	defer client.Close()
	if err != nil {
		logex.Log.Error(err)
		return 0, err
	}

	// Get user id
	userID, err := redis.Uint64(client.Get(accessToken))
	if err != nil {
		if redisex.IsErrNil(err) {
			err = errors.New("access token is invalid")
		}
		logex.Log.Error(err)
		return 0, err
	}

	return userID, nil
}

// ResetAccessTokenTTL reset access token ttl
func ResetAccessTokenTTL(accessToken string, userID uint64) error {
	// New redis client
	client, err := redisex.NewClient()
	defer client.Close()
	if err != nil {
		return err
	}

	// Reset the key ttl of access key
	if err := client.Expire(accessToken, AccessTokenTTL); err != nil {
		logex.Log.Error(err)
		return err
	}

	// Reset the key ttl of user id and access key
	key := fmt.Sprintf(userAccessTokenKeyFormat, userID, accessToken)
	if err := client.Expire(key, AccessTokenTTL); err != nil {
		logex.Log.Error(err)
		return err
	}

	return nil
}

// RevokeAccessToken revoke access token
func RevokeAccessToken(accessToken string, userID uint64) error {
	// New redis client
	client, err := redisex.NewClient()
	defer client.Close()
	if err != nil {
		return err
	}

	// New keys
	keys := []string{
		accessToken,
		fmt.Sprintf(userAccessTokenKeyFormat, userID, accessToken),
	}

	// Delete keys
	if err := client.Del(keys...); err != nil {
		logex.Log.Error(err)
		return err
	}

	return nil
}

// RevokeAccessTokensByUserID revoke access tokens by user id
func RevokeAccessTokensByUserID(userID uint64) error {
	// New redis client
	client, err := redisex.NewClient()
	defer client.Close()
	if err != nil {
		return err
	}

	// Init keys
	keys := make([]string, 0)
	// New match
	match := fmt.Sprintf(userAccessTokensKeyFormat, userID)
	// New default cursor
	var cursor uint64
	// Start scan
	for {
		// Scan by match
		values, err := redis.Values(client.Scan(cursor, match, 0))
		if err != nil {
			logex.Log.Error(err)
			return err
		}

		// Get scanned keys
		scannedKeys, err := redis.Strings(values[1], nil)
		if err != nil {
			logex.Log.Error(err)
			return err
		}

		// Set keys
		for _, scannedKey := range scannedKeys {
			// Set scannedKey
			keys = append(keys, scannedKey)
			// Split scanned key with ':'
			splits := strings.Split(scannedKey, ":")
			// Append access key to keys
			keys = append(keys, splits[1])
		}

		// Get cursor
		cursor, err = redis.Uint64(values[0], nil)
		if err != nil {
			logex.Log.Error(err)
			return err
		}
		// if cursor is 0, break for loop
		if cursor == 0 {
			break
		}
	}

	// Delete keys
	if len(keys) > 0 {
		if err := client.Del(keys...); err != nil {
			logex.Log.Error(err)
			return err
		}
	}

	return nil
}
