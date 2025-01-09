#!/bin/bash

set -euxo pipefail

(
  cd /workspace/ts/web
  bun install --frozen-lockfile
)
