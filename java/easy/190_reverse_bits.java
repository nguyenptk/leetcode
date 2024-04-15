package java.easy;

class ReverseBits {
    public int reverseBits(int n) {
        int result = 0;

        for (int i = 0; i < 32; i++) {
            result <<= 1;
            result = result | (n & 1);
            n >>= 1;
        }

        return result;
    }
}
