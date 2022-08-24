package config

import "github.com/fatih/color"

func initialMessage() string {
	green := color.New(color.FgGreen)
	blue := color.New(color.FgBlue)
	cyan := color.New(color.FgCyan)
	logo := "" +
		green.Sprint(" ______  __       ____    ____      __        __      ") + "\n" +
		green.Sprint("/\\__  _\\/\\ \\     /\\  _`\\ /\\  _`\\   /\\ \\      /\\ \\     ") + "\n" +
		green.Sprint("\\/_/\\ \\/\\ \\ \\    \\ \\ \\/\\ \\ \\ \\L\\ \\ \\_\\ \\___  \\_\\ \\___ ") + "\n" +
		cyan.Sprint("   \\ \\ \\ \\ \\ \\  __\\ \\ \\ \\ \\ \\ ,  //\\___  __\\/\\___  __\\") + "\n" +
		cyan.Sprint("    \\ \\ \\ \\ \\ \\L\\ \\\\ \\ \\_\\ \\ \\ \\\\ \\/__/\\ \\_/\\/__/\\ \\_/") + "\n" +
		blue.Sprint("     \\ \\_\\ \\ \\____/ \\ \\____/\\ \\_\\ \\_\\ \\ \\_\\     \\ \\_\\") + " \n" +
		blue.Sprint("      \\/_/  \\/___/   \\/___/  \\/_/\\/ /  \\/_/      \\/_/") + "\n" +
		"                                                      "
	return logo
}
