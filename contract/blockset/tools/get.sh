#!/bin/bash 
source env.sh

./govermentSet -u $url -addr $addr -priv $PK -get $@
