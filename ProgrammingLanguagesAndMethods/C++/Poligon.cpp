#include <iostream>
#include <vector>
#include <cmath>

using namespace std;

template <typename T, int M>
class Poligon
{
    public:
        vector<vector<T>> listCoords;
        Poligon(vector<vector<T>> listCoords);
        double perimetr();
        void addNewV(vector<T> v);
        T square();
};

template<typename T, int M>
Poligon<T, M>::Poligon (vector<vector<T>> listCoords)
{   
    this->listCoords = listCoords;
}

template<typename T, int M>
double Poligon<T, M>::perimetr()
{
    double p = 0;

    vector<T> elem1 = (this->listCoords)[0];
    vector<T> elem2 = (this->listCoords)[(this->listCoords).size()-1];

    double sum1 = 0;
    for(int j = 0; j < elem1.size(); j++)
    {
        sum1 += pow((elem1[j] - elem2[j]), 2);
    }
    p += sqrt(sum1);

    for(int i = 1; i < (this->listCoords).size(); i++)
    {
        vector<T> elem1 = (this->listCoords)[i-1];
        vector<T> elem2 = (this->listCoords)[i];
        double sum = 0;

        for(int j = 0; j < elem1.size(); j++)
        {
            sum += pow((elem1[j] - elem2[j]), 2);
        }
        p += sqrt(sum);
    }

    return p;
}

template<typename T, int M>
void Poligon<T, M>::addNewV (vector<T> v)
{   
    (this->listCoords).push_back(v);
}

template<typename T, int M>
T Poligon<T, M>::square ()
{   
    if (M == 2)
    {
        T maxX = 0;
        T minX = 10000000;
        T maxY = 0;
        T minY = 10000000;

        for(int i = 0; i < (this->listCoords).size(); i++)
        {
            T x = listCoords[i][0];
            T y = listCoords[i][1];

            maxX = max(maxX, x);
            maxY = max(maxY, y);
            minX = min(minX, x);
            minY = min(minY, y);
        }

        return (maxX - minX) * (maxY - minY);
    }
}

int main()
{
    vector<vector<int>> listCoords;
    int mn = 1;

    for(int i = 0; i < 5; i++)
    {
        vector<int> v1;

        for(int j = 0; j < 5; j++)
        {
            v1.push_back(j*mn);            
        }

        listCoords.push_back(v1);
        mn += 1;
    }

    Poligon<int, 5> myPoligon = Poligon<int, 5>(listCoords);

    cout << myPoligon.perimetr() << endl;

    cout << "----------------" << endl;

    vector<int> newV;
    for(int j = 0; j < 5; j++)
    {
        newV.push_back(j*mn);            
    }

    myPoligon.addNewV(newV);

    cout << myPoligon.perimetr() << endl;

    cout << "----------------" << endl;


    vector<vector<int>> listCoords2;

    for(int i = 0; i < 10; i++)
    {
        vector<int> v1;

        for(int j = 1; j < 3; j++)
        {
            v1.push_back(j*mn);            
        }

        listCoords.push_back(v1);
        mn += 1;
    }

    Poligon<int, 2> myPoligon2 = Poligon<int, 2>(listCoords);
    cout << myPoligon2.square() << endl;

    cout << "----------------" << endl;

    vector<int> newV2;
    for(int j = 1; j < 3; j++)
    {
        newV2.push_back(40);            
    }

    myPoligon2.addNewV(newV2);

    cout << myPoligon2.square() << endl;

    cout << "----------------" << endl;

    return 0;
}