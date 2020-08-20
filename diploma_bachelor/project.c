#include <sys/types.h>
#include <time.h>
#include <stdlib.h>
#include <stdio.h>
#include <fcntl.h>
#include <unistd.h>
#include <sys/stat.h>
#include <signal.h>
#define HOT 11
#define COLD 4
#define MAX 100

#define DELAY 1500 // Delay between reading in msec

//#define DEBUG

void sign(int sign){
    switch(sig){
	case SIGTERM:
	    //fclose(fp_cold);
	    //fclose(fp_hot);
	    puts("Process is completed.\n");
	    exit(0);
	default:
	    printf("Wrong process (%d).\n", sign);
    }
}

int gpio_init(int port){
    FILE *fd;
    char filename[MAX];

#ifndef DEBUG
    if((fd = fopen("/sys/class/gpio/export", "w" )) == NULL){
#else
    if((fd = fopen("/home/iguana98/KR/export", "w+" )) == NULL){
#endif
	fprintf(stderr, "\nError: can't init port: %d.\n", port);
	return -1;
    }

    fprintf(fd, "%d\n", port);
    fclose(fd);

#ifndef DEBUG
    snprintf(filename, sizeof(filename), "/sys/class/gpio/gpio%d/direction", port);
#else
    snprintf(filename, sizeof(filename), "/home/iguana98/gpio/gpio%d/direction", port);
#endif
    fd = fopen(filename, "w");
    if(fd == NULL)
    {
	printf("Error! Can't set direction on port: %d!\n", port);
	return -1;
    }
    fprintf(fd, "in");
    fclose(fd);
    return 0;
}

int logging(int n, int watertype){
    time_t seconds;
    FILE *file;

    seconds = time(NULL);
    if((file = fopen("monitoring.log", "a")) == NULL){
	puts("Cannot open file.");
	return -1;
    }

    fprintf(file, "%d;%s;%d\n", (int)seconds, (watertype==COLD ? "cold water" : "hot water"), n);
    puts("Done!");
    fclose(file);
}

int gpio_open(int port){
    int c;
    char ch;
    char filename[100];
    FILE* fd;

#ifndef DEBUG
    snprintf(filename, sizeof(filename), "/sys/class/gpio/gpio%d/value", port);
#else
    snprintf(filename, sizeof(filename), "/home/iguana98/gpio/gpio%d/value", port);
#endif
    if((fd = fopen(filename, "w+")) == NULL){
	printf("\nCannot open file (GPIO).\n");
	return -1;
    }

    fread(&ch, sizeof(ch), 1, fd);
    fclose(fd);
    c = ch - '0';
    return c;

}

int main(){
    int cold, currentcold = 0;
    int hot, currenthot = 0;
    int err;
    //FILE* fp_hot;
    //FILE* fp_cold;

    signal(SIGTERM, &sign);
    err  = gpio_init(COLD);
    err += gpio_init(HOT);
    if(err)
    {
	printf("Error!\n");
	return -1;
    }

    while(1){
	cold = gpio_open(COLD);
	usleep(DELAY);
	if(cold != currentcold){
	    logging(cold, COLD);
	    printf("Cold water is flow: %d\n", currentcold);
	    currentcold = cold;
	}
	hot = gpio_open(HOT);
	usleep(DELAY);
	if(hot != currenthot){
	    logging(hot, HOT);
	    printf("Hot water is flow: %d\n", currenthot);
	    currenthot = hot;
	}
    }
}