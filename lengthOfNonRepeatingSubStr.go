package main

import (
	"testing"
)




func TestLengthOfNonRepeatingSubStr(T *testing.T){
	s:="黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans:=8


		actual := lengthOfNonRepeatingSubStr(s)
		if actual!=ans{
			T.Error("cuole")
		}
	}

