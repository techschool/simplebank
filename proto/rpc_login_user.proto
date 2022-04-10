syntax = "proto3";

package pb;

import "user.proto";
import "google/protobuf/timestamp.proto";

option go_package = "github.com/techschool/simplebank/pb";

message LoginUserRequest {
    string username = 1;
    string password = 2;
}

message LoginUserResponse {
    User user = 1;
    string session_id = 2;
    string access_token = 3;
    string refresh_token = 4;
    google.protobuf.Timestamp access_token_expires_at = 5;
    google.protobuf.Timestamp refresh_token_expires_at = 6;
}
