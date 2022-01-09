package main

import (
    "os"
    "log"
    "math"
    "image"
    "image/color"
    "golang.org/x/image/bmp"
)

func getPixCol(x, y, imgPattern uint8) color.RGBA {
    var r, g, b uint8

    switch imgPattern {
    case 0: // CIRCLE PATTERN
        var distX, distY = float64(x) - 127.5, float64(y) - 127.5
        var dist = math.Sqrt(distX*distX + distY*distY)
        r, g, b = 0, uint8(math.Abs(dist)) * 4, uint8(math.Abs(dist)) * 8
    case 1: // RECURSIVE CHECKER PATTERN
        r, g, b = 20, 20, x ^ y
    case 2: // RECURSIVE CHECKER PATTERN WITH LINES AND GRADIENT
        r, g, b = 0b00001111 ^ y, x ^ y, x ^ 0b11110000
    case 3: // DIAGONAL LINES
        r, g, b = 20, 20, (x+y) * 5
    case 4: // DIAGONAL CHECKER PATTERN
        r, g, b = 20, (x+255-y) * 5, (x+y) * 5
    default:
        log.Fatal("unsupported image pattern: ", imgPattern)
    }

    return color.RGBA{r, g, b, 255}
}

func paintBmp(imgPattern uint8, imgName string) {
    const WIDTH, HEIGHT = 255, 255

    // Init the image with a red green blue alpha color mode
    img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))

    // Loop through all pixels of the image
    for x := 0; x < WIDTH ; x++ {
        for y := 0; y < HEIGHT; y++ {
            // Set the current pixel to a color from getPixCol()
            col := getPixCol(uint8(x), uint8(y), imgPattern)
            img.SetRGBA(x, y, col)
        }
    }

    // Save the image
    file, err := os.Create(imgName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    bmp.Encode(file, img)
}

func main() {
    paintBmp(0, "rings.bmp")
    paintBmp(1, "checker.bmp")
    paintBmp(2, "checker-gradient.bmp")
    paintBmp(3, "diagonals.bmp")
    paintBmp(4, "diagonals-checker.bmp")
}
