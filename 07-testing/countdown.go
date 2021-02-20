package countdown

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer, from int) {
	for i := from; i > 0; i-- {
		fmt.Fprint(out, i)
		if i != 1 {
			fmt.Fprint(out, ", ")
		}
	}
	fmt.Fprint(out, " Done!")
}
