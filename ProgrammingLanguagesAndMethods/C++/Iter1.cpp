#include <iostream>
#include <vector>
#include <cmath>

using namespace std;

class Interval
{
    public:
        double a, b;

        Interval(double a, double b)
        {
            this->a = a;
            this->b = b;
        }
};

class Iter
{
    public:
        double a, b;
        int pos, val;

        Iter(Interval i)
        {
            double a = i.a;
            double b = i.b;


            this->a = a;
            this->b = b;

            int j = 0;
            while (j < a)
            {
                j++;
            }
            
            this->val = j;
            this->pos = 0;
        }

        void operator ++()
        {
            if (val + 1 < b){
                (this->pos)++;
                (this->val)++;
            } else {
                cout<< "iterator out of range" << endl;
            }
        }

        int operator *()
        {
            return this->val;
        }
};

int main()
{
    double a = 1.235;
    double b = 12.673;

    Interval i = Interval(a, b);

    Iter it = Iter(i);

    cout << *it << endl;

    ++it;

    cout<< *it << endl;

    ++it;

    cout<< *it << endl;

    return 0;
}