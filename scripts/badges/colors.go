package main

import (
	"strconv"
	"strings"
)

const (
	colorBG     = "#1a1b27"
	colorBorder = "#414868"
	colorTitle  = "#7aa2f7"
	colorText   = "#c0caf5"
	colorSub    = "#a9b1d6"
	colorOrange = "#e0af68"
	colorGreen  = "#9ece6a"
	colorPurple = "#bb9af7"
	colorRed    = "#f7768e"
	colorCyan   = "#7dcfff"
)

func luminance(hex string) float64 {
	h := strings.TrimPrefix(hex, "#")
	r, _ := strconv.ParseInt(h[0:2], 16, 64)
	g, _ := strconv.ParseInt(h[2:4], 16, 64)
	b, _ := strconv.ParseInt(h[4:6], 16, 64)
	return (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 255
}

func textColor(hex string) string {
	if luminance(hex) > 0.5 {
		return "#000000"
	}
	return "#ffffff"
}

func esc(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}
