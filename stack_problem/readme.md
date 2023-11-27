
在用切片模拟栈时，注意事项
1. 取元素时，先判断栈是否为空否则`stack[len(stack)-1]`会成为`stack[-1]`然后报错
2. string类型for range遍历时，ch为int32类型，若栈为`map[byte]byte`需要转换
3. `var stack []byte`然后接着往里面append就行

需要取出前两个，做操作然后再放到stack中。这个过程要考虑能不能只取一个，另一个原地修改，例如：
```go
// good:取出一个直接处理
num := stack[len(stack)-1]
stack = stack[:len(stack)-1]
stack[len(stack)-1] += num

// bad:俩都取出来再放回去
sum := stack[len(stack)-1] + stack[len(stack)-2]
stack = stack[:len(stack)-2]
stack = append(stack, sum)
```

strconv库 
```go
// Itoa <==> FormatInt(int64(i), 10)
Itoa(i int) string
Atoi(s string) (int, error)
// 将i转换为base进制，作为通过string输出，base范围[2:36]
FormatInt(i int64, base int) string
```
