#!/bin/bash

serviceName = ${1}

cd app/${serviceName}
cwgo client -I ../idl --type RPC --service ${serviceName} --module github.com/bz-2021/gomall/app/${serviceName} --idl ../idl/${serviceName}.proto
cwgo server -I ../idl --type RPC --service ${serviceName} --module github.com/bz-2021/gomall/app/${serviceName} --idl ../idl/${serviceName}.proto