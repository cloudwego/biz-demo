#!/bin/bash

. scripts/list_app.sh

get_app_list

for app_path in ${app_list[*]}; do
    cd $app_path &&  go fmt.
done