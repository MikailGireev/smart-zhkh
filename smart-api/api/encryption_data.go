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

const (
	keycloakURL = "http://localhost:9080"
	realmName   = "smart-hcs"
	clientID    = "smart-hcs-client"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

type TokenClaims struct {
	Sub               string `json:"sub"`
	PreferredUsername string `json:"preferred_username"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	ID       string `json:"id"`
}

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

func ExtractClaimsFromToken(token string) (*TokenClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("невалидный токен")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("ошибка декодирования payload: %w", err)
	}

	var claims TokenClaims
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга claims: %w", err)
	}

	return &claims, nil
}
