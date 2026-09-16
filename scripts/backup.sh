#!/bin/bash
# ANGEL Backup Script
# Usage: bash scripts/backup.sh [backup_dir]

BACKUP_DIR="${1:-/opt/ANGEL/backups}"
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="$BACKUP_DIR/angel-backup-$DATE.tar.gz"

mkdir -p "$BACKUP_DIR"

echo "=== ANGEL Backup ==="
echo "Backup dir: $BACKUP_DIR"
echo "Backup file: $BACKUP_FILE"

# Backup lab data
tar -czf "$BACKUP_FILE" \
    lab/data/ \
    lab/logs/ \
    lab/implants/ \
    lab/certs/ \
    lab/configs/ \
    2>/dev/null

# Backup config files
tar -czf "$BACKUP_FILE" \
    --append \
    docker-compose.yml \
    Dockerfile.angel \
    lab.env \
    .env.example \
    Makefile \
    2>/dev/null

# Backup database (if exists)
if [ -d "lab/data/db" ]; then
    tar -czf "$BACKUP_FILE" \
        --append \
        lab/data/db/ \
        2>/dev/null
fi

# Verify backup
if [ -f "$BACKUP_FILE" ]; then
    SIZE=$(du -h "$BACKUP_FILE" | cut -f1)
    echo "Backup created: $BACKUP_FILE ($SIZE)"
else
    echo "ERROR: Backup failed"
    exit 1
fi

# Cleanup old backups (keep last 10)
cd "$BACKUP_DIR" && ls -t angel-backup-*.tar.gz 2>/dev/null | tail -n +11 | xargs -r rm -f

echo "Backup complete"