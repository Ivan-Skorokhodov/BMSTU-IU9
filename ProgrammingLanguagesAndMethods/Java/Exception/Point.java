package Exception;

import static java.lang.Math.*;

public class Point
{

    private String name;
    private double x;
    private double y;
    private double z;
    private static int n;
    public static int val;

    public Point(String argName)
    {
        System.out.println("Запущен конструктор объекта Point");
        this.name = argName;
    }

    public String getName()
    {
        return name;
    }

    public void setCoord(double varX, double varY, double varZ)
    {
        this.x=varX;
        this.y=varY;
        this.z=varZ;
    }

    public double getR()
    {
        return pow(pow(this.x,2)+pow(this.y,2)+pow(this.z,2),0.5);
    }

    public static void setMass(int valN){
        n = valN;
    }

    public void IsInCircle(Circle circle) throws myCustomException {
        if (this.getR() < circle.radius) {
            System.out.println("Старт IsInCircle");
            throw new myCustomException(); //раскомментарить, чтобы симитировать ошибку
        }
    }

}



