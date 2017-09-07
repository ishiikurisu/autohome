package autohome

import (
    "testing"
    "github.com/tarm/serial"
    "time"
)

func TestCanCommunicateToArduino(t *testing.T) {
    config := &serial.Config {
        Name: "COM4",
        Baud: 9600,
        ReadTimeout: 2 * time.Second,
    }
    arduino, oops := serial.OpenPort(config)
    defer arduino.Close()
    if oops != nil {
        t.Error("Couldn't open Arduino")
        return
    }

    time.Sleep(2 * time.Second)

    n, oops := arduino.Write([]byte("/on\n"))
    if oops != nil || n == 0 {
        t.Error("Can't write")
        return
    }

    buf := make([]byte, 128)
    n, _ = arduino.Read(buf)
    if n == 0 {
        t.Error("Can't read")
        return
    }

    time.Sleep(2 * time.Second)
    n, oops = arduino.Write([]byte("/off\n"))
    if oops != nil || n == 0 {
        t.Error("Can't write")
        return
    }

    buf = make([]byte, 128)
    n, _ = arduino.Read(buf)
    if n == 0 {
        t.Error("Can't read")
        return
    }
}
