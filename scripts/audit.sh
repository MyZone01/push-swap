#!/bin/bash
Continue() {
  echo "Press any key to continue"
  while [ true ]; do
    read -t 3 -n 1
    if [ $? = 0 ]; then
      clear
      return
    fi
  done
}

clear
echo "Try to run ./push-swap."
./build/push-swap
Continue

echo "Try to run ./push-swap '2 1 3 6 5 8' "
./build/push-swap "2 1 3 6 5 8"
echo -n "Instructions counter: " 
./build/push-swap "2 1 3 6 5 8" | wc -l
Continue

echo "Try to run ./push-swap '0 1 2 3 4 5' "
./build/push-swap "0 1 2 3 4 5"
Continue

echo "Try to run ./push-swap '0 one 2 3' "
./build/push-swap "0 one 2 3"
Continue

echo "Try to run ./push-swap '1 2 2 3' "
./build/push-swap "1 2 2 3"
Continue

echo "Try to run ./push-swap '<5 random numbers>' with 5 random numbers instead of the tag."
./pushswap.sh
Continue

echo "Try to run ./push-swap '<5 random numbers>' with 5 different random numbers instead of the tag."
./pushswap.sh
Continue

echo "Try to run ./checker and input nothing."
./build/checker
Continue

echo "Try to run ./checker '0 one 2 3' "
./build/checker "0 one 2 3"
Continue

echo "Try to run echo -e 'sa\npb\nrrr\n' | ./checker '0 9 1 8 2 7 3 6 4 5'"
echo -e "sa\npb\nrrr\n"  | ./build/checker "0 9 1 8 2 7 3 6 4 5"
Continue

echo "Try to run echo -e 'pb\nra\npb\nra\nsa\nra\npa\npa\n' | ./checker '0 9 1 8 2'"
echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./build/checker "0 9 1 8 2"
Continue

echo "Try to run ARG=('4 67 3 87 23'); ./push-swap ARG | ./checker ARG"
ARG=('4 67 3 87 23'); ./build/push-swap "$ARG" | ./build/checker "$ARG"
Continue

echo "Try to run ARG=('<100 random numbers>'); ./push-swap ARG with 100 random different numbers instead of the tag."
Continue
./pushswap.sh
Continue

echo "Try to run ARG=('<100 random numbers>'); ./push-swap ARG | ./checker ARG with 100 random different numbers instead of the tag."
./pushswap.sh
Continue