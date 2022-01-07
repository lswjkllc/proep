package join

import (
	"fmt"
	"strings"
	"time"

	us "github.com/lswjkllc/proep/src/utils"
)

func Sprintf(strs ...interface{}) string {
	st := time.Now()
	str := fmt.Sprint(strs...)
	dt := time.Since(st)
	fmt.Printf("  =>  sprintf cost: %v %d\n", dt, len(str))
	return str
}

func StringBuilder(strs ...string) string {
	st := time.Now()
	str := us.JoinStrings(strs...)
	dt := time.Since(st)
	fmt.Printf("  =>  strings builder cost: %v %d\n", dt, len(str))
	return str
}

func StringJoin(strs ...string) string {
	st := time.Now()
	str := strings.Join(strs, "")
	dt := time.Since(st)
	fmt.Printf("  =>  strings join cost: %v %d\n", dt, len(str))
	return str
}
