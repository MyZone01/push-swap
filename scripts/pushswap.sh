#!/bin/bash
echo -n "Instructions counter: "
NB=$(./random_order.sh 5); ./build/push-swap "$NB" | wc -l