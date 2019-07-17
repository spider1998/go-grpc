package main

import (
	`fmt`
	pb `go-grpc/rpc-peotobuf/costomer`
	`golang.org/x/net/context`
	`google.golang.org/grpc`
	`io`
	`log`
)

const address  = "localhost:50051"

func CreateCustomer(client pb.CustomerClient,customer *pb.CustomerRequest)  {
	resp,err := client.CreateCustomer(context.Background(),customer)
	if err != nil{
		log.Fatalf("Could not create Customer: %v",err)
	}
	if resp.Success{
		log.Printf("A new Customer has been add with id : %v",resp.Id)
	}
}

func getCustomers(client pb.CustomerClient,filter *pb.CustomerFilter)  {
	stream,err := client.GetCustomers(context.Background(),filter)
	if err != nil{
		log.Fatalf("Error on get customers: %v",err)
	}
	for{
		customer,err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("%v.GetCustomers(_) = _,%v",client,err)
		}
		log.Printf("Customer: %v",customer)
	}
}

func main()  {
	conn,err := grpc.Dial(address,grpc.WithInsecure())
	if err != nil{
		log.Fatalf("did not connect: %v",err)
	}
	defer conn.Close()

	client := pb.NewCustomerClient(conn)
	customer := &pb.CustomerRequest{
		Id:101,
		Name:"spider1998",
		Email:"2387805574@qq.com",
		Phone:"17609270263",
		Address:[]*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:"1 street",
				City:"Xian",
				State:"CA",
				Zip:"98924",
				IsShippingAddress:false,
			},
			&pb.CustomerRequest_Address{
				Street:"2 street",
				City:"XianYang",
				State:"KL",
				Zip:"98925",
				IsShippingAddress:true,
			},
		},
	}

	CreateCustomer(client,customer)
	customer = &pb.CustomerRequest{
		Id:102,
		Name:"spider1999",
		Email:"2387805575@qq.com",
		Phone:"17609270264",
		Address:[]*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:"1 street",
				City:"Xian",
				State:"CA",
				Zip:"98924",
				IsShippingAddress:true,
			},
		},
	}
	CreateCustomer(client,customer)

	fmt.Printf("xxxxxxxxxxxxxxxxxxxxx")

	filter := &pb.CustomerFilter{
		Keyword:"",
	}

	getCustomers(client,filter)

}