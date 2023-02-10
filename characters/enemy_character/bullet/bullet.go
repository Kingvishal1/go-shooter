package bullet

const DAMAGE = 1

type Bullet struct {
    damage int
}

func newBullet() *Bullet {
    return &Bullet{DAMAGE}
}
