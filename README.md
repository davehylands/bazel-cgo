## bazel/rust Integration

## Prerequisites

Install [cargo-raze](https://github.com/google/cargo-raze)

```
cargo install cargo-raze
```

Pull down crates (aka create vendored crates)

```
cd person-rs
cargo vendor --versioned-dirs cargo/vendor > .cargo/config
```

Create cargo/BUILD.bazel based on the vendored crates (also executed from within the person-rs directory).
```
cargo raze
```

Now you can cd back to the root and build/run the project.
```
cd ..
bazel run //cgorust
```

and you should see some output something like this from executing cgorust:
```
Created APerson: name: "tim", long_name: "Tim Hughes"
Hello Go rust world: My name is tim, tim hughes.
Hello Go ruat world: My name is tim, tim hughes.
Dropping APerson: name: "tim", long_name: "Tim Hughes"
do_some_work called
C.myCallback_cgo: called with num = 1
my_callback: num = 1
C.myCallback_cgo: called with num = 2
my_callback: num = 2
C.myCallback_cgo: called with num = 3
my_callback: num = 3
```
