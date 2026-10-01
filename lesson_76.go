/*
76. Упражнение: Stringer
Заставьте тип IPAddr реализовать fmt.Stringer, чтобы адрес выводился в виде точечной четверки.

Например, IPAddr{1, 2, 3, 4} должен выводиться как “1.2.3.4”
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (i IPAddr) String() string {
	str := []string{}
	for _, vl := range i {
		str = append(str, strconv.Itoa(int(vl)))
	}
	return strings.Join(str, ".")
}

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
