- Installing protobuf compiler
  - curl -OL https://github.com/google/protobuf/releases/download/v3.2.0/protoc-3.2.0-linux-x86_64.zip
  - unzip protoc-3.2.0-linux-x86_64.zip -d protoc3
  - sudo mv protoc3/bin/* /usr/local/bin/
  - sudo mv protoc3/include/* /usr/local/include/
  - sudo chwon [user] /usr/local/bin/protoc
  - sudo chwon -R [user] /usr/local/include/google
  - https://gist.github.com/sofyanhadia/37787e5ed098c97919b8c593f0ec44d8

- Installing protobuf dependencies
  - go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

- Executing
  - Inside todo/:
    - protoc --go_out=. todo.proto
  - todo add some text
  - todo list

- Inspecting mydb.pb
  - Checking the type and contents
    - file mydb.pb
    - hexdump mydb.pb
    - hexdump -c mydb.pb
  - note: the extra fields in .pb file, which are not ascii, allow protobuf to decode the file
  - Checking the contents using protoc
    - cat mydb.pb | protoc --decode_raw
    - The special characters (eg. field position numbers) are recognised by protoc and shown in the output 
