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

type Window struct {
	size   int
	window []VideoPacket
	start  int
	mu     sync.Mutex
}

func (w *Window) InitializeWindow() {
	w.window = make([]VideoPacket, w.size)
}

func (w *Window) AddPacket(packet VideoPacket) {
	w.mu.Lock()
	defer w.mu.Unlock()

	fmt.Println("Packet", packet.SequenceNumber, "received!")

	// if the packet is in rder -> play immediately
	if packet.SequenceNumber == w.start {
		w.PlayPacket(packet)
		w.start++

		// updates the window after playing the current packet and also checks for the next packets in case they arrived early
		w.UpdateWindow()
		return
	}

	// if the packet is inside the window but not in order -> add into the buffer
	if packet.SequenceNumber > w.start && packet.SequenceNumber <= w.start+w.size {
		index := packet.SequenceNumber % w.size
		w.window[index] = packet

		fmt.Println("Packet", packet.SequenceNumber, "buffered!")
		return
	}

	// if the packet is outside the window -> drop it
	fmt.Println("Packet", packet.SequenceNumber, "dropped!")
}

func (w *Window) PlayPacket(packet VideoPacket) {
	// fmt.Println("Playing video...")
	// i := w.start
	// numPacket := w.start
	// start := w.start
	// for numPacket < start+w.size {
	// 	packet := w.window[i%w.size]
	// 	fmt.Println("Playing packet number", packet.SequenceNumber, "with duration", packet.Duration, "seconds!")
	// 	time.Sleep(time.Duration(packet.Duration) * time.Second)
	// 	i = (i + 1) % w.size
	// 	numPacket++
	// 	// w.start = (w.start + 1) % w.size
	// }
	// fmt.Println("Video completed!")

	fmt.Println("Playing packet", packet.SequenceNumber, "with duration", packet.Duration, "seconds!")
	time.Sleep(time.Duration(packet.Duration) * time.Second)
	fmt.Println("Packet", packet.SequenceNumber, "played!")
}

func (w *Window) UpdateWindow() {
	for {
		index := w.start % w.size
		packet := w.window[index]

		if packet.SequenceNumber != w.start {
			return
		}

		w.PlayPacket(packet)
		w.window[index] = VideoPacket{} // clear the packet from the window after playing
		w.start++
	}
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
	window := Window{size: 5, start: 1}
	window.InitializeWindow()
	var wg2 sync.WaitGroup

	for i := 1; i <= 10; i++ {
		newPacket := VideoPacket{
			Data:           "Video packet" + strconv.Itoa(i) + "!",
			SequenceNumber: i,
			Duration:       randomTimeGenerator(),
		}
		wg2.Go(func() {
			fmt.Println("Sending packet:", newPacket.SequenceNumber)
			window.AddPacket(newPacket)
		})
	}
	wg2.Wait()
}

func randomTimeGenerator() int {
	return 1 + rand.Intn(5)
}
