#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>

int main()
{
    int s = socket(AF_INET, SOCK_STREAM, 0);
    if (s < 0)
    {
        fprintf(stderr, "\"socket\" returned error\n");
        return 1;
    }

    struct sockaddr_in addr;
    addr.sin_family = AF_INET;
    addr.sin_port = htons(6060);
    addr.sin_addr.s_addr = inet_addr("127.0.0.1");

    if (bind(s, (struct sockaddr *)&addr, sizeof(addr)) < 0)
    {
        fprintf(stderr, "\"bind\" returned error\n");
        return 1;
    }
    if (listen(s, 32) < 0)
    {
        fprintf(stderr, "\"listen\" returned error\n");
        return 1;
    }

    struct sockaddr_in client_addr;
    socklen_t client_addr_size = sizeof(client_addr);
    int s2 = accept(s, (struct sockaddr *)&client_addr, &client_addr_size);
    if (s2 < 0)
    {
        fprintf(stderr, "\"accept\" returned error\n");
        return 1;
    }

    char buf[1];
    if (recv(s2, buf, 1, 0) <= 0)
    {
        fprintf(stderr, "\"request\" data error\n");
        return 1;
    }

    printf("%c from %s:%hu\n", buf[0],
           inet_ntoa(client_addr.sin_addr),
           ntohs(client_addr.sin_port));

    if (send(s2, "B", 1, 0) < 0)
    {
        fprintf(stderr, "send data error\n");
        return 1;
    }
    if (shutdown(s2, SHUT_WR) < 0)
    {
        fprintf(stderr, "\"shutdown\" returned error\n");
        return 1;
    }
    if (close(s) < 0 || close(s2) < 0)
    {
        fprintf(stderr, "\"close\" returned error\n");
        return 1;
    }
    return 0;
}