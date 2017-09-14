#!/bin/bash
echo "Test lambda..."
sleep 1
for i in $(seq 1 100)
do
	if
 (time curl http://162.3.111.24:5002/p2) 2>&1 | awk '{print $2}' | sed -n '6p' | awk -F 'm' '{print $2}' | cut -d 's' -f 1 >>output.txt
 done
sleep 1
echo "#################################################Test without Lambda####################################"
sleep 1
for i in $(seq 1 100)
do
	(time curl http://162.3.111.24:5002/p3) 2>&1 | awk '{print $2}' | sed -n '6p' | awk -F 'm' '{print $2}' | cut -d 's' -f 1 >>output.txt
done
#time curl http://162.3.111.24:5001/p2
