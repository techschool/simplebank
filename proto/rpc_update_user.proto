syntax = "proto3";

package pb;

import "user.proto";

option go_package = "github.com/techschool/simplebank/pb";

message UpdateUserRequest {
    string username = 1;
    optional string full_name = 2;
    optional string email = 3;
    optional string password = 4;
}

message UpdateUserResponse {
    User user = 1;
}
