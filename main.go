package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/term"
)

// var (
// 	r = raspi.NewAdaptor()
// )
//
func main() {
	terminal, _ := term.Open("/dev/tty")
	term.RawMode(terminal)
	bytes := make([]byte, 3)

	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	signal.Notify(gracefulStop, syscall.SIGKILL)

	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)

		// enableA.PwmWrite(0)
		// enableB.PwmWrite(0)

		terminal.Restore()
		terminal.Close()
		os.Exit(0)
	}()

	for {
		numRead, err := terminal.Read(bytes)
		if err != nil {
			return
		}

		if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
			// Three-character control sequence, beginning with "ESC-[".

			// Since there are no ASCII codes for arrow keys, we use
			// Javascript key codes.
			switch bytes[2] {
			case 65:
				// up
				fmt.Println("UP")
			case 66:
				// down
				fmt.Println("DOWN")
			case 67:
				// RIGHT
				fmt.Println("RIGHT")
			case 68:
				// LEFT
				fmt.Println("RIGHT")
			}
		}
	}

	// enableA := gpio.NewDirectPinDriver(r, "12")
	// enableB := gpio.NewDirectPinDriver(r, "13")
	// a1 := gpio.NewDirectPinDriver(r, "31")
	// a2 := gpio.NewDirectPinDriver(r, "33")
	// b1 := gpio.NewDirectPinDriver(r, "18")
	// b2 := gpio.NewDirectPinDriver(r, "16")
	//
	// gracefulStop := make(chan os.Signal)
	// signal.Notify(gracefulStop, syscall.SIGTERM)
	// signal.Notify(gracefulStop, syscall.SIGINT)
	// signal.Notify(gracefulStop, syscall.SIGKILL)
	//
	// go func() {
	// 	sig := <-gracefulStop
	// 	fmt.Printf("caught sig: %+v", sig)
	//
	// 	enableA.PwmWrite(0)
	// 	enableB.PwmWrite(0)
	//
	// 	os.Exit(0)
	// }()
	//
	// work := func() {
	// 	fmt.Println("Initiating enas")
	//
	// 	// enableA.DigitalWrite(1)
	// 	// enableB.DigitalWrite(1)
	//
	// 	// forward
	// 	a1.DigitalWrite(1)
	// 	a2.DigitalWrite(0)
	//
	// 	b1.DigitalWrite(1)
	// 	b2.DigitalWrite(0)
	//
	// 	enableA.PwmWrite(10)
	// 	enableA.PwmWrite(250)
	//
	// 	enableB.PwmWrite(10)
	// 	enableB.PwmWrite(250)
	//
	// 	fmt.Println("Sleeping")
	// 	time.Sleep(3000 * time.Millisecond)
	//
	// 	enableA.DigitalWrite(0)
	// 	enableB.DigitalWrite(0)
	// }
	//
	// robot := gobot.NewRobot("GoCar",
	// 	[]gobot.Connection{r},
	// 	[]gobot.Device{enableA, enableB, a1, a1, b1, b2},
	// 	work)
	//
	// robot.Start()
}
