package main

import "os"
import "fmt"
import "io/ioutil"

func checkArguments(arg []string) bool {
	if len(os.Args) < 2 {
		return false
	} else {
		return true
	}
}

func main() {
	//Test#1 : s'il y au moins un argument
	if checkArguments(os.Args) == true {
		for index := 1; index <= len(os.Args)-1; index++ {
			dat, err := ioutil.ReadFile(os.Args[index])
			//Test : s'il y a une erreur
			if err != nil {
				//Test#2 : s'il s'agit d'une erreur de permission
				if os.IsPermission(err) == false {
					fmt.Println("[" + os.Args[index] + " : no file or directory.]")
				} else {
					fmt.Println("[" + os.Args[index] + " : Permission denied]")
				}
			} else {
				fmt.Println("[Reading : " + os.Args[index] + "]")
				fmt.Println(string(dat))
			}
		}
	} else {
		fmt.Print("[you should write at least one argument]")
	}
}
