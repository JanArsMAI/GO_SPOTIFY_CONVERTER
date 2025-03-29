package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"spotifyparser/internal/application"
	"spotifyparser/internal/domain/track/entity"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(".env loading error")
	}
	client := &http.Client{
		Timeout: time.Duration(time.Second * 5),
	}
	apiId := os.Getenv("CLIENT_ID")
	apiKey := os.Getenv("CLIENT_SECRET")
	baseURL := "https://api.spotify.com"
	parser := application.CreateNewParser(
		client,
		apiKey,
		baseURL,
		time.Duration(time.Second*5),
	)
	token, err := parser.GetAuthorizationToken(context.Background(), apiId, apiKey)
	if err != nil {
		log.Fatalf("Ошибка при получении токена: %v", err)
	}
	//пример id
	trackId := "4w2GLmK2wnioVnb5CPQeex"
	track, err := parser.GetTrackById(context.Background(), trackId, token.Token)
	if err != nil {
		log.Fatalf("Ошибка при получении трека: %v", err)
	}
	printTrackInfo(track)

}

func printTrackInfo(track *entity.Track) {
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
}
