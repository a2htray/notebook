#!/bin/bash

curl -X POST http://localhost:1234/rpc -H "content-type: application/json" -d '{"method": "JSONServer.GiveBookDetail", "params": [{"id": "1234"}], "id": 1}'
