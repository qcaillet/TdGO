package main 

import "testing"

func TestAddition(t *testing.T){
	var  funcReturn int

	param1 := 2
	param2 := 7

	funcReturn = monAddition(param1, param2)
	if funcReturn <= 0 {
		t.Error("mnaddition return is <= 0")
	}

}

func TestSoustraction(t *testing.T){
	var  funcReturn int

	param1 := 14
	param2 := 7

	funcReturn = maSoustraction(param1, param2)
	if funcReturn <= 0 {
		t.Error("maSoustraction return is <= 0")
	}

}

func TestSoustractionKo(t *testing.T){	
	var  funcReturn int

	param1 := 13
	param2 := 7

	funcReturn = maSoustraction(param1, param2)
	if funcReturn != 7 {
		t.Errorf("funcReturn  incorrect, got: %d, want: %d.", funcReturn, 7)
	}

}

func TestMultiplication(t *testing.T){

	var  funcReturn int

	param1 := 15
	param2 := 7

	funcReturn = maMultiplication(param1, param2)
	if funcReturn <= 0 {
		t.Error("maMultiplication return is <= 0")
	}

}

func TestMultiplicationKo(t *testing.T){

	var  funcReturn int

	param1 := 15
	param2 := 7

	funcReturn = maMultiplication(param1, param2)
	if funcReturn <= 0 {
		t.Error("maMultiplication return is <= 0")
	}

}


func TestDivision(t *testing.T){

	var  funcReturn int

	param1 := 15
	param2 := 7

	funcReturn = maDivision(param1, param2)
	if funcReturn == 0 {
		t.Error("maDivision return is = 0")
	}

}

// func TestDivisionKo(t *testing.T){

// 	var  funcReturn int

// 	param1 := 16
// 	param2 := 7

// 	funcReturn = maDivision(param1, param2)
// 	if funcReturn != 2 {
// 		t.Errorf("funcReturn  incorrect, got: %d, want: %d.", funcReturn, 2)
// 	}

// }



