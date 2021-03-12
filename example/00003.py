class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        """
        :type s: str
        :rtype: int
        """
        # 最大字符
        r = 0
        # 不重复字符串的数组
        not_repeat = []
        for i in range(len(s)):
            v = s[i]
            if v in not_repeat:
                #如果存在重复元素，删除到重复的位置，重新开始计算
                del not_repeat[0:not_repeat.index(v) + 1]
                not_repeat.append(v)
            else:
                #不存在重复元素，计算最大长度并且保存长度
                not_repeat.append(v)
                l = len(not_repeat)
                if l > r:
                    r = l
        return r
# 实现2
class Solution2:
    def lengthOfLongestSubstring(self, s: str) -> int:
        # 哈希集合，记录每个字符是否出现过
        occ = set()
        n = len(s)
        # 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
        rk, ans = -1, 0
        for i in range(n):
            if i != 0:
                # 左指针向右移动一格，移除一个字符
                occ.remove(s[i - 1])
            while rk + 1 < n and s[rk + 1] not in occ:
                # 不断地移动右指针
                occ.add(s[rk + 1])
                rk += 1
            # 第 i 到 rk 个字符是一个极长的无重复字符子串
            ans = max(ans, rk - i + 1)
        return ans


s1 = "abcabcbb"
s2 = 'pwwkew'

print(Solution().lengthOfLongestSubstring(s1))
print(Solution().lengthOfLongestSubstring(s2))
print(Solution2().lengthOfLongestSubstring(s1))
print(Solution2().lengthOfLongestSubstring(s2))

