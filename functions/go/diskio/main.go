package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func BenchmarkWriteFile() {
	f, err := os.Create("/context/test.txt")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100000; i++ {
		f.WriteString("Benchmarking...\n")
	}

	f.Close()
}

func BenchmarkWriteFileBuffered() {
	f, err := os.Create("/context/test.txt")
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)

	for i := 0; i < 100000; i++ {
		w.WriteString("Benchmarking...\n")
	}

	w.Flush()
	f.Close()
}

func BenchmarkReadFile() {
	f, err := os.Open("/context/test.txt")
	if err != nil {
		panic(err)
	}

	b := make([]byte, 10)

	_, err = f.Read(b)
	for err == nil {
		_, err = f.Read(b)
	}
	if err != io.EOF {
		panic(err)
	}

	f.Close()
}

func BenchmarkReadFileBuffered() {
	f, err := os.Open("/context/test.txt")
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)

	_, err = r.ReadString('\n')
	for err == nil {
		_, err = r.ReadString('\n')
	}
	if err != io.EOF {
		panic(err)
	}

	f.Close()
}
func main() {
	fmt.Println("Writing /context/test.txt...")
	BenchmarkWriteFileBuffered()
	fmt.Println("Reading /context/test.txt...")
	BenchmarkReadFileBuffered()
}
