echo 'export GOPATH=$HOME/Go' >> $HOME/.bashrc source $HOME/.bashrc

# protoc --go_out=. --go_opt=paths=source_relative ./protos/*
protoc --go_out=api/rpc --go_opt=paths=source_relative --go-grpc_out=api/rpc  --go-grpc_opt=paths=source_relative ./proto/payment/*

echo "Copying proto files into api/rpc"
cp -R api/rpc/proto/payment/* api/rpc

echo "Removing unused directories"
rm -rf api/rpc/proto
rm -rf api/rpc/proto/payment

echo "DONE"