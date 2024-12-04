# 生成服务代码
#
# -module 参数表明生成代码的 go mod 中的 module name，在本例中为 example_shop
# -service 参数表明我们要生成脚手架代码，后面紧跟的 example.shop.item 或 example.shop.stock 为该服务的名字。
# -use 参数表示让 kitex 不生成 kitex_gen 目录，而使用该选项给出的 import path。在本例中因为第一次已经生成 kitex_gen 目录了，后面都可以复用。
# 最后一个参数则为该服务的 IDL 文件
#
kitex -module github.com/chhz0/go-hello -service example.shop.item -use github.com/chhz0/go-hello/cloudwego/kitex/basic/example_shop/kitex_gen ../../idl/stock.thrift