package study;

/**
 * Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
 *
 * Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
 *
 *
 * Example 1:
 *
 * Input: x = 123
 * Output: 321
 * Example 2:
 *
 * Input: x = -123
 * Output: -321
 * Example 3:
 *
 * Input: x = 120
 * Output: 21
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/reverse-integer
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
public class ReverseInt {

    public int reverse(int x) {
        if (x < 0) {
            if (x == -x) {
                return 0;
            }
            return -reverse(-x);
        }
        int ret = 0;
        while (x != 0) {
            if (ret > Integer.MAX_VALUE / 10) {
                return 0;
            }
            ret *= 10;
            ret += x % 10;
            x /= 10;
        }
        return ret;
    }

}
