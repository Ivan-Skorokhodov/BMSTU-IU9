#include <iostream>
#include <cmath>
#include <vector>

using namespace std;

class Point 
{
    public:
        double x, y, mass; 
        Point(double x, double y, double mass);
};

Point::Point (double x, double y, double mass)
{   
    this->x = x;
    this->y = y;
    this->mass = mass;
}

class System
{
    private:
        vector<Point> listPoints;

    public:
        System(vector<Point> listPoints);
        double getCentrMassX();
        double getCentrMassY();
        int getCount();
        Point getPoint(int n);
        void addPoint(Point p);
        void RemoveAllZeroMass();
};

System::System (vector<Point> listPoints)
{   
    this->listPoints = listPoints;
}

double System::getCentrMassX() 
{
    double chX = 0;
    double sumMass = 0;
    
    for (int i = 0; i < listPoints.size(); i++)
    {
        chX += listPoints[i].x * listPoints[i].mass;
        sumMass += listPoints[i].mass;
    }

    double X = chX / sumMass;

    return X;
}

double System::getCentrMassY() 
{
    double chY = 0;
    double sumMass = 0;
    

    for (int i = 0; i < listPoints.size(); i++)
    {
        chY += listPoints[i].y * listPoints[i].mass;
        sumMass += listPoints[i].mass;
    }

    double Y = chY / sumMass;

    return Y;
}

int System::getCount() 
{
    return listPoints.size();
}

Point System::getPoint(int n)
{
    return listPoints[n];
}

void System::addPoint(Point p)
{
    listPoints.push_back(p);
}

void System::RemoveAllZeroMass()
{
    vector<Point> newlistPoints;
    for (int i = 0; i < listPoints.size(); i++)
    {
        if (listPoints[i].mass != 0)
        {
            newlistPoints.push_back(listPoints[i]);
        }
    }
    
    listPoints = newlistPoints;
}

int main()
{
    vector<Point> listPoints;
    for (int i = 0; i < 50; i++) {
        listPoints.push_back(Point(i/10, i/12, i/2));
    }

    System mySystem(listPoints);

    cout << mySystem.getCentrMassX() << endl;
    cout << mySystem.getCentrMassY() << endl;
    cout << mySystem.getCount() << endl;
    cout << mySystem.getPoint(25).mass << endl;

    cout << "----------------" << endl;

    Point p(10, 20, 0);
    mySystem.addPoint(p);

    cout << mySystem.getCount() << endl;

    cout << "----------------" << endl;

    mySystem.RemoveAllZeroMass();
    cout << mySystem.getCount() << endl;
    
    return 0;
}