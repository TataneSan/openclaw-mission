import io
import json
import unittest
from contextlib import redirect_stdout
from unittest.mock import patch

from json_path_get.cli import main

DOC = json.dumps({
    "users": [
        {"name": "alice", "tags": ["a", "b"]},
        {"name": "bob", "tags": []},
    ],
    "count": 2,
})


class PathTest(unittest.TestCase):
    def run_cli(self, argv, stdin_text=DOC):
        buf = io.StringIO()
        with patch("sys.stdin", io.StringIO(stdin_text)):
            with redirect_stdout(buf):
                rc = main(argv)
        return rc, buf.getvalue()

    def test_dot(self):
        rc, out = self.run_cli(["-", "users.0.name"])
        self.assertEqual(rc, 0)
        self.assertEqual(out.strip(), '"alice"')

    def test_bracket_raw(self):
        rc, out = self.run_cli(["-", "users[1].name", "--raw"])
        self.assertEqual(out.strip(), "bob")

    def test_int(self):
        rc, out = self.run_cli(["-", "count"])
        self.assertEqual(out.strip(), "2")

    def test_wildcard(self):
        rc, out = self.run_cli(["-", "users[*].name", "--compact", "--raw"])
        self.assertEqual(out.split(), ["alice", "bob"])

    def test_negative_index(self):
        rc, out = self.run_cli(["-", "users[-1].name", "--raw"])
        self.assertEqual(out.strip(), "bob")

    def test_missing(self):
        rc, out = self.run_cli(["-", "users.5.name"])
        self.assertEqual(rc, 2)

    def test_default(self):
        rc, out = self.run_cli(["-", "nope", "--default", "N/A"])
        self.assertEqual(rc, 2)
        self.assertEqual(out.strip(), "N/A")


if __name__ == "__main__":
    unittest.main()
