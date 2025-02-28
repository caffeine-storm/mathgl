package mathgl_test

import (
	"testing"

	"github.com/MobRulesGames/mathgl"
	"github.com/stretchr/testify/assert"
)

func TestTranslation(t *testing.T) {
	t.Run("mat3", func(t *testing.T) {
		var m mathgl.Mat3
		m.Translation(4, 7)

		// Column major order, remember!
		expected := []float32{
			1, 0, 0,
			0, 1, 0,
			4, 7, 1,
		}

		assert.Equal(t, mathgl.Mat3(expected), m)
	})

	t.Run("mat4", func(t *testing.T) {
		var m mathgl.Mat4
		m.Translation(4, 7, 0)

		// Column major order, remember!
		expected := []float32{
			1, 0, 0, 0,
			0, 1, 0, 0,
			0, 0, 1, 0,
			4, 7, 0, 1,
		}

		assert.Equal(t, mathgl.Mat4(expected), m)
	})
}
