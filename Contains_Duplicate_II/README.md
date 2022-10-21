## 219. Contains Duplicate II
**解題思路**
給一個數字陣列與目標 K，找到符合陣列有符合以下邏輯的 **nums[i] == nums[j] and abs(i - j) <= k** 就 return true
用一個 Map 儲存出現過的數字與其最近一次出現的位置，如果再次出現就檢查其是否小於 k。