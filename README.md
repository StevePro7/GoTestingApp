# GoTestingApp
Testing GoTestingApp
go get -u github.com/stevepro7/gotestinglib
add code
go mod tidy

either consume

01.
import (
	"fmt"
	"github.com/stevepro7/gotestinglib/adriana"
)

func main() {
	val := adriana.Add(2, 3)
    fmt.Println(val)
}


02.
import (
	"fmt"
	. "github.com/stevepro7/gotestinglib/adriana"
)

func main() {
	val := Add(2, 3)
    fmt.Println(val)
}
