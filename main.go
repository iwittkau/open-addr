package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("no args received")
		os.Exit(1)
	}

	browser := "Safari"
	address := ""

	if strings.HasPrefix(os.Args[1], "-") {
		flag := strings.TrimPrefix(os.Args[1], "-")
		switch flag {
		case "h":
			fmt.Println("*help does not exist*")
			os.Exit(1)
		case "c":
			fmt.Println("using chrome")
			browser = "Google Chrome"
		case "f":
			fmt.Println("using firefox")
			browser = "Firefox"
		case "s":
			fmt.Println("using safari")
			browser = "Safari"
		}
		if len(os.Args) < 3 {
			fmt.Println("missing address")
			os.Exit(1)
		}
		address = os.Args[2]
	} else {
		address = os.Args[1]

	}

	if strings.HasPrefix(address, ":") {
		address = "localhost" + address
	}
	if !strings.HasPrefix(address, "http") {
		address = "http://" + address
	}

	cmd := exec.Command("open", "-a", browser, address)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
