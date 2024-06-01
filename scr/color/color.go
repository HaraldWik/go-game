package Color

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

func (RGB) White() RGB {
	return RGB{
		1.0,
		1.0,
		1.0,
	}
}

func (RGB) All(number float32) RGB {
	return RGB{
		number,
		number,
		number,
	}
}
