package Monad2;

import java.math.BigInteger;
import java.util.Comparator;

public class NameComparator implements Comparator<BigInteger> {
    public int compare(BigInteger a, BigInteger b) {
        return a.compareTo(b);
    }
}
