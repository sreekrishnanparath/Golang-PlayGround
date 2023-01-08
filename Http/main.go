package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Box struct {
    headers  string
    url string
}

func main() {

	resp, err := http.Get("https://httpbin.org/get")
	if(err != nil){
		fmt.Println("Error : ", err)
	}
	fmt.Println("Response : ", resp.Body)

	//body, err := ioutil.ReadAll(resp.Body)
   	if err != nil {
		fmt.Println("Error : ", err)
   	}
	a, _ := io.Copy(os.Stdout, resp.Body)
   	//sb := string(body)
	fmt.Println(a)
}