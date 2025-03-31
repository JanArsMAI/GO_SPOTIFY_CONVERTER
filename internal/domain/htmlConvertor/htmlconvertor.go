package htmlconvertor

import (
	"fmt"
	"os"
	entity "spotifyparser/internal/domain/track/entity"
	"strings"
)

func formatDuration(durationMs int) string {
	totalSeconds := durationMs / 1000
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
func getArtists(track *entity.Track) string {
	ans := ""
	for _, artist := range track.Artists {
		ans += artist.Name + ", "
	}
	return ans[:len(ans)-2]
}

func formatData(data string) string {
	return strings.Split(data, "-")[0]
}

func ConvertTracksToHTML(filename string, tracks []entity.Track) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer file.Close()
	_, err = file.WriteString(`<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <link href="https://fonts.googleapis.com/css?family=Montserrat&display=swap" rel="stylesheet">
    <title>Название</title>
    <style>
        body {
            padding: 0;
            font-family: Arial , sans-serif;
            background: #212121;
            margin: 0;
            display: flex;
            flex-direction: column;
            align-items: center;
        }
        h1 {
            color: #d9faf7;
            text-shadow: 3px 3px 6px rgba(0, 0, 0, 0.9);
            text-align: center;
            width: 100%;
            margin-bottom: 15px;
            margin-top: 15px;
            font-family: 'Montserrat', sans-serif;
        }
        .track {
            width: 70%;
            display: flex;
            align-items: center;
            background-color: #323232;
            margin-bottom: 10px;
            padding: 10px;
            border-radius: 5px;
            text-shadow: 0 0 5px #0d7377;
            box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
        }
        .track-cover {
            width: 45px;
            height: 45px;
            margin-right: 10px;
            box-shadow: 3px 3px 6px rgba(0, 0, 0, 0.9);
        }
        .track-info {
            display: flex;
            flex-direction: column;
            justify-content: center;
            flex-basis: 40%;
        }
        .track-album{
            width: 40%;
            color:#d9faf7;
        }
        .track-title {
            font-size: 18px;
            font-weight: bold;
            color:#d9faf7;
        }
        .track-artist {
            font-size: 14px;
            color: #d9faf7;
        }
        .track-duration {
            font-size: 14px;
            color: #d9faf7;
            margin-left: 0;
        }
        .track-year{
            width: 10%;
            color: #d9faf7;
        }
    </style>
</head>
<body>
    <h1>Название</h1>
`)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл: %v", err)
	}
	for _, track := range tracks {
		trackInfo := fmt.Sprintf(`
    <div class="track">
        <img class="track-cover" src="%s">
        <div class="track-info">
            <div class="track-title">%s</div>
            <div class="track-artist">%s</div>
        </div>
        <div class="track-album">%s</div>
        <div class="track-year">%s</div>
        <div class="track-duration">%s</div>
    </div>`,
			track.Album.Images[1].URL, track.Name, getArtists(&track), track.Album.Name, formatData(track.Album.ReleaseDate), formatDuration(track.DurationMs))
		_, err = file.WriteString(trackInfo)
		if err != nil {
			return fmt.Errorf("ошибка при записи в файл: %v", err)
		}
	}
	_, err = file.WriteString(`
</body>
</html>`)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл: %v", err)
	}
	return nil
}
