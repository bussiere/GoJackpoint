package main

import (
	"fmt"
)

type Jack struct {
	id              rune
	Created         int64
	Updated         int64
	Nom             string
	Email           string
	Skill_Jack_id   []rune
	Carac_Jack_id   []rune
	Item_Jack_id    []rune
	Skillprivate_id []rune
	Itemprivate_id  []rune
	Password        string
	Key_public      string
	Key_private     string
	Message_id      []rune
	Statut          string
}

type Hand struct {
	id            rune
	Created       int64
	Updated       int64
	Skill_Jack_id []rune
	Carac_Jack_id []rune
	Item_Jack_id  []rune
	Message       string
}

type Skill struct {
	id              rune
	Created         int64
	Updated         int64
	Nom             string
	Parent_Skill_id rune
}

func main() {
	j := new(Jack)
	j.id = 2
	fmt.Printf("%d", j.id)

}
