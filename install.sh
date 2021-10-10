#!/usr/bin/env bash

# Add deps, just in case you haven't already
apt update
apt install sudo build-essential libssl-dev libcurl4-openssl-dev libjansson-dev libgmp-dev automake git wget -y
sudo wget https://golang.org/dl/go1.17.1.linux-amd64.tar.gz;
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.1.linux-amd64.tar.gz;
sudo rm -rf go1.17.1.linux-amd64.tar.gz;
export PATH=$PATH:/usr/local/go/bin;
if [ "go version" == "" ]; then
    echo "Go installation failed";
fi
cd /tmp;
git clone https://github.com/altopm/alto.git;
cd alto;
go build;
sudo cp -r /tmp/alto /usr/local/bin/;
cd ~/;
rm -rf /tmp/alto;
export PATH=$PATH:/usr/local/bin/;
printf "\n\nAlto installed successfully!\n";