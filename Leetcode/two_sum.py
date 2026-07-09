
from typing import List

def twoSum(nums: List[int], target: int) -> List[int]:
    """
    This solution works by creating a hash map of the value
    of seen numbers and their index.
    By calculating the number we need, we can check if we have 
    seen the value previously and return its index.
    This greatly reduces the time complexity.

    The Key Lesson: Hash Map
    """
    seen_numbers = {}
    for i, num in enumerate(nums):
        num_we_need = target - num
        if num_we_need in seen_numbers:
            return [seen_numbers[num_we_need], i]
        seen_numbers[num] = i