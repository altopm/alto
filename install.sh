#!/usr/bin/env bash

# Get the distro type
DISTRO_TYPE="`lsb_release -i | cut -f2`"

# Add deps, just in case you haven't already
if [ $DISTRO_TYPE == "Ubuntu" || $DISTRO_TYPE == "Debian" ]; then
    sudo apt install git wget gnupg software-properties-common -y
elif [ $DISTRO_TYPE == "CentOS" ]; then
    sudo yum install git wget -y
elif [ $DISTRO_TYPE == "Fedora" ]; then
    sudo dnf install git wget -y
elif [ $DISTRO_TYPE == "Arch" ]; then
    sudo pacman -S git wget -y
fi
sudo wget https://golang.org/dl/go1.17.1.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.1.linux-amd64.tar.gz
sudo rm -rf go1.17.1.linux-amd64.tar.gz
sudo export PATH=$PATH:/usr/local/go/bin
if [ go version == "" ]; then
    echo "Go installation failed"
fi

cd /tmp
git clone https://github.com/altopm/alto.git
cd alto
go build
sudo cp -r /tmp/alto/alto /usr/local/bin/
cd ~/
rm -rf /tmp/alto
export PATH=$PATH:/usr/local/bin/
printf "\n\nAlto installed successfully!\n"