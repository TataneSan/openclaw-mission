"""Scaffold helper for new Python CLI tools (stdlib only)."""

import os
import sys

BASE = "/root/openclaw/tools"

LICENSE = open(os.path.join(BASE, "file-oldest-newest", "LICENSE")).read()
GITIGNORE = open(os.path.join(BASE, "file-oldest-newest", ".gitignore")).read()


def pyproject(name, pkg, desc):
    return f'''[build-system]
requires = ["setuptools>=61.0"]
build-backend = "setuptools.build_meta"

[project]
name = "{name}"
version = "1.0.0"
description = "{desc}"
readme = "README.md"
requires-python = ">=3.9"
license = {{text = "MIT"}}
authors = [{{name = "TataneSan"}}]
keywords = ["cli", "tool"]
classifiers = [
    "Environment :: Console",
    "Intended Audience :: Developers",
    "License :: OSI Approved :: MIT License",
    "Programming Language :: Python :: 3",
    "Programming Language :: Python :: 3.9",
    "Programming Language :: Python :: 3.10",
    "Programming Language :: Python :: 3.11",
    "Programming Language :: Python :: 3.12",
]

[project.scripts]
{name} = {pkg}.cli:main

[tool.setuptools.packages.find]
include = ["{pkg}*"]
'''


MAIN = '''import sys

from .cli import main

if __name__ == "__main__":
    sys.exit(main())
'''


def scaffold(name, desc):
    pkg = name.replace("-", "_")
    root = os.path.join(BASE, name)
    os.makedirs(os.path.join(root, pkg), exist_ok=True)
    with open(os.path.join(root, "LICENSE"), "w") as f:
        f.write(LICENSE)
    with open(os.path.join(root, ".gitignore"), "w") as f:
        f.write(GITIGNORE)
    with open(os.path.join(root, "pyproject.toml"), "w") as f:
        f.write(pyproject(name, pkg, desc))
    with open(os.path.join(root, pkg, "__init__.py"), "w") as f:
        f.write(f'"""{desc}."""\n\n__version__ = "1.0.0"\n')
    with open(os.path.join(root, pkg, "__main__.py"), "w") as f:
        f.write(MAIN)
    return root, pkg


if __name__ == "__main__":
    print(scaffold(sys.argv[1], sys.argv[2]))
