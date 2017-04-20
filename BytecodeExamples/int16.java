/* Java does not have bytecode instructions that operate on 'short's
 * It treats them like ints, and performs a cast
 *
 * */

class int16 {
    public static short add(short x, short y) {
        return (short)(x + y);
    }
}
