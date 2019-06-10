load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "helloworld_proto",
    srcs = ["helloworld.proto"],
)

go_proto_library(
    name = "helloworld_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "proto",
    proto = ":helloworld_proto",
    visibility = ["//visibility:public"],
)
