package game

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) LoadResources() {
	g.LoadModels()
	g.LoadTextures()
	g.LoadSounds()
}

func (g *Game) LoadSounds() {
	g.sounds["chess_slide"] = g.LoadSoundEmbedded("assets/sounds/chess_slide.wav")
	g.sounds["falling_impact"] = g.LoadSoundEmbedded("assets/sounds/falling_impact.wav")
}

func (g *Game) LoadModels() {
	g.playerCastle.model = g.LoadModelEmbedded("assets/models/castle/tower.glb")
	g.plainTileModel = g.LoadModelEmbedded("assets/models/tiles/basic_ground_tile.glb")
	g.waterTileModel = g.LoadModelEmbedded("assets/models/tiles/basic_ground_tile.glb")
	pawnModel := g.LoadModelEmbedded("assets/models/enemies/pawn.glb")
	knightModel := g.LoadModelEmbedded("assets/models/enemies/knight.glb")
	bishopModel := g.LoadModelEmbedded("assets/models/enemies/bishop.glb")
	kingModel := g.LoadModelEmbedded("assets/models/enemies/king.glb")
	queenModel := g.LoadModelEmbedded("assets/models/enemies/queen.glb")
	attackPawnCardTexture := g.LoadTextureEmbedded("assets/textures/cards/attack_pawn.png")
	attackKnightCardTexture := g.LoadTextureEmbedded("assets/textures/cards/attack_knight.png")
	attackBishopCardTexture := g.LoadTextureEmbedded("assets/textures/cards/attack_bishop.png")
	attackQueenCardTexture := g.LoadTextureEmbedded("assets/textures/cards/attack_queen.png")
	attackKingCardTexture := g.LoadTextureEmbedded("assets/textures/cards/attack_king.png")
	backCardTexture := g.LoadTextureEmbedded("assets/textures/cards/card_back.png")
	g.enemyModels[EnemyTypePawn] = &pawnModel
	g.enemyModels[EnemyTypeKnight] = &knightModel
	g.enemyModels[EnemyTypeBishop] = &bishopModel
	g.enemyModels[EnemyTypeQueen] = &queenModel
	g.enemyModels[EnemyTypeKing] = &kingModel
	g.cardTextures[CardTypeAttackPawn] = &attackPawnCardTexture
	g.cardTextures[CardTypeAttackKnight] = &attackKnightCardTexture
	g.cardTextures[CardTypeAttackBishop] = &attackBishopCardTexture
	g.cardTextures[CardTypeAttackQueen] = &attackQueenCardTexture
	g.cardTextures[CardTypeAttackKing] = &attackKingCardTexture
	g.cardTextures[CardTypeBack] = &backCardTexture

	g.tiles[TileTypeClear] = Tile{
		model:    &g.plainTileModel,
		tileType: TileTypeClear,
	}
	g.tiles[TileTypeWater] = Tile{
		model:    &g.waterTileModel,
		tileType: TileTypeWater,
	}

}

func (g *Game) LoadTextures() {

}

// LoadModelEmbedded writes the embedded data to a temp file, loads it, then cleans up.
func (g *Game) LoadModelEmbedded(filename string) rl.Model {
	// 1. Read bytes from embed
	fileData, err := g.assets.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading embedded model:", err)
		return rl.Model{}
	}

	ext := ""
	if len(filename) > 4 {
		ext = filename[len(filename)-4:] // Simple check for .png, .jpg
	}
	// 2. Create a temporary file
	// We use "model-*.glb" so the temp file keeps the .glb extension.
	// Raylib NEEDS the extension to know which importer to use.
	tempFile, err := os.CreateTemp("", "raylib_model-*"+ext)
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return rl.Model{}
	}

	// 3. Write data to disk
	if _, err := tempFile.Write(fileData); err != nil {
		fmt.Println("Error writing temp file:", err)
		return rl.Model{}
	}

	// Close the file handle so Raylib can open it
	tempPath := tempFile.Name()
	tempFile.Close()

	// 4. Load the Model using the standard function
	model := rl.LoadModel(tempPath)

	// 5. Clean up! Delete the file immediately.
	// We defer this or call it now. Raylib has loaded the data into VRAM/RAM,
	// so we don't need the file on disk anymore.
	os.Remove(tempPath)

	return model
}

func (g *Game) LoadTextureEmbedded(filename string) rl.Texture2D {
	// 1. Read the file bytes from the embedded filesystem
	fileData, err := g.assets.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading embedded file:", err)
		return rl.Texture2D{}
	}

	// 2. Determine file extension (e.g., ".png")
	// Raylib needs this hint to know how to decode the bytes
	ext := ""
	if len(filename) > 4 {
		ext = filename[len(filename)-4:] // Simple check for .png, .jpg
	}

	// 3. Load Image from RAM
	img := rl.LoadImageFromMemory(ext, fileData, int32(len(fileData)))

	// 4. Upload to GPU (Texture)
	texture := rl.LoadTextureFromImage(img)

	// 5. Unload the CPU copy (Image)
	rl.UnloadImage(img)

	return texture
}

func (g *Game) LoadSoundEmbedded(filename string) rl.Sound {
	fileData, err := g.assets.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading embedded sound:", err)
		return rl.Sound{}
	}

	ext := ""
	if len(filename) > 4 {
		ext = filename[len(filename)-4:]
	}

	tempFile, err := os.CreateTemp("", "raylib_sound-*"+ext)
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return rl.Sound{}
	}

	if _, err := tempFile.Write(fileData); err != nil {
		fmt.Println("Error writing temp file:", err)
		return rl.Sound{}
	}

	tempPath := tempFile.Name()
	tempFile.Close()

	sound := rl.LoadSound(tempPath)

	os.Remove(tempPath)

	return sound
}
