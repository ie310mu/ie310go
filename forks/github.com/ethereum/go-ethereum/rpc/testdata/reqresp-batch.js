// There is no response for all-notification batches.

//这个为什么返回超时错误?  read pipe: deadline exceeded  
//原始参数中没有指定ID,现在指定了就好了
--> [{"jsonrpc":"2.0","id":1,"method":"test_echo","params":["x",99]}]
<-- [{"jsonrpc":"2.0","id":1,"result":{"String":"x","Int":99,"Args":null}}]

// This test checks regular batch calls.

--> [{"jsonrpc":"2.0","id":2,"method":"test_echo","params":[]}, {"jsonrpc":"2.0","id": 3,"method":"test_echo","params":["x",3]}]
<-- [{"jsonrpc":"2.0","id":2,"error":{"code":-32602,"message":"missing value for required argument 0"}},{"jsonrpc":"2.0","id":3,"result":{"String":"x","Int":3,"Args":null}}]
