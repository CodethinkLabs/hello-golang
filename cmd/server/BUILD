load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/CodethinkLabs/hello-golang/cmd/server",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/proto:go_default_library",
        "//pkg/server:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_binary(
    name = "main",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
