package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// httptest.NewRecorderを使うとResponseWriterインターフェースを満たす
func TestNewMux(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	sut := NewMux()
	sut.ServeHTTP(w, r)
	resp := w.Result()
	// テストが終了した際に、レスポンスのボディをクローズするクリーンアップ関数を設定。
	t.Cleanup(func() { _ = resp.Body.Close() })

	if resp.StatusCode != http.StatusOK {
		t.Error("want status code 200, but", resp.StatusCode)
	}

	// レスポンスのボディを読み取る
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	want := `{"status": "ok"}`
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}
