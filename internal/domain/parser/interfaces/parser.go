package interfaces

import (
	"context"
	TokenEntity "spotifyparser/internal/domain/token"
	TrackEntity "spotifyparser/internal/domain/track/entity"
)

type Parser interface {
	GetTrackById(ctx context.Context, id string) (*TrackEntity.Track, error)
	getAutorisationToken(ctx context.Context, id, key string) (*TokenEntity.Token, error)
}
