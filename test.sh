#!/bin/bash

max=25

for i in `seq 1 $max`; do
    val=`printf %02d $i`
    if [ -d "$val" ]; then
        is_part=false
        if [ -e "./$val/.part" ]; then
            echo "Day $val is only partly done, skipping part 2 tests..."
            is_part=true
        fi
        bash "./$val/test.sh"
        if [ "$?" != 0 ]; then
            >&2 echo "Day $val failed"
            exit 1
        fi
        if [ "$is_part" = true ]; then
            echo "Day $val part 1 passed!"
        else
            echo "Day $val passed!"
        fi
    fi
done
