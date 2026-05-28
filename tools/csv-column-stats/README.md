# csv-column-stats

Statistiques descriptives par colonne pour les fichiers CSV : min, max, moyenne, médiane, écart-type, nombre de nulls.

## Installation

```
go install github.com/TataneSan/csv-column-stats@latest
```

Ou compiler directement :

```
go build -o csv-column-stats .
```

## Usage

```
csv-column-stats <file.csv> [--json]
```

### Arguments

| Argument     | Requis | Description                          |
| -------------| ------ | ------------------------------------ |
| file.csv     | Oui    | Chemin vers le fichier CSV           |
| --json       | Non    | Sortie en format JSON au lieu du tableau |

## Exemples

### Tableau en terminal

```
csv-column-stats data.csv
```

```
COLUMN                 COUNT NULLS        MIN        MAX       MEAN     MEDIAN     STDDEV
--------------------------------------------------------------------------------------------
name                       4     0          -          -          -          -          -
age                        4     1      25.00      35.00      30.00      30.00       5.00
score                      4     1      78.30      92.00      85.27      85.50       6.85
```

### Sortie JSON

```
csv-column-stats data.csv --json
```

## Caractéristiques

- Détecte automatiquement les colonnes numériques
- Calcule min, max, moyenne, médiane, écart-type
- Compte les valeurs nulles/vides
- Ignore les colonnes non-numériques pour les stats numériques
- Sortie tableau ou JSON

## Licence

MIT
