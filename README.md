# golang_permutation_in_string

Given two strings `s1` and `s2`, return `true` *if* `s2` *contains a permutation of* `s1`*, or* `false` *otherwise*.

In other words, return `true` if one of `s1`'s permutations is the substring of `s2`.

## Examples

**Example 1:**

```
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").

```

**Example 2:**

```
Input: s1 = "ab", s2 = "eidboaoo"
Output: false

```

**Constraints:**

- `1 <= s1.length, s2.length <= 104`
- `s1` and `s2` consist of lowercase English letters.

## 解析

給定兩個字串 s1 , s2

定義如果字串 t 是 字串 s 的 permutation 

則 t, s 的字元組成相同，也就是 s, t 中的字元種類相同且個數相同

![](https://i.imgur.com/SFfXQvS.png)

要求寫一個演算法判斷 s2 中有沒有包含 s1 字串的 permutation

如果 s2 要包含 s1 的 permutation

代表 s2 長度一定要 ≥ s1 長度

當 s1 長度大於 s2 則回傳 false

如果 s2 有 s1 的permutation 代表他的字元出現頻率跟 s1 一樣且字元彼此需要相鄰

字元頻率可以透過 hashTable 來達成

而相鄰的部份則可以透過 slide window 來做限制

建立一個 HashTable freq

先把 s1 的所有字元出現次數蒐集在 freq

初始化 left = 0, count = len(s1)

從 right = 0..len(s2)-1 做以下運算

當 freq[s[right]] > 0 時，更新 count = count-1

更新 freq[s[right]] = freq[s[right]] - 1 

當 count = 0 時 代表 已經找當相同 出現次數的 直接回傳 true

當 right - left + 1 ≥ len(s1) 時 代表 slide-window 已經超出 s1 permutation 的可能值

     所以需要 把 freq[s[left]] ≥ 0 的值回補回來，更新 count = count + 1  

     因為等下左移之後需要重新計算

     更新 freq[s[left]]= freq[s[left]]+1

    更新 left = left +1

當所有 right 都跑完 但沒有找到 count = 0 的情況 代表沒有所以回傳 false

具體如下


![](https://i.imgur.com/IUulX1L.png)

## 程式碼
```go
package sol

func checkInclusion(s1 string, s2 string) bool {
	s1Len, s2Len := len(s1), len(s2)
	if s1Len > s2Len {
		return false
	}
	left, freq, count := 0, make([]int, 26), s1Len
	for idx := 0; idx < s1Len; idx++ {
		freq[s1[idx]-'a']++
	}
	for right := 0; right < s2Len; right++ {
		if freq[s2[right]-'a'] > 0 {
			count--
		}
		freq[s2[right]-'a']--
		if count == 0 {
			return true
		}
		if right-left+1 >= s1Len { // slide-window out of range and not found
			if freq[s2[left]-'a'] >= 0 {
				count++
			}
			freq[s2[left]-'a']--
			left++
		}
	}
	return false
}

```
## 困難點

1. 要理解某個字串的 permutation 定義
2. 需要知道透過 hashTable 可以檢驗 permutation
3. 需要知道怎麼透過 slide-window 來檢驗字元出現次數

## Solve Point

- [x]  初始化 left = 0, freq 是一個 hashMap 用來紀錄出現字元的個數 , s1Len = len(s1), s2Len = len(s2), count = s1Len
- [x]  if s1Len > s2Len  回傳 false
- [x]  先 loop s1 所有字元蒐集所有字元個數在 freq
- [x]  從 right = 0…s2Len - 1 做以下運算
- [x]  當 freq[s[right]] > 0 時 ， 更新 count = count - 1
- [x]  更新 freq[s[right]] = freq[s[right]]- 1
- [x]  當 count == 0 時， 代表找到該 permutation 回傳 true
- [x]  當 right - left + 1 ≥ s1Len 且 count ≠ 0, 這時代表需要把左界右移
- [x]  所以檢查當 left 位置的字元次數 ≥ 0 代表該字元需要被紀錄，所以要把 count 加回來
- [x]  更新 freq[s[left]] =  freq[s[left]] + 1, left = left + 1
- [x]  當所有 right 都檢查完 ， count ！= 0 代表找不到所以 return false