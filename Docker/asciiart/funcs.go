package asciiart

import (
	"bufio"
	"os"
	S "strings"
)

var (
	_map  = make(map[int][8]string)
	lines = [8]string{}
)

func InitMap(banner string)bool {
	var file *os.File
	var err error

	switch banner {
		
	case "standard":
		file, err = os.Open("Banners/standard.txt")
	case "shadow":
		file, err = os.Open("Banners/shadow.txt")
	case "thinkertoy":
		file, err = os.Open("Banners/thinkertoy.txt")
	case "bubble":
		file, err = os.Open("Banners/bubble.txt")
	case "soft":
		file, err = os.Open("Banners/soft.txt")
	default :
		return false
	}

	if err != nil {
		return false
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	// insert data on the map
	for i := 32; i < 127; i++ {
		_map[i] = InsertValue(scanner)
	}

	defer file.Close()
	return true
}

func InsertValue(scanner *bufio.Scanner) [8]string {
	ArtValue := [8]string{}

	for cp := 0; cp < 8 && scanner.Scan(); cp++ {
		ArtValue[cp] = scanner.Text()
	}
	scanner.Scan()
	return ArtValue
}

func IsOnlyNewLine(str []string) bool {

	for _, v := range str {
		if v != "" {
			return false
		}
	}
	return true
}

func Storing(inp string) string {
	res := ""
	inp = S.Trim(inp, " ")
	spl := S.Split(inp, "\r\n")

	if IsOnlyNewLine(spl) {
		return ""
	}
	
	for _, val := range spl {
		for _, v := range val {
			for i := 0; i < 8; i++ {
				lines[i] += _map[int(v)][i]
			}
		}
		for i := 0; i < 8; i++ {
			res += lines[i] + "\n"
			lines[i] = ""
		}
	}

	return res
}

