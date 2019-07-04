##Next Permutation

###题目
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
// 将排列中的数字重新排列成字典序中的下一个更大的排列

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).
// 若这样的重新排列是不可能的,那必须重新排列为可能的最低顺序(即升序排序)

The replacement must be in-place and use only constant extra memory.
// 重排必须在原地进行,不分配额外的内存

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2

3,2,1 → 1,2,3

1,1,5 → 1,5,1