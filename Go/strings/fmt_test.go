package strings

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func Example_genOSSDestPath() {
	p := genOSSDestPath()
	fmt.Println(p)
	// Output:
	// ""
}

func genOSSDestPath() string {
	now := time.Now().Unix()
	return fmt.Sprintf("%x-%s", now, uuid.New().String())
}
