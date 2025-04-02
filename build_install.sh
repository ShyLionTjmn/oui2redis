#!/bin/sh
cd build
go build ../ && sudo install oui2redis /usr/local/sbin/
