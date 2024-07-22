package Iterator1;

import java.util.Iterator;

class SyllableList implements Iterable<String> {
    private String[] listSyllable;
    private String[] listWords;
    public SyllableList (String[] listSyllable) {
        this.listSyllable = listSyllable;;
        this.listWords = new String[countCombinations()];
    }

    public void fillListWords() {
        int countSyllables = listSyllable.length;
        int pos  = 0;
        for (int i = 0; i < countSyllables; i++) {
            String s = "";
            s += listSyllable[i];

            for (int j = i; j < countSyllables; j++) {
                if (i != j) {
                    s += listSyllable[j];
                }

                listWords[pos] = s;
                pos++;
            }
        }
    }

    public int countCombinations() {
        int countSyllables = listSyllable.length;
        int count = 0;
        for (int i = 0; i < countSyllables; i++) {
            for (int j = i; j < countSyllables; j++) {
                count++;
            }
        }
        return count;
    }

    public Iterator iterator() {
        return new SyllableListItarator();
    }
    private class SyllableListItarator implements Iterator {
        private int pos;
        public SyllableListItarator () {
            fillListWords();
            pos = 0;

        }

        public boolean hasNext () {
            return pos < countCombinations();
        }

        public String next () {
            pos++;
            return listWords[pos - 1];
        }
    }
}