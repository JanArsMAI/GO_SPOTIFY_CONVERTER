package htmlconvertor

import (
	"fmt"
	"os"
	entity "spotifyparser/internal/domain/track/entity"
)

func formatDuration(durationMs int) string {
	totalSeconds := durationMs / 1000
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
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
    <title>Обзор треков</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            margin: 20px;
        }
        .track {
            background-color: #fff;
            margin-bottom: 10px;
            padding: 10px;
            border-radius: 5px;
            box-shadow: 0 0 5px rgba(0,0,0,0.1);
        }
        .track-title {
            font-size: 18px;
            font-weight: bold;
        }
        .track-artist, .track-duration {
            font-size: 14px;
            color: #555;
        }
    </style>
</head>
<body>
    <h1>Обзор треков</h1>
`)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл: %v", err)
	}
	for _, track := range tracks {
		trackInfo := fmt.Sprintf(`
    <div class="track">
        <div class="track-title">Трек: %s</div>
        <div class="track-artist">Исполнитель: %s</div>
        <div class="track-duration">Длительность: %s</div>
    </div>`,
			track.Name, track.Artists[0], formatDuration(track.DurationMs))
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
