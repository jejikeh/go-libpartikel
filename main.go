package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jejikeh/go-libpartikel/partikel"
)

const WindowWidth = 800
const WindowHeight = 600

var textureCircle rl.Texture2D

func main() {
	rl.InitWindow(WindowWidth, WindowHeight, "go-libpartikel demo")
	defer rl.CloseWindow()

	imageCircle := rl.GenImageGradientRadial(16, 16, 0.3, rl.White, rl.Black)
	textureCircle = rl.LoadTextureFromImage(imageCircle)
	defer rl.UnloadImage(imageCircle)
	defer rl.UnloadTexture(textureCircle)

	rl.SetTargetFPS(60)

	ps := initFlameParticleSystem()
	ps.Start()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		mousePos := rl.GetMousePosition()

		ps.SetOrigin(mousePos)
		ps.Update()
		ps.Draw()

		rl.EndDrawing()
	}
}

func initFlameParticleSystem() *partikel.ParticleSystem {
	ps := &partikel.ParticleSystem{}

	configFlame1 := partikel.EmitterConfig{
		StartSize:    rl.NewVector2(2, 2),
		EndSize:      rl.NewVector2(1, 1),
		Capacity:     100,
		EmmisionRate: 500,
		Origin: rl.Vector2{
			X: 0,
			Y: 0,
		},
		OriginAcceleration: [2]float32{
			50,
			100,
		},
		Offset: [2]float32{
			0,
			10,
		},
		Direction: rl.NewVector2(0, -1),
		DirectionAngle: [2]float32{
			-90,
			-90,
		},
		Velocity: [2]float32{
			30,
			150,
		},
		VelocityAngle: [2]float32{
			90,
			90,
		},
		StartColor: rl.NewColor(255, 20, 0, 255),
		EndColor:   rl.NewColor(255, 20, 0, 0),
		Age: [2]float32{
			1.0,
			2.0,
		},
		Texture:   textureCircle,
		BlendMode: rl.BlendAdditive,
	}

	emitterFlame1 := partikel.NewEmitter(configFlame1)

	configFlame2 := configFlame1
	configFlame2.StartSize = rl.NewVector2(1, 1)
	configFlame2.EndSize = rl.NewVector2(0, 0)
	configFlame2.Capacity = 20
	configFlame2.EmmisionRate = 20
	configFlame2.StartColor = rl.NewColor(255, 255, 255, 10)
	configFlame2.EndColor = rl.NewColor(255, 255, 255, 0)
	configFlame2.Age = [2]float32{
		0.5,
		1.0,
	}

	emitterFlame2 := partikel.NewEmitter(configFlame2)

	configSmokeEmitter := configFlame2
	configSmokeEmitter.StartSize = rl.NewVector2(2, 2)
	configSmokeEmitter.EndSize = rl.NewVector2(1, 1)
	configSmokeEmitter.Capacity = 500
	configSmokeEmitter.EmmisionRate = 100
	configSmokeEmitter.DirectionAngle = [2]float32{
		-6,
		6,
	}
	configSmokeEmitter.VelocityAngle = [2]float32{
		0,
		0,
	}
	configSmokeEmitter.OriginAcceleration = [2]float32{
		4,
		4,
	}
	configSmokeEmitter.StartColor = rl.NewColor(125, 125, 125, 30)
	configSmokeEmitter.EndColor = rl.NewColor(125, 125, 125, 10)
	configSmokeEmitter.Age = [2]float32{
		3.0,
		5.0,
	}

	smokeEmitter := partikel.NewEmitter(configSmokeEmitter)

	ps.Add(emitterFlame1)
	ps.Add(smokeEmitter)
	ps.Add(emitterFlame2)

	return ps
}
