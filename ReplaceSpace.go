package leetcode

import "fmt"

func ReplaceSpace(src string)string{
	var tmp = make([]byte, 0)
	space := []byte{'%'. '2', '0'}
	for idx, val := range(src){
		if val == ' '{
			tmp = append(tmp, space...)
		}else{
			tmp = appemd(tmp, val)
		}
	}
	return string(tmp)
}