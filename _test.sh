#!/bin/bash
echo "Test lambda..."
sleep 1
for i in $(seq 1 5)
do
	var = $(time curl http://162.3.111.24:5001/p2)
echo $var
# time curl http://162.3.111.24:5001/p2
 done
sleep 1
echo "#################################################Test without Lambda####################################"
sleep 1
for i in $(seq 1 5)
do
 #time curl http://162.3.111.24:5001/p3
done
#time curl http://162.3.111.24:5001/p2
