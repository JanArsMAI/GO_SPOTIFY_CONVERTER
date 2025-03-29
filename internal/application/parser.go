package application

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	TokenEntity "spotifyparser/internal/domain/token"
	TrackEntity "spotifyparser/internal/domain/track/entity"
	"strings"
	"time"
)

type Parser struct {
	client  *http.Client
	apiKey  string
	baseURL string
	timeout time.Duration
}

func CreateNewParser(cl *http.Client, apiKey, baseURL string, timeout time.Duration) *Parser {
	return &Parser{
		client:  cl,
		apiKey:  apiKey,
		baseURL: baseURL,
		timeout: timeout,
	}
}
func (p *Parser) GetAuthorizationToken(ctx context.Context, id, key string) (*TokenEntity.Token, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*5))
	defer cancel()
	Req := fmt.Sprintf("%s:%s", id, key)
	finalReq := base64.StdEncoding.EncodeToString([]byte(Req))
	req, err := http.NewRequestWithContext(ctx, "POST", "https://accounts.spotify.com/api/token", strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Basic "+finalReq)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error, status of request: %d", resp.StatusCode)
	}
	var token TokenEntity.Token
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}
func (p *Parser) GetTrackById(ctx context.Context, id, token string) (*TrackEntity.Track, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*5))
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/v1/tracks/%s", p.baseURL, id), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error, status of request: %d", resp.StatusCode)
	}
	var track TrackEntity.Track
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&track); err != nil {
		return nil, err
	}
	return &track, nil
}
