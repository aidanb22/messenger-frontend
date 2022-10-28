#!/bin/bash

export RepoName=`pwd`
export APIUser=$USER

sudo apt-get -y update && sudo apt-get -y upgrade
sudo apt-get install -y git-all wget build-essential

cd ~
export HomeDir=`pwd`
wget https://dl.google.com/go/go1.14.13.linux-amd64.tar.gz
sudo tar -C $HomeDir -xzf go1.14.13.linux-amd64.tar.gz

export PATH=$PATH:$RepoName/bin
export PATH=$PATH:$RepoName/pkg
export PATH=$PATH:$RepoName/src
export GOPATH=$RepoName

echo 'export PATH=$PATH:'$HOME'/go/bin' >> $HomeDir/.profile
echo 'export PATH=$PATH:'$HOME'/go/pkg' >> $HomeDir/.profile
echo 'export PATH=$PATH:'$HOME'/go/src' >> $HomeDir/.profile
echo 'export PATH=$PATH:'$RepoName'/bin' >> $HomeDir/.profile
echo 'export PATH=$PATH:'$RepoName'/pkg' >> $HomeDir/.profile
echo 'export PATH=$PATH:'$RepoName'/src' >> $HomeDir/.profile
echo 'export GOPATH='$RepoName >> $HomeDir/.profile
source $HomeDir/.profile

sudo apt install -y golang-rice

sudo sh $RepoName/install/install.sh

rm -f ~/go1.14.13.linux-amd64.tar.gz

sudo apt -y update
sudo apt install -y redis-server
sudo sed -i 's/supervised no/supervised systemd/g' /etc/redis/redis.conf
sudo sed -i 's/# requirepass foobared/requirepass Foobared!/g' /etc/redis/redis.conf
sudo systemctl restart redis.service

sudo sh $RepoName/install/config_systemd_service.sh $RepoName $APIUser