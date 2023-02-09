package player_character

type Character struct {
    healthBar int
}

func CharacterCreation() Character {
    return Character{10}
}
