package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	cases := []struct {
		name string
		size int
		want int
	}{
		{"обычный", 5, 5},
		{"ноль", 0, 0},
		{"отрицательный", -3, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := len(generateRandomElements(c.size)); got != c.want {
				t.Errorf("generateRandomElements(%d): длина %d, хотели %d", c.size, got, c.want)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	cases := []struct {
		name string
		data []int
		want int
	}{
		{"пустой", []int{}, 0},
		{"один", []int{42}, 42},
		{"макс в начале", []int{100, 1, 2}, 100},
		{"макс в конце", []int{1, 2, 100}, 100},
		{"все равны", []int{7, 7, 7}, 7},
		{"обычный", []int{3, 1, 4, 1, 5, 9, 2, 6}, 9},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := maximum(c.data); got != c.want {
				t.Errorf("maximum(%v) = %d, хотели %d", c.data, got, c.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	t.Run("совпадает с однопоточной", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 4, 2, 1, 7}
		want := maximum(data)
		if got := maxChunks(data); got != want {
			t.Errorf("maxChunks() = %d, хотели %d", got, want)
		}
	})
	t.Run("пустой слайс", func(t *testing.T) {
		if got := maxChunks([]int{}); got != 0 {
			t.Errorf("maxChunks([]) = %d, хотели 0", got)
		}
	})
	t.Run("короче чем CHUNKS", func(t *testing.T) {
		data := []int{5, 2, 8, 1}
		want := maximum(data)
		if got := maxChunks(data); got != want {
			t.Errorf("maxChunks() = %d, хотели %d", got, want)
		}
	})
	t.Run("с остатком от деления", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 99}
		want := maximum(data)
		if got := maxChunks(data); got != want {
			t.Errorf("maxChunks() = %d, хотели %d", got, want)
		}
	})
}
