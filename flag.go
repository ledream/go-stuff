package main
import "flag"

var helpMessage = "-help -- for help"

var help = flag.String("help", "", helpMessage)
var file = flag.String("file", "", helpMessage)

func main()  {
	flag.Parse();
}
