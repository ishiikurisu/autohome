package autohome

import (
    "testing"
    "fmt"
)

func TestCanSayHello(t *testing.T) {
    said := fmt.Sprintf("hello")
    if said != "hello" {
        t.Error("WTF DUDE")
    }
}
