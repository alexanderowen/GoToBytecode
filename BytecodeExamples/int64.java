class int64{
    public static long add(long x, long y) {
        return x + y; 
    }
    public static long sub(long x, long y) {
        return x - y; 
    }
    public static long mult(long x, long y) {
        return x * y; 
    }
    public static long div(long x, long y) {
        return x / y; 
    }
    public static long mod(long x, long y) {
        return x % y; 
    }

    public static long bitand(long x, long y) {
        return x & y; 
    }
    public static long bitor(long x, long y) {
        return x | y; 
    }
    public static long bitxor(long x, long y) {
        return x ^ y; 
    }
    /* In Golang: x &^ y */
    public static long bitclear(long x, long y) {
        return x & ~y;
    }

    /* In Golang, for both << and >>, y MUST be an unsigned integer 
     * */
    public static long leftshift(long x, long y) {
        return x << y;
    }
    public static long rightshift(long x, long y) {
        return x >> y;
    }

    public static boolean eq(long x, long y) {
        return x == y;
    }
    public static boolean neq(long x, long y) {
        return x != y;
    }
    public static boolean lt(long x, long y) {
        return x < y;
    }
    public static boolean lteq(long x, long y) {
        return x <= y;
    }
    public static boolean gt(long x, long y) {
        return x > y;
    }
    public static boolean gteq(long x, long y) {
        return x >= y;
    }

    public static long pos(long x) {
        return +x;
    }
    public static long neg(long x) {
        return -x;
    }
    public static long bitnot(long x) {
        return ~x;
    }
}
