package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return bool布尔型
 */
func isValid(s string) bool {
	// write code here
	length := len(s)
	if length%2 != 0 {
		return false
	}

	stack := make([]int32, 0)
	stackTop := 0

	for _, char := range s {
		// 题目保证只有六种字符
		if char == '{' || char == '[' || char == '(' {
			if stackTop == len(stack) {
				stack = append(stack, char)
			} else {
				stack[stackTop] = char
			}
			stackTop++
		} else {
			// 先判断栈里有没有左括号，没有的话直接报错
			if stackTop == 0 {
				return false
			}
			var matchFlag bool = false
			switch char {
			case '}':
				if stack[stackTop-1] == '{' {
					matchFlag = true
				}
			case ']':
				if stack[stackTop-1] == '[' {
					matchFlag = true
				}
			case ')':
				if stack[stackTop-1] == '(' {
					matchFlag = true
				}
			}

			if matchFlag == false {
				return false
			} else {
				stackTop--
			}
		}
	}

	// 栈为空才算通过
	if stackTop == 0 {
		return true
	}
	return false
}
