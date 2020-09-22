package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/okamotchan/imageconv/imgconv"
)

func main() {
	from := flag.String("from", "jpg", "変更前の拡張子")
	to := flag.String("to", "png", "変更後の拡張子")
	rm := flag.Bool("remove", false, "削除")

	flag.Parse()
	src := flag.Arg(0)
	dst := strings.Replace(flag.Arg(0), *from, *to, 1)

	err := imgconv.Convert(src, dst, *rm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
