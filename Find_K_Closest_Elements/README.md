### 658. Find K Closest Elements

 **解題思路**：
 這題的目標是找到一組陣列大小為給定的 K，而且裡面的數字必須是最接近 X 的值
 從題目可以知道這個陣列是由小排到大，所以我們可以從最左 (最小值) 跟最右 (最大值)
 來進行哪邊的值是最近 X，直到這兩個範圍內的陣列大小剛好等於 K 時，這個陣列裡面的值都是最接近 X 的值。



[LeetCode Link](https://leetcode.com/problems/find-k-closest-elements/)