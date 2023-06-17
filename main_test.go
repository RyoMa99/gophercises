package main

import "testing"

func TestReadFile(t *testing.T) {
	t.Run("ファイル読み込み成功", func(t *testing.T) {
		want := "test\n"
		got, err := readFile("test.txt")
		if err != nil {
			t.Error(err)
		}

		if want != string(got) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
