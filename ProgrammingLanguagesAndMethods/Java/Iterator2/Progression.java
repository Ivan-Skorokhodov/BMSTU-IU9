package Iterator2;

public class Progression {
    private int countNums;
    private int differense;
    private int firstNum;
    private int[] array;

    public Progression(int countNums, int differense, int firstNum) {
        this.countNums = countNums;
        this.differense = differense;
        this.firstNum = firstNum;
        this.array = new int[countNums];

        for (int i = 0; i < countNums; i++) {
            array[i] = firstNum + i * differense;
        }
    }

    public int getCountNums() {
        return countNums;
    }

    public int[] getArray() {
        return array;
    }
}
