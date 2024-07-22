package Universe;

public class Point
{
    private double x;
    private double y;
    private double mass;
    static private int count = 0;

    public Point(double x, double y, double mass) {
        this.x = x;
        this.y = y;
        this.mass = mass;
        count++;
    }

    public static int getCount()
    {
        return Point.count;
    }

    public double getXForse(Point point) {
        double XForse = 0.00667 * (this.mass + point.mass) /
                ((this.x - point.x) * (this.x - point.x));
        return XForse;
    }

    public double getYForse(Point point) {
        double Yforse = 0.00667 * (this.mass + point.mass) /
                ((this.y - point.y) * (this.y - point.y));
        return Yforse;
    }
}
