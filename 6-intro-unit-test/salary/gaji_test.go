package salary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHitungGajiPokok(t *testing.T) {
	t.Run("Test C-Level", func(t *testing.T) {
		posisi := "c-level"
		expected := 50000000
		actual := HitungGaji(posisi)
		if actual != expected {
			t.Error("hasil tidak sesuai, actual:", actual, "expected:", expected)
		}
	})

	t.Run("Test Manager", func(t *testing.T) {
		posisi := "manager"
		expected := 30000000
		actual := HitungGaji(posisi)
		// if actual != expected {
		// 	t.Error("hasil tidak sesuai, actual:", actual, "expected:", expected)
		// }
		assert.Equal(t, expected, actual, "hasil tidak sesuai")
	})

	// var posisi = []string
	// var expected = []int

	t.Run("Test Senior", func(t *testing.T) {
		posisi := "senior"
		expected := 30000000
		actual := HitungGaji(posisi)
		assert.Equal(t, expected, actual, "hasil tidak sesuai")
	})

	t.Run("Test Junior", func(t *testing.T) {
		posisi := "junior"
		expected := 10000000
		actual := HitungGaji(posisi)
		assert.Equal(t, expected, actual, "hasil tidak sesuai")
	})

	t.Run("Test karyawan", func(t *testing.T) {
		assert.Equal(t, 0, HitungGaji("karyawan"), "Hasil tidak sesuai")
	})

}

func TestSalaryTotal(t *testing.T) {
	actual := SalaryTotal(10000000, 50000000)
	expected := 60000000
	assert.Equal(t, expected, actual, "hasil tidak sesuai")
}
