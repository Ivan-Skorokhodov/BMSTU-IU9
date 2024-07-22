package Iterator1;

public class TestIterator1 {

    public static void main(String[] args) {
        String[] myArray = {"ba", "la", "gan"};

        SyllableList array = new SyllableList(myArray);

        for(String s : array) {
            System.out.println(s);
        }

        System.out.println("--------------");

        myArray[2] = "po";
        for(String s : array) {
            System.out.println(s);
        }

    }
}
