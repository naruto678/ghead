package main

import (
	"flag"
	"fmt"
	"bufio"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		var line string
		for {
			fmt.Scanf("%s\n", line)
			fmt.Println(line)
		}
	}
	var (
		c           int
		bytes       int
		n           int
		lines       int
		showHelp    bool
		showVersion bool
	)
	flag.IntVar(&c, "c", -1, "print the first num bytes")
	flag.IntVar(&bytes, "bytes", -1, "print the first num bytes")
	flag.IntVar(&n, "n", -1, "print the first num lines")
	flag.IntVar(&lines, "lines", -1, "print the first num lines")
	flag.BoolVar(&showVersion, "version", false, "output version information and exit")
	flag.BoolVar(&showHelp, "help", false, "display this and exit")
	flag.Parse()

        if showHelp {
                flag.Usage()
                os.Exit(0)
        }
        if showVersion {
                fmt.Println("Version", 1.0)
                os.Exit(0)
        }

        if len(flag.Args()) == 0 {
                flag.Usage()
                fmt.Println("filename not specified")
                os.Exit(-1)
                
        }

        for _, fileName := range flag.Args() {
                fmt.Println(fileName)
                if c!=-1 {
                        printBytes(c, fileName)
                } else if bytes!=-1 {
                        printBytes(bytes, fileName)
                } else if n!=-1 {
                        printLines(n, fileName)
                } else if lines !=-1 {
                        printLines(lines, fileName)
                }
        } 
}

func printBytes(num int, fileName string){
        buffer := make([]byte, num)
        file, err := os.Open(fileName)
        defer file.Close()
        if err!=nil {
                fmt.Println(err)
                return 
        }
        if _ , err := file.Read(buffer); err==nil {
                fmt.Println(string(buffer))
        } else {
                fmt.Println(err)
        }
}

func printLines(num int, fileName string){
        file, err := os.Open(fileName)
        if err!=nil {
                fmt.Println(err)
                return 
        }
        defer file.Close()
        scanner := bufio.NewScanner(file)
        for num > 0 && scanner.Scan() {
                line := scanner.Text()
                fmt.Println(line)
                num--
        }
        if err = scanner.Err(); err!=nil {
                fmt.Println(err)
                return 
        }
}
