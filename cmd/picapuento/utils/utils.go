package utils

import (
	"fmt"
	"regexp"
)

func ExtractToken(html string) (string, error) {
	re := regexp.MustCompile(`<input type="hidden" value="([^"]*)" id="token" />`)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", fmt.Errorf("token not found")
}

func DisplayArt() {
	fmt.Println(` ________  ___  ________  ________  ________  ___  ___  _______   ________   _________  ________     
|\   __  \|\  \|\   ____\|\   __  \|\   __  \|\  \|\  \|\  ___ \ |\   ___  \|\___   ___\\   __  \    
\ \  \|\  \ \  \ \  \___|\ \  \|\  \ \  \|\  \ \  \\\  \ \   __/|\ \  \\ \  \|___ \  \_\ \  \|\  \   
 \ \   ____\ \  \ \  \    \ \   __  \ \   ____\ \  \\\  \ \  \_|/_\ \  \\ \  \   \ \  \ \ \  \\\  \  
  \ \  \___|\ \  \ \  \____\ \  \ \  \ \  \___|\ \  \\\  \ \  \_|\ \ \  \\ \  \   \ \  \ \ \  \\\  \ 
   \ \__\    \ \__\ \_______\ \__\ \__\ \__\    \ \_______\ \_______\ \__\\ \__\   \ \__\ \ \_______\
    \|__|     \|__|\|_______|\|__|\|__|\|__|     \|_______|\|_______|\|__| \|__|    \|__|  \|_______|`)

}
