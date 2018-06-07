package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var (
	r = raspi.NewAdaptor()
)

func main() {
	enableA := gpio.NewDirectPinDriver(r, "12")
	enableB := gpio.NewDirectPinDriver(r, "13")
	a1 := gpio.NewDirectPinDriver(r, "31")
	a2 := gpio.NewDirectPinDriver(r, "33")
	b1 := gpio.NewDirectPinDriver(r, "18")
	b2 := gpio.NewDirectPinDriver(r, "16")

	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	signal.Notify(gracefulStop, syscall.SIGKILL)

	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)

		enableA.PwmWrite(0)
		enableB.PwmWrite(0)

		os.Exit(0)
	}()

	work := func() {
		fmt.Println("Initiating enas")

		// enableA.DigitalWrite(1)
		// enableB.DigitalWrite(1)

		// forward
		a1.DigitalWrite(1)
		a2.DigitalWrite(0)

		b1.DigitalWrite(1)
		b2.DigitalWrite(0)

		enableA.PwmWrite(10)
		enableA.PwmWrite(250)

		enableB.PwmWrite(10)
		enableB.PwmWrite(250)

		fmt.Println("Sleeping")
		time.Sleep(3000 * time.Millisecond)

		enableA.DigitalWrite(0)
		enableB.DigitalWrite(0)
	}

	robot := gobot.NewRobot("GoCar",
		[]gobot.Connection{r},
		[]gobot.Device{enableA, enableB, a1, a1, b1, b2},
		work)

	robot.Start()
}
