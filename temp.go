package main

import (
	"bufio"
	"fmt"
	"os"
    "time"
)
import s "strings"

func atoi(s string) int {
	ret := 0
	arr := []byte(s)
	var temp byte
	for i := 0; i < len(arr); i++ {
		ret *= 10
		temp = arr[i]
		ret += int(temp - byte('0'))
	}
	return ret
}

func compare(a, b int) int {
	if a > b {
		return 1
	}
	if a < b {
		return 2
	}
	if a == b {
		return 3
	}
	return 0
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//var data string
	stack := make([]int, 0)
	resStack := make([]int, 0)
	var dat string
	var t int
	//var ins []string
	var prog []string
	var cmp int
	for scanner.Scan() {
		prog = append(prog, scanner.Text())
	}
	for i := 0; i < len(prog); i++ {
		dat = prog[i]
		tmp := []rune(dat)
		if s.HasPrefix(dat, "push") {
			val := []rune(dat)
			op := string(val[5:])
			stack = append(stack, atoi(op))
		} else if dat == ">"{
            fmt.Scan(&t)
            stack = append(stack, t)
        } else if dat == "$" {
            if len(stack) < 1 {
                fmt.Println("Not enough elements in stack to perform the operation on line ", i+1)
                return
            }
			fmt.Println(stack[len(stack)-1])
		} else if dat == "add" {
            if len(stack) < 1 {
                fmt.Println("Not enough elements in stack to perform the operation on line ", i+1)
                return
            }
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		} else if dat == "$$" {
            if len(stack) < 1 {
                fmt.Println("Not enough elements in Rstack to perform the operation on line ", i+1)
                return
            }
			fmt.Println(resStack[len(resStack)-1])
		} else if dat == "mul" {
            if len(stack) < 1 {
                fmt.Println("Not enough elements in stack to perform the operation on line ", i+1)
                return
            }
			stack = append(stack, stack[len(stack)-1]*stack[len(stack)-2])
		} else if dat == "div" {
            if len(stack) < 1 {
                fmt.Println("Not enough elements in stack to perform the operation on line ", i+1)
                return
            }
			stack = append(stack, stack[len(stack)-1]/stack[len(stack)-2])
		} else if dat == "mod"{
            if len(stack) < 1 {
                fmt.Println("Not enough elements in stack to perform the operation on line ", i+1)
                return
            }
            stack = append(stack, stack[len(stack)-1]%stack[len(stack)-2])
        } else if dat == "sub" {
            if len(stack) < 1 {
                fmt.Println("Not enough elements in stack to perform the operation on line ", i+1)
                return
            }
			stack = append(stack, stack[len(stack)-1]-stack[len(stack)-2])
		} else if dat == "swap" {
            if len(stack) < 1 || len(resStack) < 1{
                fmt.Println("Not enough elements in stack to perform the operation on line ", i+1)
                return
            }
			t = stack[len(stack)-1]
			stack[len(stack)-1] = resStack[len(resStack)-1]
			resStack[len(resStack)-1] = t
		} else if dat == "movRS" {
            if len(stack) < 1 {
                fmt.Println("Not enough elements in Rstack to perform the operation on line ", i+1)
                return
            }
			stack = append(stack, resStack[len(resStack)-1])
		} else if dat == "movS" {
            if len(stack) < 1 {
                fmt.Println("Not enough elements in stack to perform the operation on line ", i+1)
                return
            }
			resStack = append(resStack, stack[len(stack)-1])
		} else if dat == "delS" {
            if len(stack) < 1 {
                fmt.Println("Not enough elements in stack to perform the operation on line ", i+1)
                return
            }
			stack = stack[:len(stack)-1]
		} else if dat == "delRS" {
            if len(stack) < 1 {
                fmt.Println("Not enough elements in Rstack to perform the operation on line ", i+1)
                return
            }
			resStack = resStack[:len(resStack)-1]
		} else if string(tmp[:5]) == "goto " {
			val := []rune(dat)
			op := string(val[5:])
			if atoi(op) >= len(prog) {
				fmt.Println("Bad Pointer line ", i+1)
				return
			}
			i = atoi(op)-2
		} else if dat == "comp" {
			if len(stack) < 2 {
				fmt.Println("Too small stack to compare on line ", i+1)
				return
			}
			cmp = compare(stack[len(stack)-1], stack[len(stack)-2])
		} else if s.HasPrefix(dat, "gotoM") {
			if cmp != 1 {
				continue
			}
			val := []rune(dat)
			op := string(val[6:])
			if atoi(op) >= len(prog) {
				fmt.Println("Bad Pointer line ", i)
				return
			}
			i = atoi(op)-2
		} else if s.HasPrefix(dat, "gotoL") {
			if cmp != 2 {
				continue
			}
			val := []rune(dat)
			op := string(val[6:])
			if atoi(op) >= len(prog) {
				fmt.Println("Bad Pointer line ", i)
				return
			}
			i = atoi(op)-2
		} else if s.HasPrefix(dat, "gotoE") {
			if cmp != 3 {
				continue
			}
			val := []rune(dat)
			op := string(val[6:])
			if atoi(op) >= len(prog) {
				fmt.Println("Bad Pointer line ", i)
				return
			}
			i = atoi(op)-2
		} else if s.HasPrefix(dat, "gotoNE") {
			if cmp != 3 {
				val := []rune(dat)
				op := string(val[7:])
				if atoi(op) >= len(prog) {
					fmt.Println("Bad Pointer line ", i)
					return
				}
				i = atoi(op)-2
			}
		} else if dat == "inc" {
            if len(stack) >= 1 {
                stack[len(stack)-1]++
            } else {
                fmt.Println("Not enough items to perform incrementation on line ", i)
            }
        } else if dat == "dec" {
            if len(stack) >= 1 {
                stack[len(stack)-1]--
            } else {
                fmt.Println("Not enough items to perform decrementation on line ", i)
            }
        } else if dat == "end" {
            goto end
        } else if dat == "wait" {
            time.Sleep(500*time.Millisecond)
        }
	}
    end:
	fmt.Println("Success")
}
