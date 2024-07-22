#include <iostream>
#include <cmath>
#include <vector>

using namespace std;

class Point 
{
    private:
        double x, y, mass;
        static inline int count;

    public:
        Point(double x, double y, double mass);
        static int getCount();
        double getXForce(Point point);
        double getYForce(Point point);

};

Point::Point (double x, double y, double mass)
{   
    this->x = x;
    this->y = y;
    this->mass = mass;
    count++;
}

int Point::getCount() 
{
    return count;
}

double Point::getXForce(Point point) 
{
    double XForce = 0.00667 * (mass * point.mass) / ((x - point.x) * (x - point.x));
    return XForce;
}

double Point::getYForce(Point point) 
{
    double YForce = 0.00667 * (mass * point.mass) / ((y - point.y) * (y - point.y));
    return YForce;
}


class Universe
{
    private:
        vector<Point> listPoints;

    public:
        Universe(vector<Point> listPoints);
        double getForce(Point point);
};

Universe::Universe (vector<Point> listPoints)
{   
    this->listPoints = listPoints;
}

double Universe::getForce(Point point) 
{
    double X, Y;
    double FinalX = 0;
    double FinalY = 0;

    for (int i = 0; i < listPoints.size(); i++)
    {
        X = point.getXForce(listPoints[i]);
        Y = point.getYForce(listPoints[i]);
        FinalX += X;
        FinalY += Y;
    }

    double result = sqrt((FinalX * FinalX) + (FinalY * FinalY));
    return result;
}



int main()
{
    vector<Point> listPoints;
    for (int i = 0; i < 20; i++) {
        listPoints.push_back(Point(i/10, i/10, i/2));
    }

    Universe myUniverse(listPoints);

    Point p(10, 20, 1);

    double forse = myUniverse.getForce(p);

    cout << forse << endl;
    cout << p.getCount() << endl;
    return 0;
}