#include <iostream>
#include <sys/types.h>
#include <dirent.h>
#include <string>
#include <cstring>
#include <fstream>
#include <vector>
#include <algorithm>

using namespace std;

bool isCFile(char* filename) {
    char* dot = strrchr(filename, '.');
    if (!dot || dot == filename) {
        return false;
    }
    return strcmp(dot, ".c") == 0;
}

string extractComments(string& input) {
    size_t posEnd = input.find("//");
    
    if (posEnd != std::string::npos) {
        return input.substr(0, posEnd);
    }

    return input;
}

int main()
{
    DIR* dirp = opendir("/home/skor/GoStudy/CProjects/abra");
    struct dirent* file;

    for(int i = 0; i < 6; i++)
    {
        file = readdir(dirp);

        vector<string> list;

        if (isCFile(file->d_name))
        {

            string path = "/home/skor/GoStudy/CProjects/abra/" + string(file->d_name);
            ifstream f(path);
            string line;

            while (getline(f,line))
            {
                string str = extractComments(line);

                list.push_back(str);
            }

            f.close();
        }

        string path = "/home/skor/GoStudy/CProjects/abra/" + string(file->d_name);
        ofstream f(path);

        for(int i = 0; i < list.size(); i++)
        {
            f << list[i] << endl;
        }

        f.close();
    }

    closedir(dirp);

    return 0;
}
