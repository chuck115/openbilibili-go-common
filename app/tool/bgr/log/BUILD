package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["log.go"],
    importpath = "go-common/app/tool/bgr/log",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/tool/bgr/log/color:go_default_library",
        "//vendor/golang.org/x/crypto/ssh/terminal:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//app/tool/bgr/log/color:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
