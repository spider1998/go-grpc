package main

import (
	`fmt`
	echo `go-grpc/echo/proto`
	`golang.org/x/net/context`
	`google.golang.org/grpc`
	`os`
	`strings`
)

func main()  {
	conn,err := grpc.Dial("127.0.0.1:3000",grpc.WithInsecure())
	if err != nil{
		fmt.Errorf("dail failed. err:[%v]\n",err)
		return
	}
	client := echo.NewEchoClient(conn)
	res,err := client.Echo(context.Background(),&echo.EchoReq{
		Msg:strings.Join(os.Args[1:]," "),
	})
	if err != nil{
		fmt.Errorf("client echo failed.err: [%v]",err)
		return
	}
	fmt.Printf("message from server: %v\n",res.GetMsg())
}
