package minor

import (
	"fmt"
	"strings"
	"testing"
)

func IsFloatStr(str string) int {
	return len(strings.Split(str, "."))
}

//func BigNumberAdd(strA, strB string) {
//	if IsFloatStr(strA) > 2 || IsFloatStr(strB) > 2 {
//		panic("字符串非法")
//	}
//	if IsFloatStr(strA) == 1 {
//		strA = strings.Join([]string{strA, ".0"}, "")
//	}
//	if IsFloatStr(strB) == 1 {
//		strB = strings.Join([]string{strB, ".0"}, "")
//	}
//
//	var (
//		strASplit                                  = strings.Split(strA, ".")
//		strBSplit                                  = strings.Split(strB, ".")
//		strAInt, strAFloat                         = strASplit[0], strASplit[1]
//		strBInt, strBFloat                         = strBSplit[0], strBSplit[1]
//		nStrAInt, nStrAFloat, nStrBInt, nStrBFloat = len(strAInt), len(strAFloat), len(strBInt), len(strBFloat)
//		countInt, countfloat                       int
//		isCarryInt, isCarryFloat                   bool
//		res                                        []string
//	)
//
//	for nStrAInt-countInt-1 >= 0 && nStrBInt-countInt-1 >= 0 {
//		sum := int(strA[len(strA)-1-countInt] - '0' + strB[len(strB)-1-countInt] - '0')
//		if isCarryInt {
//			sum += 1
//		}
//		res = append([]string{fmt.Sprintf("%c", sum%10)}, res...)
//		if sum >= 10 {
//			isCarryInt = true
//		} else {
//			isCarryInt = false
//		}
//		countInt++
//	}
//
//	for nStrAInt-countInt-1 >= 0 || nStrBInt-countInt-1 >= 0 {
//		if nStrAInt-countInt-1 >= 0 {
//
//		} else if nStrBInt-countInt-1 >= 0 {
//
//		}
//		countInt++
//	}
//
//}

func TestBigNumberAdd(t *testing.T) {
	//strA := "12345.6789"
	//strB := "9876.54321"
	//BigNumberAdd(strA, strB)
	fmt.Println(len(strings.Split("123", ".")))
}
