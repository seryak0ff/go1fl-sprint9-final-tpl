package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	t.Run("zero size", func(t *testing.T) {
		result := generateRandomElements(0)
		if len(result) != 0 {
			t.Errorf("expected empty slice, got length %d", len(result))
		}
	})

	t.Run("negative size", func(t *testing.T) {
		result := generateRandomElements(-5)
		if len(result) != 0 {
			t.Errorf("expected empty slice, got length %d", len(result))
		}
	})

	t.Run("positive size", func(t *testing.T) {
		size := 100
		result := generateRandomElements(size)
		if len(result) != size {
			t.Errorf("expected length %d, got %d", size, len(result))
		}
		for i, value := range result {
			if value < 0 {
				t.Errorf("element at index %d is negative: %d", i, value)
			}
		}
	})
}

func TestMaximum(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		got := maximum([]int{})
		if got != 0 {
			t.Errorf("expected 0, got %d", got)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		got := maximum(nil)
		if got != 0 {
			t.Errorf("expected 0, got %d", got)
		}
	})

	t.Run("single element", func(t *testing.T) {
		got := maximum([]int{42})
		if got != 42 {
			t.Errorf("expected 42, got %d", got)
		}
	})

	t.Run("multiple elements", func(t *testing.T) {
		got := maximum([]int{1, 5, 3, 9, 2})
		if got != 9 {
			t.Errorf("expected 9, got %d", got)
		}
	})

	t.Run("max at beginning", func(t *testing.T) {
		got := maximum([]int{100, 1, 2, 3})
		if got != 100 {
			t.Errorf("expected 100, got %d", got)
		}
	})

	t.Run("max at end", func(t *testing.T) {
		got := maximum([]int{1, 2, 3, 100})
		if got != 100 {
			t.Errorf("expected 100, got %d", got)
		}
	})

	t.Run("all equal", func(t *testing.T) {
		got := maximum([]int{7, 7, 7, 7})
		if got != 7 {
			t.Errorf("expected 7, got %d", got)
		}
	})
}

func TestMaxChunks(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		got := maxChunks([]int{})
		if got != 0 {
			t.Errorf("expected 0, got %d", got)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		got := maxChunks(nil)
		if got != 0 {
			t.Errorf("expected 0, got %d", got)
		}
	})

	t.Run("single element", func(t *testing.T) {
		got := maxChunks([]int{42})
		if got != 42 {
			t.Errorf("expected 42, got %d", got)
		}
	})

	t.Run("less than chunks", func(t *testing.T) {
		got := maxChunks([]int{1, 5, 3})
		if got != 5 {
			t.Errorf("expected 5, got %d", got)
		}
	})

	t.Run("exact chunks size", func(t *testing.T) {
		data := make([]int, CHUNKS*10)
		for i := range data {
			data[i] = i
		}
		data[len(data)-1] = 9999
		got := maxChunks(data)
		if got != 9999 {
			t.Errorf("expected 9999, got %d", got)
		}
	})

	t.Run("matches maximum", func(t *testing.T) {
		data := generateRandomElements(1000)
		want := maximum(data)
		got := maxChunks(data)
		if got != want {
			t.Errorf("maxChunks=%d, maximum=%d", got, want)
		}
	})
}
