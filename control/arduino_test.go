package autohome

import (
    "fmt"
    "testing"
    "github.com/tarm/serial"
)

func TestCanCommunicateToArduino(t *testing.T) {
    config := &serial.Config {
        Name: "COM3",
        Baud: 9600,
    }
    arduino, oops := serial.OpenPort(config)
    if oops != nil {
        t.Error("Couldn't open Arduino")
        return
    }
    defer arduino.Close()

    n, oops := arduino.Write([]byte("/on"))
    if oops != nil {
        t.Error("Can't write")
        return
    }

    buf := make([]byte, 128)
    n, oops = arduino.Read(buf)
    if oops != nil {
        t.Error("Can't read")
        return
    }
    fmt.Printf("%q", buf[:n])
}
