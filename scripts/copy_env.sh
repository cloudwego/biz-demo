#!/bin/bash

. scripts/list_app.sh

get_app_list

readonly root_path=`pwd`
for app_path in ${app_list[*]}; do
  if [[ "${app_path}" = "app/common" ]]; then
    continue
  fi
  if [[  -e "${app_path}/.env" ]]; then
      continue
    fi
  echo "copy ${app_path} env file"
  cp "${app_path}/.env.example" "${app_path}/.env"
  echo "Done! Please replace the real value"
done