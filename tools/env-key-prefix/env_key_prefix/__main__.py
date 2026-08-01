"""Module entry point for ``python -m env_key_prefix``."""

import sys

from .cli import main

if __name__ == "__main__":
    sys.exit(main())
