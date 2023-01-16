#!/bin/bash

usage() { echo "Usage: $0 [-v <version of GO (default 1.19.5)>]"; }

OPTIND=1
shift $((OPTIND-1))

version="1.19.5"
while getopts ":v:" opt; do
    case "${opt}" in
        v)
            version=${OPTARG}
            ;;
        *)
            usage
            ;;
    esac
done

if [ -z "${v}" ]; then
    usage
fi

echo "version = $version"