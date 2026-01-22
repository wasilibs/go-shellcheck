FROM debian:13-slim

RUN apt-get update && apt-get install -y binaryen build-essential curl git jq unzip xz-utils zstd

WORKDIR /ghc
# Use version before tail call extension usage
RUN curl -f -L --retry 5 https://gitlab.haskell.org/haskell-wasm/ghc-wasm-meta/-/archive/ada3b8fa0f763e4dccb2b1f6bbf2518bff2a7c6e/ghc-wasm-meta-ada3b8fa0f763e4dccb2b1f6bbf2518bff2a7c6e.tar.gz | tar xz --strip-components=1
RUN ./setup.sh

WORKDIR /shellcheck
ADD buildtools/wasm/version.txt version.txt
RUN curl -L https://github.com/koalaman/shellcheck/archive/refs/tags/$(cat version.txt).tar.gz | tar -xz --strip-components 1 -C /shellcheck

RUN ./striptests
RUN . ~/.ghc-wasm/env && wasm32-wasi-cabal update && wasm32-wasi-cabal build --allow-newer shellcheck

RUN wasm-opt -o shellcheck.wasm --flatten --rereloop --converge -O3 dist-newstyle/build/wasm32-wasi/ghc-*/ShellCheck-*/x/shellcheck/build/shellcheck/shellcheck.wasm

CMD ["cp", "shellcheck.wasm", "/out/"]
