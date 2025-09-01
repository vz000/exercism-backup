import(
	"io/ioutil"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		panic (err) // calls panic
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
