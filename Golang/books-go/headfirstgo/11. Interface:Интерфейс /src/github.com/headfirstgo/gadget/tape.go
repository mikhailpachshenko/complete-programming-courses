package gadget

import "fmt"

type TypePlayer struct {
	Batteries string
}

func (t TypePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TypePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TypeRecorder struct {
	Microphones int
}

func (t TypeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TypeRecorder) Record() {
	fmt.Println("Recording")
}

func (t TypeRecorder) Stop() {
	fmt.Println("Stopped!")
}
