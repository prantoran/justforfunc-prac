package main

import (
	"flag"
	"fmt"
	"image/color"
	"strconv"
)

type colorValue struct {
	color.Color // struct embedding
}

func (c *colorValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return fmt.Errorf("not a color: %v", err)
	}
	b := uint8(v % 256)
	g := uint8((v / 256) % 256)
	r := uint8((v / (256 * 256)) % 256)
	c.Color = color.RGBA{R: r, G: g, B: b, A: 255}
	return nil
}

func (c *colorValue) String() string {
	var r, g, b, a uint32
	if c != nil && c.Color != nil {
		r, g, b, a = c.RGBA()
		r, g, b, a = r/256, g/256, b/256, a/256
	}
	return fmt.Sprintf("rgba(%v, %v, %v, %v)", r, g, b, a)
}

func flagColor(name string, value color.Color, usage string) color.Color {
	v := &colorValue{value}
	flag.Var(v, name, usage)
	return v
}

func main() {
	// fg, bg := colorValue{color.White}, colorValue{} // color.White = {255,255,255}
	// flag.Var(&bg, "bg", "background color")
	// flag.Var(&fg, "fg", "foreground color")

	fg := flagColor("fg", color.White, "foreground color")
	bg := flagColor("bg", color.Black, "background color")

	flag.Parse()

	draw(fg, bg)
	// background := flag.String("bg", "000000", "background color")
	// foreground := flag.String("fg", "ffffff", "foreground color")
	// flag.Parse()

	// fg, err := parseColor(*foreground)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// bg, err := parseColor(*background)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("background is %s\n", *background)

	// draw(fg, bg)
}

// func parseColor(s string) (color.Color, error) {
// 	v, err := strconv.ParseInt(s, 16, 64)
// 	if err != nil {
// 		return nil, fmt.Errorf("not a color: %v", err)
// 	}
// 	b := uint8(v % 256)
// 	g := uint8((v / 256) % 256)
// 	r := uint8((v / (256 * 256)) % 256)
// 	return color.RGBA{R: r, G: g, B: b, A: 255}, nil
// }

// func printColor(c color.Color) string {
// 	r, g, b, a := c.RGBA()
// 	r, g, b, a = r/256, g/256, b/256, a/256
// 	return fmt.Sprintf("rgba(%v, %v, %v, %v)", r, g, b, a)
// }

func draw(fg, bg color.Color) {

	// fmt.Printf("drawing with foreground %s and background %s", printColor(fg), printColor(bg))
	fmt.Printf("drawing with foreground %s and background %s", fg, bg)

}
