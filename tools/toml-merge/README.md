# toml-merge

Fusionne plusieurs fichiers TOML en un seul. Les fichiers sont fusionnés de gauche à droite : les valeurs des fichiers suivants remplacent celles des précédents. La fusion est profonde (deep merge) pour les tableaux de hachage imbriqués.

## Installation

```bash
go install github.com/TataneSan/toml-merge@latest
```

Ou construire depuis la source :

```bash
git clone https://github.com/TataneSan/toml-merge.git
cd toml-merge
go build -o toml-merge .
```

## Usage

```bash
toml-merge <sortie.toml> <entree1.toml> [entree2.toml ...]
```

Les fichiers sont fusionnés dans l'ordre. Les valeurs des fichiers de droite prennent le pas sur celles de gauche.

## Exemples

### Fusionner deux fichiers

```bash
toml-merge output.toml base.toml override.toml
```

### Fusionner trois fichiers

```bash
toml-merge config.toml defaults.toml environment.toml local.toml
```

### Exemple concret

`base.toml` :
```toml
[server]
host = "localhost"
port = 8080

[database]
host = "localhost"
port = 5432
```

`override.toml` :
```toml
[server]
port = 9090
ssl = true

[logging]
level = "info"
```

Résultat (`output.toml`) :
```toml
[database]
  host = "localhost"
  port = 5432

[logging]
  level = "info"

[server]
  host = "localhost"
  port = 9090
  ssl = true
```

## Fusion profonde

Les tableaux de hachage sont fusionnés récursivement. Si les deux fichiers contiennent une même section, les clés sont combinées et les valeurs dupliquées sont remplacées par celles du fichier le plus à droite.

## Licence

MIT
