二分查找作为程序员的一项基本技能，是面试官最常使用来考察程序员基本素质的算法之一，也是解决很多查找类题目的常用方法，它可以达到O(log n)的时间复杂度。

【重点】一般在出现以下特性时考虑二分查找

1. 【有序】待查找的数组有序或者部分有序
2. 【时间复杂度】要求时间复杂度低于O(n)，或者直接要求时间复杂度为O(log n)

需要注意的是一旦数据中有重复元素，使用二分查找法返回的元素下标可能不是唯一的

## 写法一
定义 target 是在一个在左闭右闭的区间里，也就是`[left, right]`

区间的定义这就决定了二分法的代码应该如何写，因为定义target在`[left, right]`区间，所以有如下两点：

`while (left <= right)` 要使用 <= ，因为`left == right`是有意义的，所以使用 <=
if (nums[middle] > target) right 要赋值为 middle - 1，因为当前这个nums[middle]一定不是target，那么接下来要查找的左区间结束下标位置就是 middle + 1

**mid一致没找到，最后的l和r，对应的是**
1. l:比target大的位置
2. r:比target小的位置

## 写法二

如果说定义 target 是在一个在左闭右开的区间里，也就是[left, right) ，那么二分法的边界处理方式则截然不同。

while (left < right)，这里使用 < ,因为left == right在区间[left, right)是没有意义的
if (nums[middle] > target) right 更新为 middle，因为当前nums[middle]不等于target，去左区间继续寻找，而寻找区间是左闭右开区间，所以right更新为middle，即：下一个查询区间不会去比较nums[middle]
if (nums[middle] < target) left 更新为 middle + 1，因为左侧与右侧不同，是闭区间，因为当前这个nums[middle]一定不是target，直接指向middle + 1即可

## Go中自带的二分查找

```go
// 找出target在a切片中的索引下标，如果没有就返回要插入的地方(大于等于target的最小索引)
sort.SearchInts(a,target)

// 返回 [0， n) 中 f(i) 为真的最小索引 i
// 即：满足函数的最小位置
// 第一个参数是切片的长度，第二个是要满足的函数
sort.Search(len(a), func(i int) bool {})
```
