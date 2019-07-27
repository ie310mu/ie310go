// This test checks reverse calls.

//通知服务端callMeBackLater，方法名foo，参数[1]
--> {"jsonrpc":"2.0","id":2,"method":"test_callMeBackLater","params":["foo",[1]]}

//马上收到服务端返回的空结果  callMeBack中，这个是最后才会返回，而且有正确的结果，这里没有结果
<-- {"jsonrpc":"2.0","id":2,"result":null}

//又收到服务端的调用请求
<-- {"jsonrpc":"2.0","id":1,"method":"foo","params":[1]}

//模拟输出结果到服务端
--> {"jsonrpc":"2.0","id":1,"result":"my result"}