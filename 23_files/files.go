package main

import (
	"fmt"
	"os"
)

func main() {
	// read the file
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// defer f.Close()
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }
	// fmt.Println("file name", fileInfo.Name())
	// fmt.Println("file size", fileInfo.Size())
	// fmt.Println("file size", fileInfo.Mode())
	// fmt.Println("file time", fileInfo.ModTime())

	// buf := make([]byte, 11)
	// d, err := f.Read(buf)

	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	fmt.Println("data", d, string(buf[i]))
	// }

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	// read folder

	// dir, err := os.Open("../")

	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileSync, err := dir.ReadDir(-1)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, file := range fileSync {
	// 	fmt.Println(file.Name())
	// }

	// Create a file

	// file, err := os.Create("example2.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()
	// file.WriteString("hlo")
	// file.WriteString("how are you")
	// bytes := []byte("Hello Golang")
	// file.Write(bytes)

	// read from one file and write to another

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(err)
	// 	}

	// }

	// writer.Flush()
	// fmt.Println("File written successfully")

	// delete file

	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File deleted successfully")

}
