#!/bin/bash

usage() { echo "Usage: $0 [-v <version of GO (default 1.19.5)>]"; }

OPTIND=1
shift $((OPTIND-1))

GO_version="1.19.5"
while getopts ":v:" opt; do
    case "${opt}" in
        v)
            GO_version=${OPTARG}
            ;;
        *)
            usage
            ;;
    esac
done

if [ -z "${v}" ]; then
    usage
fi

echo "version = $GO_version"

GO_directory="/usr/tmp-downloads/"
GO_archive="$GOversion.tar.gz"
GO_extract="go19-extracted"
GO_install_version="go-$GO_version"
GO_install_dir="/usr/bin/go"

sudo mkdir $GO_directory

sudo wget https://go.dev/dl/go$GO_version.linux-amd64.tar.gz -O $GO_directory$GO_archive
sudo mkdir $GO_directory$GO_extract
sudo tar -xf $GO_directory$GO_archive -C $GO_directory$GO_extract

sudo rm -rf $GO_install_dir/$GO_install_version
sudo mkdir $GO_install_dir
sudo mkdir $GO_install_dir/$GO_install_version
sudo cp -R $GO_directory$GO_extract/go/* $GO_install_dir/$GO_install_version

sudo rm -rf $GO_directory$GO_extract
sudo rm -rf $GO_directory$GO_archive

while [ "$(sudo grep -n "export GOROOT=" ~/.bashrc | cut -d: -f1)" != "" ]
do 
 line_to_cut="$(sudo grep -n "export GOROOT=" ~/.bashrc | cut -d: -f1)"
 line="$(echo $line_to_cut | cut -d " " -f1)"
 sudo sed -i $line"d" ~/.bashrc 
done

while [ "$(sudo grep -n "export GOPATH=" ~/.bashrc | cut -d: -f1)" != "" ]
do 
 line_to_cut="$(sudo grep -n "export GOPATH=" ~/.bashrc | cut -d: -f1)"
 line="$(echo $line_to_cut | cut -d " " -f1)"
 sudo sed -i $line"d" ~/.bashrc 
done

while [ "$(sudo grep -n "export PATH=" ~/.profile| cut -d: -f1)" != "" ]
do 
 line_to_cut="$(sudo grep -n "export PATH=" ~/.profile| cut -d: -f1)"
 line="$(echo $line_to_cut | cut -d " " -f1)"
 sudo sed -i $line"d" ~/.profile
done

sudo echo "export GOROOT=$GO_install_dir/$GO_install_version" >> ~/.bashrc
sudo echo "export GOPATH=$GO_install_dir" >> ~/.bashrc
sudo echo "export PATH=\$PATH:$GO_install_dir/$GO_install_version/bin" >> ~/.profile
source ~/.bashrc
source ~/.profile
