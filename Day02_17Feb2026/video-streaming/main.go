package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

type VideoPacket struct {
	Data           string
	SequenceNumber int
	Duration       int
}

type Buffer struct {
	Packets []VideoPacket
}

func (b *Buffer) AddPacket(packet VideoPacket) {
	fmt.Println("Sending packet number:", packet.SequenceNumber)
	b.Packets = append(b.Packets, packet)
}

func (b Buffer) Play() {
	sort.Slice(b.Packets, func(i, j int) bool {
		return b.Packets[i].SequenceNumber < b.Packets[j].SequenceNumber
	})

	fmt.Println("Playing video...")
	for _, packet := range b.Packets {
		fmt.Println("Playing packet number", packet.SequenceNumber, "with duration", packet.Duration, "seconds!")
		time.Sleep(time.Duration(packet.Duration) * time.Second)
	}
	fmt.Println("Video completed!")
}

func (b Buffer) PlayWithoutSort() {
	fmt.Println("Playing video...")

}

type Window struct {
	size   int
	window []VideoPacket
	start  int
}

func (w *Window) InitializeWindow() {
	w.window = make([]VideoPacket, w.size)
}

func (w *Window) AddPacket(packet VideoPacket) {
	fmt.Println("Sending packet number:", packet.SequenceNumber)
	if packet.SequenceNumber <= w.start+w.size {
		index := (packet.SequenceNumber - 1) % w.size
		w.window[index] = packet
	}
}

func (w Window) Play() {
	fmt.Println("Playing video...")
	i := w.start
	numPacket := w.start
	start := w.start
	for numPacket < start+w.size {
		packet := w.window[i%w.size]
		fmt.Println("Playing packet number", packet.SequenceNumber, "with duration", packet.Duration, "seconds!")
		time.Sleep(time.Duration(packet.Duration) * time.Second)
		i = (i + 1) % w.size
		numPacket++
		// w.start = (w.start + 1) % w.size
	}
	fmt.Println("Video completed!")
}

func main() {
	buffer := Buffer{}
	var wg sync.WaitGroup

	// method 1: using a buffer and sorting it before playing
	for i := 1; i <= 5; i++ {
		newPacket := VideoPacket{
			Data:           "Video packet" + strconv.Itoa(i) + "!",
			SequenceNumber: i,
			Duration:       randomTimeGenerator(),
		}
		wg.Go(func() {
			buffer.AddPacket(newPacket)
		})
	}
	wg.Wait()
	buffer.Play()

	// method 2: using a window
	window := Window{size: 5, start: 2}
	window.InitializeWindow()
	var wg2 sync.WaitGroup

	for i := 3; i <= 8; i++ {
		newPacket := VideoPacket{
			Data:           "Video packet" + strconv.Itoa(i) + "!",
			SequenceNumber: i,
			Duration:       randomTimeGenerator(),
		}
		wg2.Go(func() {
			window.AddPacket(newPacket)
		})
	}
	wg2.Wait()
	window.Play()

}

func randomTimeGenerator() int {
	return 1 + rand.Intn(5)
}
