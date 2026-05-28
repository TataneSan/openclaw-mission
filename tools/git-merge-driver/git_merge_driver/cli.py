"""CLI for git-merge-driver: configure git merge drivers for file types."""

import argparse
import os
import shutil
import subprocess
import sys
from pathlib import Path


SCRIPTS_DIR = Path.home() / ".git-merge-drivers"

BUILTIN_DRIVERS = {
    "union": {
        "cmd": 'cat "$2" > "$MERGED"',
        "desc": "Keep both versions (append theirs to ours)",
    },
    "ours": {
        "cmd": 'cp "$1" "$MERGED"',
        "desc": "Always keep our version",
    },
    "theirs": {
        "cmd": 'cp "$2" "$MERGED"',
        "desc": "Always keep their version",
    },
    "concat": {
        "cmd": 'cat "$1" > "$MERGED" && echo "" >> "$MERGED" && cat "$2" >> "$MERGED"',
        "desc": "Concatenate both versions",
    },
    "empty": {
        "cmd": 'echo "" > "$MERGED"',
        "desc": "Replace with empty file",
    },
}


def ensure_scripts_dir() -> Path:
    SCRIPTS_DIR.mkdir(exist_ok=True)
    return SCRIPTS_DIR


def add_driver(name: str, script: str | None, pattern: str | None) -> None:
    scripts_dir = ensure_scripts_dir()
    driver_file = scripts_dir / f"{name}.sh"

    if script:
        driver_file.write_text(f"#!/bin/sh\n{script}\n", encoding="utf-8")
    elif script is None:
        # Read from stdin
        content = sys.stdin.read().strip()
        driver_file.write_text(f"#!/bin/sh\n{content}\n", encoding="utf-8")

    os.chmod(driver_file, 0o755)
    print(f"Merge driver '{name}' saved to {driver_file}")

    # Register with git
    subprocess.run(
        ["git", "config", "--global", f"merge.{name}.name", f"Custom merge driver: {name}"],
        check=True, capture_output=True,
    )
    subprocess.run(
        ["git", "config", "--global", f"merge.{name}.driver", f"{driver_file} %O %A %B"],
        check=True, capture_output=True,
    )
    print(f"Registered as git merge driver 'merge.{name}'")

    if pattern:
        subprocess.run(
            ["git", "config", "--global", f"merge.{name}.driver", f"{driver_file} %O %A %B"],
            check=True, capture_output=True,
        )
        # Write to .gitattributes template
        attrs_file = scripts_dir / "gitattributes"
        content = attrs_file.read_text(encoding="utf-8") if attrs_file.exists() else ""
        # Remove old entry for this pattern
        lines = [l for l in content.splitlines() if not l.startswith(f"{pattern} merge={name}")]
        lines.append(f"{pattern} merge={name}")
        attrs_file.write_text("\n".join(lines) + "\n", encoding="utf-8")
        print(f"Added pattern '{pattern}' -> driver '{name}' to {attrs_file}")


def add_builtin(name: str, pattern: str | None) -> None:
    if name not in BUILTIN_DRIVERS:
        print(f"Error: unknown builtin driver '{name}'", file=sys.stderr)
        print(f"Available: {', '.join(BUILTIN_DRIVERS)}", file=sys.stderr)
        sys.exit(1)

    driver = BUILTIN_DRIVERS[name]
    scripts_dir = ensure_scripts_dir()
    driver_file = scripts_dir / f"{name}.sh"
    driver_file.write_text(f"#!/bin/sh\n{driver['cmd']}\n", encoding="utf-8")
    os.chmod(driver_file, 0o755)

    subprocess.run(
        ["git", "config", "--global", f"merge.{name}.name", driver["desc"]],
        check=True, capture_output=True,
    )
    subprocess.run(
        ["git", "config", "--global", f"merge.{name}.driver", f"{driver_file} %O %A %B"],
        check=True, capture_output=True,
    )
    print(f"Builtin merge driver '{name}' registered: {driver['desc']}")

    if pattern:
        attrs_file = scripts_dir / "gitattributes"
        content = attrs_file.read_text(encoding="utf-8") if attrs_file.exists() else ""
        lines = [l for l in content.splitlines() if not l.startswith(f"{pattern} merge={name}")]
        lines.append(f"{pattern} merge={name}")
        attrs_file.write_text("\n".join(lines) + "\n", encoding="utf-8")
        print(f"Added pattern '{pattern}' -> driver '{name}' to {attrs_file}")


def list_drivers() -> None:
    scripts_dir = ensure_scripts_dir()
    drivers = sorted(scripts_dir.glob("*.sh"))
    if not drivers:
        print("No custom merge drivers. Builtins:", ", ".join(BUILTIN_DRIVERS))
        return
    for d in drivers:
        name = d.stem
        result = subprocess.run(
            ["git", "config", "--global", f"merge.{name}.name"],
            capture_output=True, text=True,
        )
        desc = result.stdout.strip() or "custom"
        print(f"  {name:20s} {desc}")


def install_gitattributes() -> None:
    scripts_dir = ensure_scripts_dir()
    attrs_file = scripts_dir / "gitattributes"
    if not attrs_file.exists():
        print("No gitattributes rules defined yet.", file=sys.stderr)
        sys.exit(1)

    repo_attrs = Path(".gitattributes")
    content = attrs_file.read_text(encoding="utf-8")
    if repo_attrs.exists():
        existing = repo_attrs.read_text(encoding="utf-8")
        # Avoid duplicates
        existing_lines = existing.splitlines()
        new_lines = [l for l in content.splitlines() if l not in existing_lines]
        if new_lines:
            prefix = "\n" if existing and not existing.endswith("\n") else ""
            repo_attrs.write_text(existing + prefix + "\n".join(new_lines) + "\n", encoding="utf-8")
    else:
        repo_attrs.write_text(content, encoding="utf-8")

    print(f"Installed merge driver rules to {repo_attrs}")


def main() -> None:
    parser = argparse.ArgumentParser(prog="git-merge-driver", description="Configure git merge drivers for file types.")
    sub = parser.add_subparsers(dest="command")

    p_add = sub.add_parser("add", help="Add a custom merge driver")
    p_add.add_argument("name", help="Driver name")
    p_add.add_argument("-s", "--script", help="Merge script (shell commands)")
    p_add.add_argument("-p", "--pattern", help="File pattern for .gitattributes")

    p_builtin = sub.add_parser("builtin", help="Register a builtin merge driver")
    p_builtin.add_argument("name", choices=BUILTIN_DRIVERS, help="Builtin driver name")
    p_builtin.add_argument("-p", "--pattern", help="File pattern for .gitattributes")

    sub.add_parser("list", help="List registered merge drivers")

    sub.add_parser("install", help="Install .gitattributes rules into current repo")

    args = parser.parse_args()

    if args.command == "add":
        add_driver(args.name, args.script, args.pattern)
    elif args.command == "builtin":
        add_builtin(args.name, args.pattern)
    elif args.command == "list":
        list_drivers()
    elif args.command == "install":
        install_gitattributes()
    else:
        parser.print_help()


if __name__ == "__main__":
    main()
