#include <iostream>
#include <cmath>

using namespace std;

class Point 
{
    public:
        double x, y;
        double r();
        Point(double x, double y);
        virtual ~Point();
};

Point :: Point (double x, double y)
{   
    this->x = x;
    this->y = y;
    cout << "(" << x << "," << y << ")" << endl;
}

double Point :: r() 
{
    return sqrt(x*x + y*y);
}

Point :: ~Point()
{
    cout << "self destruction" << endl;
}

int main()
{
    cout << "help" << endl;
    
    Point A(1.0, 1.0);
    cout << A.r() << endl;
    return 0;
}