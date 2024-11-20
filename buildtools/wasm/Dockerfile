FROM debian:12-slim

RUN apt-get update && apt-get install -y binaryen build-essential curl jq unzip xz-utils zstd

RUN curl https://gitlab.haskell.org/ghc/ghc-wasm-meta/-/raw/master/bootstrap.sh | sh

WORKDIR /shellcheck
ADD buildtools/wasm/version.txt version.txt
RUN curl -L https://github.com/koalaman/shellcheck/archive/refs/tags/$(cat version.txt).tar.gz | tar -xz --strip-components 1 -C /shellcheck

RUN ./striptests
RUN . ~/.ghc-wasm/env && wasm32-wasi-cabal update && wasm32-wasi-cabal build --allow-newer shellcheck

RUN wasm-opt -o shellcheck.wasm --flatten --rereloop --converge -O3 dist-newstyle/build/wasm32-wasi/ghc-*/ShellCheck-*/x/shellcheck/build/shellcheck/shellcheck.wasm

CMD ["cp", "shellcheck.wasm", "/out/"]
