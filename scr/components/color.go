package component

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

func (color RGB) All(number float32) RGB {
	color = NewRGB(number, number, number)

	return color
}

func (color RGB) White() RGB {
	color = NewRGB(1.0, 1.0, 1.0)

	return color
}

func (color RGB) Black() RGB {
	color = NewRGB(0.0, 0.0, 0.0)

	return color
}

func (color RGB) Red() RGB {
	color = NewRGB(1.0, 0.0, 0.0)

	return color
}

func (color RGB) Green() RGB {
	color = NewRGB(0.0, 1.0, 0.0)

	return color
}

func (color RGB) Blue() RGB {
	color = NewRGB(0.0, 0.0, 1.0)

	return color
}

func (color RGB) Yellow() RGB {
	color = NewRGB(1.0, 1.0, 0.0)

	return color
}

func (color RGB) ToHex() string {
	// *Ensure values are within the expected range [0, 1]
	color.R = float32(math.Max(0, math.Min(1, float64(color.R))))
	color.G = float32(math.Max(0, math.Min(1, float64(color.G))))
	color.B = float32(math.Max(0, math.Min(1, float64(color.B))))

	// *Scale the float values to [0, 255]
	rInt := int(math.Round(float64(color.R * 255)))
	gInt := int(math.Round(float64(color.G * 255)))
	bInt := int(math.Round(float64(color.B * 255)))

	// *Format the integers as a hex string
	return fmt.Sprintf("#%02X%02X%02X", rInt, gInt, bInt)
}

func (color RGB) HexToRGB(hex string) RGB {
	// *Remove the '#' character if it exists
	hex = strings.TrimPrefix(hex, "#")

	// *Ensure the hex string has the correct length
	if len(hex) != 6 {
		fmt.Println(err_.WARNING+"Invalid hex:", hex)
		return color.All(0.0)
	}

	// *Parse the hex string to integer values
	rInt, err := strconv.ParseInt(hex[0:2], 16, 64)
	if err != nil {
		fmt.Println(err_.WARNING, err)
		return color.All(0.0)
	}

	gInt, err := strconv.ParseInt(hex[2:4], 16, 64)
	if err != nil {
		fmt.Println(err_.WARNING, err)
		return color.All(0.0)
	}
	bInt, err := strconv.ParseInt(hex[4:6], 16, 64)
	if err != nil {
		fmt.Println(err_.WARNING, err)
		return color.All(0.0)
	}

	// *Convert the integer values to floats in the range [0, 1]
	r := float32(rInt) / 255.0
	g := float32(gInt) / 255.0
	b := float32(bInt) / 255.0

	return NewRGB(r, g, b)
}
