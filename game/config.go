package game

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameConfig struct {
	lastModTime time.Time
	configPath  string
	GameName    string
	Window      struct {
		Width           int32 `json:"width"`
		Height          int32 `json:"height"`
		BackgroundColor struct {
			R float32 `json:"r"`
			G float32 `json:"g"`
			B float32 `json:"b"`
			A float32 `json:"a"`
		} `json:"background_color"`
		TargetFPS int32 `json:"target_fps"`
	} `json:"window"`
	Camera struct {
		MoveSpeed float32 `json:"move_speed"`
	} `json:"camera"`
}

func (g *Game) ApplyConfig() {
	fmt.Println("Applying new config")
	rl.SetWindowSize(int(g.Config.Window.Width), int(g.Config.Window.Height))
	rl.SetTargetFPS(g.Config.Window.TargetFPS)
}

func (g *Game) CheckAndLoadConfig(isFirstLoad bool) {
	g.Config.configPath = "config/config.json"
	info, err := os.Stat(g.Config.configPath)
	if err != nil {
		fmt.Println("EEERRROORRR")
		return
	}
	if info.ModTime().After(g.Config.lastModTime) {
		fmt.Println("Reloading config...")

		// Read file
		data, err := os.ReadFile(g.Config.configPath)
		if err != nil {
			fmt.Printf("Error reading config: %v\n", err)
			return
		}

		// Decode
		var newConfig GameConfig
		err = json.Unmarshal(data, &newConfig)
		if err != nil {
			fmt.Printf("JSON syntax error: %v\n", err)
			return
		}

		// Update live config
		g.Config = newConfig
		g.Config.lastModTime = info.ModTime()
		fmt.Println("Config updated succesfully")
		if !isFirstLoad {
			g.ApplyConfig()
		}
	}
}
