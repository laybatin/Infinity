package main

import "testing"

func TestAdd(t *testing.T) {

	num1 := NewBigInteger("91960556760854814937033526730808847279867183356182295073407492165")
	num2 := NewBigInteger("88308583794966823785307881355544062885439970862651882782669518103")
	expected := NewBigInteger("180269140555821638722341408086352910165307154218834177856077010268")

	result := num1.Add(num2)
	if result.CompareTo(expected) != 0 {
		t.Error("기대값과 계산값이 다릅니다.")
	}
}

var newTestData = []string{
	"1024",
	"10240123",
	"10242389",
	"-1",
}

type compareTestData struct {
	value1, value2 string
	retValue int
}
var newTestDataCompare = []compareTestData {
	{"0", "0", 0},
	{"0", "1", 1},
	{"1", "0", -1},
	{"102400", "102400", 0},
	{"102400", "102500", 1},
	{"102500", "102400", -1},
	{"000000", "000000", 0},
	{"91960556760854814937033526730808847279867183356182295073407492165", "88308583794966823785307881355544062885439970862651882782669518103", -1},
	{"180269140555821638722341408086352910165307154218834177856077010268", "180269140555821638722341408086352910165307154218834177856077010268", 0},
}

func TestNewBigInteger(t *testing.T) {

	var newBigInteger []bigInteger

	num1 := NewBigInteger("1024")
	num2 := NewBigInteger("10240123")
	num3 := NewBigInteger("10242389")
	num4 := NewBigInteger("-1")

	newBigInteger = append(newBigInteger, *num1)
	newBigInteger = append(newBigInteger, *num2)
	newBigInteger = append(newBigInteger, *num3)
	newBigInteger = append(newBigInteger, *num4)

	for idx, d := range newTestData {
		if d != newBigInteger[idx].ToString() {
			t.Errorf("결과값 [%s]가 기대값[%s]과 같지 않습니다.", d, newBigInteger[idx].ToString())
		}
	}
}

func TestCompareTo(t *testing.T) {
	
	for _, value := range newTestDataCompare {
		val1 := NewBigInteger(value.value1)
		val2 := NewBigInteger(value.value2)
		cmpValue := val1.CompareTo(val2)
		if value.retValue != cmpValue  {
			t.Errorf("결과값[%d]이 기대값[%d]와 다릅니다. [%s] [%s]\n", cmpValue, value.retValue, val1.ToString(), val2.ToString())
		}
	}
	
}