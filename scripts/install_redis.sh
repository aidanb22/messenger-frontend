#!/bin/bash

sudo apt -y update
sudo apt install -y redis-server
sudo sed -i 's/supervised no/supervised systemd/g' /etc/redis/redis.conf
sudo sed -i 's/# requirepass foobared/requirepass Foobared!/g' /etc/redis/redis.conf
sudo systemctl restart redis.service