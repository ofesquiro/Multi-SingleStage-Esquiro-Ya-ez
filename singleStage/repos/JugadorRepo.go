package repos

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"app/models"
)

var (
	file_path = "/repos/database.json"
)

func GetExecutablePath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(exePath), nil
}

// LoadDatabase loads the database.json file and unmarshals its content.
func LoadDatabase() (*models.Database, error) {
	exeDir, err := GetExecutablePath()
	if err != nil {
		return nil, err
	}

	dbPath := filepath.Join(exeDir, file_path)
	data, err := os.ReadFile(dbPath)
	if err != nil {
		return nil, err
	}

	var database models.Database
	if err := json.Unmarshal(data, &database); err != nil {
		return nil, err
	}

	return &database, nil
}

func ReadJugadoresFromDB() ([]models.Jugador, error) {
	db, err := LoadDatabase()
	if err != nil {
		fmt.Println("Error loading database:", err)
		return nil, err
	}
	var jugadores []models.Jugador
	for _, jugador := range db.Jugadores {
		jugadores = append(jugadores, jugador)
	}
	return jugadores, nil
}
