#!/bin/bash

cd demo
runTests=$(go test)
runRes=$runTests
echo $runRes
failStr="FAIL"
if [[ $runRes == *$failStr* ]]; then
    curlRes=$(curl -X POST -H "Content-Type: application/json" -d '{"msg_type":"text","content":{"text":"run test failed"}}' https://open.larksuite.com/open-apis/bot/v2/hook/7796654a-f7dd-40a9-9190-6b1d40525638)
fi
echo "-------------curlRes-------------"
echo $curlRes