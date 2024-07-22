package Monad2;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Optional;
import java.util.stream.Stream;

public class SetFractions {
    ArrayList<Fraction> array;
    static int count = 0;

    SetFractions() {
        array = new ArrayList<>();
    }

    static boolean compareBigIntegers(BigInteger a, BigInteger b) {
        if (a.compareTo(b) == -1) {
            return true;
        }
        return false;
    }
    public void add(Fraction frac) {
        array.add(frac);
        count++;
    }

    public Stream<BigInteger> nameStream() {
        ArrayList<BigInteger> result = new ArrayList<>();
        array.stream()
                .filter(x -> compareBigIntegers(x.ch, x.zn))
                .forEach(x -> result.add(x.ch));

        return result.stream();
    }

    static boolean isDel(BigInteger num, ArrayList<Fraction> array) {
        BigInteger zero = new BigInteger("0");
        for (Fraction frac : array) {
            if (zero.compareTo(num.mod(frac.zn)) != 0){
                return false;
            }
        }
        return true;
    }

    public Optional<Fraction> getMaxFraction() {
        Optional<Fraction> result = Optional.empty();

        BigInteger maxNum = new BigInteger("0");
        for (Fraction i : array) {
            if (isDel(i.zn, array) && i.zn.compareTo(maxNum) == 1) {
                maxNum = i.zn;
            }
        }

        final BigInteger max = maxNum;

        Optional<Fraction> tmp = array.stream()
                .filter(x -> x.zn.compareTo(max) == 0)
                .findFirst();

        if (tmp.isPresent()) {
            result = Optional.ofNullable(tmp.get());
        }
        return result;
    }
}
