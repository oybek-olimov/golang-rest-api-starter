package jwt

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	SecretKey     string
	TokenDuration time.Duration
}

func NewManager(secretKey string, duration time.Duration) *Manager {
	return &Manager{
		SecretKey:     secretKey,
		TokenDuration: duration,
	}
}

// Token yaratish
func (m *Manager) Generate(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(m.TokenDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.SecretKey))
}

// Tokenni tekshirish
func (m *Manager) Verify(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		// Faqat HMAC (HS256) tokenlarni qabul qilamiz
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(m.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, jwt.ErrInvalidKey
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, jwt.ErrInvalidKey
	}

	return int(userIDFloat), nil
}

//
// --- Middleware qismi ---
//

type contextKey string

const userIDKey contextKey = "user_id"

// Middleware — JWT tokenni tekshiradi va user_id ni kontekstga qo‘shadi
func (m *Manager) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		// "Bearer <token>" formatini ajratamiz
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}

		userID, err := m.Verify(parts[1])
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// token'dan userID ni kontekstga saqlaymiz
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Context'dan user_id ni chiqarib olish
func GetUserIDFromContext(ctx context.Context) (int, error) {
	userID, ok := ctx.Value(userIDKey).(int)
	if !ok {
		return 0, errors.New("user_id not found in context")
	}
	return userID, nil
}
