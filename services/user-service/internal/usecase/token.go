package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/yodzafar/food-marketpalce/user-service/internal/domain"
)

type claims struct {
	UserID string      `json:"userId"`
	Role   domain.Role `json:"role"`
	Exp    int64       `json:"exp"`
}

func generateToken(user *domain.User) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	c := claims{
		UserID: user.ID,
		Role:   user.Role,
		Exp:    time.Now().Add(time.Hour * 24).Unix(),
	}

	payload, err := json.Marshal(c)

	if err != nil {
		return "", err
	}

	header := base64.RawStdEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	body := base64.RawURLEncoding.EncodeToString(payload)

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(header + "." + body))

	return fmt.Sprintf("%s.%s.%x", header, body, mac.Sum(nil)), nil
}
