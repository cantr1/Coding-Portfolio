#!/bin/bash

MENU="1.) Generate Hash / API Key
2.) Write String to File
3.) View System Resources
q.) Exit the program
"

VALID_OPTIONS=( "1" "2" "3" "q" )

function collect_user_input()
{
    while true; do
        read -rp "$MENU
Select an option: " user_input

        if validate_user_input "$user_input"; then
            echo "$user_input"
            return 0
        else
            echo -e "invalid option...\n" >&2
        fi
    done
}

function validate_user_input()
{
    user_input="${1:-""}"
    for opt in "${VALID_OPTIONS[@]}"; do
        if [ "$user_input" = "$opt" ]; then
            return 0
        fi
    done
    return 1
}

function create_api_key()
{
    key=$(openssl rand -hex 24)
    echo -e "Here is your key: $key\n"
}

function write_string()
{
    read -rp "Enter file location: " file_location
    read -rp "Enter string: " string
    echo "$string" > "$file_location"
}

function get_system_resources()
{
    echo "Storage"
    df -h
    echo ""
    echo "Machine Info"
    uname -a
    echo ""
}

function main()
{
    while true; do
        opt=$(collect_user_input)

        case "$opt" in
            "1") create_api_key;;
            "2") write_string;;
            "3") get_system_resources;;
            "q") exit 0;;
        esac
    done
}

main