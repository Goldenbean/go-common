package common

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// HTTPPost :
func HTTPPost(url string, data []byte) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	respx, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer respx.Body.Close()

	fmt.Println("Response Status:", respx.Status)
	bodyx, _ := ioutil.ReadAll(respx.Body)
	fmt.Println("Response Body:", string(bodyx))
}

// HTTPPostWithHeader :
func HTTPPostWithHeader(url string, data []byte, headers map[string]string) []byte {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	for k, v := range headers {
		//req.Header.Set("Content-Type", "application/json")
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	respx, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer respx.Body.Close()

	fmt.Println("response Status:", respx.Status)
	body, _ := ioutil.ReadAll(respx.Body)
	fmt.Println("response Body:", string(body))
	return body
}

// HTTPGet :
func HTTPGet(url string) string {

	res, err := http.Get(url)
	if err != nil {
		return ""
		//panic(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
		// handle error
	}

	//	fmt.Println(string(body))
	//	fmt.Println(res.Status)

	return string(body)
}

// HTTPGetBinary :
func HTTPGetBinary(url string) []byte {

	if url == "" {
		return nil
	}

	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		//return ""
		panic(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// return ""
		panic(err.Error())
	}

	//	fmt.Println(string(body))
	//	fmt.Println(res.Status)

	return body
}

// HTTPGetBinaryToFile :
func HTTPGetBinaryToFile(url string, path string) (string, bool) {

	fmt.Printf("url: [%s], path: [%s]\n", url, path)

	res, err := http.Get(url)
	if err != nil {
		return err.Error(), false
	}
	defer res.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return "", true
}

// HTTPGetWithHeaders :
func HTTPGetWithHeaders(url string, headers map[string]string) string {

	fmt.Println("HTTPGetWithHeaders: ", url)

	req, err := http.NewRequest("GET", url, nil)

	for k, v := range headers {
		//req.Header.Set("Content-Type", "application/json")
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return ""
		// panic(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
		// handle error
	}
	ret := string(body)
	fmt.Println(ret)
	fmt.Println(res.Status)

	return ret
}

// GetIP :
func GetIP() []string {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	ret := []string{}
	for _, addr := range addrs {
		// fmt.Println(addr.String())
		match, _ := regexp.MatchString(`^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+/[0-9]+$`, addr.String())
		if !match {
			continue
		}
		slit := strings.Split(addr.String(), "/")
		ret = append(ret, slit[0])
	}

	return ret
}

// PrintIP :
func PrintIP() {

	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			match, _ := regexp.MatchString(`^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+/[0-9]+$`, addr.String())
			if !match {
				continue
			}
			slit := strings.Split(addr.String(), "/")
			fmt.Println(i.Name, " , ", i.Flags, " , ", slit)
		}
	}
}
