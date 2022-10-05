### 623. Add One Row to Tree
**解題思路**
找到指定的 depth，並把指定的 Value 塞到那個 depth，所以用遞迴的方式找到指定的 depth
新的 node 指向原本在那個 depth 的 node 的 child， 並把 node 的 child 指向新的 node 
需要考慮 depth 是 root 的情境