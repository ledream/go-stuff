package main
import "flag"

var b = flag.String("b", "", "number non-blank output lines")
var e = flag.String("e", "", "display $ at the end of each line (implies -v)")
var n = flag.String("n", "", "number all output lines")
var s = flag.String("s", "", "squeeze multiple blank lines into one")
var t = flag.String("t", "", "display tab as ^I (implies -v)")
var u = flag.String("u", "", "do not buffer output")
var v = flag.String("v", "", "display non-printing chars as ^X or M-a")


func main()  {
	flag.Parse();
}
