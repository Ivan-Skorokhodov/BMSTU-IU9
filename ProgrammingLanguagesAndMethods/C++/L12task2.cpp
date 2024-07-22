#include <iostream>
#include <sys/types.h>
#include <dirent.h>
#include <string>
#include <cstring>
#include <fstream>
#include <vector>
#include <algorithm>

using namespace std;

bool isHtmlFile(char* filename) {
    char* dot = strrchr(filename, '.');
    if (!dot || dot == filename) {
        return false;
    }
    return strcmp(dot, ".html") == 0;
}

string extractRef(string& input) {
    size_t posStart = input.find("=") + 2;
    size_t posEnd = input.find(">", posStart, 1) - 1;
    
    if (posStart != std::string::npos && posEnd != std::string::npos) {
        return input.substr(posStart, posEnd - posStart);
    }

    return "";
}

int main()
{
    DIR* dirp = opendir("/home/skor/GoStudy/CProjects/abra");
    struct dirent* file;

    vector<string> listLinks;

    for(int i = 0; i < 7; i++)
    {
        file = readdir(dirp);

        if (isHtmlFile(file->d_name))
        {
            string path = "/home/skor/GoStudy/CProjects/abra/" + string(file->d_name);
            ifstream f(path);
            string line;

            cout << file->d_name << endl; // !!!!!!!!!!!!!!!!!!!!!!!!!

            while (getline(f,line))
            {
                string str = extractRef(line);
                bool flag = true;
                for(int i = 0; i < listLinks.size(); i++)
                {
                    if (str == listLinks[i])
                    {
                        flag = false;
                        break;
                    }
                }
                
                if (flag)
                {
                    listLinks.push_back(str);
                }

                cout << line << endl;
                cout << str << endl;
            }

            f.close();
        }
    }

    sort(listLinks.begin(), listLinks.end());

    ofstream fLinks("/home/skor/GoStudy/CProjects/abra/links.txt");

    for(int i = 0; i < listLinks.size(); i++)
    {
        fLinks << listLinks[i] << endl;
    }

    fLinks.close();

    closedir(dirp);

    return 0;
}