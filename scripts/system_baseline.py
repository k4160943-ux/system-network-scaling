#!/usr/bin/env python3

import json
import os
import platform
import socket
import sys
from datetime import datetime, timezone
from pathlib import Path


def collect_baseline():
    return {
        "timestamp_utc": datetime.now(timezone.utc).isoformat().replace("+00:00", "Z"),
        "host": {
            "hostname": socket.gethostname(),
            "fqdn": socket.getfqdn(),
            "os": platform.system(),
            "os_release": platform.release(),
            "os_version": platform.version(),
            "architecture": platform.machine(),
        },
        "python": {
            "version": sys.version,
            "executable": sys.executable,
        },
        "environment": {
            "cwd": os.getcwd(),
            "user": os.getenv("USERNAME") or os.getenv("USER"),
        },
    }


def write_snapshot(data):
    output_dir = Path("output")
    output_dir.mkdir(exist_ok=True)

    ts = datetime.now(timezone.utc).strftime("%Y%m%dT%H%M%SZ")
    path = output_dir / f"system_baseline_{ts}.json"

    with path.open("w", encoding="utf-8") as f:
        json.dump(data, f, indent=2)

    return path


def main():
    baseline = collect_baseline()
    path = write_snapshot(baseline)
    print(f"[baseline] snapshot written to {path}")


if __name__ == "__main__":
    main()
