package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	baseTokenUrl = "http://169.254.169.254/metadata/identity/oauth2/token"
)

type KeyVaultClient struct {
	name       string
	httpClient http.Client
	token      string
}

func NewKeyVaultClient(name string) *KeyVaultClient {
	return &KeyVaultClient{
		name: name,
		httpClient: http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (kvc *KeyVaultClient) Authorize(identity string) error {
	token, err := kvc.retrieveAccessToken(identity)
	if err != nil {
		return err
	}
	kvc.token = token
	return nil
}

func (kvc *KeyVaultClient) GetSecret(secretName string) (string, error) {

	url, err := buildSecretUrl(kvc.name, secretName)
	if err != nil {
		return "", fmt.Errorf("failed to build URL: %w", err)
	}

	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", kvc.token))

	res, err := kvc.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("http request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("key vault server responded with status: %s", res.Status)
	}

	var secret struct {
		Id    string `json:"id"`
		Value string `json:"value"`
	}

	if err := json.NewDecoder(res.Body).Decode(&secret); err != nil {
		return "", fmt.Errorf("failed to decode the response body: %w", err)
	}

	return secret.Value, nil
}

func (kvc *KeyVaultClient) retrieveAccessToken(identity string) (string, error) {

	url, err := buildTokenUrl(identity)
	if err != nil {
		return "", fmt.Errorf("failed to build URL: %w", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Metadata", "true")

	res, err := kvc.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("http request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("authorization server responded with status: %s", res.Status)
	}

	return parseToken(res)
}

func parseToken(res *http.Response) (string, error) {

	var tokens struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    string `json:"expires_in"`
		ExpiresOn    string `json:"expires_on"`
		NotBefore    string `json:"not_before"`
		Resource     string `json:"resource"`
		TokenType    string `json:"token_type"`
	}

	if err := json.NewDecoder(res.Body).Decode(&tokens); err != nil {
		return "", fmt.Errorf("failed to decode the response body: %w", err)
	}

	if len(tokens.AccessToken) == 0 {
		return "", fmt.Errorf("access token is empty")
	}

	return tokens.AccessToken, nil
}

func buildTokenUrl(identity string) (string, error) {

	var u, err = url.Parse(baseTokenUrl)
	if err != nil {
		return "", err
	}

	var q = u.Query()
	q.Set("api-version", "2018-02-01")
	q.Set("resource", "https://vault.azure.net")

	if len(identity) > 0 {
		q.Set("client_id", identity)
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}

func buildSecretUrl(keyVaultName string, secretName string) (string, error) {

	u, err := url.Parse(
		fmt.Sprintf("https://%s.vault.azure.net/secrets/%s", keyVaultName, secretName))
	if err != nil {
		return "", err
	}

	var q = u.Query()
	q.Set("api-version", "2016-10-01")
	u.RawQuery = q.Encode()

	return u.String(), nil
}
