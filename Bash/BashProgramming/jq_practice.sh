#!/bin/bash

FILE_PATH="./data.json"

find_developers () {
    echo "Here are the current active dev team members"
    jq -r '.users[] | select((.roles[] == "developer")  and (.active == true)) | .name' $FILE_PATH
}

find_source () {
    echo "Current Repo Source:"
    jq '.meta.source' $FILE_PATH
}

main () {
    while ((1 == 1)); do
        read -p "What would you like to parse? Devs or Source " choice
        case "$choice" in
            "Devs")
                find_developers
                ;;
            "Source")
                find_source
                ;;
            *)
                echo "Unknown option..."

        esac
    done
}

main

