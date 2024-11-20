package runner

import (
	"context"
	"crypto/rand"
	"io"
	"log"
	"os"
	"strings"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/experimental/sysfs"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	wzsys "github.com/tetratelabs/wazero/sys"
)

func Run(name string, args []string, wasm []byte, stdin io.Reader, stdout io.Writer, stderr io.Writer, cwd string) int {
	ctx := context.Background()

	rt := wazero.NewRuntime(ctx)

	wasi_snapshot_preview1.MustInstantiate(ctx, rt)

	args = append([]string{name}, args...)

	root := sysfs.DirFS(cwd)

	cfg := wazero.NewModuleConfig().
		WithSysNanosleep().
		WithSysNanotime().
		WithSysWalltime().
		WithStderr(stderr).
		WithStdout(stdout).
		WithStdin(stdin).
		WithRandSource(rand.Reader).
		WithArgs(args...).
		WithFSConfig(wazero.NewFSConfig().(sysfs.FSConfig).WithSysFSMount(root, "/"))
	for _, env := range os.Environ() {
		k, v, _ := strings.Cut(env, "=")
		if k == "PWD" {
			// Causes Haskell runtime to try chdir which isn't supported in wasip1
			continue
		}
		cfg = cfg.WithEnv(k, v)
	}

	_, err := rt.InstantiateWithConfig(ctx, wasm, cfg)
	if err != nil {
		if sErr, ok := err.(*wzsys.ExitError); ok { //nolint:errorlint
			return int(sErr.ExitCode())
		}
		log.Fatal(err)
	}
	return 0
}
