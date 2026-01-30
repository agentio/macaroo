#!/bin/bash

nonce=alfabravocharlie
secret=12345678

m0=$(macaroo create $nonce $secret)
m1=$(macaroo extend $m0 alfa)
m2=$(macaroo extend $m1 bravo)
m3=$(macaroo extend $m2 charlie)

echo
echo "m0 ($(echo $m0 | wc -c) characters)"
echo $m0

echo
echo "m1 ($(echo $m1 | wc -c) characters)"
echo $m1

echo
echo "m2 ($(echo $m2 | wc -c) characters)"
echo $m2

echo
echo "m3 ($(echo $m3 | wc -c) characters)"
echo $m3

echo
echo "verifying m3"
macaroo verify $m3 $secret

echo
echo "evaluating m3"
macaroo evaluate $m3
