package world

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNewWorld(t *testing.T) {
    t.Run("", func(t *testing.T) {
        world := NewWorld()
        assert.IsType(t, &World{}, world)
    })
}
