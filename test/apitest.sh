#!/bin/sh

curl -X POST -H "Content-Type: application/json" -d '{ "method": "APIService.RegisterClient", "params": [ { "Hostname": "localhost3", "Role": "controller" } ], "id": 1 }' http://localhost:8080/rpc

curl -X POST -H "Content-Type: application/json" -d '{ "method": "APIService.RegisterClient", "params": [ { "Hostname": "localhost", "Role": "controller" } ], "id": 1 }' http://localhost:8080/rpc

curl -X POST -H "Content-Type: application/json" -d '{ "method": "APIService.RegisterClient", "params": [ { "Hostname": "localhost2", "Role": "controller" } ], "id": 1 }' http://localhost:8080/rpc

curl -X POST -H "Content-Type: application/json" -d '{ "method": "APIService.ListClients", "params": [ { "count": "1" } ], "id": 1 }' http://localhost:8080/rpc
