import csv
import io
import json
import subprocess
import sys

import pytest

from csv_column_stats.cli import main


DATA = """name,age,score,notes
alice,30,88.5,
bob,25,92.1,x
carol,41,75.0,
dan,,60.5,broken
"""

DATA_NO_HEADER = "1,2\n3,4\n"


def run_cli(argv, exit_ok=True):
    rc = main(argv)
    return rc


def test_all_columns_table(capsys):
    with open("/tmp/stats.csv", "w") as fp:
        fp.write(DATA)
    rc = main(["/tmp/stats.csv"])
    out = capsys.readouterr().out
    assert rc == 0
    # header
    assert "name" in out
    assert "age" in out
    # age row: count 3 (dan has empty age), min 25, max 41
    assert "3" in out
    assert "25" in out
    assert "41" in out


def test_selected_columns(capsys):
    with open("/tmp/stats.csv", "w") as fp:
        fp.write(DATA)
    rc = main(["/tmp/stats.csv", "-c", "age"])
    out = capsys.readouterr().out
    assert rc == 0
    assert "age" in out
    # score should not appear since not selected
    first_line = out.splitlines()[0]
    assert "score" not in first_line


def test_json_output(capsys):
    with open("/tmp/stats.csv", "w") as fp:
        fp.write(DATA)
    rc = main(["/tmp/stats.csv", "--json"])
    out = capsys.readouterr().out
    assert rc == 0
    report = json.loads(out)
    assert report["rows"] == 4
    cols = {c["name"]: c for c in report["columns"]}
    assert cols["age"]["count"] == 3
    assert cols["age"]["min"] == 25.0
    assert cols["age"]["max"] == 41.0
    assert abs(cols["age"]["mean"] - (30 + 25 + 41) / 3) < 1e-9
    assert cols["age"]["empty"] == 1
    assert cols["notes"]["count"] == 0
    assert cols["notes"]["empty"] == 2
    assert cols["notes"]["non_numeric"] == 2


def test_check_pass_and_fail(capsys):
    with open("/tmp/stats.csv", "w") as fp:
        fp.write(DATA)
    # passes
    rc = main(["/tmp/stats.csv", "--check", "age:mean>=30"])
    assert rc == 0
    # fails
    rc = main(["/tmp/stats.csv", "--check", "score:min>=90"])
    assert rc == 2


def test_check_json(capsys):
    with open("/tmp/stats.csv", "w") as fp:
        fp.write(DATA)
    rc = main(["/tmp/stats.csv", "--check", "age:count==3", "--json"])
    out = capsys.readouterr().out
    assert rc == 0
    report = json.loads(out)
    assert report["checks"][0]["ok"] is True


def test_check_by_index(capsys):
    with open("/tmp/stats.csv", "w") as fp:
        fp.write(DATA)
    rc = main(["/tmp/stats.csv", "--check", "2:max==41"])
    assert rc == 0


def test_stdev_flag(capsys):
    with open("/tmp/stats.csv", "w") as fp:
        fp.write(DATA)
    rc = main(["/tmp/stats.csv", "--stdev", "--json"])
    out = capsys.readouterr().out
    assert rc == 0
    report = json.loads(out)
    cols = {c["name"]: c for c in report["columns"]}
    assert "stdev" in cols["age"]


def test_no_header(capsys):
    with open("/tmp/nh.csv", "w") as fp:
        fp.write(DATA_NO_HEADER)
    rc = main(["/tmp/nh.csv", "--no-header", "--json"])
    out = capsys.readouterr().out
    assert rc == 0
    report = json.loads(out)
    cols = {c["name"]: c for c in report["columns"]}
    assert cols["1"]["max"] == 3.0
    assert cols["2"]["max"] == 4.0


def test_unknown_column(capsys):
    with open("/tmp/stats.csv", "w") as fp:
        fp.write(DATA)
    rc = main(["/tmp/stats.csv", "-c", "nope"])
    assert rc == 1


def test_missing_file(capsys):
    rc = main(["/tmp/definitely-missing-file.csv"])
    assert rc == 1


def test_multiple_checks(capsys):
    with open("/tmp/stats.csv", "w") as fp:
        fp.write(DATA)
    rc = main(
        [
            "/tmp/stats.csv",
            "--check",
            "age:count==3,age:min==25",
            "--check",
            "score:max==92.1",
        ]
    )
    assert rc == 0


def test_empty_csv(capsys):
    with open("/tmp/empty.csv", "w") as fp:
        fp.write("")
    rc = main(["/tmp/empty.csv"])
    assert rc == 1
