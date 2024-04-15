// https://leetcode.com/problems/longest-palindromic-substring/

package java.medium;

class LongestPalindromicSubstring {
    private int maxLen = 0;

    public String longestPalindrome(String s) {
        int n = s.length();

        if (n < 2) {
            return s;
        }

        String result = "";

        for (int i = 0; i < n; i++) {
            // odd length -> single character center
            result = count(s, i, i, result);

            // even length -> consecutive character center
            result = count(s, i, i+1, result);

        }

        return result;
    }

    private String count(String s, int left, int right, String result) {
        while (left >= 0 && right < s.length() && s.charAt(left) == s.charAt(right)) {
            if (right - left + 1 > maxLen) {
                result = s.substring(left, right+1);
                maxLen = right - left + 1;
            }
            left--;
            right++;
        }

        return result;
    }
}
