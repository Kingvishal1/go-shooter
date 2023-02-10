package spaceship

const SHIP_HEALTH = 20

type Spaceship struct {
    healthBar int
}

func NewSpaceship() *Spaceship {
    return &Spaceship{healthBar: SHIP_HEALTH}
}
