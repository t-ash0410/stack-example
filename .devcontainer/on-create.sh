#!/bin/bash

set -euxo pipefail

(
  cd /workspace/ts/web
  bun install --frozen-lockfile
)
(
  cd /workspace/ts/proxy
  bun install --frozen-lockfile
)
