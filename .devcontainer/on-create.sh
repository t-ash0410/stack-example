#!/bin/bash

set -euxo pipefail

(
  cd /workspace/ts
  bun install --frozen-lockfile
)
