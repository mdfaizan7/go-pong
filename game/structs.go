package game

// Color is a Pixel struct
type Color struct {
	R byte
	G byte
	B byte
}

// Pos describes the position
type Pos struct {
	X float32
	Y float32
}

// Ball represent the pong ball
type Ball struct {
	Pos
	Radius float32 // radius of the ball
	XV     float32 // X velocity
	YV     float32 // Y velocity
	Color  Color
}

// Paddle represents the pong paddle
type Paddle struct {
	Pos
	W     float32 // Width
	H     float32 // Height
	Color Color
	Score int
}
