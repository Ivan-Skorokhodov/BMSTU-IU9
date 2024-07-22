package Monad1;

import java.util.ArrayList;
import java.lang.Math;
import java.util.Optional;
import java.util.stream.Stream;

public class MySet {
    ArrayList<Integer> array;

    MySet() {
        array = new ArrayList<>();
    }

    public void add(int number) {
        array.add(number);
    }

    public static  boolean CompletePattern(int number, String pattern) {
        char[] helper = pattern.toCharArray();
        ArrayList<Integer> listOfNums = new ArrayList<>();

        for (int i = (int)Math.pow(10, (pattern.length() - 1)); i < Math.pow(10, pattern.length()); i++){
            String numStr = String.valueOf(i);

            boolean check = true;
            for (int j = 0; j < pattern.length(); j++) {
                if (helper[j] == '?') {
                    continue;
                }
                if (helper[j] != numStr.charAt(j)) {
                    check = false;
                    break;
                }
            }

            if (check) {
                listOfNums.add(i);
            }
        }

        for (int i : listOfNums) {
            if (i == number) {
                return true;
            }
        }
        return false;
    }

    public Stream<Integer> nameStream(String pattern) {
        ArrayList<Integer> result = new ArrayList<>();
        array.stream()
                .filter(x -> CompletePattern(x, pattern))
                .forEach(x -> result.add(x));

        return result.stream();
    }

    public Optional<Integer> getMaxInt(String pattern) {
        Optional<Integer> result = Optional.empty();

        int maxNum = 0;
        for (int i : array) {
            if (CompletePattern(i, pattern) && i > maxNum) {
                maxNum = i;
            }
        }

        final int max = maxNum;

        Optional<Integer> tmp = array.stream()
                .filter(x -> x == max)
                .findFirst();

        if (tmp.isPresent()) {
            result = Optional.ofNullable(tmp.get());
        }
        return result;
    }

}
