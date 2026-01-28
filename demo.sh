#!/bin/bash

secret=12345678

m1=$(macaroo create $secret)
m2=$(macaroo extend $m1 alfa)
m3=$(macaroo extend $m2 bravo)
m4=$(macaroo extend $m3 charlie)

echo $m1
echo $m2
echo $m3
echo $m4

macaroo verify $m4 $secret
