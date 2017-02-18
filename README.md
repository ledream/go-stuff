# go-stuff
## Features
###Cat1.go
Cat1.go is an exercise that reproduces the functionality of the Unix command "cat" using the Go language.
- Read a file
- Read multiple files
- -b  -- number non-blank output lines
- -e  -- display $ at the end of each line (implies -v)
- -n  -- number all output lines
- -s  -- squeeze multiple blank lines into one
- -t  -- display tab as ^I (implies -v)
- -u  -- do not buffer output
- -v  -- display non-printing chars as ^X or M-a

###Librairies
- fmt
- os
- io/ioutil
- flag

##Instalation
```
git clone https://github.com/ledream/go-stuff.git
```

##Usage
- Read a file:
```
go run cat1.go samples/file1.txt
```
- Read multiple files :
```
go run cat1.go samples/file1.txt samples/file2.txt
```
