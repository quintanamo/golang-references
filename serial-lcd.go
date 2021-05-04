package references

import (
	"github.com/jacobsa/go-serial/serial"
)

func SerialLCDOutput([]byte message) {
	// options for writing to ADM1602U 16x2 Serial LCD Screen
	// outputs with Raspberry Pi UART pin
	options := serial.OpenOptions{
        PortName: "/dev/serial0",
        BaudRate: 9600,
        DataBits: 8,
        StopBits: 1,
        MinimumReadSize: 4,
    }

	// open the port, check for errors
	port, err := serial.Open(options)
    if err != nil {
        fmt.Println(err)
    }
    defer port.Close()

	// output the message
	port.Write(message)
}