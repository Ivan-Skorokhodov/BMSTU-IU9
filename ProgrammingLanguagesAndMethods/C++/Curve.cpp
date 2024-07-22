#include <iostream>
#include <vector>
#include <cmath>

using namespace std;

template <typename T>
class Curve
{
    public:
        vector<pair<bool, int>> list;
        Curve(bool Bool);
        Curve(vector<pair<bool, int>> list);
        
        Curve<T> operator +(Curve<T> other)
        {
            vector<pair<bool, int>> newList;
            
            for(int i = 0; i < (this->list).size(); i++)
            {
                newList.push_back((this->list)[i]);
            }

            for(int i = 0; i < other.list.size(); i++)
            {
                newList.push_back(other.list[i]);
            }

            return Curve(newList);
        }

        Curve<T> operator -(Curve<T> other)
        {
            vector<pair<bool, int>> newList;
            
            for(int i = 0; i < (this->list).size(); i++)
            {
                newList.push_back((this->list)[i]);
            }

            for(int i = 0; i < other.list.size(); i++)
            {
                pair<bool, int> invPair = other.list[i];
                invPair.second *= -1;
                newList.push_back(invPair);
            }

            return Curve(newList);
        }

        Curve<T> operator *(int k)
        {
            vector<pair<bool, int>> newList;
            
            for(int i = 0; i < (this->list).size(); i++)
            {   
                pair<bool, int> kPair = this->list[i];
                kPair.second *= k;

                newList.push_back(kPair);
            }

            return Curve(newList);
        }

        Curve<T> operator !()
        {
            vector<pair<bool, int>> newList;

            for(int i = 0; i < this->list.size(); i++)
            {
                pair<bool, int> invPair = this->list[i];
                if (invPair.first)
                {
                    invPair.first = false;
                } else {
                    invPair.first = true;
                }

                newList.push_back(invPair);
            }

            return Curve(newList);
        }

        Curve<T> operator -()
        {
            vector<pair<bool, int>> newList;

            for(int i = 0; i < this->list.size(); i++)
            {
                pair<bool, int> invPair = this->list[i];
                invPair.second *= -1;

                newList.push_back(invPair);
            }

            return Curve(newList);
        }

        float operator ()(float x)
        {
            float ans = 0;

            for(int i = 0; i < this->list.size(); i++)
            {
                pair<bool, int> Pair = this->list[i];

                if (Pair.first)
                {
                    ans += sin(x) * Pair.second;
                } else {
                    ans += cos(x) * Pair.second;
                }
            }

            return ans;
        }

};

template<typename T>
Curve<T>::Curve (bool Bool)
{   
    pair<bool, int> firstPair = pair<bool, int>(Bool, 1);
    vector<pair<bool, int>> newList;
    newList.push_back(firstPair);
    this->list = newList;
}

template<typename T>
Curve<T>::Curve (vector<pair<bool, int>> list)
{   
    this->list = list;
}

int main()
{
    Curve<float> func = Curve<float>(true);

    cout << func(10) << endl;

    cout << (func + func)(10) << endl;

    cout << (func - func * 2)(10) << endl;

    cout << (func * 10)(10) << endl;

    cout << (-func)(10) << endl;

    cout << (!(func))(10) << endl;

    return 0;
}