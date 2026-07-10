from typing import List

class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        """
        In this solution, I initial made seen_nums a list,
        the problem with that is that when doing the check
        if the number has been seen before, it creates O(n^2)

        By using a dict / hash table, the lookup is extremely
        fast.

        Key Lesson: Hash Table beats list everytime
        """
        seen_nums = {}
        for num in nums:
            if num in seen_nums:
                return True
            else:
                seen_nums[num] = 1
        
        return False
        # a crazy solution - compare the lengths of a set of nums
        # and nums itself
        # return len(nums) == len(set(nums))
    
    def containsDuplicateBruteForce(self, nums: List[int]) -> bool:
        """
        Brute force approach, checks every value in the list

        Time complexity of O(n^2) because you have to check each num
        twice
        """
        n = len(nums)
        for i in range(n - 1):
            for j in range(i + 1, n):
                if nums[i] == nums[j]:
                    return True
        return False