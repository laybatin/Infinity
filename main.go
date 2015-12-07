// Infinity project main.go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 *  bigInteger는 int16의 슬라이스로 구성되어 값을 표현한다.
 *  12345678   을 사용자가 입력하면
 *   1234       5678        String
 *  value[1]    value[0] 에 저장되도록 구성된다
 *  연산 시 슬라이스의 0번 요소부터 덧셈하는게 올림, 내림의 구현이 쉽기 때문에 위와 같이 구현한다
 *  TODO : 슬라이스 사이즈에 맞게 New를 할 수 있도록 구현
 */
type bigInteger struct {
	value []int16
	sign  int //음수면 -1, 0이면 0, 양수면 1
}

func StringToUin164Slice(str string, array []int16) {
	var i, j int
	str_length := len(str)
	j = str_length

	//맨 뒤에서 네자리씩 잘라서 넣는다.
	for ; i < str_length/4; i, j = i+1, j-4 {
		value, _ := strconv.Atoi(str[j-4 : j])
		array[i] = int16(value)
	}

	//만약 문자열의 길이가 남았다면 해당 나머지 문자열을 복사한다.
	if str_length%4 != 0 {
		value, _ := strconv.Atoi(str[:j])
		array[i] = int16(value)
	}
}

func InvalidCheck(s string) bool {
	for _, c := range s {
		if !(c-'0' >= 0 && c-'0' <= 9) {
			return false
		}
	}
	return true
}

func Abs(value int16) int16 {
	if value < 0 {
		return -1 * value
	}
	return value
}

func NewBigIntegerSmallNum(c string) *bigInteger {
	slice := make([]int16, 1)
	intValue, _ := strconv.Atoi(c)
	value := int16(intValue)
	sign := 0
	if value < 0 {
		sign = -1
	} else if value == 0 {
		sign = 0
	} else {
		sign = 1
	}

	slice[0] = Abs(value)
	return &bigInteger{slice, sign}
}

func NewBigInteger(c string) *bigInteger {
	if len(c) == 0 {
		panic("값이 없습니다.")
	}

	if c == "0" {
		return NewBigIntegerSmallNum(c)
	}

	//첫번째 문자가 0이라면 00001 이렇게 입력했을 수 있음
	//그래서 왼쪽을 모두 지워줌
	if c[0] == '0' {
		c = strings.TrimLeft(c, "0")
		fmt.Println(c)
	}

	var sign, checkRange int
	length := len(c)
	if length == 0 {
		return NewBigIntegerSmallNum("0")
	}

	if c[0] == '-' {
		sign = -1
		checkRange = 1
		length = length - 1
	} else if c[0] == '0' {
		sign = 0
	} else {
		sign = 1
	}

	//숫자가 맞는지 체크
	isNumber := InvalidCheck(c[checkRange:])
	if isNumber == false {
		fmt.Println("숫자가 아님요")
		return nil
	}

	//Slice Make
	value := make([]int16, length/4+1)
	StringToUint16slice(c[checkRange:], value)

	return &bigInteger{value, sign}
}

func NewBigIntegerBySlice(intSlice []int16, sign int) (b *bigInteger) {
	return &bigInteger{intSlice, sign}
}

func StringToUint16slice(str string, array []int16) {
	var i, j int
	str_length := len(str)
	j = str_length

	for ; i < str_length/4; i, j = i+1, j-4 {
		value, _ := strconv.Atoi(str[j-4 : j])
		array[i] = int16(value)
	}

	if str_length%4 != 0 {
		value, _ := strconv.Atoi(str[:j])
		array[i] = int16(value)
	}
}

//00000000000 을 0000 으로 만들어주기 위해 구현
func SliceLeftTrim(x []int16) []int16 {
	if len(x) == 0 {
		return nil
	}

	for idx, value := range x {
		if value != 0 {
			return x[idx:]
		}
	}
	value := make([]int16, 1)
	return value
}

func (b *bigInteger) AddByString(c string) {
	if b.value == nil {
		panic("value is nil")
	}

	if c == "" {
		panic("argument length zero")
	}
}

func AddBySlice(x, y []int16) *bigInteger {
	var bigSlice []int16
	var smallSlice []int16
	bigLength := 0
	smallLength := 0

	if len(x) > len(y) {
		bigSlice = x
		smallSlice = y
		bigLength = len(x)
		smallLength = len(y)
	} else {
		bigSlice = y
		smallSlice = x
		bigLength = len(y)
		smallLength = len(x)
	}

	var ret = make([]int16, bigLength)

	for i, value := range bigSlice {
		if smallLength <= i {
			ret[i] = ret[i] + value
		} else {
			ret[i] = ret[i] + value + smallSlice[i]
			if ret[i] >= 10000 {
				ret[i] = ret[i] - 10000
				ret[i+1] += 1
			}
		}
	}

	return &bigInteger{ret, 1}
}

//덧셈에 대한 절대값
func (x *bigInteger) Add(y *bigInteger) *bigInteger {
	if x.sign == 0 {
		return y
	}

	if y.sign == 0 {
		return x
	}

	//부호가 같다면
	if x.sign == y.sign {
		//새 메모리를 생성해서 반환
		//덧셈연산 후 부호를 붙인다.
		temp := AddBySlice(x.value, y.value)
		temp.sign = x.sign
		return temp
	}

	var bigSlice []int16
	var smallSlice []int16
	bigLength := 0
	smallLength := 0

	if len(x.value) > len(y.value) {
		bigSlice = x.value
		smallSlice = y.value
		bigLength = len(x.value)
		smallLength = len(y.value)
	} else {
		bigSlice = y.value
		smallSlice = x.value
		bigLength = len(y.value)
		smallLength = len(x.value)
	}

	var ret = make([]int16, bigLength)

	for i, value := range bigSlice {
		if smallLength <= i {
			ret[i] = ret[i] + value
		} else {
			ret[i] = ret[i] + value + smallSlice[i]
			if ret[i] >= 10000 {
				ret[i] = ret[i] - 10000
				ret[i+1] += 1
			}
		}
	}

	return &bigInteger{ret, x.CompareTo(y)}
}

func (x *bigInteger) ToString() string {
	if x.sign == 0 {
		return "0"
	}

	retString := []string{}

	if x.sign == -1 {
		retString = append(retString, "-")
	}

	s := x.value
	if s[len(s)-1] != 0 {
		firstDigit := fmt.Sprintf("%d", s[len(s)-1])
		retString = append(retString, firstDigit)
	}

	for i := len(s) - 2; i >= 0; i-- {
		digit := fmt.Sprintf("%04d", s[i])
		retString = append(retString, digit)
	}

	return strings.Join(retString, "")
}

//TODO 큰수를 찾는 부분 구현
//     절대값 후 비교부분을 구현
//     -555555 + 555555 = 0
//    매개변수가 크면 1
//    같으면 0
//    this가 크면 -1
func (a *bigInteger) CompareTo(b *bigInteger) int {
	//value is zero
	if a.sign == 0 && b.sign == 0 {
		return 0
	}

	if a.sign == 1 && b.sign == -1 {
		return 1
	}

	if a.sign == -1 && b.sign == 1 {
		return -1
	}

	//sign Equals
	cmpValue := int16SliceCmp(a.value, b.value)

	//sign minus
	if a.sign == -1 && b.sign == -1 {
		return -1 * cmpValue
	}

	return cmpValue

}

func int16SliceCmp(args1, args2 []int16) int {

	a := SliceLeftTrim(args1)
	b := SliceLeftTrim(args2)

	if len(a) > len(b) {
		return -1
	} else if len(a) < len(b) {
		return 1
	}

	bigSlice := a
	smallSlice := b


	for idx := range bigSlice {
		if bigSlice[len(bigSlice) -1 -idx] > smallSlice[len(smallSlice) -1 -idx] {
			return -1
		} else if bigSlice[len(bigSlice) -1 -idx] < smallSlice[len(smallSlice) -1 -idx] {
			return 1
		}
	}

	return 0
}

func (x *bigInteger) Sub(y *bigInteger) *bigInteger {
	if x.sign == 0 {
		return y
	}

	if y.sign == 0 {
		return x
	}

	cmpValue := x.CompareTo(y)
	
	if x.sign != y.sign {
		if cmpValue == 0 {
			return NewBigIntegerSmallNum("0")
		}
	}
	
	//Sing Equals
	retValue := x.Add(y)
	return retValue
	
}

func PrintData(printValue *bigInteger) {
	if printValue.sign == -1 {
		fmt.Printf("-")
	}

	s := printValue.value
	if s[len(s)-1] != 0 {
		fmt.Printf("%d", s[len(s)-1])
	}
	for i := len(s) - 2; i >= 0; i-- {
		fmt.Printf("%04d", s[i])
	}
}

func main() {
	var xValue, yValue, zValue string
	var bigArr []bigInteger
	
	fmt.Scanf("%s", &xValue)
	fmt.Scanf("%s", &yValue)
	fmt.Scanf("%s", &zValue)
	//fmt.Scanf("%s", &eqValue)

	num1 := NewBigInteger(xValue)
	num2 := NewBigInteger(yValue)
	num3 := NewBigInteger(zValue)
	//eqNum := NewBigInteger(eqValue)

	//fmt.Printf("%s\n", PrintData(num1))
	//fmt.Printf("%s\n", PrintData(num2))

	PrintData(num1)
	fmt.Printf("[%s]\n", num1.ToString())
	fmt.Println()
	PrintData(num2)
	fmt.Printf("[%s]\n", num2.ToString())
	fmt.Println()

//	ret := num1.Add(num2)
//	PrintData(ret)
//	fmt.Println()

	fmt.Printf("ret가 [%d]합니다\n", num1.CompareTo(num2))
	
	
	bigArr = append(bigArr, *num1)
	bigArr = append(bigArr, *num2)
	bigArr = append(bigArr, *num3)
	
}
