#include <iostream>
#include <vector>
#include <cmath>

using namespace std;

bool isPow2(int n)
{
    int k = 1;
    while (1 << k <= n)
    {
        if (1 << k == n)
        {
            return true;
        } else {
            k++;
        }
    }
    return false;
}

class Seq
{
    public:
        vector<int> list;

        Seq(vector<int> list)
        {
            this->list = list;
        }

        int operator [](int index)
        {
            return (this->list)[index];
        }
};

class Iter
{
    public:
        vector<int> list;
        int pos, val;

        Iter(Seq i)
        {
            this->list = i.list;

            int val = 0;
            int pos = 0;

            for(int j = 0; j < i.list.size(); j++)
            {
                if (isPow2(i.list[j]))
                {
                    val = i.list[j];
                    pos = j;
                    break;
                }
            }

            if (val == 0)
            {
                cout << "error seq" << endl;
                this->val = -1;
                this->pos = -1;
            } else {
                this->val = val;
                this->pos = pos;
            }
        }

        void operator ++()
        {
            bool hasNumber = false;

            for(int j = this->pos + 1; j < (this->list).size(); j++)
            {
                if (isPow2((this->list)[j]))
                {
                    this->val = (this->list)[j];
                    this->pos = j;
                    hasNumber = true;
                    break;
                }
            }

            if (!hasNumber)
            {
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
    vector<int> v;

    for(int i = 1; i < 20; i++){
        v.push_back(i);
    }

    Seq seq = Seq(v);

    Iter it = Iter(seq);

    cout << *it << endl;

    ++it;

    cout<< *it << endl;

    ++it;

    cout<< *it << endl;

    cout<< "----------" << endl;

    cout<< seq[10] << endl;

    return 0;
}