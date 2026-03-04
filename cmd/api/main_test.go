package main

import "testing"

// 這個測試是故意寫成「會失敗」，方便你觀察 CI / go test 的失敗情況
func TestIntentionalFail(t *testing.T) {
	t.Fatalf("這是一個故意失敗的測試，用來練習 CI 失敗狀況")
}

