package fmt

import (
	"fmt"
	"path/filepath"
)

func Example_quote() {
	s := "(/|$)(.*)"
	fmt.Println(quote(s))
	// Output:
	// "(/|$)(.*)"
}

func dbFilePath(dir string, id uint64) string {
	return filepath.Join(dir, fmt.Sprintf("%016x.snap.db", id))
}

func Example_dbFilePath() {
	s := dbFilePath("/var/lib/etcd/member", 3700038)
	fmt.Println(s)
	// Output:
	// bad
}
