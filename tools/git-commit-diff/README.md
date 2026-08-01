# git-commit-diff

Affiche le diff d'un commit spécifique de manière formatée. Supporte les mêmes références que git (SHA, branch, tag, HEAD~N).

## Installation

```bash
go install github.com/TataneSan/git-commit-diff@latest
```

Ou télécharger le binaire depuis [Releases](https://github.com/TataneSan/git-commit-diff/releases).

## Usage

```bash
git-commit-diff [OPTIONS] <commit>
```

### Options

| Option | Description |
|---|---|
| `-s, --stat` | Affiche uniquement les statistiques de diff |
| `-n, --name-only` | Affiche uniquement les noms de fichiers modifiés |
| `-N, --name-status` | Affiche les noms et statut (A/M/D) des fichiers |
| `-p, --patch` | Affiche le diff complet (défaut) |
| `-U, --context N` | Nombre de lignes de contexte (défaut: 3) |
| `-c, --color` | Force la coloration du diff |
| `-f, --files` | Affiche la liste des fichiers modifiés |
| `-h, --help` | Affiche l'aide |

## Exemples

```bash
# Diff complet du commit
git-commit-diff abc1234

# Diff du commit actuel
git-commit-diff HEAD

# Diff du parent de HEAD
git-commit-diff HEAD~1

# Statistiques du commit
git-commit-diff abc1234 --stat

# Fichiers modifiés
git-commit-diff abc1234 --name-only

# Fichiers avec statut (A/M/D)
git-commit-diff abc1234 --name-status

# 5 lignes de contexte
git-commit-diff abc1234 --context 5

# Diff coloré
git-commit-diff abc1234 --color
```

## Sortie

L'outil affiche d'abord les métadonnées du commit (SHA, auteur, date, sujet), puis le diff dans le format demandé.

```
$ git-commit-diff abc1234
Commit: abc12340
Auteur: John Doe <john@example.com>
Date:   2024-01-15
Sujet:  feat: add user authentication
------------------------------------------------------------
diff --git a/auth.go b/auth.go
new file mode 100644
index 0000000..a1b2c3d
--- /dev/null
+++ b/auth.go
@@ -0,0 +1,15 @@
+package main
+
+import (
+    "fmt"
+    ...
```

## Build

```bash
go build -o git-commit-diff .
```

## Licence

MIT
