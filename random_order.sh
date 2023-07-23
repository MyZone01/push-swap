#!/bin/bash

if [ $# -eq 0 ]; then
  echo "Usage: $0 n"
  exit 1
fi

n=$1

# Generate numbers from 1 to n
numbers=$(seq 1 $n)

# Shuffle the numbers randomly
random_order=$(shuf -e $numbers)

# Print the numbers separated by space
echo $random_order
