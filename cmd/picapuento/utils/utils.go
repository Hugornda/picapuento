package utils

import (
	"fmt"
	"log"
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
	log.Println(`
________  ___  ________  ________                 ________  ________  ________   _________  ________     
|\   __  \|\  \|\   ____\|\   __  \               |\   __  \|\   __  \|\   ___  \|\___   ___\\   __  \    
\ \  \|\  \ \  \ \  \___|\ \  \|\  \  ____________\ \  \|\  \ \  \|\  \ \  \\ \  \|___ \  \_\ \  \|\  \   
 \ \   ____\ \  \ \  \    \ \   __  \|\____________\ \   ____\ \  \\\  \ \  \\ \  \   \ \  \ \ \  \\\  \  
  \ \  \___|\ \  \ \  \____\ \  \ \  \|____________|\ \  \___|\ \  \\\  \ \  \\ \  \   \ \  \ \ \  \\\  \ 
   \ \__\    \ \__\ \_______\ \__\ \__\              \ \__\    \ \_______\ \__\\ \__\   \ \__\ \ \_______\
    \|__|     \|__|\|_______|\|__|\|__|               \|__|     \|_______|\|__| \|__|    \|__|  \|_______|
    `)

}
