package Monad1;

public class Test {
    public static void main(String[] args) {
        MySet set = new MySet();
        set.add(123);
        set.add(524);
        set.add(829);
        set.add(121);
        set.add(182);
        set.add(956);

        set.nameStream("?2?").sorted(Integer::compareTo).forEach(System.out::println);

        System.out.println("------------------------");

        System.out.println(set.getMaxInt("?2?").get());
    }
}
