package utils

import (
	"bytes"
	"fmt"
)

func HexByte(data []byte) []byte {
	var buffer bytes.Buffer
	for i := 0; i < len(data); i++ {
		buffer.WriteString("\\x")
		buffer.WriteString(string("0123456789abcdef"[data[i]>>4]))
		buffer.WriteString(string("0123456789abcdef"[data[i]&0xf]))
	}
	return buffer.Bytes()
}

func HexDump(data []byte) []byte {
	var buffer bytes.Buffer
	var i int
	for i = 0; i < len(data); i++ {
		if i > 0 && i%16 == 0 {
			buffer.WriteString("\n")
		} else if i > 0 && i%8 == 0 {
			buffer.WriteString("  ")
		}
		buffer.WriteString(fmt.Sprintf("%02X ", data[i]))
		if i%16 == 15 {
			buffer.WriteString(" ")

			for j := i - 15; j <= i; j++ {
				if data[j] >= 32 && data[j] <= 126 {
					buffer.WriteByte(data[j])
				} else {
					buffer.WriteString(".")
				}
			}
		}
	}
	if i > 0 && i%16 != 0 {
		for j := 0; j < 16-i%16; j++ {
			buffer.WriteString("   ")
		}
		buffer.WriteString(" ")
		if i%16 <= 8 {
			buffer.WriteString("  ")
		}
		for j := i - i%16; j < i; j++ {
			if data[j] >= 32 && data[j] <= 126 {
				buffer.WriteByte(data[j])
			} else {
				buffer.WriteString(".")
			}
		}
	}
	return buffer.Bytes()
}

func ProcessMessageErr(err error, message string) {
	if err != nil {
		fmt.Println("\033[31m\033[1m[x] Error: \033[0m", message, err.Error())
	}
}

func ShowMessageInfo(message string) {
	fmt.Println("\033[32m\033[1m[+] Info: \033[0m", message)
}
