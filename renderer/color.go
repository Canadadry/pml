package renderer

import (
    "fmt"
    "image/color"
)

func parseHexColor(s string) (color.RGBA, error) {

    c := color.RGBA{
        A: 0xff,
    }

    var err error = nil

    switch len(s) {
    case 6:
        _, err = fmt.Sscanf(s, "%02x%02x%02x", &c.R, &c.G, &c.B)
    default:
        err = fmt.Errorf("invalid length, must be 6")

    }
    return c, err
}
