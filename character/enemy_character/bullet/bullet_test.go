package bullet

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNewBullet(t *testing.T) {
    bullet := newBullet()
    assert.IsType(t, bullet, &Bullet{})
}
