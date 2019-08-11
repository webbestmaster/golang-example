package main

import (
	"fmt"
	"github.com/micmonay/keybd_event"
	"os"
	"runtime"
	"time"
)

func main() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	args := os.Args[1:]

	fmt.Println(args[0])

	// For linux, it is very important wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	time.Sleep(2 * time.Second)

	//set keys
	kb.SetKeys(keybd_event.VK_A, keybd_event.VK_B)
	kb.AddKey(keybd_event.VK_A)
	kb.AddKey(keybd_event.VK_B)
	kb.AddKey(keybd_event.VK_ENTER)

	//set shif is pressed
	kb.HasSHIFT(true)

	//launch
	err = kb.Launching()
	if err != nil {
		panic(err)
	}
	//Ouput : AB
}
