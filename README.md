环境搭建相关的命令：

docker run -u root --volume ~/Developer:/code -it golang:bullseye

进入docker 容器之后再安装

apt update
apt install -y gcc-10-riscv64-linux-gnu qemu-user
ln -sf /usr/bin/riscv64-linux-gnu-gcc-10 /usr/bin/riscv64-linux-gnu-gcc




make test 运行验证