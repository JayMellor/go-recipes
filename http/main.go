package main

import "net/http"
import "fmt"
import "os"
import "io"

type logWriter struct{}

func main() {
	response, error := http.Get("https://www.google.com")
	if error != nil {
		fmt.Println("Error from get", error)
		os.Exit(1)
	}

	// Manual approach
	// var body = make([]byte, 99999)
	// bytesRead := response.Body.Read(body)
	// fmt.Println(bytesRead, string(body[:bytesRead]))

	// Using io.Copy
	logWriter := logWriter{}
	io.Copy(logWriter, response.Body)

}

func (logWriter) Write(bytes []byte) (numWritten int, error error) {
	fmt.Println(string(bytes))
	fmt.Printf("Wrote %v bytes\n", len(bytes))
	return len(bytes), nil
}
