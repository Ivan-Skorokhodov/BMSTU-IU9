package Universe;

public class TestUniverse {
    public static void main(String[] args) {

        Point[] listPoints = new Point[20];
        for (int i = 0; i < 20; i++) {
            listPoints[i] = new Point(i/10, i/10, i/2);
        }

        Universe myUniverse = new Universe(listPoints);

        Point p = new Point(10, 20, 1);

        double forse = myUniverse.getForse(p);

        System.out.println(forse); // print forse
        System.out.println(Point.getCount()); // print number of initialized points
    }
}
