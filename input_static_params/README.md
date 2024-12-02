1. 程序内定义变量
2. `go build` 的时候使用参数 `-ldflags` 传入变量的值，注意传入的时候的，变量路径是相对的整个包module的路径，main内的可以直接使用 main.xxx=aabb 来传入
3. 例子
    ```
   go build -ldflags="-X 'github.com/qoqozhang/go-basic-test.git/input_static_params/version.Version=$(git rev-parse --short HEAD)' -X main.mainParam=123123"
   
   .\input_static_params.exe
   main:  123123
   GitCommit:  6561fd0
   ```