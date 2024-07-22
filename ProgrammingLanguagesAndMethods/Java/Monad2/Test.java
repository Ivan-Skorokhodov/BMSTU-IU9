package Monad2;

import java.math.BigInteger;

public class Test {
    public static void main(String[] args) {
        BigInteger ch1 = new BigInteger("4");
        BigInteger zn1 = new BigInteger("10");
        Fraction f1 = new Fraction(ch1, zn1);

        BigInteger ch2 = new BigInteger("5");
        BigInteger zn2 = new BigInteger("20");
        Fraction f2 = new Fraction(ch2, zn2);

        BigInteger ch3 = new BigInteger("45256");
        BigInteger zn3 = new BigInteger("50");
        Fraction f3 = new Fraction(ch3, zn3);

        BigInteger ch4 = new BigInteger("6");
        BigInteger zn4 = new BigInteger("100");
        Fraction f4 = new Fraction(ch4, zn4);

        SetFractions set = new SetFractions();

        set.add(f1);
        set.add(f2);
        set.add(f3);
        set.add(f4);

        set.nameStream().sorted(new NameComparator()).forEach(System.out::println);

        System.out.println("------------------------");

        System.out.println(set.getMaxFraction().get().zn);
    }
}
