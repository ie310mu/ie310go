// This test checks reverse calls.
//反向调用，让服务端调用自己的某个方法
//callMeBack：通知服务端调用自己的某个方法，成功后返回

//通知服务端callMeBack，方法名foo，参数[1]
--> {"jsonrpc":"2.0","id":2,"method":"test_callMeBack","params":["foo",[1]]}

//收到服务端返回调用自己的方法的请求
<-- {"jsonrpc":"2.0","id":1,"method":"foo","params":[1]}

//模拟向服务端返回数据
--> {"jsonrpc":"2.0","id":1,"result":"my result"}

//服务端又将数据返回到客户端
<-- {"jsonrpc":"2.0","id":2,"result":"my result"}
