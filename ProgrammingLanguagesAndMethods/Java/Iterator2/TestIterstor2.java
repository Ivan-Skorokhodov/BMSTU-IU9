package Iterator2;

public class TestIterstor2 {
    public static void main(String[] args) {

        Progression p1 = new Progression(3, 5, 0);
        Progression p2 = new Progression(4, 10, 0);
        Progression p3 = new Progression(5, 15, 0);

        Progression[] arrayProgressions = {p1, p2};
        ProgressionList array = new ProgressionList(arrayProgressions);

        for(int i : array) {
            System.out.println(i);
        }

        System.out.println("--------------");

        arrayProgressions[1] = p3;
        for(int i : array) {
            System.out.println(i);
        }
    }
}
