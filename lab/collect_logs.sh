#!/bin/bash
# Collect logs from Docker containers to lab/logs/
mkdir -p lab/logs

echo "Collecting logs..."

docker compose logs angel-teamserver 2>&1 > lab/logs/teamserver.log &
docker compose logs angel-console 2>&1 > lab/logs/console.log &
docker compose logs angel-rules 2>&1 > lab/logs/rules.log &

echo "Logs collected to lab/logs/"
ls -la lab/logs/