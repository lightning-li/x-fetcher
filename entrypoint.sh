#!/bin/sh
set -ex

# 启动 socat
exec socat TCP4-LISTEN:9222,fork TCP4:127.0.0.1:9223 &

# 启动 headless-shell
/headless-shell/headless-shell \
  --no-sandbox \
  --use-gl=angle \
  --use-angle=swiftshader \
  --remote-debugging-address=0.0.0.0 \
  --remote-debugging-port=9223 \
  "$@" &

# 等待 headless-shell 启动完成（可选）
sleep 2

# 运行自定义命令
exec /app/fetch_x