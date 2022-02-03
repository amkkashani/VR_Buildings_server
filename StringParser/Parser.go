package StringParser

import (
	"fmt"
	"regexp"
	"strings"
)

func StringParser(input string) (string, string, string, string) { // name x y z
	// GNNG N E H
	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.Contains(line, "location data") && strings.Contains(line, "position") {
			//its my log data
			regexName := regexp.MustCompile(`¦\w*`)
			name := string(regexName.Find([]byte(line)))
			regexVector := regexp.MustCompile(`=-?\d*`)
			vectorPos := (regexVector.FindAll([]byte(line), 3))
			fmt.Println("***")
			print(line)
			fmt.Println(name[1:], string(vectorPos[0][1:]), string(vectorPos[1][1:]), string(vectorPos[2][1:]))
			x := string(vectorPos[0][1:])
			y := string(vectorPos[1][1:])
			z := string(vectorPos[2][1:])
			return name[2:], x, y, z
		} else if strings.Contains(line, "$GNGGA") {
			n, e, a := "", "", ""
			strArr := strings.Split(line, ",")
			for i, str := range strArr {
				if str == "N" || str == "S" {
					n = strArr[i-1]
				} else if str == "E" || str == "W" {
					e = strArr[i-1]
				} else if str == "M" && a == "" { // hight
					a = strArr[i-1]
				}
			}
			return "$GNGGA", n, e, a
		}

	}
	return "", "", "", ""
}

//log example
//uality=100}, 00C9 distance=Distance{length=681, quality=100}, 0907 distance=Distance{length=895, quality=100}
//2708.041: 0xâ€¦0024 location data: position: x=251 y=88 z=104 q=53; distances: 21C0 distance=Distance{length=0, quality=100}, 1522 distance=Distance{length=460, quality=100}, 0907 distance=Distance{length=927, quality=100}, 00C9 distance=Distance{length=831, quality=100}
//2708.042: 0xâ€¦2D92 location data: position: x=686 y=0 z=18 q=67; distances: 1522 distance=Distance{length=0, quality=100}, 21C0 distance=Distance{length=362, quality=100}, 00C9 distance=Distance{length=770, quality=100}, 0907 distance=Distance{length=887, quality=100}
