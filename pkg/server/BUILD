load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "concat.go",
        "message.go",
        "yo.go",
    ],
    importpath = "github.com/CodethinkLabs/hello-golang/pkg/server",
    visibility = ["//visibility:public"],
    deps = ["//pkg/proto:go_default_library"],
)
