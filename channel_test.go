package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// ********************************* MEMBUAT CHANNEL ***********************************
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		// note: mengirim data ke channel
		channel <- "Abdul Wahid Kahar"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println("Ini adalah data yang diambil dari channel :", data)

	time.Sleep(5 * time.Second)

	// Note: pastikan jika membuat channel harus ada yang mengirim dan harus ada yang mengambil data tersebut
}

// ********************************* CHANNEL SEBAGAI PARAMTER ***********************************
func GiveMeResponse(channelFunc chan string) {
	time.Sleep(2 * time.Second)
	channelFunc <- "Daeng Sura"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	fmt.Println("Selesai mengirim data ke channel")

	data := <-channel
	fmt.Println("Ini adalah data yang diambil dari channel :", data)

	time.Sleep(5 * time.Second)
}

//************************************** CHANNEL IN OUT ***********************************

// Note: ini adalah function hanya untuk mengirim data ke channel (tidak bisa menerima data dari channel)
func OnlyIn(channelFunc chan<- string) {
	time.Sleep(2 * time.Second)
	channelFunc <- "Nama saya Wahid"
	//data := <-channelFunc | ini tidak bisa dilakukan
}

// Note: ini adalah function hanya untuk menerima data dari channel (tidak bisa mengirim data ke channel)
func OnlyOut(channelFunc <-chan string) {
	data := <-channelFunc
	//channelFunc <- "Nama saya Wahid" | ini tidak bisa dilakukan
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// ************************************** BUFFERED CHANNEL ***********************************
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 4) // chan string, 4 = "4" adalah capacity buffered channel
	defer close(channel)

	// mengirim data ke channel
	go func() {
		channel <- "Abdul"
		channel <- "Wahid"
		channel <- "Kahar"
		channel <- "Sura"
	}()

	// menerima data dari channel
	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(cap(channel))
		fmt.Println(len(channel))
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

// ************************************** RANGE CHANNEL ***********************************
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	// Mengirim data ke channel menggunakan for atau perulangan
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	// Menerima data dari channel menggunakan for atau perulangan
	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")

	// RESULT :
	//Menerima data Perulangan ke 0
	//Menerima data Perulangan ke 1
	//Menerima data Perulangan ke 2
	//Menerima data Perulangan ke 3
	//Menerima data Perulangan ke 4
	//Menerima data Perulangan ke 5
	//Menerima data Perulangan ke 6
	//Menerima data Perulangan ke 7
	//Menerima data Perulangan ke 8
	//Menerima data Perulangan ke 9
}

// ************************************** SELECT CHANNEL ***********************************

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data1 := <-channel1:
			fmt.Println("Menerima data dari channel 1", data1)
			counter++
		case data2 := <-channel2:
			fmt.Println("Menerima data dari channel 2", data2)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

// ************************************** SELECT DEFAULT CHANNEL ***********************************
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data1 := <-channel1:
			fmt.Println("Menerima data dari channel 1", data1)
			counter++
		case data2 := <-channel2:
			fmt.Println("Menerima data dari channel 2", data2)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
