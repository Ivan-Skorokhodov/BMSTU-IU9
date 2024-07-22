#include <iostream>
#include <vector>
#include <cmath>

using namespace std;

vector<string> splitString(const string &str)
{
    vector<string> words;
    string word;
    for (char c : str)
    {
        if (c == ' ')
        {
            if (!word.empty())
            {
                words.push_back(word);
                word.clear();
            }
        }
        else
        {
            word += c;
        }
    }
    if (!word.empty())
    {
        words.push_back(word);
    }
    return words;
}

class Sentence
{
public:
    vector<string> list;

    Sentence(string str)
    {
        this->list = splitString(str);
    }

    string operator[](int x)
    {
        return list[x];
    }
};

class Iter
{
public:
    vector<string> list;
    int pos;
    string val;

    Iter(Sentence s)
    {
        this->list = s.list;
        this->val = list[0] + list[1];
        this->pos = 0;
    }

    void operator++()
    {
        if (pos + 1 < list.size())
        {
            (this->pos)++;
            (this->val) = list[pos] + list[pos + 1];
        }
        else
        {
            cout << "iterator out of range" << endl;
        }
    }

    void operator--()
    {
        if (pos - 1 >= 0)
        {
            (this->pos)--;
            (this->val) = list[pos] + list[pos + 1];
        }
        else
        {
            cout << "iterator out of range" << endl;
        }
    }

    string operator*()
    {
        return this->val;
    }
};

int main()
{
    string str = "abcd efg hijk l";

    Sentence s = Sentence(str);

    Iter it = Iter(s);

    cout << *it << endl;

    ++it;

    cout << *it << endl;

    ++it;

    cout << *it << endl;

    cout << "----------------" << endl;

    --it;

    cout << *it << endl;

    return 0;
}
