package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func runHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	// Note: jika ingin menjalankan function dengan goroutine kita perlu menambahkan "go" didepan function
	go runHelloWorld()
	fmt.Println("Upssss")

	// Note: time.Sleep untuk memberitahu aplikasi bahwa harus berhenti sejenak 1 detik untuk mengakali goroutine
	time.Sleep(1 * time.Second)
}
