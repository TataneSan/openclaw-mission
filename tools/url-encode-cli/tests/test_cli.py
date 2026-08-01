import io
import json
import sys
import unittest
from contextlib import redirect_stdout

from url_encode_cli.cli import main, has_errors


class TestUrlEncode(unittest.TestCase):
    def run_cli(self, argv, stdin=None):
        buf = io.StringIO()
        old = sys.stdin
        if stdin is not None:
            sys.stdin = io.StringIO(stdin)
        try:
            with redirect_stdout(buf):
                code = main(argv)
        finally:
            sys.stdin = old
        return code, buf.getvalue()

    def test_encode_basic(self):
        code, out = self.run_cli(["hello world"])
        self.assertEqual(code, 0)
        self.assertEqual(out.strip(), "hello%20world")

    def test_encode_query_chars(self):
        _, out = self.run_cli(["a=b&c=d"])
        self.assertEqual(out.strip(), "a%3Db%26c%3Dd")

    def test_preset_path(self):
        _, out = self.run_cli(["/a b/c?d", "--preset", "path"])
        self.assertEqual(out.strip(), "/a%20b/c%3Fd")

    def test_preset_form(self):
        _, out = self.run_cli(["a b+c", "--preset", "form"])
        self.assertEqual(out.strip(), "a+b%2Bc")

    def test_preset_strict(self):
        _, out = self.run_cli(["~ok", "--preset", "strict"])
        self.assertEqual(out.strip(), "~ok")

    def test_decode(self):
        _, out = self.run_cli(["-d", "hello%20world"])
        self.assertEqual(out.strip(), "hello world")

    def test_decode_form(self):
        _, out = self.run_cli(["-d", "--preset", "form", "a+b"])
        self.assertEqual(out.strip(), "a b")

    def test_decode_utf8(self):
        _, out = self.run_cli(["-d", "%C3%A9"])
        self.assertEqual(out.strip(), "é")

    def test_decode_errors_replace(self):
        _, out = self.run_cli(["-d", "--errors", "replace", "%FF"])
        self.assertIn("�", out)

    def test_stdin_batch(self):
        _, out = self.run_cli([], stdin="a b\nc d\n")
        self.assertEqual(out.splitlines(), ["a%20b", "c%20d"])

    def test_json(self):
        _, out = self.run_cli(["a b", "--json"])
        obj = json.loads(out.strip())
        self.assertEqual(obj, {"input": "a b", "output": "a%20b"})

    def test_check_encode_ok(self):
        code, _ = self.run_cli(["a b", "--check"])
        self.assertEqual(code, 0)

    def test_check_decode_malformed(self):
        code, _ = self.run_cli(["-d", "--check", "%GG"])
        self.assertEqual(code, 2)

    def test_check_decode_clean(self):
        code, _ = self.run_cli(["-d", "--check", "%C3%A9"])
        self.assertEqual(code, 0)

    def test_has_errors(self):
        self.assertTrue(has_errors("100% sure"))
        self.assertTrue(has_errors("%G1"))
        self.assertFalse(has_errors("%20 %c3%a9"))
        self.assertFalse(has_errors("plain"))

    def test_no_input(self):
        from contextlib import redirect_stderr
        buf = io.StringIO()
        with redirect_stderr(buf):
            code = main([])
        self.assertEqual(code, 1)


if __name__ == "__main__":
    unittest.main()
