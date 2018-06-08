package main

import (
	"fmt"

	"github.com/pkg/term"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var (
	r       = raspi.NewAdaptor()
	enableA = gpio.NewDirectPinDriver(r, "12")
	enableB = gpio.NewDirectPinDriver(r, "13")
	aOne    = gpio.NewDirectPinDriver(r, "31")
	aTwo    = gpio.NewDirectPinDriver(r, "33")
	bOne    = gpio.NewDirectPinDriver(r, "18")
	bTwo    = gpio.NewDirectPinDriver(r, "16")
)

func main() {

	terminal, _ := term.Open("/dev/tty")
	term.RawMode(terminal)
	bytes := make([]byte, 3)

	// Everything disabled to start off
	enableA.DigitalWrite(0)
	enableB.DigitalWrite(0)
	aOne.DigitalWrite(0)
	aTwo.DigitalWrite(0)
	bOne.DigitalWrite(0)
	bTwo.DigitalWrite(0)

	for {
		numRead, err := terminal.Read(bytes)
		if err != nil {
			return
		}

		if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
			// Three-character control sequence, beginning with "ESC-[".

			// Since there are no ASCII codes for arrow keys, we use Javascript key codes.
			switch bytes[2] {
			// up forward
			case 65:
				aOne.DigitalWrite(0)
				aTwo.DigitalWrite(1)
				bOne.DigitalWrite(0)
				bTwo.DigitalWrite(1)
			// down reverse
			case 66:
				fmt.Println("DOWN")
				aOne.DigitalWrite(1)
				aTwo.DigitalWrite(0)
				bOne.DigitalWrite(1)
				bTwo.DigitalWrite(0)
			// RIGHT
			case 67:
				fmt.Println("RIGHT")
				aOne.DigitalWrite(0)
				aTwo.DigitalWrite(1)
				bOne.DigitalWrite(1)
				bTwo.DigitalWrite(0)
			// LEFT
			case 68:
				fmt.Println("LEFT")
				aOne.DigitalWrite(1)
				aTwo.DigitalWrite(0)
				bOne.DigitalWrite(0)
				bTwo.DigitalWrite(1)
			}
		} else if numRead == 1 {
			switch int(bytes[0]) {
			// q is for quit!
			case 113:
				enableA.DigitalWrite(0)
				enableB.DigitalWrite(0)

				aOne.DigitalWrite(0)
				aTwo.DigitalWrite(0)
				bOne.DigitalWrite(0)
				bTwo.DigitalWrite(0)

				terminal.Restore()
				terminal.Close()
				// s Start Engine
			case 115:
				enableA.DigitalWrite(1)
				enableB.DigitalWrite(1)
				// b break
			case 98:
				aOne.DigitalWrite(0)
				aTwo.DigitalWrite(0)
				bOne.DigitalWrite(0)
				bTwo.DigitalWrite(0)
			}
		}
	}
}
