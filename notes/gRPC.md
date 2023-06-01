# gRPC Golang
// protoc --go_out=. --go-grpc_out=. chat.proto <- command to compile for grpc
1. Download the right Protocol Buffers binary for your OS. At this time the zip is called `protoc-3.20.0-win64.zip`.
    Here is a link: [Protobuf](https://github.com/protocolbuffers/protobuf/releases/tag/v3.20.0)

2. Once downloaded I created a folder in my `C:/Projects` folder called `protobuf` so `C:/Projects/protobuf` is where I stored the protobuf `bin` folder and `include` folder for safe keepin.

3. Next is setting up the environment variable. I had to put the protobuf `bin` folder in my `C:/` drive, so `C:/bin` is where I had to store this folder in order to get global access with through my "Path" environment variable. Then after that I was able to setup my workspace env.

4. Make your `.proto` file then create you a `go mod init` file. Then install the support packages for protobuf-go:
    ```
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
    ```

5. Be sure to include the `include` folder in the build path, then build your protobuf stuff by going to the dir where the `.proto` file is located and then using `protoc --go_out=. --go_opt_grpc=. {name of protofile}` to create your gRPC stuff 

6. Write your code then `go mod tidy` to update imports
