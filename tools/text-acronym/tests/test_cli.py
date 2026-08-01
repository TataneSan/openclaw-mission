import io
import sys
import unittest
from contextlib import redirect_stdout, redirect_stderr

from text_acronym import cli


def run(argv, stdin=""):
    out, err = io.StringIO(), io.StringIO()
    old = sys.stdin
    sys.stdin = io.StringIO(stdin)
    try:
        with redirect_stdout(out), redirect_stderr(err):
            code = cli.main(argv)
    finally:
        sys.stdin = old
    return code, out.getvalue(), err.getvalue()


class TestBuild(unittest.TestCase):
    def test_nasa(self):
        r = cli.build_acronym("National Aeronautics and Space Administration")
        self.assertEqual(r["acronym"], "NASA")
        self.assertEqual(r["skipped"], ["and"])

    def test_stop_words(self):
        r = cli.build_acronym("as soon as possible")
        self.assertEqual(r["acronym"], "ASP")

    def test_all_flag(self):
        r = cli.build_acronym("as soon as possible", include_stop_words=True)
        self.assertEqual(r["acronym"], "ASAP")

    def test_accents(self):
        r = cli.build_acronym("Réalité Virtuelle")
        self.assertEqual(r["acronym"], "RV")

    def test_french(self):
        r = cli.build_acronym("Organisation des Nations Unies")
        self.assertEqual(r["acronym"], "ONU")

    def test_hyphenated(self):
        r = cli.build_acronym("Command-Line Interface")
        self.assertEqual(r["acronym"], "CLI")


class TestExpand(unittest.TestCase):
    def test_known(self):
        self.assertEqual(cli.expand("nasa"), "National Aeronautics and Space Administration")

    def test_unknown(self):
        self.assertIsNone(cli.expand("ZZZZ"))


class TestCli(unittest.TestCase):
    def test_single(self):
        code, out, _ = run(["National", "Aeronautics", "and", "Space", "Administration"])
        self.assertEqual(code, 0)
        self.assertEqual(out.strip(), "NASA")

    def test_stdin_batch(self):
        code, out, _ = run([], stdin="Command Line Interface\nSecure Shell\n")
        self.assertEqual(code, 0)
        self.assertIn("CLI", out)
        self.assertIn("SSH", out)

    def test_reverse(self):
        code, out, _ = run(["--reverse", "NASA"])
        self.assertEqual(code, 0)
        self.assertIn("National Aeronautics", out)

    def test_reverse_miss(self):
        code, _, _ = run(["--reverse", "ZZZZ"])
        self.assertEqual(code, 2)

    def test_lower(self):
        code, out, _ = run(["-l", "Secure Shell"])
        self.assertEqual(out.strip(), "ssh")

    def test_json(self):
        code, out, _ = run(["--json", "Secure Shell"])
        self.assertIn('"acronym": "SSH"', out)

    def test_list(self):
        code, out, _ = run(["--list"])
        self.assertIn("ASCII", out)


if __name__ == "__main__":
    unittest.main()
