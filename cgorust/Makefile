MANIFEST_OPT = --manifest-path=../person-rs/Cargo.toml

all: main run

main: lib
	go build

lib:
	cargo build ${MANIFEST_OPT}

test:
	cargo test  ${MANIFEST_OPT}

clean:
	go clean
	cargo clean --release  ${MANIFEST_OPT}
	cargo clean  ${MANIFEST_OPT}

run:
	./cgorust
