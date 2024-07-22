package IntDel;

import java.util.Arrays;

public class TestIntDel {
    public static void main(String[] args) {
    IntDel num1 = new IntDel(30);
    IntDel num2 = new IntDel(5);
    IntDel num3 = new IntDel(24);

    IntDel[] array = new IntDel[3];
    array[0] = num1;
    array[1] = num2;
    array[2] = num3;

    Arrays.sort(array);
    for (IntDel elem : array){
        System.out.println(elem);
    }

    }
}
