package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	//func ReadFile(filename string) ([]byte, error)
	by, err := ioutil.ReadFile("/proc/net/tcp")
	if err != nil {
		fmt.Println(err)
		return
	}
	//out := hex.Dump(by)
	reader := bytes.NewReader(by)
	scan := bufio.NewScanner(reader)
	var count int
	for scan.Scan() {
		count++

		vals := scan.Text()
		//fmt.Println(vals)
		if count == 1 {
			fmt.Println(vals)
			continue
		}
		//fmt.Println(hex.Dump([]byte(vals)))
		fields := strings.Fields(vals)
		fmt.Printf("%s ", fields[0])
		fields = fields[1:]
		for _, val := range fields {
			//v :=fmt.Sprintf("%+x", val)
			if strings.Index(val, ":") > -1 {
				dd := strings.Split(val, ":")
				if len(dd) > 1 && len(dd[1]) == 4 { //ip and port
					ip, err := hex.DecodeString(dd[0])
					if err != nil {
						fmt.Printf("%s ", dd[0])
						continue
					}
					fmt.Printf("%v.%v.%v.%v:", ip[3], ip[2], ip[1], ip[0])

					port, _ := strconv.ParseInt("0x"+dd[1], 0, 64)
					fmt.Printf("%d ", port)
					continue
				}
				for i, data := range dd {
					number, _ := strconv.ParseInt("0x"+data, 0, 64)
					if i == 0 {
						fmt.Printf("%d:", number)
					} else {
						fmt.Printf("%d ", number)
					}
				}
				fmt.Printf(" ")
				continue
			}
			//out, err := hex.DecodeString(val)
			//if err != nil {
			//	fmt.Printf("%s ", val)
			//	continue
			//}
			//fmt.Printf("%v ", out)
			out, _ := strconv.ParseInt("0x"+val, 0, 64)
			fmt.Printf("%d  ", out)
		}
		fmt.Printf("\n")
	}
}
