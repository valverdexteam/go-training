package main

import (
	"fmt"
)

func main() {
	oldDvDPlayer := &DvdPlayer{
		Name: "Old classic Dvd Player",
	}

	amazingCdPlayer := &CdPlayer{
		Name: "Old fashioned cd player",
	}

	devices := []MediaPlayer{oldDvDPlayer, amazingCdPlayer}

	for _, device := range devices {
		fmt.Println(device.Play())
	}

}

type DvdPlayer struct {
	Name string
}

type MediaPlayer interface {
	Play() string
}

func (dvd *DvdPlayer) Play() string {
	return fmt.Sprintf("dvd %s is playing some good movie", dvd.Name)
}

type CdPlayer struct {
	Name string
}

func (cd *CdPlayer) Play() string {
	return fmt.Sprintf("cd %s is playing some good music", cd.Name)
}
