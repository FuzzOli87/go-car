package main

import (
	"github.com/pkg/term"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/particle"
)

var (
	core    = particle.NewAdaptor("390044000f47363333343437", "4f04cecd40b14fa5b77b3c472d9a09a3cfb89c62")
	enableA = gpio.NewDirectPinDriver(core, "D0")
	enableB = gpio.NewDirectPinDriver(core, "A0")
	aOne    = gpio.NewDirectPinDriver(core, "D1")
	aTwo    = gpio.NewDirectPinDriver(core, "D2")
	bOne    = gpio.NewDirectPinDriver(core, "A1")
	bTwo    = gpio.NewDirectPinDriver(core, "A2")
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
				bOne.DigitalWrite(1)
				bTwo.DigitalWrite(0)
			// down reverse
			case 66:
				aOne.DigitalWrite(0)
				aTwo.DigitalWrite(0)
				bOne.DigitalWrite(0)
				bTwo.DigitalWrite(0)

				aOne.DigitalWrite(1)
				aTwo.DigitalWrite(0)
				bOne.DigitalWrite(0)
				bTwo.DigitalWrite(1)
			// RIGHT
			case 67:
				aOne.DigitalWrite(0)
				aTwo.DigitalWrite(1)
				bOne.DigitalWrite(1)
				bTwo.DigitalWrite(0)
			// LEFT
			case 68:
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
				// enableB.DigitalWrite(0)

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
