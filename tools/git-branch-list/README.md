# git-branch-list

Liste les branches Git (locales et distantes) avec le dernier commit, l'auteur, la date relative et le statut ahead/behind.

## Installation

```bash
go install github.com/TataneSan/git-branch-list@latest
```

Ou construire depuis la source :

```bash
git clone https://github.com/TataneSan/git-branch-list.git
cd git-branch-list
go build -o git-branch-list .
```

## Usage

```
git-branch-list [OPTIONS]
```

### Options

| Option | Description | Défaut |
|--------|-------------|--------|
| `-f` | Format de sortie : `table` ou `json` | `table` |
| `-repo` | Chemin du répertoire Git | répertoire courant |

### Exemples

Lister les branches du répertoire courant :
```bash
git-branch-list
```

Sortie JSON :
```bash
git-branch-list -f json
```

Spécifier un autre dépôt :
```bash
git-branch-list -repo /chemin/vers/repo
```

## Sortie

### Format table

```
Branches locales
--------------------------------------------------------------------------------
* main                        a1b2c3d4  Jean Dupont  il y a 2 heures
  feature/login               e5f6g7h8  Marie Martin  il y a 3 jours  ahead 5, behind 2
  develop                     i9j0k1l2  Jean Dupont  il y a 1 semaine

Branches distantes
--------------------------------------------------------------------------------
  origin/main                 a1b2c3d4  Jean Dupont  il y a 2 heures
  origin/feature/login        e5f6g7h8  Marie Martin  il y a 3 jours

Total: 3 locale(s), 2 distante(s)
```

- `*` indique la branche courante (en vert)
- Les branches sont triées par date de dernier commit (plus récent en premier)
- Le statut ahead/behind est affiché quand un upstream est configuré

### Format JSON

```json
[
  {"name": "main", "hash": "a1b2c3d4", "subject": "fix: correct login", "author": "Jean Dupont", "date": "2026-05-29T14:00:00Z", "is_current": true, "is_remote": false},
  {"name": "feature/login", "hash": "e5f6g7h8", "subject": "feat: add login", "author": "Marie Martin", "date": "2026-05-26T10:00:00Z", "is_current": false, "is_remote": false, "ahead": 5, "behind": 2}
]
```

## Licence

MIT
