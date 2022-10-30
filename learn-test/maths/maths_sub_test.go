package maths

import "testing"

func TestSeq(t *testing.T) {
	t.Run("test +ve numbers", func(t *testing.T) {
		if s := Add(5, 5); s != 10 {
			t.Errorf("Sum incorrect!! got: %d, want: %d\n", s, 10)
		}

		if s := Add(15, 5); s != 20 {
			t.Errorf("Sum incorrect!! got: %d, want: %d\n", s, 20)
		}

		if s := Add(0, 5); s != 5 {
			t.Errorf("Sum incorrect!! got: %d, want: %d\n", s, 5)
		}
	})

	t.Run("test -ve numbers", func(t *testing.T) {
		if s := Add(-5, -5); s != -10 {
			t.Errorf("Sum incorrect!! got: %d, want: %d\n", s, -10)
		}

		if s := Add(-15, -5); s != -20 {
			t.Errorf("Sum incorrect!! got: %d, want: %d\n", s, -20)
		}

		if s := Add(0, -5); s != -5 {
			t.Errorf("Sum incorrect!! got: %d, want: %d\n", s, -5)
		}
	})
}

func TestParallel(t *testing.T) {
	t.Run("test +ve numbers", func(t *testing.T) {
		t.Parallel()
		if s := Add(5, 5); s != 10 {
			t.Errorf("Sum incorrect!! got: %d, want: %d\n", s, 10)
		}

		if s := Add(15, 5); s != 20 {
			t.Errorf("Sum incorrect!! got: %d, want: %d\n", s, 20)
		}

		if s := Add(0, 5); s != 5 {
			t.Errorf("Sum incorrect!! got: %d, want: %d\n", s, 5)
		}
	})

	t.Run("test -ve numbers", func(t *testing.T) {
		t.Parallel()
		if s := Add(-5, -5); s != -10 {
			t.Errorf("Sum incorrect!! got: %d, want: %d\n", s, -10)
		}

		if s := Add(-15, -5); s != -20 {
			t.Errorf("Sum incorrect!! got: %d, want: %d\n", s, -20)
		}

		if s := Add(0, -5); s != -5 {
			t.Errorf("Sum incorrect!! got: %d, want: %d\n", s, -5)
		}
	})
}
