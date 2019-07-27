// This test checks basic subscription support.

//对nftest的someSubscription进行订阅，并传入参数，让期模拟5个发布
--> {"jsonrpc":"2.0","id":1,"method":"nftest_subscribe","params":["someSubscription",5,1]}

//得到订阅的id
<-- {"jsonrpc":"2.0","id":1,"result":"0x1"}

//得到5次通知
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":1}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":2}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":3}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":4}}
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":5}}

--> {"jsonrpc":"2.0","id":2,"method":"nftest_echo","params":[11]}
<-- {"jsonrpc":"2.0","id":2,"result":11}


//订阅abc
--> {"jsonrpc":"2.0","id":3,"method":"nftest_subscribe","params":["abc","abccccccccccccccccc"]}
//返回订阅信息
<-- {"jsonrpc":"2.0","id":3,"result":"0x2"}
//得到通知
<-- {"jsonrpc":"2.0","method":"nftest_subscription","params":{"subscription":"0x1","result":1}}
