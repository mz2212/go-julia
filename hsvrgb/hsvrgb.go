package hsvrgb

type RgbColor struct {
	R, G, B int
}
type HsvColor struct {
	H, S, V int
}

func Hsv2Rgb(hsv HsvColor) RgbColor {
	var rgb RgbColor
	var region, remainder, p, q, t int

	if hsv.S == 0 {
		rgb.R = hsv.V
		rgb.G = hsv.V
		rgb.B = hsv.V
		return rgb
	}

	region = hsv.H / 43
	remainder = (hsv.H - (region * 43)) * 6

	p = (hsv.V * (255 - hsv.S)) >> 8
	q = (hsv.V * (255 - ((hsv.S * remainder) >> 8))) >> 8
	t = (hsv.V * (255 - ((hsv.S * (255 - remainder)) >> 8))) >> 8

	switch region {
	case 0:
		rgb.R, rgb.G, rgb.B = hsv.V, t, p
		break
	case 1:
		rgb.R, rgb.G, rgb.B = q, hsv.V, p
		break
	case 2:
		rgb.R, rgb.G, rgb.B = p, hsv.V, t
		break
	case 3:
		rgb.R, rgb.G, rgb.B = p, q, hsv.V
		break
	case 4:
		rgb.R, rgb.G, rgb.B = t, p, hsv.V
		break
	default:
		rgb.R, rgb.G, rgb.B = hsv.V, p, q
		break
	}

	return rgb
}
