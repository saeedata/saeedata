package main

import (
	"fmt"
	"os"
)

type badge struct {
	path    string
	message string
	color   string
	style   string
}

var badges = []badge{
	// flat style — profile header row
	{"badges/website.svg", "nimcode.me", "44cc11", "flat"},
	{"badges/linkedin.svg", "LinkedIn", "0080ff", "flat"},
	{"badges/android-dev.svg", "Senior Android Dev", "3DDC84", "flat"},
	{"badges/kotlin-flat.svg", "Kotlin", "7F52FF", "flat"},
	{"badges/go-flat.svg", "Go", "00ADD8", "flat"},
	{"badges/ai-flat.svg", "AI / ML", "FF6F00", "flat"},
	// for-the-badge style — tech stack section
	{"badges/android.svg", "Android", "3DDC84", "for-the-badge"},
	{"badges/kotlin.svg", "Kotlin", "7F52FF", "for-the-badge"},
	{"badges/java.svg", "Java", "007396", "for-the-badge"},
	{"badges/jetpack-compose.svg", "Jetpack Compose", "4285F4", "for-the-badge"},
	{"badges/go.svg", "Go", "00ADD8", "for-the-badge"},
	{"badges/gradle.svg", "Gradle", "02303A", "for-the-badge"},
	{"badges/openai.svg", "OpenAI", "412991", "for-the-badge"},
	{"badges/langchain.svg", "LangChain", "1C3C3C", "for-the-badge"},
	{"badges/git.svg", "Git", "F05032", "for-the-badge"},
	{"badges/github-actions.svg", "GitHub Actions", "2088FF", "for-the-badge"},
	{"badges/docker.svg", "Docker", "2496ED", "for-the-badge"},
}

func main() {
	if err := os.MkdirAll("assets/badges", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "mkdir: %v\n", err)
		os.Exit(1)
	}
	for _, b := range badges {
		path := "assets/" + b.path
		content := badgeSVG(b.message, b.color, b.style)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "write %s: %v\n", path, err)
			os.Exit(1)
		}
		fmt.Printf("  wrote %s\n", path)
	}
	fmt.Println("Done.")
}
