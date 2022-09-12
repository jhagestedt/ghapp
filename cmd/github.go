package cmd

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

var client = resty.New()
var privateKeyPrefix = "-----BEGIN RSA PRIVATE KEY-----"
var privateKeySuffix = "-----END RSA PRIVATE KEY-----"

type GitHubError struct {
	Message string `json:"message"`
}

type GitHubInstallationToken struct {
	Token string `json:"token"`
}

func CreateToken(id string, privateKey string) (string, error) {
	if !strings.HasPrefix(privateKey, privateKeyPrefix) {
		return "", fmt.Errorf("private-key should have prefix %s", privateKeyPrefix)
	}
	if !strings.HasSuffix(strings.TrimSuffix(privateKey, "\n"), privateKeySuffix) {
		return "", fmt.Errorf("private-key should have suffix %s", privateKeySuffix)
	}
	dec, _ := pem.Decode([]byte(privateKey))
	rsa, err := x509.ParsePKCS1PrivateKey(dec.Bytes)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": id,
		"iat": time.Now().Add(-60 * time.Second).Unix(),
		"exp": time.Now().Add(10 * 60 * time.Second).Unix(),
	})
	return token.SignedString(rsa)
}

func CreateInstallationToken(installId string, token string) (string, error) {
	gitHubInstallationToken := &GitHubInstallationToken{}
	gitHubError := &GitHubError{}
	res, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/vnd.github+json").
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		SetResult(&gitHubInstallationToken).
		SetError(&gitHubError).
		SetPathParams(map[string]string{
			"installId": installId,
		}).
		Post("https://api.github.com/app/installations/{installId}/access_tokens")
	if err != nil {
		return "", err
	}
	if res.IsError() {
		return "", fmt.Errorf("GitHub API error: %s - %s", res.Status(), gitHubError.Message)
	}
	return gitHubInstallationToken.Token, nil
}
