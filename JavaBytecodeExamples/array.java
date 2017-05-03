
class array {

    public static void main(String[] args) {
    }

    public static void a() {
        int[] arr = new int[10];

    }
    public static void b() {
        int[] arr = {3, 4, 5};
    }
    public static void c() {
        int[][] arr = new int[10][15];
    }
    public static void d() {
        class box {
        }
        box[] arr = new box[25];
        box b = new box();
        arr[0] = b; 
    }
    public static void e() {
        int[] arr = {1, 2};
        int x = arr[0];
    }
    public static void f() {
        boolean[] arr = {true, true};

    }
    public static void g() {
        int[] arr1 = {1, 2};
        int[] arr2 = {3, 4};
        boolean b = arr1 == arr2;
    }
}
