#!/bin/bash
go build ./src/spa_app/cmd/app
sudo systemctl start spaapp
sudo systemctl enable spaapp
