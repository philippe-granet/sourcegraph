load("@npm//:mocha/package_json.bzl", "bin")

def mocha_test(name, srcs, args = [], data = [], env = {}, **kwargs):
    bin.mocha_test(
        name = name,
        args = [
            native.package_name() + "/**/*.test.js",
        ] + args,
        data = data + srcs + [
            "//:mocha_config",
            # "//examples/macro:mocha_reporters.json",
            # "//examples/macro:node_modules/mocha-multi-reporters",
            # "//examples/macro:node_modules/mocha-junit-reporter",
        ],
        env = dict(env, **{
            # Add environment variable so that mocha writes its test xml
            # to the location Bazel expects.
            "MOCHA_FILE": "$$XML_OUTPUT_FILE",
        }),
        tags = ["manual"],
        **kwargs
    )
