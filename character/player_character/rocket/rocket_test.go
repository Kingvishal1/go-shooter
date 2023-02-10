package rocket

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNewRocket(t *testing.T) {
    t.Run("", func(t *testing.T) {
        rocket := NewRocket()
        assert.IsType(t, &Rocket{}, rocket)
    })
}
