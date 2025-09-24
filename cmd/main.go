package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"spotifyparser/internal/application"
	htmlconv "spotifyparser/internal/domain/htmlConvertor"
	parcer "spotifyparser/internal/domain/parser"
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
	parser := parcer.CreateNewParser(
		client,
		apiKey,
		baseURL,
		time.Duration(time.Second*5),
	)
	SpotifyService, err := application.CreateNewSpotifyService(context.Background(), parser, apiKey, apiId)
	if err != nil {
		log.Fatalf("Ошибка при создании сервиса: %v", err)
	}
	var option string
	var choice string
	for option != "-1" {
		fmt.Println("options for working with service:")
		fmt.Println("1 - Добавить трек в очередь")
		fmt.Println("2 - Удалить очередь")
		fmt.Println("3 - Показать очередь")
		fmt.Println("4 - Конвертировать очередь в SVG")
		fmt.Println("-1 - Выход")
		fmt.Scanln(&option)
		switch option {
		case "1":
			fmt.Println("Введите Id трека")
			fmt.Scanln(&choice)
			SpotifyService.AddTrackToList(choice)
		case "2":
			SpotifyService.ClearAllTracks()
			fmt.Println("Очередь удалена")
		case "3":
			SpotifyService.PrintInfoOfList()
		case "4":
			fmt.Println("введите имя файла, в который делается запись")
			fmt.Scanln(&choice)
			htmlconv.ConvertTracksToHTML(choice, SpotifyService.TrackList)
		}
	}
}
