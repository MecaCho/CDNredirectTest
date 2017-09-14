#!/bin/bash
echo "Test lambda..."
sleep 1
for i in $(seq 1 100)
do
 time curl http://162.3.111.24:5002/p2
 time curl http://162.3.111.24:5002/p2
 done
sleep 1
echo "#################################################Test without Lambda####################################"
sleep 1
for i in $(seq 1 100)
do
 time curl http://162.3.111.24:5002/p3
done
#time curl http://162.3.111.24:5001/p2
