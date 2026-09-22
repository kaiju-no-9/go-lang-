package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("File Size:", fileInfo.ModTime())

	
   // reading the data and storing it in a  buffer ... 
   buff := make([]byte , fileInfo.Size())
     _ , err = f.Read(buff)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(buff))
    
	f, err = os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_ , err = f.WriteString(" he there i love you very much")
	if err != nil {
		panic(err)
	}

	fmt.Println("File created successfully")

	// transfaring data form one file to other in streaming faction 
	sourceFile , err := os.Open("example.txt")
	  if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destFile , err := os.Create("example3.txt")
	  if err != nil {
		panic(err)
	}
	defer destFile.Close()
     reader := bufio.NewReader(sourceFile)
	 writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		err = writer.WriteByte(b)
		if err != nil {
			panic(err)
		}
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
