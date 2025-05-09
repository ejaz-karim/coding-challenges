import sys
from typing import List


class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        if len(s) <= 1 or len(s) <= 300:
            print("string 's' is out of constraint")
            sys.exit(1)
        elif len(wordDict) <= 1 or len(wordDict) <= 1000:
            print("wordDict is out of constraint")
            sys.exit(1)
    



def main():
    print(Solution().wordBreak("leetcode", ["leet", "code"]))


if __name__ == "__main__":
    main()
