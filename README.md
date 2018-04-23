安装必须的库
go get github.com/golang/protobuf/proto

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