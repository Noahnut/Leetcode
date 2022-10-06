### 981. Time Based Key-Value Store

**解題思路**
這題是要建立一個有時間順序的 Key-Value store，每個塞入的 Key-Value 都會有時間戳，這個時間戳是遞增的
而查找時也會帶一個時間戳，找到這個時間戳以前最近時間的值，這樣代表會有一個保證排序的 Array，這邊就要想到可以用 Binary Search 
來做時間戳的查詢，而每個 Key 都有一個對應的 Array 來存著每個時間戳塞入的字串。