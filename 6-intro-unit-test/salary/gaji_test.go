package salary

import "testing"

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
		if actual != expected {
			t.Error("hasil tidak sesuai, actual:", actual, "expected:", expected)
		}
	})

	// var posisi = []string
	// var expected = []int

	t.Run("Test Senior", func(t *testing.T) {
		posisi := "senior"
		expected := 30000000
		actual := HitungGaji(posisi)
		if actual != expected {
			t.Error("hasil tidak sesuai, actual:", actual, "expected:", expected)
		}
	})

	t.Run("Test Junior", func(t *testing.T) {
		posisi := "junior"
		expected := 10000000
		actual := HitungGaji(posisi)
		if actual != expected {
			t.Error("hasil tidak sesuai, actual:", actual, "expected:", expected)
		}
	})

}

func TestSalaryTotal(t *testing.T) {
	actual := SalaryTotal(10000000, 50000000)
	expected := 60000000
	if actual != expected {
		t.Error("hasil tidak sesuai, actual:", actual, "expected:", expected)
	}
}
