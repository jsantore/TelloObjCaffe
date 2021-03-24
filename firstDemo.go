package main

import (
	"fmt"

	//	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"time"
)

func main() {
	drone := tello.NewDriver("8888")

	work := func() {
		err := drone.TakeOff()
		if err != nil {
			fmt.Println(err)
		}
		gobot.After(3*time.Second, func() {
			drone.Left(50)
			time.Sleep(time.Second * 3)
			drone.Right(50)
			time.Sleep(time.Second * 3)
			drone.Land()
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}
