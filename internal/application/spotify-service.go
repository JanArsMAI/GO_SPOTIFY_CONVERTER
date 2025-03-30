package application

import (
	"context"
	"errors"
	"fmt"
	parcer "spotifyparser/internal/domain/parser"
	TokenEntity "spotifyparser/internal/domain/token"
	"spotifyparser/internal/domain/track/entity"
	"sync"
	"time"
)

type SpotifyService struct {
	Parser     *parcer.Parser
	TrackList  []entity.Track
	token      TokenEntity.Token
	lastUpdate time.Time
	apiKey     string
	apiId      string
	ctx        context.Context
	mutex      sync.RWMutex
}

func CreateNewSpotifyService(ctx context.Context, parser *parcer.Parser, apikey string, apiid string) (*SpotifyService, error) {

	token, err := parser.GetAuthorizationToken(ctx, apiid, apikey)
	if err != nil {
		return nil, errors.New("error. can not get token and create service")
	}
	return &SpotifyService{
		Parser:     parser,
		token:      *token,
		lastUpdate: time.Now(),
		apiKey:     apikey,
		apiId:      apiid,
		ctx:        ctx,
	}, nil
}

func (s *SpotifyService) UpdateToken() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	token, err := s.Parser.GetAuthorizationToken(s.ctx, s.apiId, s.apiKey)
	if err != nil {
		return err
	}
	s.token = *token
	s.lastUpdate = time.Now()
	return nil
}

func (s *SpotifyService) CheckTimeDurationToken() bool {
	return time.Since(s.lastUpdate) < time.Hour
}

func (s *SpotifyService) AddTrackToList(idTrack string) {
	if !s.CheckTimeDurationToken() {
		s.mutex.Lock()
		if !s.CheckTimeDurationToken() {
			s.UpdateToken()
		}
		s.mutex.Unlock()
	}
	track, err := s.Parser.GetTrackById(s.ctx, idTrack, s.token.Token)
	if err != nil {
		return
	}
	s.mutex.Lock()
	s.TrackList = append(s.TrackList, *track)
	s.mutex.Unlock()
}

func (s *SpotifyService) ClearAllTracks() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.TrackList = s.TrackList[:0]
}

func (s *SpotifyService) PrintInfoOfList() {
	c := 0
	fmt.Println("-------------------------------------------------")
	for _, track := range s.TrackList {
		c++
		fmt.Printf("Track Name: %s\n", track.Name)
		fmt.Printf("Artist(s): ")
		for i, artist := range track.Artists {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(artist.Name)
		}
		fmt.Printf("\nAlbum: %s\n", track.Album.Name)
		fmt.Printf("Release Date: %s\n", track.Album.ReleaseDate)
		fmt.Printf("Popularity: %d\n", track.Popularity)
		if track.PreviewURL != "" {
			fmt.Printf("Picture: %s\n", track.PreviewURL)
		}
		if track.Album.Images[1].URL != "" {
			fmt.Println(track.Album.Images[1].URL)
		}
		fmt.Println("-------------------------------------------------")
	}
}
