package gadget

import "fmt"

// Tape Player
type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

// Tape Recorder
type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapeRecorder) Record() { //Same method as TapePlayer
	fmt.Println("Recording")
}
func (t TapeRecorder) Stop() { //Same method as TapePlayer
	fmt.Println("Stopped!")
}
