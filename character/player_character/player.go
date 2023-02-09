package player_character

const HEALTHBAR = 10

type Character struct {
    healthBar int
}

func CharacterCreation() *Character {
    return &Character{HEALTHBAR}
}
