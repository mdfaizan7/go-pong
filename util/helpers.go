package util

// WinWidth is the window width
var WinWidth int = 800

// WinHeight is the window height
var WinHeight int = 600

// ClearPixels clears the previous occupied pixels
func ClearPixels(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}
