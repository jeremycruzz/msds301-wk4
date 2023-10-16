#!/bin/bash

# Check if an argument is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <number of runs>"
    exit 1
fi

# Number of times the programs should run
num_runs=$1

total_time_go=0
total_time_go_fast=0
total_time_python=0
total_time_r=0

for i in $(seq 1 $num_runs)
do
   # Go
   start_time_go=$(date +%s%N)
   ./houseanalyzer.exe ./data/housesInput.csv ./data/benchmark/housesOutputGo.txt
   end_time_go=$(date +%s%N)
   total_time_go=$((total_time_go + end_time_go - start_time_go))

   # Go Fast
   start_time_go_fast=$(date +%s%N)
   ./houseanalyzerfast.exe
   end_time_go_fast=$(date +%s%N)
   total_time_go=$((total_time_go_fast + end_time_go - start_time_go))

   # Python
   start_time_python=$(date +%s%N)
   py ./python/runhouses.py
   end_time_python=$(date +%s%N)
   total_time_python=$((total_time_python + end_time_python - start_time_python))

   # R
   start_time_r=$(date +%s%N)
   Rscript ./r/runHouses.r
   end_time_r=$(date +%s%N)
   total_time_r=$((total_time_r + end_time_r - start_time_r))
done

# Calculate the averages
avg_time_go=$((total_time_go / num_runs))
avg_time_go_fast=$((total_time_go_fast / num_runs))
avg_time_python=$((total_time_python / num_runs))
avg_time_r=$((total_time_r / num_runs))

echo "Average time for Go: $avg_time_go nanoseconds"
echo "Average time for Go (fast): $avg_time_go_fast nanoseconds"
echo "Average time for Python: $avg_time_python nanoseconds"
echo "Average time for R: $avg_time_r nanoseconds"