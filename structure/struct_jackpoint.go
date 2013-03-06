package struct_jackpoint

type Jack struct {
	Id                    rune
	Created               int64
	Updated               int64
	Nom                   string
	Email                 string
	Skill_Jack_id         []rune
	Carac_Jack_id         []rune
	Item_Jack_id          []rune
	Skill_Jack_Private_id []rune
	Item_Jack_Private_id  []rune
	Password              string
	Key_public            string
	Key_private           string
	Message_id            []rune
	Statut                string
	Avatar                string
}

type Admin struct {
	id      rune
	Jack_id rune
	Created int64
	Updated int64
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
	id          rune
	Created     int64
	Updated     int64
	Nom         string
	Description string
}

type Filiation_Skill struct {
	id              rune
	Created         int64
	Updated         int64
	Parent_Skill_id rune
	Enfant_Skill_id rune
}

type Carac struct {
	id          rune
	Created     int64
	Updated     int64
	Nom         string
	Description string
}

type Item struct {
	id          rune
	Created     int64
	Updated     int64
	Nom         string
	Description string
}

type Item_Carac struct {
	id       rune
	Created  int64
	Updated  int64
	Carac_id rune
	Item_id  rune
}

type Item_Skill struct {
	id       rune
	Created  int64
	Updated  int64
	Skill_id rune
	Item_id  rune
}

type Skill_Jack struct {
	id       rune
	Created  int64
	Updated  int64
	Skill_id rune
	Jack_id  rune
}
type Item_Jack struct {
	id      rune
	Created int64
	Updated int64
	Item_id rune
	Jack_id rune
}
type Carac_Jack struct {
	id      rune
	Created int64
	Updated int64
	Item_id rune
	Jack_id rune
}
