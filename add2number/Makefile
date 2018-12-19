CC=g++
CCFLAGS= -c -g -Wall -std=c++11
OBJECT=main.o add2num.o

add2number: ${OBJECT}
	$(CC) -o $@ ${OBJECT} 

main.o: main.cpp
	$(CC) $(CCFLAGS) $<  -o $@

add2num.o: add2num.cpp
	$(CC) $(CCFLAGS)  $< -o $@


clear:
	rm *.o add2number
