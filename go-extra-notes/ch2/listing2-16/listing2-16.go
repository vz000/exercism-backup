import (
	"io/ioutil"
	"os"
	"fmt"
)

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.RealAll(f)
}

/*
Notes:
* The defer statement postpones the execution of a function until the
surrounding function returns.
It's commonly used for cleanup operations like closing files or unlocking mutexes
* defer postpones a function call until after the surronding function completes.
* The deferred call executes after all non-deferred statements in the function
*/

func main() {
	defer fmt.Println("This prints last")
	fmt.Println("This prints first")
	fmt.Println("This prints second")
}
