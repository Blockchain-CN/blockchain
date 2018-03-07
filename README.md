# blockchain
Just a blockchain, Nothing in particular.

## 二．General design

	main.go // 入口
	server  // 传输层入口 
		- http      // HTTP server to support remote operation
		- command line // standard inputs to support local operation
	handlers        // 函数入口层
		- transaction
        - append
	models
		- block
	protocal
	    - status machine
	common
	    - idl
		- errno
		- const
		- util