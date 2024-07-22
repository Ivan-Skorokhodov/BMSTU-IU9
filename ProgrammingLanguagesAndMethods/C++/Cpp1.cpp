#include <iostream>
#include <sys/types.h>
#include <dirent.h>
#include <string>
#include <cstring>
#include <fstream>
#include <vector>
#include <algorithm>

using namespace std;

bool isTxtFile(char *filename)
{
    char *dot = strrchr(filename, '.');
    if (!dot || dot == filename)
    {
        return false;
    }
    return strcmp(dot, ".txt") == 0;
}

int main()
{
    DIR *dirp = opendir("/home/skor/GoStudy/CProjects/abra");
    struct dirent *file;

    vector<string> listLinks;

    for (int i = 0; i < 7; i++)
    {
        file = readdir(dirp);

        vector<string> list;

        if (isTxtFile(file->d_name))
        {
            string path = "/home/skor/GoStudy/CProjects/abra/" + string(file->d_name);
            ifstream f(path);
            string line;

            cout << file->d_name << endl; // !!!!!!!!!!!!!!!!!!!!!!!!!

            while (getline(f, line))
            {
                list.push_back(line);
            }

            f.close();
        }

        if (isTxtFile(file->d_name))
        {
            string path = "/home/skor/GoStudy/CProjects/abra/" + string(file->d_name);
            ofstream f(path);

            for (int i = list.size() - 1; i >= 0; i--)
            {
                f << list[i] << endl;
            }

            f.close();
        }
    }

    closedir(dirp);

    return 0;
}