package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//f, err := os.Create("text.txt")

	f, err := os.Open("F:\\IDM\\tailwind.png")

	if err != nil {
		panic("file not created")
	}

	writeToFile(f, []byte{'1', '2', '3'})

}

func writeToFile(reader io.Reader, data []byte) {

	k := make([]byte, 1024)
	//f, err := os.OpenFile("./text.txt", os.O_CREATE, 0644)
	/*	if err != nil {
		panic("text.file not found")
	}*/

	for {

		n, err := reader.Read(k)
		if err != nil {
			panic("failed read")
		}

		/*	fmt.Println(n)
			var encodedByte = make([]byte, 1024)

			eb := base64.Encoding.AppendEncode(encodedByte, k)
			writen, err2 := f.Write(encodedByte)

			if err2 != nil {
				panic(err2.Error())
			}
			fmt.Println("witten ", writen)*/

		fmt.Println(n)
		//fmt.Println(k)

		for _, v := range k[:n] {

			fmt.Print(string(v))

		}

		if err != nil {

			if err == io.EOF {
				fmt.Println("FILE ENDED")

				break
			}

			fmt.Println("SOMETHING HAPPEN")

			break

		}

	}

	/*	n, err := writer.Write(data)

		fmt.Println("writen to file", n)

		if err != nil {
			panic(err)
		}*/

}
