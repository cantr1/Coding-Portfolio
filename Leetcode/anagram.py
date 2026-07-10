class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        """
        This solution uses a frequency hash map to count the
        number of occurances of a character then compares
        the counts between the two strings.

        The Key Lesson: Hash Map can be used to track frequency (frequency map)
        """
        if len(s) != len(t):
            return False

        collected_letters_s = {}
        collected_letters_t = {}

        for i in range(0, len(s)):
            if s[i] in collected_letters_s:
                collected_letters_s[s[i]] += 1
            else:
                collected_letters_s[s[i]] = 1

            if t[i] in collected_letters_t:
                collected_letters_t[t[i]] += 1
            else:
                collected_letters_t[t[i]] = 1

        for key, value in collected_letters_s.items():
            if key not in collected_letters_t:
                return False
            if value != collected_letters_t[key]:
                return False
    
        return True