package main

import (
	"fmt"
	"strings"
)

func badgeSVG(message, colorHex, style string) string {
	colorHex = strings.TrimPrefix(colorHex, "#")

	var (
		H, RX      int
		charW, pad int
		text       string
		fontWeight string
	)
	const fontSize = 11

	if style == "for-the-badge" {
		H, RX = 28, 0
		charW, pad = 8, 20
		text = strings.ToUpper(message)
		fontWeight = "bold"
	} else {
		H, RX = 20, 3
		charW, pad = 7, 16
		text = message
		fontWeight = "normal"
	}

	W := len(text)*charW + pad
	fg := textColor("#" + colorHex)
	cx := W / 2
	cy := H/2 + 4

	return fmt.Sprintf(
		`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d">
  <rect width="%d" height="%d" rx="%d" fill="#%s"/>
  <text x="%d" y="%d" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" font-size="%d" font-weight="%s" fill="%s" text-anchor="middle">%s</text>
</svg>`,
		W, H,
		W, H, RX, colorHex,
		cx, cy, fontSize, fontWeight, fg, esc(text),
	)
}
