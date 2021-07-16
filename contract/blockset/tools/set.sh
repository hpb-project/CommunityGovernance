#!/bin/bash
source env.sh

k=${1:-""}
v=${2:-"0"}

./govermentSet -u $url -addr $addr -priv $PK -set $k,$v
