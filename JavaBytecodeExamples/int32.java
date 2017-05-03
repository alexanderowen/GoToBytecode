class int32 {
    public static int add(int x, int y) {
        return x + y; 
    }
    public static int sub(int x, int y) {
        return x - y; 
    }
    public static int mult(int x, int y) {
        return x * y; 
    }
    public static int div(int x, int y) {
        return x / y; 
    }
    public static int mod(int x, int y) {
        return x % y; 
    }

    public static int bitand(int x, int y) {
        return x & y; 
    }
    public static int bitor(int x, int y) {
        return x | y; 
    }
    public static int bitxor(int x, int y) {
        return x ^ y; 
    }
    /* In Golang: x &^ y */
    public static int bitclear(int x, int y) {
        return x & ~y;
    }

    /* In Golang, for both << and >>, y MUST be an unsigned integer
     * */
    public static int leftshift(int x, int y) {
        return x << y;
    }
    public static int rightshift(int x, int y) {
        return x >> y;
    }

    public static boolean eq(int x, int y) {
        return x == y;
    }
    public static boolean neq(int x, int y) {
        return x != y;
    }
    public static boolean lt(int x, int y) {
        return x < y;
    }
    public static boolean lteq(int x, int y) {
        return x <= y;
    }
    public static boolean gt(int x, int y) {
        return x > y;
    }
    public static boolean gteq(int x, int y) {
        return x >= y;
    }

    public static int pos(int x) {
        return +x;
    }
    public static int neg(int x) {
        return -x;
    }
    public static int bitnot(int x) {
        return ~x;
    }
}
