
字符串这块一个重点是KMP，一般情况下可以尝试双指针去做，也可以考虑直接用第三方库

首先记住些常用的库和方法
1. `[]byte(s)` 字符串转换为字符切片
2. `string(s)`字符切片转换为字符串
3. `strings`库下的常用方法`import "strings"`
   + `Replace(s, old, new string, n int) string` n代表替换的个数，小于0则全部替换
   + `Index(s, substr string) int` 返回s中substr的第一个实例的索引，如果s中没有substr，则返回-1
   + `Count(s, substr string) int`
   + `Contains(s, substr string) bool` (Count和Contains底层都是调用的Index方法)
   + 里面对应有区分rune byte any，可以根据需有去标准库里面找对应函数

KMP (字符串匹配算法)

直接调用库的话就是`Index(s, substr string) int`
