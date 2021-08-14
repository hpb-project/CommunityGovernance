#!/bin/bash 
source env.sh

./govermentSet -u $url -proxy $proxy -priv $PK -get $@
