import json
import os
import subprocess
import tempfile
import unittest
from io import StringIO
from contextlib import redirect_stdout

from git_tag_list.cli import main


def git(repo, *args_):
    env = dict(os.environ,
               GIT_AUTHOR_NAME="t", GIT_AUTHOR_EMAIL="t@t",
               GIT_COMMITTER_NAME="t", GIT_COMMITTER_EMAIL="t@t")
    subprocess.run(["git", "-C", repo] + list(args_), check=True,
                   capture_output=True, env=env)


def make_repo(tags):
    d = tempfile.mkdtemp()
    git(d, "init", "-q", "-b", "main")
    with open(os.path.join(d, "f.txt"), "w") as fh:
        fh.write("x")
    git(d, "add", ".")
    git(d, "commit", "-q", "-m", "init")
    for name, annotated in tags:
        if annotated:
            git(d, "tag", "-a", name, "-m", "release")
        else:
            git(d, "tag", name)
    return d


class TagListTest(unittest.TestCase):
    def run_cli(self, argv):
        buf = StringIO()
        with redirect_stdout(buf):
            rc = main(argv)
        return rc, buf.getvalue()

    def test_lists_and_sorts_semver(self):
        repo = make_repo([("v1.10.0", True), ("v1.9.0", False), ("v1.1.0", True)])
        rc, out = self.run_cli([repo])
        self.assertEqual(rc, 0)
        names = [l.split()[0] for l in out.strip().splitlines()]
        self.assertEqual(names, ["v1.1.0", "v1.9.0", "v1.10.0"])

    def test_json(self):
        repo = make_repo([("v1.1.0", True), ("v1.4.0", True), ("v2.0.0-rc.1", True)])
        rc, out = self.run_cli([repo, "--json", "--gaps"])
        data = json.loads(out)
        self.assertEqual(data["count"], 3)
        self.assertTrue(data["gaps"])
        self.assertEqual(data["tags"][2]["prerelease"], "rc.1")

    def test_no_tags(self):
        repo = make_repo([])
        rc, _ = self.run_cli([repo])
        self.assertEqual(rc, 2)

    def test_limit_reverse(self):
        repo = make_repo([("v1.0.0", False), ("v1.1.0", False), ("v1.2.0", False)])
        rc, out = self.run_cli([repo, "--limit", "2"])
        names = [l.split()[0] for l in out.strip().splitlines()]
        self.assertEqual(names, ["v1.1.0", "v1.2.0"])
        rc, out = self.run_cli([repo, "--limit", "1", "--reverse"])
        names = [l.split()[0] for l in out.strip().splitlines()]
        self.assertEqual(names, ["v1.0.0"])


if __name__ == "__main__":
    unittest.main()
