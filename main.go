/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"embed"
	"fmt"
	"log"

	"github.com/HrithikSawant/go-ccwc/cmd"
)

//go:embed assets/banner.txt
var banner embed.FS

func printBannerFromFile() {
	// Read the embedded banner file
	data, err := banner.ReadFile("assets/banner.txt")
	if err != nil {
		log.Println("Error: could not read embedded banner:", err)
		return
	}

	// Print the banner content
	fmt.Println(string(data))
}

func main() {
	printBannerFromFile()
	cmd.Execute()
}
