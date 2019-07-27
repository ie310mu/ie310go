// This test checks that an error response is sent for calls
// with named parameters. 
//不支持命名参数
--> {"jsonrpc":"2.0","method":"test_echo","params":{"int":23},"id":3}
<-- {"jsonrpc":"2.0","id":3,"error":{"code":-32602,"message":"non-array args"}}
