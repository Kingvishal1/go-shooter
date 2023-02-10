package spaceship

const SHIP_HEALTH = 20
const MAX_BULLETS = 100

type Spaceship struct {
    healthBar       int
    numberOfBullets int
}

func NewSpaceship() *Spaceship {
    return &Spaceship{healthBar: SHIP_HEALTH, numberOfBullets: MAX_BULLETS}
}
