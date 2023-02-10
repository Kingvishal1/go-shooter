package rocket

const DAMAGE = 1

type Rocket struct {
    damage int
}

func NewRocket() *Rocket {
    return &Rocket{DAMAGE}
}
