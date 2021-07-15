#!/bin/bash

k=${1:-""}
v=${2:-"0"}

./govermentSet -u https://hpbnode.com -addr 0x1c478d99da808290f7734a588c5187d8ccae2d10 -priv $PK -set $k,$v
