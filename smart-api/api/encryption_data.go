package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Константы Keycloak
const (
	keycloakURL = "http://localhost:9080"
	realmName   = "smart-hcs"
	clientID    = "smart-hcs-client" // client_id именно тот, который ты создавал в Keycloak
)

// Структуры
type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

type TokenClaims struct {
	PreferredUsername string `json:"preferred_username"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

// Получаем токен у Keycloak
func GetToken(name, pass string) (string, error) {
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("username", name)
	data.Set("password", pass)
	data.Set("grant_type", "password")

	resp, err := http.PostForm(fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", keycloakURL, realmName), data)
	if err != nil {
		return "", fmt.Errorf("ошибка при отправке запроса токена: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка при чтении ответа токена: %w", err)
	}

	var tokenResp TokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return "", fmt.Errorf("ошибка при парсинге ответа токена: %w", err)
	}

	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("токен не получен: %s", string(body))
	}

	return tokenResp.AccessToken, nil
}

// Декодируем payload из access_token
func ExtractUsernameFromToken(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("невалидный токен")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("ошибка декодирования payload: %w", err)
	}

	var claims TokenClaims
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга claims: %w", err)
	}

	return claims.PreferredUsername, nil
}

// Основной хендлер логина

