package combine

import "testing"

func TestHelloWorld(t *testing.T) {
	f, d, c := Combine([]string{"1.wav", "2.wav"})
	WriteToFile("combine.wav", f, d, c)
}

// Go file
/***********************************************
# Copyright (c) 2020, Shanghai
# All rights reserved.
#
# @Filename: combineWav_test.go
# @Version：V1.0
# @Author: line - line.xj.lin@gmail.com
# @Description: ---
# @Create Time: 2020-07-01 11:19:23
# @Last Modified: 2020-07-01 11:19:23
***********************************************/
