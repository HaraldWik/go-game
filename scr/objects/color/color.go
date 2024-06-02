package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	err_ "github.com/HaraldWik/go-game/scr/err"
)

type RGB struct {
	R, G, B float32
}

func NewRGB(red float32, green float32, blue float32) RGB {
	return RGB{
		R: red,
		G: green,
		B: blue,
	}
}

func All(number float32) RGB {
	return RGB{number, number, number}
}

func White() RGB {
	return RGB{1.0, 1.0, 1.0}
}

func Black() RGB {
	return RGB{0.0, 0.0, 0.0}
}

func Red() RGB {
	return RGB{1.0, 0.0, 0.0}
}

func Green() RGB {
	return RGB{0.0, 1.0, 0.0}
}

func Blue() RGB {
	return RGB{0.0, 0.0, 1.0}
}

func Yellow() RGB {
	return RGB{1.0, 1.0, 0.0}
}

func RGBToHex(r, g, b float32) string {
	// *Ensure values are within the expected range [0, 1]
	r = float32(math.Max(0, math.Min(1, float64(r))))
	g = float32(math.Max(0, math.Min(1, float64(g))))
	b = float32(math.Max(0, math.Min(1, float64(b))))

	// *Scale the float values to [0, 255]
	rInt := int(math.Round(float64(r * 255)))
	gInt := int(math.Round(float64(g * 255)))
	bInt := int(math.Round(float64(b * 255)))

	// *Format the integers as a hex string
	return fmt.Sprintf("#%02X%02X%02X", rInt, gInt, bInt)
}

func HexToRGB(hex string) RGB {
	// *Remove the '#' character if it exists
	hex = strings.TrimPrefix(hex, "#")

	// *Ensure the hex string has the correct length
	if len(hex) != 6 {
		fmt.Println(err_.WARNING+"Invalid hex:", hex)
		return All(0)
	}

	// *Parse the hex string to integer values
	rInt, err := strconv.ParseInt(hex[0:2], 16, 64)
	if err != nil {
		fmt.Println(err_.WARNING, err)
		return All(0)
	}

	gInt, err := strconv.ParseInt(hex[2:4], 16, 64)
	if err != nil {
		fmt.Println(err_.WARNING, err)
		return All(0)
	}
	bInt, err := strconv.ParseInt(hex[4:6], 16, 64)
	if err != nil {
		fmt.Println(err_.WARNING, err)
		return All(0)
	}

	// *Convert the integer values to floats in the range [0, 1]
	r := float32(rInt) / 255.0
	g := float32(gInt) / 255.0
	b := float32(bInt) / 255.0

	return NewRGB(r, g, b)
}
