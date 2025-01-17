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

    struct sockaddr_in peer;
    peer.sin_family = AF_INET;
    peer.sin_port = htons(6060);
    peer.sin_addr.s_addr = inet_addr("127.0.0.1");

    if (connect(s, (struct sockaddr *)&peer, sizeof(peer)) < 0)
    {
        fprintf(stderr, "\"connect\" returned error\n");
        return 1;
    }

    if (send(s, "A", 1, 0) <= 0)
    {
        fprintf(stderr, "send data error\n");
        return 1;
    }

    if (shutdown(s, SHUT_WR) < 0)
    {
        fprintf(stderr, "\"shutdown\" returned error\n");
        return 1;
    }

    char buf[1];
    if (recv(s, buf, 1, 0) <= 0)
    {
        fprintf(stderr, "request data error\n");
        return 1;
    }

    printf("%c\n", buf[0]);

    if (close(s) < 0)
    {
        fprintf(stderr, "\"close\" returned error\n");
        return 1;
    }

    return 0;
}