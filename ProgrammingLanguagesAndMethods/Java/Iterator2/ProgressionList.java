package Iterator2;

import java.util.Arrays;
import java.util.Iterator;

class ProgressionList implements Iterable<Integer> {
    private Progression[] ProgressionsArray;
    private int[] listNums;
    public ProgressionList (Progression[] ProgressionsArray) {
        this.ProgressionsArray = ProgressionsArray;
        this.listNums = new int[countCombinations()];
    }

    public void fillListNums() {
        int pos = 0;
        for (int i = 0; i < ProgressionsArray.length; i++) {
            for (int num : ProgressionsArray[i].getArray()) {
                listNums[pos] = num;
                pos++;
            }
        }

        Arrays.sort(listNums);
    }

    public int countCombinations() {
        int count = 0;
        for (int i = 0; i < ProgressionsArray.length; i++) {
            count += ProgressionsArray[i].getCountNums();
        }
        return count;
    }

    public Iterator iterator() {
        return new ProgressionListIterator();
    }
    private class ProgressionListIterator implements Iterator {
        private int pos;
        public ProgressionListIterator () {
            listNums = new int[countCombinations()];
            fillListNums();
            pos = 0;

        }

        public boolean hasNext () {
            return pos < countCombinations();
        }

        public Integer next() {
            pos++;
            return listNums[pos - 1];
        }
    }
}