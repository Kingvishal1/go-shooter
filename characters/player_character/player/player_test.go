package player

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestCharacterCreation(t *testing.T) {
    t.Run("", func(t *testing.T) {
        player := CharacterCreation()
        assert.Equal(t, 10, player.healthBar)
    })
}
