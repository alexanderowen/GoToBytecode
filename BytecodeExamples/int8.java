/* Java does not have bytecode instructions to operate on 'byte's. 
 * It treats them as ints, and performs a cast.
 * */

class int8 {
    public static byte add(byte x, byte y) {
        return (byte)(x + y);
    }
}
