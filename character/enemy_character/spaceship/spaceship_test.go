package spaceship

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNewSpaceship(t *testing.T) {
    t.Run("", func(t *testing.T) {
        spaceShip := NewSpaceship()
        assert.IsType(t, &Spaceship{}, spaceShip)
    })
}
