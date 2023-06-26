package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Floor struct {
	LeftOfCharacter  *FloorElement
	OnCharacter      *FloorElement
	RightOfCharacter *FloorElement
	Reusable         *FloorElement
	MoveSpeed        float64
}

type FloorElement struct {
	Width, Height int
	PosX, PosY    float64
	Next          *FloorElement
	Color         color.Color
}

func (f *Floor) Update(cWidth int) {

	elements := []*FloorElement{f.LeftOfCharacter, f.OnCharacter, f.RightOfCharacter}

	var all *FloorElement

	for _, element := range elements {
		for element != nil {
			element.Update(f.MoveSpeed)
			tmp := element.Next
			element.Next = all
			all = element
			element = tmp
		}
	}

	for all != nil {
		tmp := all.Next
		end := all.PosX + float64(all.Width)
		if end < 0 {
			all.Next = f.Reusable
			f.Reusable = all
		} else if end < float64(gCharacterXPos)-float64(cWidth)/2 {
			all.Next = f.LeftOfCharacter
			f.LeftOfCharacter = all
		} else if all.PosX <= float64(gCharacterXPos)+float64(cWidth)/2 {
			all.Next = f.OnCharacter
			f.OnCharacter = all
		} else {
			all.Next = f.RightOfCharacter
			f.RightOfCharacter = all
		}
		all = tmp
	}

}

func (f *FloorElement) Update(MoveSpeed float64) {
	f.PosX -= MoveSpeed
}

func (f Floor) Draw(screen *ebiten.Image) {

	elements := []*FloorElement{f.LeftOfCharacter, f.OnCharacter, f.RightOfCharacter}

	for _, element := range elements {
		for element != nil {
			element.Draw(screen)
			element = element.Next
		}
	}

}

func (e FloorElement) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(e.PosX), float32(e.PosY), float32(e.Width), float32(e.Height), e.Color, false)
}
