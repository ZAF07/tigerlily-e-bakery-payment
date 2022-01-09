echo 'export GOPATH=$HOME/Go' >> $HOME/.bashrc source $HOME/.bashrc

# protoc --go_out=. --go_opt=paths=source_relative ./protos/*
protoc --go_out=api/rpc --go_opt=paths=source_relative --go-grpc_out=api/rpc  --go-grpc_opt=paths=source_relative ./proto/*



echo "DONE"