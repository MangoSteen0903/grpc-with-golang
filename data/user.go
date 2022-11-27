package data

import (
	pb "github.com/MangoSteen0903/grpc-go-tutorial/user"
)

var UserData = []*pb.UserMessage{
	{
		UserId:      "1",
		Name:        "Milky",
		PhoneNumber: "010-321232-123",
		Age:         22,
	},
	{
		UserId:      "2",
		Name:        "Kim",
		PhoneNumber: "2313-123",
		Age:         12,
	},
}
