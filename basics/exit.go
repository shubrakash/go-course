package basics

import (
	"fmt"
	"os"
)

func exitMain() {

	/// doesn't get executed as well
	defer fmt.Println("a Deferred statement")

	fmt.Println("starting main function")

	/// exit with status code 1
	os.Exit(1)

	/// never gets executed
	fmt.Println("End of main function")

}
