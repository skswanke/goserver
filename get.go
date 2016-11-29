package get

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://http-methods.appspot.com/Hungary/?v=true", nil)
	if err != nil {
		log.Fatal(err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resBody := res.Body
	buff := new(bytes.Buffer)
	buff.ReadFrom(resBody)
	fmt.Println(buff.String())
	resBody.Close()
	if res.StatusCode == 404 {
		fmt.Println(res.Status)
	}
}
