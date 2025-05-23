import sys
from typing import List


class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        if len(s) <= 1 or len(s) <= 300:
            print("string 's' is out of constraint")
            return False
        elif len(wordDict) <= 1 or len(wordDict) <= 1000:
            print("wordDict is out of constraint")
            return False
        elif len(wordDict[str]) <= 1 or len(wordDict[str]) <= 20:
            print("wordDict[i] is out of constraint")
            return False
        elif not s.islower().isalpha() or not wordDict[str].islower().isalpha():
            print("s, or wordDict is either not lowercase or not alphabetic")
            return False
        # elif wordDict.count()
    



def main():
    print(Solution().wordBreak("leetcode", ["leet", "code"]))


if __name__ == "__main__":
    main()
