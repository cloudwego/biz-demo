#!/bin/bash

. scripts/list_app.sh

get_app_list

readonly root_path=`pwd`
for app_path in ${app_list[*]}; do
    go vet ${root_path}/${app_path}
done