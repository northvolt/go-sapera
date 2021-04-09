package main

import (
	"fmt"
	"time"

	"github.com/northvolt/go-sapera"
)

const (
	serverName = "your-capture-device-goes-here"
	configFile = "./your-config-file-here.ccf"
)

func main() {
	doneChan := make(chan bool)

	fmt.Println("Initializing SaperaLT...")
	sapera.Init()

	fmt.Println("starting, sapera.GetServerCount() = ", sapera.GetServerCount())
	go scan(doneChan)

	select {
	case <-doneChan:
	}

	fmt.Println("done.")
}

func scan(doneChan chan bool) {
	loc := sapera.NewSapLocation(serverName, 0)
	defer loc.Delete()

	acq := sapera.NewSapAcquisitionForLocation(loc, configFile)
	defer acq.Delete()

	acq.Create()
	defer acq.Destroy()

	buf := sapera.NewSapBufferForAcquisition(1, acq)
	defer buf.Delete()

	buf.Create()
	defer buf.Destroy()

	context := sapera.NewSapXferContext()
	defer context.Delete()

	context.SetID(1)
	context.SetBuffer(buf)

	xfer := sapera.NewSapAcqToBuf(acq, buf, context)
	defer xfer.Delete()

	height, ok := acq.GetParameterInt32(sapera.CORACQ_PRM_CROP_HEIGHT)
	if !ok {
		fmt.Println("Could not get height parameter")
	}
	fmt.Println("height = ", height)
	width, ok := acq.GetParameterInt32(sapera.CORACQ_PRM_CROP_WIDTH)
	if !ok {
		fmt.Println("Could not get width parameter")
	}
	fmt.Println("width = ", width)
	format, ok := acq.GetParameterInt32(sapera.CORACQ_PRM_OUTPUT_FORMAT)
	if !ok {
		fmt.Println("Could not get format parameter")
	}
	fmt.Println("format = ", format)

	xfer.Create()
	defer xfer.Destroy()

	time.Sleep(3 * time.Second)

	xfer.Grab()
	fmt.Println("grabbing...")

	start := time.Now()
	for {
		if time.Since(start) > 5*time.Second {
			break
		}

		time.Sleep(time.Second)
	}

	xfer.Freeze()

	if !xfer.Wait(3000) {
		xfer.Abort()
	}

	doneChan <- true
}
