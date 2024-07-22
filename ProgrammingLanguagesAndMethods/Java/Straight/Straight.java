package Straight;

public class Straight {
    private double a, b, c;
    public Straight(double a, double b, double c) {
        this.a = a;
        this.b = b;
        this.c = c;
    }

    public Straight Perpendicular(double x, double y){
        var perpenA = this.b;
        var perpenB = -this.a;
        var perpenC = - (perpenA * x) - (perpenB * y);

        Straight perpendicular = new Straight(perpenA, perpenB, perpenC);

        return perpendicular;
    }

    public String toString() {
        return a + "x + " + b + "y + " + c + " = 0";
    }
}
