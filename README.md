yum -y install gcc gcc-c++

g++ -fPIC -c -o ./libs/ecdh.o ecdh.cpp -I./

ar cr ./libs/libecdh.a ./libs/ecdh.o

g++ -shared -fPIC -o ./libs/libecdh_x64.so -Wl,--whole-archive ./libs/libecdh.a ./libs/libssl.a ./libs/libcrypto.a -Wl,--no-whole-archive

安装必须的库 go get github.com/golang/protobuf/proto

cp /usr/src/wxjqr/wxjqr.service /usr/lib/systemd/system/wxjqr.service

chmod 754 /usr/lib/systemd/system/wxjqr.service

systemctl enable wxjqr.service

systemctl disable wxjqr.service

systemctl status wxjqr.service

systemctl list-units --type=service

systemctl daemon-reload

systemctl start wxjqr.service

systemctl stop wxjqr.service

systemctl reload wxjqr.service