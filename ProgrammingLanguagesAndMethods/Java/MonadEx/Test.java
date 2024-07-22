package MonadEx;

public class Test {
    public static void main(String[] args) {
        ProductTable t = new ProductTable();
        t.add("a", 10, 10);
        t.add("b", 6, 10);
        t.add("c", 2, 5);
        t.add("d", 7, 110);
        t.nameStream(90).sorted(new NameComparator()).forEach(System.out::println);

        System.out.println(t.getProduct().get().name);
    }
}