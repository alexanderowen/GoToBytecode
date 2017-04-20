class bool {
    /* Functions t and f are created so that the Java 
     * compiler doesn't optimize code involving only 
     * 'true' and 'false'
     * */
    public static boolean t() {
        return true;
    }
    public static boolean f() {
        return false;
    }

    public static boolean eq() {
        return t() == f();
    }
    public static boolean neq() {
        return t() != f();
    }

    public static boolean and() {
        return t() && f();
    }
    public static boolean or() {
        return t() || f();
    }
    public static boolean not() {
        return !t();
    }

}
