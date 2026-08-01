# OpenClaw — File d'attente des outils

## RÈGLE
Chaque outil = son propre repo Git sur github.com/TataneSan.
JAMAIS mentionner IA/agent dans le code ou les commits.
Push automatique après chaque outil.

## Outil publié — GPU Research
- [x] modal-research-arena : orchestration d'expériences reproductibles avec planification de coût, budgets verrouillés, exécution locale déterministe et backend Modal distant opt-in ✓ 2026-08-01

## Vague 238 — CLI Tools (à définir)
- [x] csv-column-stats : statistiques numériques par colonne CSV (min/max/mean/median, colonnes ciblées, --check CI, JSON) ✓ 2026-08-01
- [x] json-diff-patch : génère un patch JSON (RFC 6902-like simplifié) entre deux documents (--apply, --check CI, JSON) ✓ 2026-08-01
- [x] text-number-lines : numérote les lignes (start/step, largeur fixe, skip blank, --check CI, JSON) ✓ 2026-08-01
- [x] url-set-params : ajoute/remplace des paramètres query dans des URLs en flux (k=v, --replace, --check CI, JSON) ✓ 2026-08-01
- [x] file-hardlink-detect : liste les fichiers partageant le même inode dans une arborescence (--min-links, --check CI, JSON) ✓ 2026-08-01

## Vague 239 — CLI Tools (à définir)
- [x] csv-fill-header : complète les en-têtes manquantes d'un CSV (col_N, détection, --check CI, JSON) ✓ 2026-08-01
- [x] json-merge-objects : fusionne des objets JSON (deep/shallow, arrays append/replace, --check CI, JSON) ✓ 2026-08-01
- [x] text-strip-prefix : retire un préfixe commun ou fixe de chaque ligne (auto-detect, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-params : extrait les paramètres query d'URLs en clés/valeurs tabulées (--unique CI, JSON) ✓ 2026-08-01
- [x] file-size-outliers : liste les fichiers dont la taille dévie de la médiane d'une arborescence (--factor CI, JSON) ✓ 2026-08-01

## Vague 245 — CLI Tools (à définir)
- [x] csv-count-distinct : compte les valeurs distinctes par colonne CSV (top-N, --check CI, JSON) ✓ 2026-08-01
- [x] json-select-fields : garde seulement certains champs d'objets JSONL (dot-paths whitelist, --check CI, JSON) ✓ 2026-08-01
- [x] text-align-columns : aligne des colonnes whitespace en tableau à largeur fixe (left/right, --check CI, JSON) ✓ 2026-08-01
- [x] url-strip-query : retire tout ou partie de la query string des URLs (keep-list, --check CI, JSON) ✓ 2026-08-01
- [x] file-line-count-report : rapport nombre de lignes par fichier d'une arborescence (total, moyenne, --check CI, JSON) ✓ 2026-08-01

## Vague 250 — CLI Tools (échantillon CSV, chemins JSON par regex, fréquence mots, style redirect URL, magic numbers)
- [x] csv-sample-rows : échantillonne N lignes d'un CSV (seed, reservoir, --require CI, JSON) ✓ 2026-08-01
- [x] json-extract-paths-matching : extrait les paires chemin=valeur dont le chemin matche un regex (--invert, --require CI, JSON) ✓ 2026-08-01
- [x] text-frequency-words : fréquence des mots d'un texte (top-N, min-length, stop-words fichier, --check CI, JSON) ✓ 2026-08-01
- [x] url-detect-redirect-style : classe les URLs par style de redirection (?url=, /redirect/, shorteners, --check CI, JSON) ✓ 2026-08-01
- [x] file-extension-mismatch : détecte les fichiers dont le magic number contredit l'extension (png/jpg/pdf/zip/gz, --check CI, JSON) ✓ 2026-08-01

## Vague 251 — CLI Tools (à définir)
- [x] csv-replace-values : remplace des valeurs exactes ou regex dans des colonnes CSV ciblées (mapping fichier, --dry-run CI, JSON) ✓ 2026-08-01
- [x] json-array-flatten : aplatit les tableaux imbriqués d'un document JSON/JSONL à une profondeur donnée (--depth, --check CI, JSON) ✓ 2026-08-01
- [x] text-number-words : convertit les nombres en toutes lettres et inversement (en/fr, plage, --check CI, JSON) ✓ 2026-08-01
- [x] url-compare-parts : diff structurel de deux URLs (scheme/host/port/path/query/fragment, --same-host CI, JSON) ✓ 2026-08-01
- [x] file-permission-report : rapport sur les bits de permission inhabituels d'une arborescence (world-writable, suid, --check CI, JSON) ✓ 2026-08-01

## Vague 252 — CLI Tools (à définir)
- [x] csv-transpose-header-column : permute la première ligne (header) avec la première colonne (--check CI, JSON) ✓ 2026-08-01
- [x] json-merge-arrays-zip : fusionne des tableaux JSONL par index en tuples [a,b] (--strict-len CI, JSON) ✓ 2026-08-01
- [x] text-remove-empty-parens : retire les paires de parenthèses vides et espaces résiduels ((), [], {}, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-file-extension : extrait l'extension de fichier du path d'URLs (stats par extension, --check CI, JSON) ✓ 2026-08-01
- [x] file-naming-convention-lint : vérifie kebab/snake/camel case des noms de fichiers (--style, --fix rename, --check CI, JSON) ✓ 2026-08-01

## Vague 253 — CLI Tools (à définir)
- [x] csv-fill-median : remplace les cellules vides des colonnes numériques par la médiane de la colonne (--require CI, JSON) ✓ 2026-08-01
- [x] json-extract-comments : récupère les champs/notes de type "//" ou "comment_" dans un JSONL (--check CI, JSON) ✓ 2026-08-01
- [x] text-dedupe-lines-global : supprime toutes les lignes dupliquées en gardant la 1re occurrence (whole-file, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-port : extrait le port effectif des URLs (défaut par scheme, --check CI, JSON) ✓ 2026-08-01
- [x] file-symlink-report : inventaire des symlinks d'une arborescence (target exists/missing, relative/absolute, --check CI, JSON) ✓ 2026-08-01

## Vague 254 — CLI Tools (à définir)
- [x] csv-group-by-count : agrège un CSV par une colonne clé et compte les occurrences (top-N, --check CI, JSON) ✓ 2026-08-01
- [x] json-extract-types : rapport des types JSON par chemin dans un JSONL (histogramme MIXED, --check CI, JSON) ✓ 2026-08-01
- [x] text-repeat-pattern : détecte les motifs répétés consécutifs dans chaque ligne (--min-times, --check CI, JSON) ✓ 2026-08-01
- [x] url-detect-language : heuristique de langue d'une URL par TLD/path/query (--check LANGS CI, JSON) ✓ 2026-08-01
- [x] file-hash-chain : chaîne de hash SHA-256 incrémentale par fichier (tip tamper-evident, --check TIP/manifeste CI, JSON) ✓ 2026-08-01

## Vague 255 — CLI Tools (checksums de lignes CSV, nombres JSON, segments quotés, sous-domaines, préfixe timestamp)
- [x] csv-add-checksum-column : ajoute une colonne hash d'intégrité par ligne CSV (md5/sha256, colonnes ciblées, --check manifeste CI, JSON) ✓ 2026-08-01
- [x] json-extract-numbers : extrait tous les nombres d'un JSON/JSONL avec leur chemin (min/max/sum/mean, --check bornes CI, JSON) ✓ 2026-08-01
- [x] text-extract-quoted : extrait les segments entre guillemets de chaque ligne (apostrophe-aware, quotes non refermées, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-subdomain : extrait le sous-domaine de chaque URL (co.uk/com.br/composite suffixes, --check CI, JSON) ✓ 2026-08-01
- [x] file-timestamp-prefix : préfixe les fichiers de leur mtime YYYYMMDD-HHMMSS (idempotent, --dry-run/--apply/--check CI, JSON) ✓ 2026-08-01

## Vague 256 — CLI Tools (CSV split par colonne, census JSON par profondeur, classes caractères, normalisation query vide, rapport ELF)
- [x] csv-split-by-column : éclate un CSV en un fichier par valeur de colonne clé (sortie dossier, --check CI, JSON) ✓ 2026-08-01
- [x] json-count-deep : compte les noeuds d'un JSON par type et profondeur (histogramme, --check CI, JSON) ✓ 2026-08-01
- [x] text-detect-chars : rapport des classes de caractères présentes (unicode categories, non-ASCII, --check ASCII CI, JSON) ✓ 2026-08-01
- [x] url-normalize-query-empty : uniformise les paramètres query vides (a= vs a, --check CI, JSON) ✓ 2026-08-01
- [x] file-elf-report : rapport sur les binaires ELF d'une arborescence (arch, type, lien, --check CI, JSON) ✓ 2026-08-01

## Vague 257 — CLI Tools (corrélation CSV, JSON schema, emails, TLD, permissions cascade)
- [x] csv-column-correlation : corrélation de Pearson entre paires de colonnes numériques CSV (top-N, --check r² CI, JSON) ✓ 2026-08-01
- [x] json-schema-infer : infère un JSON Schema simplifié depuis un JSONL (types, required, enums, --check conformité CI, JSON) ✓ 2026-08-01 (déjà publié Vague 215)
- [x] text-extract-emails : extrait et valide les adresses email d'un texte (dedup, domain-stats, --check CI, JSON) ✓ 2026-08-01 (déjà publié)
- [x] url-extract-tld : extrait le TLD effectif respectant public-suffix-list embarquée (com.co.uk correct, --check CI, JSON) ✓ 2026-08-01 (déjà publié Vague 243)
- [x] file-cascade-permissions : vérifie que chaque parent d'un fichier est bien traversable par son owner (--check CI, JSON) ✓ 2026-08-01

## Vague 258 — CLI Tools (z-scores CSV, JSON DFS walk, wrap quotes-aware, path depth, totals par extension)
- [x] csv-column-z-score : ajoute une colonne z-score par colonne numérique (mean/stdev, --threshold CI outlier, JSON) ✓ 2026-08-01
- [x] json-depth-first-walk : émet chaque noeud d'un JSONL en DFS avec profondeur et chemin (--max-depth CI, JSON) ✓ 2026-08-01
- [x] text-wrap-quotes-aware : wrap de texte qui préserve les guillemets ouvrants/fermants sur les lignes coupées (--check CI, JSON) ✓ 2026-08-01
- [x] url-extract-site-path-depth : profondeur de path par URL (/, /a, /a/b... histogramme, --max-depth CI, JSON) ✓ 2026-08-01
- [x] file-group-by-extension-total : agrège taille et count par extension, avec sous-total par dossier (--top, --threshold CI, JSON) ✓ 2026-08-01

## Vague 259 — CLI Tools (rolling window CSV, chaîne de merge patches RFC 7386, fenêtre la plus modifiée, labels host, dossiers par âge)
- [x] csv-window-functions : fonctions glissantes sur colonnes CSV (rolling mean/sum, window N, --check CI, JSON) ✓ 2026-08-01
- [x] json-merge-patch-multi : applique une série de RFC 7386 merge patches en chaîne (--exit-on-fail CI, JSON) ✓ 2026-08-01
- [x] text-diff-rolling-window : détecte la fenêtre N lignes la plus modifiée entre deux fichiers (--check CI, JSON) ✓ 2026-08-01
- [x] url-parse-host-labels : éclate le host en labels (tld/sld/sub) avec counts statistiques (--min-sub-depth CI, JSON) ✓ 2026-08-01
- [x] file-dirs-by-age-total : regroupe les dossiers par âge moyen des fichiers qu'ils contiennent (--check CI, JSON) ✓ 2026-08-01

## Vague 260 — CLI Tools (delta CSV, replay patches JSON, listes markdown, creds faibles, audit permissions)
- [x] csv-sliding-diff : différence ligne-à-ligne d'une colonne numérique CSV (delta, pct-change, --threshold CI, JSON) ✓ 2026-08-01
- [x] json-time-travel : rejoue une série de patches et extrait l'état à l'étape K (--at CI, JSON) ✓ 2026-08-01
- [x] text-markdown-list-extract : extrait les items de listes markdown (bullets/ordered/task, --checked CI, JSON) ✓ 2026-08-01
- [x] url-extract-default-creds-risk : détecte les URLs utilisant des credentials triviaux (admin:admin, root:root, --check CI, JSON) ✓ 2026-08-01
- [x] file-permission-audit : rapport setuid/setgid/world-writable et sticky bits anormaux (--check CI, JSON) ✓ 2026-08-01

## Vague 262 — CLI Tools (médiane glissante CSV, renommage regex chemins JSON, décalage indentation, sous-paths regex URL, doublons de contenu)
- [x] csv-rolling-median : médiane glissante d'une colonne CSV (fenêtre N, --check CI, JSON) ✓ 2026-08-01
- [x] json-path-rename-regex : renomme les chemins d'un JSON par regex capture groups (dry-run, --check CI, JSON) ✓ 2026-08-01
- [x] text-indent-shift : décale l'indentation de tout un fichier (N espaces ±, tab-aware, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-subpath-regex : extrait les segments de path matchant un regex par position (--check CI, JSON) ✓ 2026-08-01
- [x] file-duplicate-content-groups : groupe les fichiers par contenu identique (hash complet, --min-size CI, JSON) ✓ 2026-08-01

## Vague 263 — CLI Tools (renommage regex colonnes CSV, permutation tableaux JSON, espaces finaux, tracking params URL, fichiers récents)
- [x] csv-column-rename-regex : renomme des colonnes CSV par regex capture groups (--dry-run, --check CI, JSON) ✓ 2026-08-01
- [x] json-array-shuffle-detect : détecte si deux tableaux JSON sont permutations l'un de l'autre (multiset canonique, --check CI, JSON) ✓ 2026-08-01
- [x] text-trailing-space-report : rapport des lignes avec espaces finaux (space/tab/mixed, --fix in-place, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-tracking-params : détecte/strip les paramètres de tracking (utm_*, *_clid, hsa_*, --extra, --check CI, JSON) ✓ 2026-08-01
- [x] file-recent-changes-since : liste les fichiers modifiés depuis timestamp/fenêtre (--since ISO/ts, --within 6h, --check CI, JSON) ✓ 2026-08-01

## Vague 264 — CLI Tools (regex capture CSV, collisions clés JSON, histogramme runs vides, open-redirect URL, diff permissions)
- [x] csv-column-regex-extract : extrait les groupes de capture d'un regex appliqué à une colonne CSV en nouvelles colonnes (named groups, --only-matching, --check CI, JSON) ✓ 2026-08-01
- [x] json-key-collision-report : détecte les clés qui changeraient après normalisation (casefold/underscore/snake/strip-punct, --check CI, JSON) ✓ 2026-08-01
- [x] text-blank-run-histogram : histogramme des longueurs de runs de lignes vides consécutives (--whitespace-blank, --runs, --max-run CI, JSON) ✓ 2026-08-01
- [x] url-detect-open-redirect-param : repère les paramètres contenant des URLs de redirection (url=, next=, redirect=, fragment SPA, multi-decode, --check CI, JSON) ✓ 2026-08-01
- [x] file-permission-diff-manifest : compare les permissions actuelles d'une arborescence à un manifeste octal (--dump, --apply, relative/absolute, --check CI, JSON) ✓ 2026-08-01

## Vague 265 — CLI Tools (tri numérique CSV, collecte chemin JSON, subst regex groupes, mapping host URL, restauration permissions)
- [x] csv-row-sort-by-numeric : trie un CSV par une colonne numérique (desc, stable, --check CI, JSON) ✓ 2026-08-01
- [x] json-collect-by-path : agrège les valeurs d'un chemin donné à travers JSONL (stats, uniques, --check CI, JSON) ✓ 2026-08-01
- [x] text-line-regex-replace-groups : substitution regex par groupes de capture sur chaque ligne (--count, --check CI, JSON) ✓ 2026-08-01
- [x] url-replace-host-mapping : remplace hosts via fichier mapping old=new (multi-URLs, port préservé, --check CI, JSON) ✓ 2026-08-01
- [x] file-hard-permission-restore : restaure les modes 0644/0755 standards récursivement (--exclude, --dry-run, --check CI, JSON) ✓ 2026-08-01

## Vague 266 — CLI Tools (numérotation CSV, strip nulls JSON, slice colonnes, slash final URL, dossiers vides)
- [x] csv-add-row-number : ajoute une colonne numéro de ligne à un CSV (start/step, header custom, --check séquence CI, JSON) ✓ 2026-08-01
- [x] json-strip-null-fields : retire les clés à valeur null des objets JSONL (récursif, empty-dict prune, --check CI, JSON) ✓ 2026-08-01
- [x] text-column-slice : extrait une tranche de colonnes caractères N:M par ligne (1-based, négatifs, --check width CI, JSON) ✓ 2026-08-01
- [x] url-normalize-trailing-slash : uniformise le slash final des paths URL (add/remove, racine conservée, --check CI, JSON) ✓ 2026-08-01
- [x] file-empty-dir-detect : liste les dossiers vides (ou récursivement vides) d'une arborescence (--prune option, --check CI, JSON) ✓ 2026-08-01

## Vague 267 — CLI Tools (milieu CSV, casse clés JSON, préfixe commun, dots URL, anciens/récents fichiers)
- [x] csv-middle-rows : extrait les N lignes du milieu d'un CSV (ou tranche centrée %, --check CI, JSON) ✓ 2026-08-01
- [x] json-key-case-convert : convertit les clés d'objets JSONL en snake_case/camelCase/kebab (récursif, collisions report, --check CI, JSON) ✓ 2026-08-01
- [x] text-longest-common-prefix-lines : calcule le préfixe commun le plus long de toutes les lignes (--ignore-case, --check CI, JSON) ✓ 2026-08-01
- [x] url-path-normalize-dots : résout les segments . et .. dans les paths d'URLs (RFC 3986 remove_dot_segments, --check CI, JSON) ✓ 2026-08-01
- [x] file-oldest-newest-report : rapport des fichiers les plus anciens/récents par arborescence (top-N mtime, --check CI, JSON) ✓ 2026-08-01

## Vague 268 — CLI Tools (percentiles CSV, compte clés JSON, strip ANSI, schemes URL, profondeur dossiers)
- [x] csv-percentile-column : calcule les percentiles (P50/P90/P99) d'une colonne numérique CSV (interpolation linéaire, --check CI, JSON) ✓ 2026-08-01
- [x] json-key-count-report : rapport du nombre de clés par objet d'un JSONL (min/max/mean/histogramme, --check CI, JSON) ✓ 2026-08-01
- [x] text-strip-ansi-codes : retire les séquences d'échappement ANSI (couleurs, CSI) de chaque ligne (--check CI, JSON) ✓ 2026-08-01
- [x] url-extract-scheme-report : rapport des schemes d'URLs (comptage, scheme manquant détecté, --check CI, JSON) ✓ 2026-08-01
- [x] file-dir-depth-histogram : histogramme des profondeurs de répertoires d'une arborescence (--max-depth CI, JSON) ✓ 2026-08-01

## Vague 269 — CLI Tools (casse headers CSV, index tableaux JSON, wrap suffixe, host:port, fichiers cachés)
- [x] csv-header-case-convert : convertit les en-têtes d'un CSV en snake/camel/kebab (collision report, --check CI, JSON) ✓ 2026-08-01
- [x] json-array-index-extract : extrait l'élément à l'index N de chaque tableau d'un JSONL (négatifs, --check CI, JSON) ✓ 2026-08-01
- [x] text-word-wrap-suffix : wrap de texte avec suffixe de continuation custom sur chaque ligne coupée (--check CI, JSON) ✓ 2026-08-01
- [x] url-host-port-mapping : rapport des combinaisons host:port uniques d'un flux d'URLs (défauts par scheme, --check CI, JSON) ✓ 2026-08-01
- [x] file-hidden-detect : liste les fichiers et dossiers cachés (dotfiles, attributs) d'une arborescence (--check CI, JSON) ✓ 2026-08-01

## Vague 270 — CLI Tools (à définir)
- [x] csv-collapsed-dupe-rows : fusionne les lignes CSV dupliquées avec compteur (key cols, --check CI, JSON) ✓ 2026-08-01
- [x] json-array-element-counts : rapport du nombre d'éléments par tableau d'un JSONL (histogramme, --check CI, JSON) ✓ 2026-08-01
- [x] text-column-justify-right : aligne à droite les colonnes d'un tableau whitespace (width auto, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-userinfo : extrait/rapporte la présence de userinfo dans des URLs (masquage, --check CI, JSON) ✓ 2026-08-01
- [x] file-extension-counts-by-dir : compte les extensions par dossier immédiat d'une arborescence (--check CI, JSON) ✓ 2026-08-01

## Vague 271 — CLI Tools (paires CSV, index tableaux JSON, swap colonnes, plages ports, parité mtime)
- [x] csv-row-pair-diff : diff ligne par ligne de deux CSV alignés (clé optionnelle, side-by-side, --check CI, JSON) ✓ 2026-08-01
- [x] json-array-index-of : trouve l'index des éléments matchant un prédicat dans les tableaux d'un JSONL (equals/regex/number-range, --check CI, JSON) ✓ 2026-08-01
- [x] text-column-swap-first-last : échange la première et la dernière colonne whitespace de chaque ligne (--check CI, JSON) ✓ 2026-08-01
- [x] url-extract-port-range : classe les ports d'URLs par plage (well-known/registered/dynamic, --check CI, JSON) ✓ 2026-08-01
- [x] file-mtime-parity-report : rapport des fichiers dont le mtime est pair/impair par dossier (histogramme, --check CI, JSON) ✓ 2026-08-01

## Vague 272 — CLI Tools (à définir)
- [x] csv-nth-occurrence : garde la N-ième occurrence de chaque clé dans un CSV (keep-nth, drop-nth, --check CI, JSON) ✓ 2026-08-01
- [x] json-set-path : écrit une valeur à un dot-path dans des docs JSONL (types auto, create-missing, --check CI, JSON) ✓ 2026-08-01
- [x] text-reverse-fields : inverse l'ordre des champs whitespace par ligne (sep préservé global, --check CI, JSON) ✓ 2026-08-01
- [x] url-fragment-to-query : déplace les paramètres k=v du fragment vers la query string (--check CI, JSON) ✓ 2026-08-01
- [x] file-size-histogram-log : histogramme log2 des tailles de fichiers d'une arborescence (buckets puissances de 2, --check CI, JSON) ✓ 2026-08-01

## Vague 273 — CLI Tools (à définir)
- [x] csv-row-bucket : assigne chaque ligne CSV à un bucket par valeur de colonne (labels, --check CI, JSON) ✓ 2026-08-01
- [x] json-path-exists : teste si un dot-path existe dans des docs JSONL (--fail-if-missing CI, JSON) ✓ 2026-08-01
- [x] text-dedupe-words-line : retire les mots dupliqués à l'intérieur de chaque ligne (ordre préservé, --check CI, JSON) ✓ 2026-08-01
- [x] url-resolve-relative : résout les URLs relatives par rapport à une base (--base, --check CI, JSON) ✓ 2026-08-01
- [x] file-mtime-shift : applique un décalage horaire aux mtimes d'une arborescence (±Nh, --dry-run, --check CI, JSON) ✓ 2026-08-01

## Vague 274 — CLI Tools (à définir)
- [x] csv-strip-bom-rows : retire les U+FEFF BOM des cellules en début de ligne CSV (--check CI, JSON) ✓ 2026-08-01
- [x] json-path-stats : statistiques par dot-path (type distribution, count, --check CI, JSON) ✓ 2026-08-01
- [x] text-first-letters : extrait le premier caractère de chaque mot (acrostiche, --upper CI, JSON) ✓ 2026-08-01
- [x] url-scheme-canonical : normalise scheme+host et retire le default port d'URLs (--check CI, JSON) ✓ 2026-08-01
- [x] file-dupe-basename-groups : groupe les fichiers par basename partagé entre dossiers (--check CI, JSON) ✓ 2026-08-01

## Vague 275 — CLI Tools (à définir)
- [x] csv-normalize-delimiter : réécrit un CSV avec un délimiteur cible uniforme (quoting minimal, --check CI, JSON) ✓ 2026-08-01
- [x] json-flatten-max-depth : aplatit un JSON en limitant la profondeur d'expansion (truncate marker, --check CI, JSON) ✓ 2026-08-01
- [x] text-second-letters : extrait le deuxième caractère de chaque mot (--upper CI, JSON) ✓ 2026-08-01
- [x] url-https-upgrade : récrit les URLs http:// en https:// sauf exceptions fichier (--check CI, JSON) ✓ 2026-08-01
- [x] file-largest-per-ext : plus gros fichier par extension d'une arborescence (top-N, --check CI, JSON) ✓ 2026-08-01

## Vague 276 — CLI Tools (à définir)
- [x] csv-column-shift : décale les valeurs d'une colonne CSV vers le haut/bas (fill vide, --check CI, JSON) ✓ 2026-08-01
- [x] json-pluck-unique : extrait les valeurs uniques d'un chemin dot-path à travers JSONL (comptage, --check CI, JSON) ✓ 2026-08-01
- [x] text-strip-markdown-format : retire le formatage markdown (**, *, _, ``, ~~) en gardant le texte (--check CI, JSON) ✓ 2026-08-01
- [x] url-parse-userinfo : extrait user/pass des URLs et les masque (show/partial, --check CI, JSON) ✓ 2026-08-01
- [x] file-permission-classes : rapport des classes de permissions ugo par dossier (rwx histogramme, --check CI, JSON) ✓ 2026-08-01

## Vague 277 — CLI Tools (à définir)
- [x] csv-row-digest : calcule un hash par ligne CSV (colonnes choisies, algo md5/sha1/sha256, --check manifeste CI, JSON) ✓ 2026-08-01
- [x] json-coalesce : remplace les valeurs null/vides à un dot-path par la première valeur par défaut fournie (--values-list, --check CI, JSON) ✓ 2026-08-01
- [x] text-strip-html-tags : retire les balises HTML en gardant le texte (entités décodées, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-id-pattern : extrait des IDs numériques/uuid des segments path d'URLs (regex embarqué, --check CI, JSON) ✓ 2026-08-01
- [x] file-magic-header-dump : dump hex des N premiers octets par fichier d'une arborescence (--bytes N, --check CI, JSON) ✓ 2026-08-01

## Vague 278 — CLI Tools (pivot CSV, conflits types JSONL, accents, diff URLs en masse, permissions par défaut)
- [x] csv-pivot-count : pivot simple CSV rows->cols avec compteurs (rows=colA, cols=colB, --check CI, JSON) ✓ 2026-08-01
- [x] json-detect-conflicts : trouve les clés avec types différents à travers JSONL (--strict CI, JSON) ✓ 2026-08-01
- [x] text-remove-diacritics : retire les accents/diacritiques unicode (NFKD strip, --check CI, JSON) ✓ 2026-08-01 (déjà publié Vague 186 — upgrade --check/ligatures/JSON)
- [x] url-diff-batch : diff structurel en masse de paires d'URLs (ligne paire gauche/droite, --check CI, JSON) ✓ 2026-08-01
- [x] file-permission-set-default : applique des umask par default (0644 files / 0755 dirs) sur une arborescence (--dry-run, --check CI, JSON) ✓ 2026-08-01

## Vague 279 — CLI Tools (pivot somme CSV, ordre clés JSON, translittération, normalisation path URL, lignes partagées)
- [x] csv-pivot-sum : pivot CSV avec agrégation somme d'une colonne numérique (--sum-col, --check CI, JSON) ✓ 2026-08-01
- [x] json-normalize-key-order : trie récursivement les clés d'objets JSON/JSONL (canonique, --check CI, JSON) ✓ 2026-08-01
- [x] text-transliterate-ascii : translittération unicode->ASCII agressive (unidecode-like maison, --check CI, JSON) ✓ 2026-08-01
- [x] url-path-join-normalize : joint et normalise des segments path d'URL (.., ., //, --check CI, JSON) ✓ 2026-08-01
- [x] file-dup-lines-across : détecte les lignes apparaissant dans plusieurs fichiers d'un dossier (--min-files CI, JSON) ✓ 2026-08-01

## Vague 280 — CLI Tools (tableaux JSON stats, suffixe commun, âge domaine URL, manifeste arborescence)
- [x] json-array-stats : statistiques sur les valeurs d'un tableau à un dot-path (min/max/sum/mean, --check CI, JSON) ✓ 2026-08-01
- [x] text-common-suffix-lines : calcule le suffixe commun le plus long de toutes les lignes (--ignore-case, --check CI, JSON) ✓ 2026-08-01
- [x] url-domain-age-hint : score heuristique d'âge d'un domaine dans une URL (motifs, tld, --check CI, JSON) ✓ 2026-08-01
- [x] file-tree-manifest : génère un manifeste path+size+mtime+sha256 d'une arborescence (--verify CI, JSON) ✓ 2026-08-01

## Vague 281 — CLI Tools (buckets CSV, priorité clés JSON, répétitions, swap segments URL, types MIME)
- [x] csv-add-bucket-label : ajoute une colonne de label de bucket par seuils numériques (--thresholds, --labels, --check CI, JSON) ✓ 2026-08-01
- [x] json-merge-key-precedence : fusionne deux objets JSON avec ordre de priorité des clés documenté (left/right, --check CI, JSON) ✓ 2026-08-01
- [x] text-collapse-repeats : réduit les répétitions consécutives d'un caractère à N occurrences (--char, --max, --check CI, JSON) ✓ 2026-08-01
- [x] url-path-swap-segments : échange deux segments positionnels du path d'URLs (--i, --j, --check CI, JSON) ✓ 2026-08-01
- [x] file-content-type-report : rapport des types MIME détectés par magic numbers dans une arborescence (histogramme, --check CI, JSON) ✓ 2026-08-01

## Vague 282 — CLI Tools (histogramme CSV, prune containers JSON, squeeze whitespace, params vers headers, MIME vs extension)
- [x] csv-bucket-counts : compte les lignes par bucket de valeurs numériques (histogramme fixe, --check CI, JSON) ✓ 2026-08-01
- [x] json-remove-empty-containers : retire les objets/tableaux vides récursivement d'un JSON (keep-null, --check CI, JSON) ✓ 2026-08-01
- [x] text-collapse-whitespace-runs : réduit les runs d'espaces/tabs à un seul séparateur (style tr -s, --check CI, JSON) ✓ 2026-08-01
- [x] url-params-to-headers : déplace des paramètres query vers des pseudo-headers X-Param (mode debug, --check CI, JSON) ✓ 2026-08-01
- [x] file-mime-extension-mismatch : signale les fichiers dont le magic number contredit l'extension déclarée (rapport, --check CI, JSON) ✓ 2026-08-01

## Vague 283 — CLI Tools (fill par mapping CSV, tri tableaux JSONL, EOL mélangées, strip index URL, doublons taille+ext)
- [x] csv-column-fill-from-map : remplit une colonne vide depuis un mapping clé=valeur par clé CSV (--on missing, --check CI, JSON) ✓ 2026-08-01
- [x] json-sort-array-of-objects : trie les tableaux d'objets d'un JSONL par une clé numérique (desc, stable, --check CI, JSON) ✓ 2026-08-01
- [x] text-detect-mixed-eol : détecte les fins de ligne mélangées CRLF/LF par fichier (rapport, --check CI, JSON) ✓ 2026-08-01
- [x] url-strip-index-files : retire les segments /index.html /default.aspx finaux des URLs (--check CI, JSON) ✓ 2026-08-01
- [x] file-size-dupes-by-ext : liste les groupes de fichiers de même taille et même extension (candidats doublons, --check CI, JSON) ✓ 2026-08-01

## Vague 284 — CLI Tools (à faire)
- [ ] csv-median-column : calcule la médiane d'une colonne numérique CSV (interpolation paire, --check bornes CI, JSON)
- [ ] json-paths-as-table : rend tous les chemins feuilles d'un JSON en tableau colonnes séparées (orienté diff, --check CI, JSON)
- [ ] text-strip-word-wrap : défait un word-wrap et rejoint les lignes en paragraphes (heuristique, --check CI, JSON)
- [ ] url-extract-base-domain : extrait le domaine enregistré approximatif (eTLD+1 simple) de chaque URL (--check CI, JSON)
- [ ] file-count-lines-per-ext : compte les lignes des fichiers texte par extension dans une arborescence (Exclure binaires, --check CI, JSON)


## Vague 261 — CLI Tools (filtrage expr CSV, arbre Merkle JSON, dédup floue, breadcrumb URL, encodage fichiers)
- [x] csv-row-filter-expr : filtre les lignes CSV par expression simple (=, !=, <, >, contains, --check-min CI, JSON) ✓ 2026-08-01
- [x] json-hash-tree : construit un arbre de Merkle SHA-256 sur les chemins/valeurs d'un JSON (--verify CI, JSON) ✓ 2026-08-01
- [x] text-fuzzy-dedupe : déduplique les lignes par similarité (threshold Jaccard/normalisation, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-breadcrumb : décompose le path en breadcrumb key=value (segments positionnels, --check CI, JSON) ✓ 2026-08-01
- [x] file-content-encoding-detect : rapport d'encodage (ascii/utf8/utf16/latin1 heuristic) par fichier (--check CI, JSON) ✓ 2026-08-01

## Vague 249 — CLI Tools (à définir)
- [x] csv-rename-columns-positional : renomme les colonnes CSV par position (1=newname, partial, --check CI, JSON) ✓ 2026-08-01
- [x] json-depth-report : rapport de profondeur max/moyenne des documents JSONL (histogramme, --max-depth CI, JSON) ✓ 2026-08-01
- [x] text-wrap-after-delimiter : coupe les lignes après chaque occurrence d'un délimiteur (;, |, --keep-delim CI, JSON) ✓ 2026-08-01
- [x] url-extract-credentials : détecte les credentials user:pass@ dans des URLs (masquage, --redact CI, JSON) ✓ 2026-08-01
- [x] file-owner-report : rapport par propriétaire (uid/user) des fichiers d'une arborescence (counts, tailles, --check CI, JSON) ✓ 2026-08-01

## Vague 248 — CLI Tools (à définir)
- [x] csv-strict-width : vérifie que chaque ligne CSV a le même nombre de champs que le header (rapport écarts, --check CI, JSON) ✓ 2026-08-01
- [x] json-merge-arrays-concat : fusionne des tableaux JSONL de mêmes index en concat (mode pairwise/first, --check CI, JSON) ✓ 2026-08-01
- [x] text-blank-line-stats : statistiques sur les lignes vides (total, runs, run le plus long, positions, --check CI, JSON) ✓ 2026-08-01
- [x] url-append-path : ajoute un segment de path à des URLs (trailing slash gérée, --index insert, --check CI, JSON) ✓ 2026-08-01
- [x] file-ctime-report : rapport d'âge par ctime (metadata change) des fichiers (buckets, --stale-days CI, JSON) ✓ 2026-08-01

## Vague 247 — CLI Tools (colonnes vides CSV, index JSON par clé, trim leading, segments path URL, rapport atime)
- [x] csv-drop-empty-columns : supprime les colonnes entièrement vides d'un CSV (header conservé, --check CI, JSON) ✓ 2026-08-01
- [x] json-netaddr-index : indexe des objets JSONL par un champ clé vers un tableau (clé dupliquée agrégée, --check CI, JSON) ✓ 2026-08-01
- [x] text-trim-leading : retire les espaces/tabulations en début de ligne (--in-place, --min-indent CI, JSON) ✓ 2026-08-01
- [x] url-extract-path-segments : extrait les segments d'URL path en colonnes (index, join char, --check CI, JSON) ✓ 2026-08-01
- [x] file-atime-report : rapport d'âge par atime (last access) des fichiers d'une arborescence (buckets, --check CI, JSON) ✓ 2026-08-01

## Vague 246 — CLI Tools (normalisation nuls CSV, wrap/unwrap clé JSON, squeeze lignes vides, fragments URL, snapshots mtime)
- [x] csv-normalize-nulls : uniformise les valeurs nulles d'un CSV (NA/N/A/null/"" -> "", --check CI, JSON) ✓ 2026-08-01
- [x] json-wrap-objects-in : enveloppe des objets JSONL sous une clé donnée (--unwrap inverse, --check CI, JSON) ✓ 2026-08-01
- [x] text-squeeze-blank-runs : réduit les runs de lignes vides à N max (style cat -s, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-fragments : extrait les fragments #... des URLs (clé=valeur pairs, --unique CI, JSON) ✓ 2026-08-01
- [x] file-mtime-snapshots : snapshot des mtimes d'un dossier + diff added/changed/deleted (--check CI, JSON) ✓ 2026-08-01

## Vague 244 — CLI Tools (dédup lignes CSV, aplatissement JSON, wrap paragraphes, normalisation casse URL, lignes partagées entre fichiers)
- [x] csv-dedupe-rows : supprime les lignes dupliquées d'un CSV (clés ciblées, keep first/last, --check CI, JSON) ✓ 2026-08-01
- [x] json-flatten-paths : aplatit un JSON en paires chemin=valeur (sep custom, arrays indexées, --check CI, JSON) ✓ 2026-08-01
- [x] text-para-wrap : reformate des paragraphes à largeur fixe (détection blocs vides, indent préservé, --check CI, JSON) ✓ 2026-08-01
- [x] url-normalize-case : met en minuscules scheme+host des URLs en flux (path/query préservés, --check CI, JSON) ✓ 2026-08-01
- [x] file-dupe-lines-across : détecte les lignes identiques partagées entre fichiers d'un dossier (--min-files CI, JSON) ✓ 2026-08-01

## Vague 243 — CLI Tools (swap colonnes CSV, feuilles JSON, trim pré/suffixe, TLD URLs, rapport BOM)
- [x] csv-swap-columns : échange deux colonnes d'un CSV (noms/indices, --check CI, JSON) ✓ 2026-08-01
- [x] json-extract-leaves : liste seulement les valeurs feuilles d'un JSON (skips clés intermédiaires, --check CI, JSON) ✓ 2026-08-01
- [x] text-trim-lines : retire préfixe AND suffixe fixes des lignes (longueur ou chaîne, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-tld : extrait TLD/domaine enregistré minimal (heuristique suffixe publique, --check CI, JSON) ✓ 2026-08-01
- [x] file-bom-report : détecte les BOM dans une arborescence de fichiers texte (--types, --check CI, JSON) ✓ 2026-08-01

## Vague 242 — CLI Tools (colonne CSV, compte-chemins JSON, split lignes longues, drop scheme URL, etags MD5 fichiers)
- [x] csv-extract-column : extrait une seule colonne CSV en flux brut (index/nom, --check CI, JSON) ✓ 2026-08-01
- [x] json-path-count : compte les occurrences de chaque chemin dot/crochet dans des objets JSONL (--check CI, JSON) ✓ 2026-08-01
- [x] text-split-long-lines : coupe les lignes plus longues que N en segments (continuation \ ou indent, --check CI, JSON) ✓ 2026-08-01
- [x] url-drop-scheme : passe des URLs https://vers scheme-relative ou path-only (--check CI, JSON) ✓ 2026-08-01
- [x] file-content-etag : calcule un etag md5 par fichier et le compare à un manifeste (--check CI, JSON) ✓ 2026-08-01

## Vague 241 — CLI Tools (BOM CSV, wrap JSONL, histo longueurs, ports HTTP par défaut, âges de fichiers)
- [x] csv-strip-bom : retire/pointe les BOM UTF-8/UTF-16 des fichiers CSV (--check CI, JSON) ✓ 2026-08-01
- [x] json-wrap-array : enveloppe des objets JSONL dans un tableau JSON (--unwrap, --check CI, JSON) ✓ 2026-08-01
- [x] text-line-histogram : histogramme de longueurs de lignes (buckets, --check CI, JSON) ✓ 2026-08-01
- [x] url-strip-default-port : retire les ports par défaut des URLs (80/443/21..., --check CI, JSON) ✓ 2026-08-01
- [x] file-age-buckets : répartit les fichiers par tranches d'âge (1h/1d/1w/1m, --check CI, JSON) ✓ 2026-08-01

## Vague 240 — CLI Tools (totaux CSV, renommage clés JSON, suffixes de lignes, tri paramètres URL, plus récents par dossier)
- [x] csv-sum-columns : somme/moyenne par colonne numérique CSV (totaux ligne/footer, --check CI, JSON) ✓ 2026-08-01
- [x] json-rename-keys : renomme des clés JSON par mapping fichier (dot-paths, récursif, --check CI, JSON) ✓ 2026-08-01
- [x] text-suffix-lines : ajoute un suffixe fixe ou dynamique à chaque ligne (pattern, --check CI, JSON) ✓ 2026-08-01
- [x] url-sort-params : trie les paramètres query d'URLs alphabétiquement (stable dupes, --check CI, JSON) ✓ 2026-08-01
- [x] file-newest-per-dir : liste le fichier le plus récent de chaque sous-dossier (mtime, --check CI, JSON) ✓ 2026-08-01

## Vague 237 — CLI Tools (réordonnage colonnes CSV, tri clés JSON récursif, collapse répétitions, strip fragments URL, symlinks cassés)
- [x] csv-reorder-columns : réordonne les colonnes par liste de noms/indices (reste append/prepend, --check CI, JSON) ✓ 2026-08-01
- [x] json-sort-keys-recursive : trie toutes les clés d'un JSON récursivement (fold case, --check CI, JSON) ✓ 2026-08-01
- [x] text-collapse-runs : réduit les répétitions de caractères consécutifs (classes ciblées, char unique, --max, --check CI, JSON) ✓ 2026-08-01
- [x] url-strip-fragment : retire les fragments #... des URLs en flux (--keep-list, --check CI, JSON) ✓ 2026-08-01
- [x] file-symlink-broken : liste les liens symboliques cassés d'une arborescence (excludes, --delete, --check CI, JSON) ✓ 2026-08-01

## Vague 236 — CLI Tools (quotage CSV, strip commentaires JSONC, swap de mots, dedupe params URL, fichiers vides)
- [x] csv-quote-fields : force/retire les guillemets autour des champs CSV (minimal/all/non-numeric/none, --check CI, JSON) ✓ 2026-08-01
- [x] json-strip-comments : retire les commentaires //ligne et /*bloc*/ d'un JSONC vers JSON strict (strings préservées, --validate, --check CI, JSON) ✓ 2026-08-01
- [x] text-swap-words : échange deux mots/tokens dans chaque ligne (simultané, regex, occurrence N, --check CI, JSON) ✓ 2026-08-01
- [x] url-dedupe-params : déduplique les paramètres répétés dans des URLs (keep first/last, --check CI, JSON) ✓ 2026-08-01
- [x] file-empty-detect : liste les fichiers vides (0 octet) d'une arborescence (excludes, --delete, --check CI, JSON) ✓ 2026-08-01

## Vague 235 — CLI Tools (premières phrases, délimiteur CSV, clés JSON par défaut, conversion de casse, fichiers récents)
- [x] text-first-sentences : garde les N premières phrases de chaque paragraphe (N global, --min-words CI, JSON) ✓ 2026-08-01
- [x] csv-detect-delimiter : détecte le délimiteur d'un CSV (virgule/point-virgule/tab/pipe, échantillon, --require CI, JSON) ✓ 2026-08-01
- [x] json-ensure-keys : ajoute les clés manquantes avec valeur par défaut à des objets JSON/JSONL (dot-paths, --check CI, JSON) ✓ 2026-08-01
- [x] text-to-camel-case : convertit chaque ligne en camelCase/PascalCase/snake_case/kebab-case (--style, --no-split-digits, --check CI, JSON) ✓ 2026-08-01
- [x] file-recent-modified : liste les fichiers modifiés depuis N minutes/heures (mtime, excludes glob, --count CI, JSON) ✓ 2026-08-01

## Vague 234 — CLI Tools (trim fin de ligne, pivot CSV, mini-schéma JSON, base64 ligne par ligne, expansion CIDR)
- [x] text-trim-trailing : retire les espaces/tabulations en fin de ligne (fichier ou stdin, --in-place, --check CI, JSON) ✓ 2026-08-01
- [x] csv-pivot-simple : pivot léger CSV (lignes->colonnes sur clé/valeur, --check CI, JSON) ✓ 2026-08-01
- [x] json-validate-schema-mini : valide un JSON contre un mini-schéma (types, required, enum, --check CI, JSON) ✓ 2026-08-01
- [x] text-base64-lines : encode/décode chaque ligne en base64 (--decode, URL-safe, --check CI, JSON) ✓ 2026-08-01
- [x] ip-cidr-expand : étend un CIDR en liste d'IPs ou le résume (limites, --count CI, JSON) ✓ 2026-08-01

## Vague 233 — CLI Tools (phrases, renommage colonnes CSV, JSON unflatten, strip ANSI, doublons hash)
- [x] text-sentence-split : découpe un texte en phrases (une par ligne, abréviations ignorées, --min-words CI, JSON) ✓ 2026-08-01
- [x] csv-rename-columns : rename les colonnes CSV via mapping old=new (regex possible, --check CI, JSON) ✓ 2026-08-01
- [x] json-unflatten : reconstruit un JSON imbriqué depuis des paires chemin=valeur (sep custom, types auto, --check CI, JSON) ✓ 2026-08-01
- [x] text-strip-ansi-blocks : supprime les blocs délimités par des séquences ANSI (couleurs, codes curseur, --keep-colors, --check CI, JSON) ✓ 2026-08-01
- [x] file-dup-finder-hash : trouve les fichiers en double par hash (taille min, recursive, --delete interactive, --check CI, JSON) ✓ 2026-08-01

## Vague 232 — CLI Tools (rewrap indenté, fill-down CSV, échantillon JSON, swap host URL, renommage regex)
- [x] text-wrap-smart-indent : rewrap des paragraphes en préservant l'indentation de bloc (listes, quotes >, --check CI, JSON) ✓ 2026-08-01
- [x] csv-fill-down : propage les valeurs vides depuis la ligne précédente (colonnes ciblées, --check CI, JSON) ✓ 2026-08-01
- [x] json-pick-random : échantillonne N éléments d'un tableau JSON (seed, --check CI, JSON) ✓ 2026-08-01
- [x] url-swap-host : remplace le host des URLs en flux (port conservé, mapping custom, --check CI, JSON) ✓ 2026-08-01
- [x] text-regex-rename-files : renomme des fichiers via regex capture groups (dry-run, --check CI, JSON) ✓ 2026-08-01

## Vague 231 — CLI Tools (cipher sélectif, casse en-têtes CSV, tri tableau JSON, rotation colonnes, rapport extensions)
- [x] text-rot13-select : applique ROT13/Caesar seulement aux segments entre délimiteurs (--between quotes/brackets, --check CI, JSON) ✓ 2026-08-01
- [x] csv-header-case : uniformise la casse des en-têtes CSV (snake/camel/kebab/upper, --dry-run, --check CI, JSON) ✓ 2026-08-01
- [x] json-sort-array-by : trie un tableau JSON d'objets par champ (multi-clés, numérique, --check CI, JSON) ✓ 2026-08-01
- [x] text-column-rotate : décale les colonnes whitespace vers la gauche/droite (--shift N, --check CI, JSON) ✓ 2026-08-01
- [x] file-extension-report : rapport de fréquence d'extensions dans une arborescence (top-N, tailles, --threshold CI, JSON) ✓ 2026-08-01

## Vague 230 — CLI Tools (diff JSON structurel, dédupe adjacente, slice CSV, extraction domaine URL, répétition lignes)
- [x] json-pretty-diff : diff structurel de deux documents JSON (chemins modifiés, added/removed/changed, --check CI, JSON) ✓ 2026-08-01
- [x] text-dedupe-adjacent : supprime les lignes dupliquées consécutives (style uniq, --count, --ignore-case, --check CI, JSON) ✓ 2026-08-01
- [x] csv-slice-rows : extrait une tranche de lignes CSV (start:stop:step style Python, négatifs, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-domain : extrait schéma/host/port/path des URLs en champs tabulés (subdomain splitting, --check CI, JSON) ✓ 2026-08-01
- [x] text-repeat-lines : répète chaque ligne N fois (N global ou champ préfixe, --interleave, --check CI, JSON) ✓ 2026-08-01

## Vague 229 — CLI Tools (aplatissement JSON feuilles, sélection colonnes, dédupe CSV, normalisation URL, padding colonnes)
- [x] json-flatten-leaf : aplatit un JSON imbriqué en paires chemin=valeur (feuilles seules, sep custom, --check CI, JSON) ✓ 2026-08-01
- [x] text-column-select : sélectionne des colonnes whitespace-separated (comme cut -f, multi-ranges 1-3,5, --invert, --check CI, JSON) ✓ 2026-08-01
- [x] csv-dedupe-rows : supprime les lignes dupliquées d'un CSV (clés colonnes subsets, keep first/last, --check CI, JSON) ✓ 2026-08-01
- [x] url-normalize-cli : normalise des URLs (lowercase host, default ports, slash trailing, sort params, --check CI, JSON) ✓ 2026-08-01
- [x] text-pad-columns : aligne des colonnes whitespace en padding (left/right/center par colonne, --min-width, --check CI, JSON) ✓ 2026-08-01

## Vague 228 — CLI Tools (renommage clés JSON, wrap largeur, numérotation CSV, paramètres tracking URL, tri naturel)
- [x] json-normalize-keys : renomme les clés JSON via map/regex/case (snake/camel/kebab, récursif, --preview, --check CI, JSON) ✓ 2026-08-01
- [x] text-wrap-hard : coupe les lignes à une largeur fixe (CJK-aware, --indent continuation, --check-width CI, JSON) ✓ 2026-08-01
- [x] csv-add-rownumber : ajoute une colonne de numérotation (start/step/pad, --group-by, --position, --check/--last CI, JSON) ✓ 2026-08-01
- [x] url-strip-tracking : retire les paramètres de tracking (utm_*, fbclid, gclid, globs custom, --keep, --check-count CI, JSON) ✓ 2026-08-01
- [x] text-sort-natural : tri naturel avec nombres intégrés (v2 < v10, --field, --numeric, --reverse, --unique, --check CI, JSON) ✓ 2026-08-01

## Vague 227 — CLI Tools (extraction typée JSON, sections texte, transposition CSV, diff query URL, miroir graphemes)
- [x] json-values-by-type : extrait les valeurs JSON d'un type donné (string/number/bool/null, récursif, --check CI, JSON) ✓ 2026-08-01
- [x] text-split-at-marker : découpe un texte en sections délimitées par un marker regex (header extrait, --index N, --count CI, JSON) ✓ 2026-08-01
- [x] csv-transpose-cli : transpose un CSV (lignes ⇄ colonnes, --fill ragged, --strict, --dimensions CI, JSON) ✓ 2026-08-01
- [x] url-query-diff : compare les query strings de deux URLs (params communs/ajoutés/retirés/modifiés, batch stdin, --diff-only CI, JSON) ✓ 2026-08-01
- [x] text-mirror-pairs : retourne chaque ligne en miroir (grapheme-aware: accents, ZWJ emoji, drapeaux, --check-palindrome CI, JSON) ✓ 2026-08-01

## Vague 226 — CLI Tools (statistiques JSON, jointure lignes, taille répertoires, échappement URL, largeur CJK)
- [x] json-stats-fields : statistiques numériques par champ d'une collection JSON/JSONL (min/max/mean/median/stdev, --check CI, JSON) ✓ 2026-08-01
- [x] text-join-by-column : jointure de deux fichiers lignes par colonne clé (1-based, inner/left, --check CI, JSON) ✓ 2026-08-01
- [x] dir-size-report : rapport de taille par sous-dossier (top-N, excludes glob, --threshold CI, JSON) ✓ 2026-08-01
- [x] url-encode-cli : encode/décode composants URL (percent-encoding, form-urlencoded, batch stdin, --check CI, JSON) ✓ 2026-08-01
- [x] text-display-width : largeur d'affichage des lignes en tenant compte des caractères larges CJK/east-asian (--limit CI, JSON) ✓ 2026-08-01

## Vague 225 — CLI Tools (opérations SET, scoring, commentaires CSV, colonnes JSON, shuffle)
- [x] text-set-ops : opérations ensemblistes sur fichiers de lignes (union/intersection/différence --a-only --b-only --xor via stdin, --check CI, JSON) ✓ 2026-08-01
- [x] text-keyword-score : score la densité de mots-clés par ligne/document (pondérations mot:valeur, --top N, --threshold CI, JSON) ✓ 2026-08-01
- [x] csv-strip-comments : retire les lignes commentées (# prefix custom) d'un fichier CSV (header préservé, --check CI, JSON) ✓ 2026-08-01
- [x] json-pluck-columns : extrait une liste de champs clé/imbriqués (dot-paths) d'une liste d'objets JSON/JSONL (--keep-structure, CSV output, --check CI, JSON) ✓ 2026-08-01
- [x] text-shuffle-lines : mélange aléatoirement les lignes d'un fichier/flux (seed reproductible, --sample N reservoir, --check CI, JSON) ✓ 2026-08-01

## Vague 224 — CLI Tools (rapport ANSI, dates CSV ISO, comptage JSON, collapse espaces, conversions hex)
- [x] text-ansi-report : rapport sur les séquences ANSI présentes (couleurs SGR, curseur, styles, --check CI, JSON) ✓ 2026-08-01
- [x] csv-date-normalize : normalise les dates d'une colonne CSV vers ISO 8601 (formats multiples, --check CI, JSON) ✓ 2026-08-01
- [x] json-count-values : histogramme des valeurs d'un champ JSON/JSONL (top-N, --check CI, JSON) ✓ 2026-08-01
- [x] text-collapse-space : remplace les runs d'espaces multiples par un séparateur custom (tab, --squeeze, --check CI, JSON) ✓ 2026-08-01
- [x] hex-to-dec-cli : convertit hex <-> decimal/octet/bin pour flux ou args (prefix, padding, --check CI, JSON) ✓ 2026-08-01

## Vague 223 — CLI Tools (dédoublonnage JSON, entêtes CSV, wrapping URL, checksums texte, swapping colonnes TSV)
- [x] json-array-dedupe : déduplique les éléments d'un tableau JSON (clé de comparaison, stable, JSONL, --check CI, JSON) ✓ 2026-08-01
- [x] csv-drop-columns : supprime des colonnes CSV par nom/index (keep-list inverse, --check CI, JSON) ✓ 2026-08-01
- [x] url-defang : neutralise/réarme URLs et IPs pour partage sûr (hxxp, [.], --refang inverse, --check CI, JSON) ✓ 2026-08-01
- [x] text-line-checksum : ajoute/vérifie un hash par ligne (md5/sha256, format NEC-like, --verify CI, JSON) ✓ 2026-08-01
- [x] tsv-to-csv : convertit TSV ⇄ CSV (--to-tsv inverse, quoting minimal, --check CI, JSON) ✓ 2026-08-01

## Vague 221 — CLI Tools (indentation, colonnes CSV, timestamps texte, dupes JSON, padding CSV)
- [x] text-indent-detect : détecte le style d'indentation d'un fichier (tabs vs espaces, largeur, mixte, --require CI, JSON) ✓ 2026-08-01
- [x] csv-swap-columns : échange ou réordonne des colonnes CSV par nom/index (--swap, --order, --check CI) ✓ 2026-08-01
- [x] text-epoch-embed : remplace les timestamps Unix d'un texte par des dates ISO (--to-epoch inverse, --tz, --check CI, JSON) ✓ 2026-08-01
- [x] json-dup-keys : détecte les clés dupliquées dans un objet JSON (object_pairs_hook, positions, --check CI) ✓ 2026-08-01
- [x] csv-align-columns : réaligne la largeur des champs d'un CSV en sortie texte/table (markdown, unicode, --max-width) ✓ 2026-08-01

## Vague 222 — CLI Tools (portée lignes, JSON, casse URL, BOM, effective CIDR)
- [x] text-excerpt : extrait un extrait centré sur une ligne/pattern (contexte +/-N, --mark, JSON) ✓ 2026-08-01
- [x] json-key-case : uniformise la casse des clés JSON (snake/camel/kebab/upper, récursif, --check CI) ✓ 2026-08-01
- [x] url-lowercase-host : normalise la casse scheme/host des URLs (domaine seul, --check CI, JSON) ✓ 2026-08-01
- [x] text-strip-bom : retire BOM UTF-8/UTF-16 de fichiers (--add inverse, in-place, --check CI, JSON) ✓ 2026-08-01
- [x] ip-network-of : calcule network/broadcast d'une IP+CIDR (supernet, --contains, JSON, CI) ✓ 2026-08-01

## Vague 220 — CLI Tools (quoting CSV, hyperliens OSC 8, query string, entourage lignes, stat git)
- [x] csv-quote-all : force le quoting de toutes (ou certaines) cellules CSV (--all/--columns, --none inverse, --check CI) ✓ 2026-08-01
- [x] text-hyperlink-wrap : encadre des URLs/texte en hyperliens OSC 8 de terminal (--strip inverse, --check CI, JSON) ✓ 2026-08-01
- [x] json-to-querystring : convertit un objet JSON en query string URL (nested brackets, --sort, --decode inverse, --check CI) ✓ 2026-08-01
- [x] text-surround-lines : encadre chaque ligne avec préfixe/suffixe (délimiteurs, --pattern, --check CI, JSON) ✓ 2026-08-01
- [x] git-branch-age : classe les branches git locales par âge/activité (dernier commit, auteur, ahead/behind vs main, --stale CI, JSON) ✓ 2026-08-01

## Vague 219 — CLI Tools (largeurs lignes, colonnes CSV, JSON, query params, continuations)
- [x] text-wide-quality : lint de largeur de lignes (distribution, percentiles, --limit CI, JSON, ignore-glob) ✓ 2026-08-01
- [x] csv-column-prefix : préfixe/suffixe les noms de colonnes CSV (--prefix/--suffix/--only, dry-run, --check CI) ✓ 2026-08-01
- [x] json-array-flatten : aplatit les tableaux imbriqués d'un JSON (--depth, mode scalars, JSONL, --check CI) ✓ 2026-08-01
- [x] url-add-params : ajoute/remplace/supprime des paramètres de query d'URLs (batch stdin, --sort, JSON, --check CI) ✓ 2026-08-01
- [x] text-backslash-join : joint les lignes en continuation backslash (style shell, --join-with, --split inverse, --check CI) ✓ 2026-08-01

## Vague 218 — CLI Tools (lint JSON, escape CSV, path URL, blank lines, sharding CSV)
- [x] json-required-keys : lint CI sur clés requises d'un JSON (dot-paths, indices, --non-empty, JSON) ✓ 2026-08-01
- [x] text-csv-escape : échappe du texte libre en champ CSV RFC 4180 (quote double, --unescape, --check roundtrip) ✓ 2026-08-01
- [x] url-path-clean : normalise des paths d'URL lexicalement (//, ., .., trailing, query/fragment préservés, --check CI) ✓ 2026-08-01
- [x] text-blank-lines-report : rapport des lignes vides (runs, leading/trailing, limites CI --max-*, JSON) ✓ 2026-08-01
- [x] csv-row-hash-select : sélection déterministe de lignes CSV par hash de clé (--mod/--bucket sharding, --fraction, md5/sha1/sha256) ✓ 2026-08-01

## Vague 217 — CLI Tools (stats .env, espaces, cellules CSV, IPv6, nulls JSON)
- [x] env-count : compte variables/commentaires/lignes vides d'un .env (dupes, exported, --min CI, JSON) ✓ 2026-08-01
- [x] text-space-normalize : réduit les runs d'espaces/tabs à un seul espace (indent préservé, trim trailing, --check CI, JSON) ✓ 2026-08-01
- [x] csv-cell-extract : extrait des cellules/plages rectangulaires d'un CSV (rows/cols 1-based, noms d'header, CSV/lines/JSON) ✓ 2026-08-01
- [x] text-ipv6-expand : compresse/étend les adresses IPv6 dans du texte (ipaddress stdlib, --extract, --strict, --check CI, JSON) ✓ 2026-08-01
- [x] json-strip-nulls : retire récursivement les valeurs null/vides d'un JSON (empty-strings/arrays/objects, JSONL, --check CI) ✓ 2026-08-01

## Vague 216 — CLI Tools (diacritiques, délimiteur CSV, RFC 7386, timestamps, quoting shell)
- [x] text-diacritic-fold : retire les accents/diacritiques du texte (Latin-1/A, ligatures, --check CI, JSON) ✓ 2026-08-01
- [x] csv-detect-delimiter : détecte le séparateur d'un CSV (quote-aware, scoring, JSON report, exit 2) ✓ 2026-08-01
- [x] json-rfc7386 : applique des JSON Merge Patches RFC 7386 (multi-patches, compact/indent, tests Appendix A) ✓ 2026-08-01
- [x] ts-convert : convertit timestamps Unix ⇆ ISO 8601 (s/ms/µs/ns par magnitude, --layout custom, --tz) ✓ 2026-08-01
- [x] shell-quote : échappe une chaîne pour POSIX sh / PowerShell / cmd.exe (anti-injection, stdin batch, JSON) ✓ 2026-08-01

## Vague 215 — CLI Tools (jointure CSV, liens Markdown, schéma JSON, fréquences lignes, quotes CSV)
- [x] csv-join : jointure relationnelle de deux CSV sur colonnes clés (inner/left/right/outer, clés composites, délimiteurs par fichier, --check CI, JSON) ✓ 2026-08-01
- [x] text-markdown-links : extrait et valide les liens Markdown (inline, référence, autolinks, URLs brutes, --urls-only, --check CI, JSON) ✓ 2026-08-01
- [x] json-schema-infer : infère un schéma JSON Schema 2020-12 depuis JSON/JSONL (types, required, formats, enum, bornes, --validate CI) ✓ 2026-08-01
- [x] file-line-frequency : table de fréquences des lignes (top-N, trim/case, pourcentages, CSV/JSON, --max-share CI) ✓ 2026-08-01
- [x] csv-quote-fix : normalise le style de quoting CSV (minimal/all/non-numeric/none, dialectes entrée, --check CI, stats JSON) ✓ 2026-08-01

## Vague 214 — CLI Tools (stopwords, records CSV, truncate, ménage fichiers, durées)
- [x] text-stopword-remove : retire les stopwords d'un texte (listes en/fr embarquées, listes custom, --check CI, JSON) ✓ 2026-08-01
- [x] csv-to-records : rend un CSV en blocs key: value (ligne unique --row, alignement, JSON, délimiteur custom) ✓ 2026-08-01
- [x] text-smart-truncate : tronque un texte sur les bords de mots (ellipse, milieu, --whole, --check CI) ✓ 2026-08-01
- [x] file-temp-cleanup : trouve/supprime les fichiers temporaires (globs, âge, taille, dry-run, JSON) ✓ 2026-08-01
- [x] duration-humanize : convertit secondes ⇆ durées humaines (1h30m, ISO 8601, long, batch stdin) ✓ 2026-08-01

## Vague 213 — CLI Tools (Unicode, hash CSV, chemins, crochets, masquage IP)
- [x] text-unicode-lookup : inspecte caractères Unicode (codepoint, nom officiel, catégorie, --search, JSON) ✓ 2026-08-01
- [x] csv-add-hashcol : ajoute une colonne de hash à un CSV (md5/sha*, colonnes ciblées, --short, --check CI) ✓ 2026-08-01
- [x] file-path-normalize : normalise des chemins lexicalement (., .., //, séparateurs, --base, --check CI) ✓ 2026-08-01
- [x] text-bracket-balance : vérifie l'équilibre des crochets/quotes (strings et commentaires ignorés, positions, JSON) ✓ 2026-08-01
- [x] ip-masking : masque les IPv4/IPv6 dans les logs (octets/hash déterministe, CIDR conservés, --check CI) ✓ 2026-08-01

## Vague 212 — CLI Tools (acronymes, lignes, sous-réseaux, romains, voyelles)
- [x] text-acronym-expand : scanne et étend les acronymes dans un texte (dictionnaire embarqué, dict custom, --expand/--mark, JSON, CI) ✓ 2026-08-01
- [x] file-line-picker : extrait/supprime des lignes par numéro/range (1-based, négatifs, --drop, --check CI, JSON) ✓ 2026-08-01
- [x] ip-subnet-planner : découpe un bloc CIDR IPv4 en sous-réseaux (--count/--prefix/--hosts, netmask, hostrange, JSON) ✓ 2026-08-01
- [x] text-roman-numerals : convertit chiffres romains ⇆ entiers (1..3999, --check canonique, batch stdin, JSON) ✓ 2026-08-01
- [x] text-vowel-count : compte voyelles/consonnes par ligne ou global (accents foldés, cyrillique, ratio, --require CI, JSON) ✓ 2026-08-01

## Vague 211 — CLI Tools (UUID, SemVer, Luhn, syllabes, n-grams)
- [x] uuid-validate : valide et inspecte des UUID (toutes versions/variantes, formes canoniques/URN/braced, --extract, --check CI, JSON) ✓ 2026-08-01
- [x] semver-sort : trie des versions avec précédence SemVer 2.0.0 complète (prereleases, --lenient, --unique, --check CI, JSON) ✓ 2026-08-01
- [x] luhn-check : valide et génère des chiffres de contrôle Luhn (cartes, IMEI, SIREN ; schémas détectés, --generate, masquage, JSON) ✓ 2026-08-01
- [x] text-syllable-count : estime les syllabes en anglais par mot et par ligne (exceptions intégrées, --max/--exact poésie, JSON) ✓ 2026-08-01
- [x] text-ngrams-cli : n-grammes caractères ou mots (fréquences, --unique, --require/--forbid CI, JSON, NUL) ✓ 2026-08-01

## Vague 210 — CLI Tools (émojis, phrases, IPv4, commentaires, palindromes)
- [x] text-emoji-scan : détecte et rapporte les émojis d'un texte (codepoint, nom Unicode, catégorie, occurrences, --check CI, --strip, JSON) ✓ 2026-08-01
- [x] text-sentence-split : découpe un texte en phrases (abréviations/décimales gérées, --stats, JSON, numérotation) ✓ 2026-08-01
- [x] ip-to-integer : convertit IPv4 ⇆ entier 32 bits (hex, binaire, batch stdin, JSON) ✓ 2026-08-01
- [x] text-strip-comments : retire les commentaires (#, //, /* */) en préservant les chaînes littérales (--check CI, styles, JSON) ✓ 2026-08-01
- [x] text-palindrome-check : vérifie si des textes sont des palindromes (accents/ponctuation/casse ignorés, --check, JSON) ✓ 2026-08-01

## Vague 209 — CLI Tools (MAC, mots de passe, CSV colonnes, HTML, nombres)
- [x] mac-vendor-lookup : identifie le fabricant d'une adresse MAC (base OUI embarquée, normalisation, JSON) ✓ 2026-08-01
- [x] text-password-strength : évalue la robustesse d'un mot de passe (entropie, classes, motifs, JSON) ✓ 2026-08-01
- [x] csv-pick-columns : projection de colonnes CSV par nom/index/range (exclude, ordre, --check CI) ✓ 2026-08-01
- [x] text-html-escape : encode/décode les entités HTML (named/decimal/hex, --check roundtrip, JSON) ✓ 2026-08-01
- [x] number-to-words : convertit des nombres en toutes lettres (fr/en, ordinals, JSON) ✓ 2026-08-01

## Vague 208 — CLI Tools (CIDR, whitespace guillemets, CSV filtering, ROT, case)
- [x] ip-cidr-tools : calcules CIDR (network/broadcast/hosts, contains, split, JSON) ✓ 2026-08-01
- [x] text-straight-quotes : convertit guillemets/apostrophes typographiques en ASCII (--check CI) ✓ 2026-08-01
- [x] csv-filter-not-null : ne garde que les lignes sans cellules vides (colonnes ciblées, JSON) ✓ 2026-08-01
- [x] text-split-camel : découpe camelCase/PascalCase en mots (snake/kebab/space, JSON) ✓ 2026-08-01
- [x] csv-row-reverse : inverse l'ordre des lignes de données (header conservé, --check CI) ✓ 2026-08-01

## Vague 207 — CLI Tools (padding, diff CSV, unflatten, rot, invisible chars)
- [x] text-pad-lines : pousse les lignes à une largeur cible (left/right/center, fill char, truncate, --check CI) ✓ 2026-08-01
- [x] csv-diff-rows : compare deux CSV par colonnes clés (added/removed/changed, JSON, --exit-code CI) ✓ 2026-08-01
- [x] json-unflatten : inverse de json-flatten (clés dotted + [idx], JSONL, --check CI) ✓ 2026-08-01
- [x] text-rot-cipher : ROT13/ROT47/Caesar arbitraire (shift, digits, JSON) ✓ 2026-08-01
- [x] text-invisible-chars : détecte/retire les caractères Unicode invisibles (Cf, bidi, ZWSP, BOM mid-text, --check CI) ✓ 2026-08-01
- [x] csv-count-occurrences : tables de fréquences des valeurs de colonnes CSV (top-N, pourcentages, JSON/CSV) ✓ 2026-08-01

## Vague 206 — CLI Tools (tabs, regex CSV, sampling, ANSI, tally)
- [x] text-expand-tabs : convertit tabs <-> espaces (expand/unexpand, tab size, --check CI) ✓ 2026-08-01
- [x] csv-row-filter-regex : filtre les lignes CSV par regex sur une colonne (invert, count, exit codes grep) ✓ 2026-08-01
- [x] text-random-sample : échantillonne des lignes au hasard (reservoir sampling, seed, fraction/compte) ✓ 2026-08-01
- [x] text-ansi-strip : retire les séquences ANSI (couleurs, curseur, hyperliens, --check CI) ✓ 2026-08-01
- [x] csv-tally : table de contingence entre deux colonnes CSV (totaux, JSON) ✓ 2026-08-01

## Vague 205 — CLI Tools (swap CSV, pig latin, censure, julian, profondeur JSON)
- [x] csv-swap-columns : échange deux colonnes d'un CSV (nom/index, no-header, --check CI) ✓ 2026-08-01
- [x] text-pig-latin : traduit du texte en Pig Latin (règles voyelle/consonne, casse préservée, JSON) ✓ 2026-08-01
- [x] text-censor-words : censure des mots dans un texte (masque, bords préservés, --check CI, JSON) ✓ 2026-08-01
- [x] julian-date : conversion dates calendaires <-> jours juliens (MJD, dates ordinales, JSON) ✓ 2026-08-01
- [x] json-depth-count : histogramme des valeurs JSON par profondeur et type (--max-depth lint, JSONL) ✓ 2026-08-01

## Vague 204 — CLI Tools (group-by CSV, numérotation, stdev, chunks, shuffle)
- [x] csv-group-by : groupe les lignes CSV par clé et agrège (count/sum/mean/min/max, distinct, JSON) ✓ 2026-08-01
- [x] text-number-lines : numérote les lignes style nl (styles, largeur, séparateur, start/incr) ✓ 2026-08-01
- [x] csv-stdev : écart-type et variance des colonnes CSV numériques (population/échantillon, JSON) ✓ 2026-08-01
- [x] text-chunk-lines : découpe un texte en chunks de N lignes (split -l, suffixes alpha/num, dry-run) ✓ 2026-08-01
- [x] csv-row-shuffle : mélange les lignes de données d'un CSV (seed reproductible, header conservé) ✓ 2026-08-01

## Vague 203 — CLI Tools (tri CSV, wrap texte, doublons fichiers, Soundex, transpose/md, title case)
- [x] csv-sort-rows : trie les lignes CSV par colonnes (numérique, naturel, reverse, ignore-case, --check CI, JSON) ✓ 2026-08-01
- [x] text-word-wrap : reformatage de paragraphes à largeur cible (indents, préfixe, --break-long, --check CI) ✓ 2026-08-01
- [x] file-dup-detect : détecte les fichiers dupliqués par hash (pré-filtre taille+hash partiel, --delete, dry-run, NUL) ✓ 2026-08-01
- [x] text-soundex : codes Soundex américains (accents foldés, --compare paires stdin, --check CI, JSON) ✓ 2026-08-01
- [x] csv-transpose-md : transpose un CSV et/ou le rend en table Markdown (alignements, échappement pipes, JSON) ✓ 2026-08-01
- [x] text-title-case : Title Case anglais AP/Chicago (mots mineurs, acronymes, composés à tiret, --check CI) ✓ 2026-08-01

## Vague 202 — CLI Tools (CSV fill, Damerau, MIME, JSON path, hexdump)
- [x] csv-fill-blanks : remplit les cellules vides CSV (valeur globale, mapping par colonne, check CI, JSON) ✓ 2026-08-01
- [x] text-damerau-levenshtein : distance Damerau-Levenshtein avec transpositions (stdin tab/comma, --max CI, JSON) ✓ 2026-08-01
- [x] file-mime-detect : détecte le type MIME par magic bytes (images, archives, docs, exécutables, --require CI) ✓ 2026-08-01
- [x] json-get-path : extraction de valeurs JSON par chemins dot/bracket (wildcards, index négatifs, --default, --required) ✓ 2026-08-01
- [x] text-hex-dump : dump hexadécimal fichiers/stdin (offset hex, longueur, largeur, ASCII sidebar) ✓ 2026-08-01

## Vague 201 — CLI Tools (CSV types, Hamming, entropy, JSON merge, base64url)
- [x] csv-column-type : infère les types des colonnes CSV (int/float/bool/date/email/uuid/url, confiance, strict, JSON) ✓ 2026-08-01
- [x] text-hamming-distance : distance de Hamming entre chaînes (stdin tab/comma, --max CI, JSON) ✓ 2026-08-01
- [x] file-entropy : entropie de Shannon de fichiers (classification, --min/--max CI, JSON, stdin) ✓ 2026-08-01
- [x] json-merge-patch : applique un RFC 7386 JSON Merge Patch (stdin, --check CI, compact, -o) ✓ 2026-08-01
- [x] text-base64-url : encode/décode base64url RFC 4648 §5 (padding optionnel, strict, JSON) ✓ 2026-08-01

## Vague 200 — CLI Tools (CSV rows, levenshtein, JSON keys, gitignore, URL)
- [x] csv-drop-rows : supprime les lignes CSV par conditions (eq/ne/regex/empty/gt/lt, AND/OR, invert, check CI) ✓ 2026-08-01
- [x] text-levenshtein-distance : distance d'édition Levenshtein + ratio (paires stdin, --within CI, JSON) ✓ 2026-08-01
- [x] json-pick-keys : projection de clés JSON/JSONL (dot paths, pick/omit, --require, compact) ✓ 2026-08-01
- [x] file-gitignore-match : teste des chemins contre un .gitignore (glob/**/!/ancrage, exit codes grep) ✓ 2026-08-01
- [x] url-parse-decompose : décompose des URLs (scheme/host/port/params, --get, --require, JSON) ✓ 2026-08-01

## Vague 199 — CLI Tools (CSV slices, redaction, fichiers vides, headers, base91)
- [x] csv-slice : extraction de tranches CSV (ranges de lignes 1-based/-n:, colonnes nom/index, inversion, JSONL) ✓ 2026-08-01
- [x] text-redact : masque données sensibles (email, IP, URL, CB, JWT, clés AWS, patterns custom, check CI) ✓ 2026-08-01
- [x] file-empty-files : trouve et supprime les fichiers vides ou whitespace-only (glob, dry-run, NUL, CI) ✓ 2026-08-01
- [x] csv-header-normalize : normalise les en-têtes CSV (snake/slug/camel/pascal/upper, accents, doublons, check CI) ✓ 2026-08-01
- [x] base91-cli : encode/décode basE91 (binaire-safe, wrap, strict, roundtrip vérifié) ✓ 2026-08-01

## Vague 198 — CLI Tools (pivot CSV, Morse, JSON paths, query strings, git tags)
- [x] csv-pivot-table : tableaux croisés depuis un CSV (sum/count/mean/min/max, totaux, tri, JSON) ✓ 2026-08-01
- [x] text-morse-code : encode/décode le code Morse international (symboles custom, strict, JSON) ✓ 2026-08-01
- [x] json-path-get : extraction de valeurs JSON par chemins dot/bracket (wildcards, index négatifs, --default) ✓ 2026-08-01
- [x] http-query-build : construit des query strings URL (paires key=value, JSON, env, tri, encodage) ✓ 2026-08-01
- [x] git-tag-list : liste et analyse les tags git (semver, dates, prereleases, gaps, JSON) ✓ 2026-08-01

## Vague 197 — CLI Tools (CSV mapping, templates, substr, JSON analytics, MAJ CSV)
- [x] csv-rename-columns : renomme colonnes CSV en masse (map JSON, OLD=NEW, regex, --snake/--lower, dry-run) ✓ 2026-08-01
- [x] text-expand-vars : expansion ${VAR}/${VAR:-default} depuis env/.env/JSON (strict, list, sortie fichier) ✓ 2026-08-01
- [x] text-substring-lines : extraction de sous-chaînes par slice, range 1-based, largeur, ou regex ✓ 2026-08-01
- [x] json-tree-depth : analyse structure JSON/JSONL (profondeur max, nœuds, histogramme clés, --max-depth) ✓ 2026-08-01
- [x] csv-update-rows : mise à jour ciblée de cellules CSV par condition (eq/regex, templating {col}, dry-run) ✓ 2026-08-01

## Vague 196 — CLI Tools (CSV, whitespace, line endings, similarité)
- [x] csv-move-columns : déplace/réordonne les colonnes d'un CSV par nom (before/after, order, drop) ✓ 2026-08-01
- [x] text-trailing-fix : détecte/corrige espaces en fin de ligne et lignes vides finales (CI check, JSON) ✓ 2026-08-01
- [x] line-endings-cli : détecte/convertit les fins de ligne LF/CRLF/CR (mixed, check CI, récursif) ✓ 2026-08-01
- [x] text-jaro-winkler : similarité Jaro / Jaro-Winkler (score, similar top-N, JSON) ✓ 2026-08-01
- [x] csv-row-dedupe : supprime les lignes CSV dupliquées par colonnes clés (ordre préservé, keep-last, check) ✓ 2026-08-01

## Vague 195 — CLI Tools (génération texte, CSV, HTTP, JSONL, ASCII art)
- [x] text-lorem-gen : générateur de texte lorem ipsum (mots/phrases/paragraphes, seed reproductible, JSON) ✓ 2026-08-01
- [x] csv-add-rownum : ajoute une colonne de numérotation à un CSV (nom, position, début, pas, padding) ✓ 2026-08-01
- [x] http-range-fetch : télécharge des plages d'octets via HTTP Range (probe, suffix, resume, hex) ✓ 2026-08-01
- [x] file-json-lines : conversion texte <-> JSON Lines avec validation ligne par ligne ✓ 2026-08-01
- [x] text-ascii-banner : bannières ASCII art avec police bitmap 5x5 intégrée ✓ 2026-08-01

## Vague 194 — CLI Tools (texte, csv, fichiers, tri)
- [x] text-column-align : aligne du texte délimité en colonnes monospacées (auto-détection, alignements, header, JSON) ✓ 2026-08-01
- [x] csv-grep-rows : grep pour CSV avec conditions par colonne (regex, numériques, AND/OR, JSON, exit codes grep) ✓ 2026-08-01
- [x] file-hash-watch : snapshot de hashes + détection added/modified/deleted (check, watch, CI-friendly) ✓ 2026-08-01
- [x] text-sort-natural : tri naturel (versions, IPs, numérique, random seedé, --check) ✓ 2026-08-01
- [x] file-bom-tool : détecte/ajoute/supprime les BOM Unicode (UTF-8/16/32, check CI, récursif) ✓ 2026-08-01

## Vague 185 — CLI Tools
- [x] percent-change : calc des variations % (simple, série, chaînée) ✓ 2026-07-28
- [x] list-dedup : supprime doublons de lignes préservant l'ordre (i/t/n/d) ✓ 2026-07-28
- [x] domain-expiry : quiis whois + jours restants avant expiration ✓ 2026-07-28
- [x] text-metrics : indices de lisibilité (Flesch, FK, SMOG, CLI, ARI, LIX) ✓ 2026-07-28
- [x] env-shell : export .env vers sh/ps1/json/list/check ✓ 2026-07-28

## Vague 186 — CLI Tools
- [x] csv-row-count : compte lignes + champs CSV (multi-fichiers, JSON) ✓ 2026-07-28
- [x] file-ext-count : stats fichiers par extension (count/bytes, JSON) ✓ 2026-07-28
- [x] port-knock : client de port-knocking TCP (séquence de ports) ✓ 2026-07-28
- [x] text-diff-lite : diff unifié entre deux fichiers (stat, JSON) ✓ 2026-07-28
- [x] number-base-convert : conversion bases 2-64 (auto-détection préfixes) ✓ 2026-07-28

## Vague 187 — CLI Tools
- [x] text-to-slug-batch : slugs URL-friendly par ligne (batch, stdin, JSON) ✓ 2026-07-28
- [x] markdown-to-rst : convertisseur Markdown vers reStructuredText (headers, lists, links, code) ✓ 2026-07-28
- [x] http-post-data : envoie requêtes HTTP avec body JSON/form/raw (méthode custom, headers, timing) ✓ 2026-07-28
- [x] file-perm-copy : copie/sets permissions fichiers (mode, uid, gid, -r) ✓ 2026-07-28
- [x] json-merge-deep : fusion profonde de JSON (récursive, -concat, -q, -indent) ✓ 2026-07-28

## Vague 188 — CLI Tools
- [x] text-dup-lines : trouve et rapporte les lignes dupliquées (i/t/n, count-only, JSON) ✓ 2026-07-28
- [x] http-delay-test : benchmark de latence HTTP (percentiles p50-p99, concurrence, JSON) ✓ 2026-07-28
- [x] file-empty-dirs : trouve et supprime les dossiers vides (ignore-hidden, .gitkeep, multi-pass delete) ✓ 2026-07-28
- [x] json-key-rename : renomme les clés JSON récursivement (mapping, regex, camel/snake, in-place) ✓ 2026-07-28
- [x] text-case-convert : conversion entre conventions de nommage (camel, snake, kebab, pascal...) ✓ 2026-07-28

## Vague 189 — CLI Tools
- [x] unicode-inspect : inspecte caractères Unicode (codepoints, names, blocks, NFKC, détecte confusables) ✓ 2026-07-28
- [x] service-banner : grab banners + identifie services sur ports TCP (SSH/HTTP/FTP/MySQL/Redis...) ✓ 2026-07-28
- [x] tmux-session-save : sauvegarde/restaure sessions tmux (windows, panes, layouts, cwd) ✓ 2026-07-28
- [x] cron-html-report : audit tout crons + systemd timers → rapport HTML/JSON avec prochaines exécutions ✓ 2026-07-28
- [x] ssh-known-hosts-audit : audite ~/.ssh/known_hosts (algos faibles, doublons, classes) ✓ 2026-07-28

## Vague 193 — CLI Tools (CSV, git, JSON, texte, HTTP)
- [x] csv-stats-full : statistiques descriptives complètes par colonne CSV (min/max/mean/stddev/quartiles ou top catégories) ✓ 2026-07-28
- [x] git-repo-hotspots : classe les fichiers git par fréquence de changement et churn (lignes ajoutées+supprimées) ✓ 2026-07-28
- [x] json-key-prefix : ajoute/retire un préfixe sur les clés JSON récursivement (in-place, stdin) ✓ 2026-07-28
- [x] text-word-frequency : compteur de fréquences de mots (multi-fichiers, ppm, JSON) ✓ 2026-07-28
- [x] http-head-check : requêtes HEAD avec chaîne de redirections complète, timings et headers ✓ 2026-07-28

## Vague 192 — CLI Tools (JSON, CSV, HTTP, fichiers)
- [x] json-fuzz-validate : fuzz-validation structurelle d'un JSON (profondeur, tailles, nulls, vides, key patterns, nums, dup-keys) ✓ 2026-07-28
- [x] csv-multi-sort : tri CSV multi-colonnes avec direction et type par colonne (str/num/date) ✓ 2026-07-28
- [x] text-uniq-adjacent : supprime les lignes dupliquées adjacentes (style uniq -c -i -d -u) ✓ 2026-07-28
- [x] http-basic-client : client HTTP minimal (méthodes, headers, body, basic auth, redirects, timing, -o) ✓ 2026-07-28
- [x] file-size-distribution : histogramme des tailles de fichiers (buckets, top extensions, biggest) ✓ 2026-07-28

## Vague 191 — CLI Tools (markdown & texte)
- [x] markdown-title-case : convertit les titres Markdown en title case (APA/Chicago) ou sentence case ✓ 2026-07-28
- [x] json-path-exists : vérifie l'existence de chemins dans un JSON (dot notation, [idx], wildcards) ✓ 2026-07-28
- [x] file-strings : extrait les chaînes imprimables d'un binaire (ASCII/UTF-16, offsets, JSON) ✓ 2026-07-28
- [x] http-options-cli : requêtes OPTIONS + audit CORS pre-flight (Allow, Access-Control-*, probes) ✓ 2026-07-28
- [x] text-reverse-chars : inverse les caractères (par ligne, mot, ordre des mots, stream) ✓ 2026-07-28

## Vague 190 — CLI Tools
- [x] file-hardlinks : trouve et rapporte les groupes de hardlinks (inode, paths, JSON) ✓ 2026-07-28
- [x] text-wrap : formate du texte en colonnes (justify, prefix, width) ✓ 2026-07-28
- [x] json-stats : statistiques sur un JSON (types, profondeur, tailles) ✓ 2026-07-28
- [x] http-etag-check : vérifie ETag/Last-Modified et retourne 304 si match ✓ 2026-07-28
- [x] csv-column-reorder : réordonne colonnes CSV (noms/indices, ranges, drop, in-place) ✓ 2026-07-28

## Terminés
- [x] foundry-autopilot-dashboard : dashboard Go temps réel pour cycles, revues, PR, historique et journal systemd ✓ 2026-07-27
- [x] percent-change : calc des variations % (simple, série, chaînée) ✓ 2026-07-28
- [x] list-dedup : supprime doublons de lignes préservant l'ordre (i/t/n/d) ✓ 2026-07-28
- [x] domain-expiry : quiis whois + jours restants avant expiration ✓ 2026-07-28
- [x] text-metrics : indices de lisibilité (Flesch, FK, SMOG, CLI, ARI, LIX) ✓ 2026-07-28
- [x] env-shell : export .env vers sh/ps1/json/list/check ✓ 2026-07-28
- [x] csv-row-count : compte lignes + champs CSV (multi-fichiers, JSON) ✓ 2026-07-28
- [x] file-ext-count : stats fichiers par extension (count/bytes, JSON) ✓ 2026-07-28
- [x] port-knock : client de port-knocking TCP (séquence de ports) ✓ 2026-07-28
- [x] text-diff-lite : diff unifié entre deux fichiers (stat, JSON) ✓ 2026-07-28
- [x] number-base-convert : conversion bases 2-64 (auto-détection préfixes) ✓ 2026-07-28
(mettre à jour ici chaque outil fini)

## En cours

## Vague 28 — CLI Tools
- [x] csv-header : affiche, ajoute, supprime ou renomme les headers CSV ✓ 2026-05-28
- [x] json-to-markdown : convertisseur JSON vers tableau Markdown ✓ 2026-05-28
- [x] csv-head : affiche les N premières lignes d'un CSV ✓ 2026-05-28
- [x] git-diff-files : liste les fichiers modifiés entre deux commits ✓ 2026-05-28
- [x] toml-lint : vérifie la syntaxe et le style TOML ✓ 2026-05-28

## Vague 29 — CLI Tools
- [x] env-list : liste les variables d'environnement triées ✓ 2026-05-28
- [x] file-age : affiche l'âge des fichiers (création, modification) ✓ 2026-05-28
- [x] json-schema-gen : génère un JSON Schema depuis un fichier JSON ✓ 2026-05-28
- [x] http-headers : affiche les headers HTTP d'une URL ✓ 2026-05-28
- [x] ini-to-json : convertisseur INI vers JSON ✓ 2026-05-28

## Vague 30 — CLI Tools
- [x] yaml-to-toml : convertisseur YAML vers TOML ✓ 2026-05-28
- [x] json-validate : valide la syntaxe JSON d'un fichier ✓ 2026-05-28
- [x] csv-to-tsv : convertisseur CSV vers TSV ✓ 2026-05-28
- [x] git-branch-list : liste les branches git avec dernier commit ✓ 2026-05-28
- [x] url-encode : encode/decode des chaînes URL ✓ 2026-05-28

## Vague 26 — CLI Tools
- [x] ncat-lite : client/serveur TCP léger pour le debug réseau ✓ 2026-05-28
- [x] csv-cut : extrait des colonnes spécifiques depuis des fichiers CSV ✓ 2026-05-28
- [x] git-repo-size : affiche la répartition de la taille d'un repo git ✓ 2026-05-28
- [x] json-to-tsv : convertit des tableaux JSON en format TSV ✓ 2026-05-28
- [x] file-touch : crée/mise à jour des timestamps de fichiers en batch ✓ 2026-05-28

## File d'attente — CLI Tools
- [x] snippet-cli : gestionnaire de snippets de commandes shell (sauvegarde, recherche, exécution) ✓ 2026-05-28
- [x] tcp-dump-lite : capture et analyse simplifiée de trafic réseau ✓ 2026-05-28
- [x] curl-format : formateur de requêtes curl avec coloration syntaxique ✓ 2026-05-20
- [x] port-scanner : scanner de ports TCP rapide en Go ✓ 2026-05-20
- [x] env-check : vérifie les variables d'env manquantes pour un projet ✓ 2026-05-20
- [x] log-tail : tail avec coloration par niveau (ERROR=rouge, WARN=jaune, etc.) ✓ 2026-05-20
- [x] git-stat : stats git avancées (contributeurs, lignes/jour, heatmap) ✓ 2026-05-20
- [x] file-watcher : surveille les fichiers et lance des commandes (comme nodemon) ✓ 2026-05-20
- [x] hash-gen : générateur de hash multi-algorithme (md5, sha256, blake2, etc.) ✓ 2026-05-20
- [x] base64-cli : encode/decode base64, base32, hex avec détection auto ✓ 2026-05-20
- [x] qr-gen : génère des QR codes en terminal (ASCII ou image) ✓ 2026-05-20

## File d'attente — APIs & Services
- [x] screenshot-api : API pour capturer des screenshots d'URLs ✓ 2026-05-20
- [x] markdown-api : convertit markdown en HTML/PDF via API ✓ 2026-05-20
- [x] geo-api : géocodage inverse + distance entre coordonnées ✓ 2026-05-20
- [x] currency-api : taux de change en temps réel ✓ 2026-05-20
- [x] whois-api : lookup whois pour domaines ✓ 2026-05-20
- [x] dns-api : résolution DNS avancée (MX, TXT, CNAME, etc.) ✓ 2026-05-20
- [x] image-compress-api : compresse images (jpeg, png, webp) ✓ 2026-05-20
- [x] text-summarizer-api : résumé de texte extractif (sans LLM) ✓ 2026-05-20
- [x] url-shortener : raccourcisseur d'URLs avec stats de clics ✓ 2026-05-20
- [x] paste-service : service de paste type pastebin avec expiration ✓ 2026-05-20

## File d'attente — Bots
- [x] github-notifier-bot : bot Telegram qui notifie les events GitHub ✓ 2026-05-20
- [x] reminder-bot : bot Telegram de rappels avec cron ✓ 2026-05-20
- [x] crypto-price-bot : bot Telegram prix crypto en temps réel ✓ 2026-05-20
- [x] weather-bot : bot Telegram météo par localisation ✓ 2026-05-20
- [x] quote-bot : bot Telegram citations aléatoires ✓ 2026-05-20

## File d'attente — Outils Web
- [x] color-palette : extracteur de palettes de couleurs depuis une URL ✓ 2026-05-20
- [x] link-checker : vérifie les liens morts d'un site ✓ 2026-05-20
- [x] sitemap-gen : génère un sitemap.xml en crawlant un site ✓ 2026-05-20
- [x] meta-extractor : extrait meta tags (og:, twitter:, etc.) ✓ 2026-05-20
- [x] speed-test : mesure le temps de chargement d'une page ✓ 2026-05-20

## File d'attente — Outils DevOps
- [x] docker-clean : nettoie les images/containers/volumes Docker inutilisés ✓ 2026-05-20
- [x] ssh-manager : gestionnaire de connexions SSH avec fuzzy search ✓ 2026-05-20
- [x] cron-parser : parse et explique les expressions cron ✓ 2026-05-20
- [x] cert-check : vérifie l'expiration des certificats SSL ✓ 2026-05-20
- [x] disk-usage : visualisation disk usage en treemap terminal ✓ 2026-05-20

## File d'attente — Productivité
- [x] pomodoro-cli : timer pomodoro en terminal ✓ 2026-05-20
- [x] todo-cli : gestionnaire de tâches en CLI avec persistance SQLite ✓ 2026-05-20
- [x] expense-tracker : suivi de dépenses en CLI ✓ 2026-05-20
- [x] password-gen : générateur de mots de passe sécurisés ✓ 2026-05-20
- [x] notebook-cli : prise de notes rapide en terminal avec recherche ✓ 2026-05-20

## Vague 2 — CLI Tools
- [x] csv-to-json : convertisseur CSV vers JSON avec détection de séparateur ✓ 2026-05-27
- [x] dup-finder : trouve les fichiers dupliqués par hash ✓ 2026-05-27
- [x] text-stats : statistiques de texte (mots, caractères, lignes, lecture) ✓ 2026-05-27
- [x] json-pretty : formatteur JSON rapide en Go ✓ 2026-05-27
- [x] ip-info : infos IP publiques et géolocalisation ✓ 2026-05-27

## Vague 3 — CLI Tools
- [x] url-parser : parse et analyse des URLs (protocol, host, path, query params) ✓ 2026-05-27
- [x] markdown-lint : vérifie le style markdown et signale les problèmes ✓ 2026-05-27
- [x] dotenv-cli : lit, écrit et valide des fichiers .env ✓ 2026-05-27
- [x] byte-size : affiche la taille de fichiers en format humain (KB, MB, GB) ✓ 2026-05-27
- [x] changelog-gen : génère un CHANGELOG.md depuis les commits git ✓ 2026-05-27

## Vague 4 — CLI Tools
- [x] yaml-to-json : convertisseur YAML vers JSON ✓ 2026-05-27
- [x] semver-bump : incrémente des versions semver dans des fichiers ✓ 2026-05-27
- [x] diff-stats : statistiques de diff entre deux commits/branches ✓ 2026-05-28
- [x] env-diff : compare deux fichiers .env ou environnements ✓ 2026-05-28
- [x] git-aliases : gestionnaire d'alias git (ajout, liste, export) ✓ 2026-05-28

## Vague 5 — CLI Tools
- [x] hex-convert : convertisseur entre hex, décimal, binaire, octal ✓ 2026-05-28
- [x] markdown-preview : prévisualise du markdown en terminal (coloré) ✓ 2026-05-28
- [x] port-check : vérifie si un port est ouvert sur un hôte distant ✓ 2026-05-28
- [x] regex-extract : extrait des données avec regex depuis des fichiers ✓ 2026-05-28
- [x] string-manip : utilitaires de manipulation de chaînes (upper, lower, slug, etc.) ✓ 2026-05-28

## Vague 6 — CLI Tools
- [x] json-merge : fusionne plusieurs fichiers JSON ✓ 2026-05-28
- [x] yaml-lint : vérifie la syntaxe et le style YAML ✓ 2026-05-28
- [x] http-status : vérifie le code HTTP d'une URL ✓ 2026-05-28
- [x] file-type : détecte le type de fichier par magic bytes ✓ 2026-05-28
- [x] time-format : convertit entre formats de date/heure ✓ 2026-05-28

## Vague 7 — CLI Tools
- [x] uuid-gen : générateur de UUID v4 et v7 en Go ✓ 2026-05-28
- [x] json-schema-validate : valide du JSON contre un schema JSON Schema ✓ 2026-05-28
- [x] tree-cli : affiche un arbre de répertoires en terminal ✓ 2026-05-28
- [x] slugify : génère des slugs URL-friendly depuis du texte ✓ 2026-05-28
- [x] json-path : extrait des valeurs JSON via des expressions path ✓ 2026-05-28

## Vague 8 — CLI Tools
- [x] xml-to-json : convertisseur XML vers JSON ✓ 2026-05-28
- [x] json-diff : compare deux fichiers JSON et affiche les différences ✓ 2026-05-28
- [x] env-template : remplit des templates avec des variables d'environnement ✓ 2026-05-28
- [x] netstat-cli : affiche les connexions réseau actives filtrées ✓ 2026-05-28
- [x] chmod-cli : visualise et convertit les permissions Unix (symbolique/octal) ✓ 2026-05-28

## Vague 8 — APIs & Services
- [x] healthcheck-api : API de monitoring qui vérifie une liste d'URLs ✓ 2026-05-28
- [x] webhook-relay : proxy de webhooks avec retry et logs ✓ 2026-05-28
- [x] rate-limiter-api : API de rate limiting en mémoire ✓ 2026-05-28
- [x] cache-api : mini API de cache key/value avec TTL ✓ 2026-05-28
- [x] diff-api : API qui compare deux textes et retourne un diff unifié ✓ 2026-05-28

## Vague 8 — DevOps
- [x] log-rotate : outil de rotation de logs configurable ✓ 2026-05-28
- [x] service-check : vérifie l'état de processus/services système ✓ 2026-05-28
- [x] backup-cli : sauvegarde incrémentale de dossiers avec compression ✓ 2026-05-28
- [x] deploy-hook : exécute des commandes post-deploy avec rollback ✓ 2026-05-28
- [x] systemd-gen : génère des unit files systemd depuis un template ✓ 2026-05-28

## Vague 8 — Productivité
- [x] calc-cli : calculatrice scientifique en CLI ✓ 2026-05-28
- [x] unit-convert : convertisseur d'unités (longueur, poids, température, etc.) ✓ 2026-05-28
- [x] calendar-cli : affiche un calendrier mensuel en terminal ✓ 2026-05-28
- [x] habit-tracker : suivi d'habitudes en CLI avec streaks ✓ 2026-05-28
- [x] quick-survey : crée et analyse des sondages rapides en CLI ✓ 2026-05-28

## Vague 9 — CLI Tools
- [x] snippet-cli : gestionnaire de snippets de commandes shell (sauvegarde, recherche, exécution) ✓ 2026-05-28
- [x] tcp-dump-lite : capture et analyse simplifiée de trafic réseau ✓ 2026-05-28
- [x] process-tree : affiche l'arbre des processus en terminal ✓ 2026-05-28
- [x] git-log-visual : visualisation du historique git en ASCII art ✓ 2026-05-28
- [x] env-export : exporte les variables d'un .env dans le shell courant ✓ 2026-05-28

## Vague 9 — APIs & Services
- [x] validate-api : API de validation d'email, téléphone, CPF/CNPJ ✓ 2026-05-28
- [x] notify-api : API de notifications multi-canaux (email, webhook, Slack) ✓ 2026-05-28
- [x] queue-api : API de file d'attente simple avec priorité ✓ 2026-05-28
- [x] auth-api : API d'authentification JWT avec refresh tokens ✓ 2026-05-28
- [x] file-upload-api : API d'upload de fichiers avec validation ✓ 2026-05-28

## Vague 9 — DevOps
- [x] nginx-gen : génère des configs Nginx depuis un fichier YAML ✓ 2026-05-28
- [x] log-analyzer : analyse de logs Apache/Nginx avec stats ✓ 2026-05-28
- [x] vm-stats : collecte et affiche les stats système (CPU, RAM, disk, net) ✓ 2026-05-28
- [x] docker-logs-cli : agrège et filtre les logs de containers Docker ✓ 2026-05-28
- [x] release-notes : génère des release notes depuis git tags ✓ 2026-05-28

## Vague 9 — Productivité
- [x] journal-cli : journal quotidien en CLI avec tags et recherche ✓ 2026-05-28
- [x] goal-tracker : suivi d'objectifs avec sous-tâches et progression ✓ 2026-05-28
- [x] flashcards-cli : système de répétition espacée en terminal ✓ 2026-05-28
- [x] budget-cli : gestion budgétaire mensuelle en CLI ✓ 2026-05-28
- [x] meeting-notes : template et gestion de notes de réunion ✓ 2026-05-28

## Vague 10 — CLI Tools
- [x] jq-lite : lightweight JSON query tool (dot notation, arrays, filtering) ✓ 2026-05-28
- [x] env-export : exporte les variables d'un .env dans le shell courant ✓ 2026-05-28
- [x] csv-filter : filtre et transforme des fichiers CSV en CLI ✓ 2026-05-28

## Vague 11 — CLI Tools
- [x] jsonl-tools : outils pour lire, filtrer et transformer des fichiers JSONL ✓ 2026-05-28
- [x] env-merge : fusionne plusieurs fichiers .env avec priorité configurable ✓ 2026-05-28
- [x] git-blame-stats : stats de code ownership par fichier/blame ✓ 2026-05-28
- [x] httpie-lite : client HTTP léger en CLI (GET, POST, PUT, DELETE) ✓ 2026-05-28
- [x] markdown-toc : génère une table des matières pour des fichiers Markdown ✓ 2026-05-28

## Vague 11 — APIs & Services
- [x] proxy-api : proxy HTTP configurable avec rewriting d'URL ✓ 2026-05-28
- [x] cron-api : API pour gérer des tâches planifiées simples ✓ 2026-05-28
- [x] template-api : API de rendu de templates (text, JSON, YAML) ✓ 2026-05-28
- [x] webhook-monitor : monitoring de webhooks avec alertes ✓ 2026-05-28
- [x] redirect-api : API de redirection avec stats et A/B testing ✓ 2026-05-28

## Vague 11 — DevOps
- [x] log-stream : agrège et stream des logs de multiples sources ✓ 2026-05-28
- [x] config-validator : valide des fichiers de config (JSON, YAML, TOML) ✓ 2026-05-28
- [x] health-dashboard : dashboard de santé d'applications en terminal ✓ 2026-05-28
- [x] deploy-cli : outil de déploiement simple avec rollback ✓ 2026-05-28
- [x] secret-manager : gestionnaire de secrets local chiffré ✓ 2026-05-28

## Vague 11 — Productivité
- [x] cli-dictionary : dictionnaire et thésaurus en CLI ✓ 2026-05-28
- [x] time-tracker : suivi de temps par projet/tâche en CLI ✓ 2026-05-28
- [x] checklist-cli : gestionnaire de checklists réutilisables ✓ 2026-05-28
- [x] cli-calendar : agenda personnel en CLI avec rappels ✓ 2026-05-28
- [x] reading-list : gestionnaire de liens/articles à lire ✓ 2026-05-28

## Vague 12 — CLI Tools
- [x] markdown-pdf : convertit des fichiers Markdown en PDF via CLI ✓ 2026-05-28
- [x] jsonql : requête JSON avec syntaxe simplifiée (dot notation, filtres) ✓ 2026-05-28
- [x] env-schema : valide des fichiers .env contre un schema YAML ✓ 2026-05-28
- [x] git-branch-cleanup : nettoie les branches locales/fondues ✓ 2026-05-28
- [x] httpbin-cli : serveur HTTP de test local (comme httpbin mais en CLI) ✓ 2026-05-28

## Vague 12 — APIs & Services
- [x] echo-api : API de test qui retourne les données reçues (headers, body, method) ✓ 2026-05-28
- [x] delay-api : API qui simule des délais (pour tests de timeout) ✓ 2026-05-28
- [x] counter-api : API de compteur incrémental avec reset ✓ 2026-05-28
- [x] echo-webhook : endpoint qui logue et retourne les webhooks reçus ✓ 2026-05-28
- [x] mock-api : API de mock configurable via JSON ✓ 2026-05-28

## Vague 12 — DevOps
- [x] log-filter : filtre des logs par niveau, pattern, timeframe ✓ 2026-05-28
- [x] port-forward : gestionnaire de port forwarding simplifié ✓ 2026-05-28
- [x] docker-ps-extended : docker ps avec infos supplémentaires (disk, env) ✓ 2026-05-28
- [x] systemd-once : exécute une commande une fois avec logging systemd ✓ 2026-05-28
- [x] env-encrypt : chiffrement/déchiffrement de variables d'env ✓ 2026-05-28

## Vague 12 — Productivité
- [x] cli-pomodoro : timer pomodoro avec notifications sonores ✓ 2026-05-28
- [x] expense-report : génère des rapports de dépenses depuis un CSV ✓ 2026-05-28
- [x] cli-wishlist : liste de souhaits avec priorités et liens ✓ 2026-05-28
- [x] quote-collect : collectionne et organise des citations ✓ 2026-05-28

## Vague 13 — CLI Tools
- [x] ini-parser : lit et écrit des fichiers de config INI ✓ 2026-05-28
- [x] json-to-csv : convertisseur JSON vers CSV ✓ 2026-05-28
- [x] net-latency : mesure la latence réseau vers des hôtes ✓ 2026-05-28
- [x] git-repo-info : affiche des infos sur un repo git (taille, commits, branches) ✓ 2026-05-28
- [x] env-validate : valide un .env contre une liste de clés requises ✓ 2026-05-28

## Vague 13 — APIs & Services
- [x] ip-api : API qui retourne des infos sur l'IP du client ✓ 2026-05-28
- [x] random-api : API de génération de nombres/textes aléatoires ✓ 2026-05-28
- [x] ping-api : API de ping HTTP vers des URLs configurées ✓ 2026-05-28
- [x] log-api : API de journalisation simple avec filtrage ✓ 2026-05-28
- [x] status-api : API de statut de service configurable ✓ 2026-05-28

## Vague 13 — DevOps
- [x] git-hooks : installe des hooks git pré-configurés ✓ 2026-05-28
- [x] log-compress : compresse et archive des fichiers de logs ✓ 2026-05-28
- [x] env-backup : sauvegarde et restaure des environnements .env ✓ 2026-05-28
- [x] docker-stats-cli : affiche les stats Docker en temps réel ✓ 2026-05-28
- [x] uptime-check : vérifie l'uptime de services et affiche un rapport ✓ 2026-05-28

## Vague 13 — Productivité
- [x] cli-biography : notes biographiques rapides sur des personnes ✓ 2026-05-28
- [x] project-log : journal de projet avec milestones ✓ 2026-05-28
- [x] idea-box : boîte à idées avec tags et vote ✓ 2026-05-28
- [x] reading-notes : notes de lecture par livre/chapitre ✓ 2026-05-28
- [x] skill-tracker : suivi de compétences avec niveaux et progression ✓ 2026-05-28
## Vague 24 — CLI Tools
- [x] file-rename-batch : renommage massif de fichiers (regex, case, remplacement) ✓ 2026-05-28
- [x] json-compact : minifie du JSON en une seule ligne ✓ 2026-05-28
- [x] csv-header : affiche, ajoute, supprime ou renomme les headers CSV ✓ 2026-05-28
- [x] ini-to-json : convertisseur INI vers JSON ✓ 2026-05-28
- [x] toml-to-csv : convertisseur TOML vers CSV ✓ 2026-05-28
- [x] json-to-toml : convertisseur JSON vers TOML ✓ 2026-05-28

## Vague 25 — CLI Tools
- [x] csv-sort : trie des fichiers CSV par une ou plusieurs colonnes ✓ 2026-05-28
- [x] csv-merge : fusionne plusieurs fichiers CSV avec la même structure ✓ 2026-05-28
- [x] csv-count : compte les lignes dans des fichiers CSV ✓ 2026-05-28
- [x] json-unique : supprime les doublons dans un tableau JSON ✓ 2026-05-28
- Dernière mise à jour: 2026-05-28
- Priorité: CLI tools d'abord (rapides à coder)
- Chaque outil: max 1-2h de dev

## Vague 14 — CLI Tools
- [x] github-trending-cli : affiche les repos tendance GitHub en terminal ✓ 2026-05-28
- [x] markdown-table : convertit CSV/TSV/data en tableaux Markdown ✓ 2026-05-28
- [x] json-flatten : aplatit du JSON imbriqué en paires clé-valeur ✓ 2026-05-28
- [x] ssh-key-gen : générateur de clés SSH avec options ✓ 2026-05-28
- [x] git-ignore : génère des .gitignore depuis des templates ✓ 2026-05-28

## Vague 15 — CLI Tools
- [x] markdown-checklist : gestionnaire de checklists dans des fichiers Markdown ✓ 2026-05-28
- [x] color-convert : convertisseur de couleurs (hex, rgb, hsl, cmyk) ✓ 2026-05-28
- [x] http-file-server : serveur HTTP statique simple en Go ✓ 2026-05-28
- [x] pdf-text-extract : extrait le texte de fichiers PDF ✓ 2026-05-28
- [x] git-merge-conflict-finder : détecte les conflits de merge dans un repo git ✓ 2026-05-28

## Vague 16 — CLI Tools
- [x] docker-compose-validator : valide la syntaxe des fichiers docker-compose.yml ✓ 2026-05-28
- [x] git-submodule-list : liste et affiche les infos des submodules git ✓ 2026-05-28
- [x] env-to-json : convertit un fichier .env en JSON ✓ 2026-05-28

## Vague 17 — CLI Tools
- [x] json-sort : trie les clés d'un fichier JSON de manière récursive ✓ 2026-05-28
- [x] toml-to-json : convertisseur TOML vers JSON ✓ 2026-05-28
- [x] file-count : compte les fichiers par type/extension dans un dossier ✓ 2026-05-28
- [x] git-last-commit : affiche les infos du dernier commit de manière formatée ✓ 2026-05-28
- [x] path-join : outil de manipulation de chemins de fichiers (join, dirname, basename, etc.) ✓ 2026-05-28

## Vague 17 — APIs & Services
- [x] cors-proxy : proxy CORS simple pour contourner les restrictions ✓ 2026-05-28
- [x] timestamp-api : API qui retourne le timestamp actuel dans différents formats ✓ 2026-05-28
- [x] uuid-api : API de génération de UUIDs (v4, v7) ✓ 2026-05-28
- [x] hash-api : API de hachage (sha256, md5, blake2) de texte ou fichiers ✓ 2026-05-28
- [x] paginate-api : API de pagination pour données JSON ✓ 2026-05-28

## Vague 17 — DevOps
- [x] log-level-filter : filtre des logs par niveau de sévérité ✓ 2026-05-28
- [x] env-docs : génère la documentation d'un .env en Markdown ✓ 2026-05-28
- [x] git-tag-manager : gestionnaire de tags git (création, liste, suppression) ✓ 2026-05-28
- [x] docker-image-size : affiche la taille des images Docker triées ✓ 2026-05-28
- [x] crontab-lint : vérifie la syntaxe et signale les problèmes dans un crontab ✓ 2026-05-28

## Vague 17 — Productivité
- [x] cli-contacts : carnet d'adresses simple en CLI ✓ 2026-05-28
- [x] recipe-cli : gestionnaire de recettes de cuisine en CLI ✓ 2026-05-28
- [x] cli-bingo : générateur de cartes Bingo en terminal ✓ 2026-05-28
- [x] word-counter : compteur de mots avec historique par fichier ✓ 2026-05-28
- [x] cli-trivia : quiz de culture générale en terminal ✓ 2026-05-28

## Vague 18 — CLI Tools
- [x] json-minify : minifie des fichiers JSON (retire espaces, newlines) ✓ 2026-05-28
- [x] csv-split : divise un gros CSV en plusieurs fichiers plus petits ✓ 2026-05-28
- [x] git-stash-cleanup : liste et nettoie les stashes git obsolètes ✓ 2026-05-28
- [x] env-to-toml : convertit un fichier .env en TOML ✓ 2026-05-28
- [x] line-count : compte les lignes de code par langage dans un projet ✓ 2026-05-28

## Vague 18 — APIs & Services
- [x] captcha-api : API de génération de CAPTCHAs simples (math, texte) ✓ 2026-05-28
- [x] seed-api : API de seed/random reproductible avec seed ✓ 2026-05-28
- [x] batch-api : API qui exécute des requêtes HTTP en batch ✓ 2026-05-28
- [x] transform-api : API de transformation de données (map, filter, sort) ✓ 2026-05-28
- [x] merge-api : API de merge de fichiers JSON/YAML ✓ 2026-05-28

## Vague 18 — DevOps
- [x] log-sampler : échantillonne des logs pour analyse (1%, 10%, etc.) ✓ 2026-05-28
- [x] env-to-yaml : convertit un fichier .env en YAML ✓ 2026-05-28
- [x] git-reflog-clean : nettoie le reflog git pour réduire la taille ✓ 2026-05-28
- [x] docker-rename : renomme des containers Docker en batch ✓ 2026-05-28
- [x] symlink-manager : gestionnaire de symlinks (création, vérification, nettoyage) ✓ 2026-05-28

## Vague 18 — Productivité
- [x] cli-bookmarks : gestionnaire de favoris/shortcuts en CLI ✓ 2026-05-28
- [x] cli-dice : lanceur de dés (d4, d6, d8, d10, d12, d20) en terminal ✓ 2026-05-28
- [x] cli-cookbook : recettes de cuisine rapides en CLI ✓ 2026-05-28
- [x] cli-meditation : timer de méditation avec instructions en terminal ✓ 2026-05-28
- [x] cli-joke : générateur de blagues en CLI ✓ 2026-05-28

## Vague 19 — CLI Tools
- [x] json-rename-keys : renomme les clés JSON de manière récursive ✓ 2026-05-28
- [x] csv-join : jointure de deux fichiers CSV (inner, left, right) ✓ 2026-05-28
- [x] git-file-tree : affiche les fichiers modifiés dans un commit en arbre ✓ 2026-05-28
- [x] csv-statistics : statistiques basiques sur chaque colonne d'un CSV ✓ 2026-05-28

## Vague 20 — CLI Tools
- [x] dedup-lines : supprime les lignes dupliquées dans des fichiers ✓ 2026-05-28
- [x] json-subset : extrait un sous-ensemble de clés depuis un JSON ✓ 2026-05-28
- [x] git-history : affiche l'historique git formaté avec stats ✓ 2026-05-28
- [x] env-sanitize : sanitize un fichier .env (retire valeurs sensibles) ✓ 2026-05-28
- [x] dir-compare : compare le contenu de deux dossiers ✓ 2026-05-28

## Vague 21 — CLI Tools
- [x] glob-match : teste si des chemins matchent un pattern glob ✓ 2026-05-28
- [x] json-group-by : regroupe des objets JSON par une clé ✓ 2026-05-28
- [x] csv-append : ajoute des lignes à un fichier CSV existant ✓ 2026-05-28
- [x] git-author-stats : stats par auteur (commits, fichiers touchés, LOC) ✓ 2026-05-28
- [x] env-to-ini : convertit un fichier .env en format INI ✓ 2026-05-28

## Vague 22 — CLI Tools
- [x] json-flatten-keys : aplatit des clés JSON imbriquées en une seule couche ✓ 2026-05-28
- [x] csv-dedup : supprime les lignes dupliquées dans un CSV ✓ 2026-05-28
- [x] git-branch-age : affiche l'âge de chaque branche git ✓ 2026-05-28
- [x] json-to-yaml : convertisseur JSON vers YAML ✓ 2026-05-28
- [x] path-exists : vérifie l'existence de chemins fichiers/dossiers ✓ 2026-05-28

## Vague 23 — CLI Tools
- [x] yaml-to-csv : convertisseur YAML vers CSV ✓ 2026-05-28
- [x] json-filter : filtre des objets JSON par conditions ✓ 2026-05-28
- [x] csv-transpose : transpose les lignes et colonnes d'un CSV ✓ 2026-05-28
- [x] git-commit-count : compte les commits par période (jour, semaine, mois) ✓ 2026-05-28
- [x] env-to-markdown : génère un tableau Markdown depuis un .env ✓ 2026-05-28

## Vague 27 — CLI Tools
- [x] stopwatch-cli : chronomètre terminal avec laps et historique ✓ 2026-05-28
- [x] json-to-ini : convertisseur JSON vers INI ✓ 2026-05-28
- [x] csv-pivot : pivote un CSV (lignes↔colonnes agrégées) ✓ 2026-05-28
- [x] git-remote-info : affiche les URLs et statuts des remotes git ✓ 2026-05-28
- [x] env-to-csv : convertit un fichier .env en CSV ✓ 2026-05-28
- [x] json-to-xml : convertisseur JSON vers XML ✓ 2026-05-28
- [x] csv-validate : valide la structure d'un CSV (colonnes, types, nulls) ✓ 2026-05-28
- [x] git-tag-info : affiche les infos détaillées d'un tag git ✓ 2026-05-28
- [x] toml-to-ini : convertisseur TOML vers INI ✓ 2026-05-28

## Vague 31 — CLI Tools
- [x] json-to-sql : génère des requêtes INSERT SQL depuis un JSON ✓ 2026-05-28
- [x] csv-diff : compare deux fichiers CSV et affiche les différences ✓ 2026-05-28
- [x] markdown-link-check : vérifie les liens dans des fichiers Markdown ✓ 2026-05-28
- [x] env-to-xml : convertit un fichier .env en XML ✓ 2026-05-28
- [x] git-commit-graph : affiche un graphique ASCII du historique de commits ✓ 2026-05-28

## Vague 32 — CLI Tools
- [x] xml-to-csv : convertisseur XML vers CSV ✓ 2026-05-28
- [x] json-to-ini : convertisseur JSON vers INI ✓ 2026-05-28
- [x] csv-format : formate et normalise des fichiers CSV ✓ 2026-05-28
- [x] git-changelog : génère un changelog depuis les tags git ✓ 2026-05-28
- [x] env-to-yaml : convertit un fichier .env en YAML ✓ 2026-05-28

## Vague 33 — CLI Tools
- [x] toml-to-yaml : convertisseur TOML vers YAML ✓ 2026-05-28
- [x] json-to-markdown-table : convertit un tableau JSON en tableau Markdown ✓ 2026-05-28
- [x] csv-to-jsonl : convertisseur CSV vers JSONL ✓ 2026-05-28
- [x] git-file-stats : stats par fichier (LOC, commits, auteurs) ✓ 2026-05-28
- [x] env-to-ini : convertit un fichier .env en format INI ✓ 2026-05-28

## Vague 34 — CLI Tools
- [x] ini-to-csv : convertisseur INI vers CSV ✓ 2026-05-28
- [x] json-to-env : convertisseur JSON vers .env ✓ 2026-05-28
- [x] csv-to-yaml : convertisseur CSV vers YAML ✓ 2026-05-28
- [x] git-stash-list : liste les stashes git avec details ✓ 2026-05-28
- [x] toml-validate : valide la syntaxe d'un fichier TOML ✓ 2026-05-28

## Vague 35 — CLI Tools
- [x] ini-to-yaml : convertisseur INI vers YAML ✓ 2026-05-28
- [x] json-to-csv-nested : convertisseur JSON imbriqué vers CSV (aplatissement) ✓ 2026-05-28
- [x] git-log-summary : résumé du historique git (commits, fichiers, auteurs) ✓ 2026-05-28
- [x] xml-validate : valide la syntaxe XML d'un fichier ✓ 2026-05-28
- [x] env-to-tsv : convertit un fichier .env en TSV ✓ 2026-05-28

## Vague 36 — CLI Tools
- [x] csv-to-sqlite : importe des fichiers CSV dans une base SQLite ✓ 2026-05-28
- [x] sqlite-query : interroge des bases SQLite en CLI (table, CSV, JSON) ✓ 2026-05-28
- [x] markdown-to-html : convertit du Markdown en HTML (GFM, standalone) ✓ 2026-05-28

## Vague 37 — CLI Tools
- [x] markdown-to-text : supprime le formatage markdown et affiche du texte brut ✓ 2026-05-28
- [x] file-integrity-check : calcule et vérifie les checksums de fichiers (sha256, md5, sha512) ✓ 2026-05-28
- [x] kv-store : stockage clé-valeur simple avec persistance SQLite ✓ 2026-05-28
- [x] json-to-sqlite : importe des données JSON dans des bases SQLite ✓ 2026-05-28
- [x] http-load-test : outil de load testing HTTP avec concurrence et rapports de latence ✓ 2026-05-28

## Vague 38 — CLI Tools
- [x] hex-view : visualiseur hexadécimal pour fichiers binaires ✓ 2026-05-28
- [x] html-minify : minifie du code HTML ✓ 2026-05-28
- [x] json-pointer : extrait des valeurs JSON via JSON Pointer (RFC 6901) ✓ 2026-05-28
- [x] git-find-repo : trouve les repos git dans un arbre de répertoires ✓ 2026-05-28
- [x] file-checksum : calcule les checksums de fichiers (sha256, md5, sha1, crc32) ✓ 2026-05-28

## Vague 39 — CLI Tools
- [x] json-to-tsv : convertisseur JSON vers TSV ✓ 2026-05-28
- [x] csv-to-markdown : convertisseur CSV vers tableau Markdown ✓ 2026-05-28
- [x] git-commit-msg : affiche les messages de commits récents formatés ✓ 2026-05-28
- [x] env-list-keys : liste uniquement les clés d'un fichier .env ✓ 2026-05-28
- [x] file-line-count : compte les lignes de fichiers en batch ✓ 2026-05-28

## Vague 40 — CLI Tools
- [x] csv-to-ini : convertisseur CSV vers INI avec regroupement par section ✓ 2026-05-28
- [x] yaml-flatten : aplatit du YAML imbriqué en paires clé-valeur ✓ 2026-05-28
- [x] git-tag-list : liste les tags git avec détails (date, type, auteur, message) ✓ 2026-05-28
- [x] json-to-html : convertit des tableaux JSON en tableaux HTML ✓ 2026-05-28
- [x] yaml-validate : valide la syntaxe YAML avec mode strict ✓ 2026-05-28

## Vague 41 — CLI Tools
- [x] file-watcher : surveille les fichiers et lance des commandes automatiquement (style nodemon) ✓ 2026-05-28
- [x] color-convert : convertisseur de couleurs (hex, RGB, HSL, CMYK) ✓ 2026-05-28
- [x] markdown-toc : génère une table des matières pour des fichiers Markdown ✓ 2026-05-28

## Vague 42 — CLI Tools
- [x] go-struct-gen : génère des structs Go depuis du JSON ✓ 2026-05-28
- [x] jsonl-to-csv : convertisseur JSONL vers CSV ✓ 2026-05-28
- [x] sql-from-csv : génère des requêtes SQL (CREATE + INSERT) depuis un CSV ✓ 2026-05-28
- [x] toml-to-env : convertisseur TOML vers .env ✓ 2026-05-28

## Vague 44 — CLI Tools
- [x] tsv-to-json : convertisseur TSV vers JSON ✓ 2026-05-29
- [x] yaml-to-sql : génère des requêtes SQL (CREATE + INSERT) depuis du YAML ✓ 2026-05-29
- [x] csv-to-env : convertisseur CSV vers .env ✓ 2026-05-29
- [x] git-untracked : liste les fichiers non suivis dans un repo git ✓ 2026-05-29
- [x] ini-validate : valide la syntaxe d'un fichier INI ✓ 2026-05-29

## Vague 43 — CLI Tools
- [x] go-struct-gen : génère des structs Go depuis du JSON ✓ 2026-05-28
- [x] json-to-go : génère des structs Go depuis du JSON (avec type hints) ✓ 2026-05-28
- [x] yaml-to-go : génère des structs Go depuis du YAML ✓ 2026-05-28
- [x] sql-migrate : système de migrations SQL simple (up/down) ✓ 2026-05-28
- [x] go-mod-tidy-check : vérifie si go.mod est à jour ✓ 2026-05-28

## Vague 44 — CLI Tools
- [x] json-to-rust : génère des structs Rust depuis du JSON ✓ 2026-05-29
- [x] yaml-to-sql : génère des requêtes SQL (CREATE + INSERT) depuis YAML ✓ 2026-05-29
- [x] git-commit-lint : vérifie le format des messages de commit ✓ 2026-05-29
- [x] env-mask : masque les valeurs sensibles dans un .env ✓ 2026-05-29

## Vague 44 — CLI Tools
- [x] go-vet-check : lance go vet et formate le rapport ✓ 2026-05-28
- [x] json-to-toml-nested : convertisseur JSON profond vers TOML ✓ 2026-05-28
- [x] csv-to-sql : génère CREATE TABLE + INSERT SQL depuis un CSV ✓ 2026-05-28
- [x] git-untracked : liste les fichiers untracked avec tailles ✓ 2026-05-28
- [x] env-to-jsonl : convertit un fichier .env en JSONL ✓ 2026-05-28

## Vague 45 — CLI Tools
- [x] json-to-sqlite : importe des données JSON dans une base SQLite ✓ 2026-05-28
- [x] csv-column-stats : statistiques par colonne d'un CSV (min, max, mean, count) ✓ 2026-05-28
- [x] git-commit-frequency : histogramme des commits par heure/jour de la semaine ✓ 2026-05-28
- [x] env-sanitize : remplace les valeurs sensibles d'un .env par des placeholders ✓ 2026-05-28
- [x] toml-to-markdown : convertit un fichier TOML en tableau Markdown ✓ 2026-05-28

## Vague 46 — CLI Tools
- [x] json-to-yaml : convertisseur JSON vers YAML ✓ 2026-05-28
- [x] csv-to-toml : convertisseur CSV vers TOML ✓ 2026-05-28
- [x] git-merge-base : trouve le dernier ancêtre commun entre deux branches ✓ 2026-05-28
- [x] env-to-sql : génère des requêtes SQL INSERT depuis un fichier .env ✓ 2026-05-28
- [x] markdown-lint : vérifie le style markdown et signale les problèmes ✓ 2026-05-28

## Vague 47 — CLI Tools
- [x] json-to-tsv : convertisseur JSON vers TSV ✓ 2026-05-28
- [x] csv-to-sql : génère des requêtes SQL depuis un CSV ✓ 2026-05-28
- [x] git-unpushed : liste les commits non pushés sur le remote ✓ 2026-05-28
- [x] env-to-yaml : convertit un fichier .env en YAML ✓ 2026-05-28
- [x] xml-to-json : convertisseur XML vers JSON ✓ 2026-05-28

## Vague 48 — CLI Tools
- [x] sql-from-json : génère des requêtes SQL (CREATE + INSERT) depuis du JSON ✓ 2026-05-28
- [x] csv-to-xml : convertisseur CSV vers XML ✓ 2026-05-28
- [x] git-branch-delete : supprime les branches locales et/ou remote fusionnées ✓ 2026-05-28
- [x] env-to-sqlite : importe un fichier .env dans une base SQLite ✓ 2026-05-28
- [x] json-to-sql-schema : génère un schéma SQL (CREATE TABLE) depuis du JSON ✓ 2026-05-28

## Vague 49 — CLI Tools
- [x] git-commit-lint : linter de messages de commits (Conventional Commits) ✓ 2026-05-28
- [x] env-to-toml : convertit un fichier .env en TOML ✓ 2026-05-28
- [x] csv-to-json-nested : convertisseur CSV vers JSON imbriqué ✓ 2026-05-28
- [x] git-rebase-helper : assistant de rebase interactif en CLI ✓ 2026-05-28
- [x] json-to-protobuf : génère des fichiers .proto depuis du JSON ✓ 2026-05-28

## Vague 50 — CLI Tools
- [x] yaml-to-sql : génère des requêtes SQL depuis du YAML ✓ 2026-05-28
- [x] csv-to-env : convertisseur CSV vers .env ✓ 2026-05-28
- [x] git-branch-rename : renomme des branches git en batch ✓ 2026-05-28
- [x] json-to-tsv-nested : convertisseur JSON imbriqué vers TSV ✓ 2026-05-28
- [x] toml-to-sql : génère des requêtes SQL depuis un TOML ✓ 2026-05-28

## Vague 51 — CLI Tools
- [x] go-fmt-check : vérifie si les fichiers Go sont bien formatés avec un rapport clair ✓ 2026-05-28
- [x] markdown-count : compte les éléments dans des fichiers Markdown (headings, links, images, etc.) ✓ 2026-05-28
- [x] git-ignored : liste les fichiers ignorés par .gitignore ✓ 2026-05-28
- [x] tsv-to-json : convertisseur TSV vers JSON ✓ 2026-05-28
- [x] env-to-protobuf : convertit un fichier .env en définition Protocol Buffers ✓ 2026-05-28

## Vague 52 — CLI Tools
- [x] go-test-coverage : affiche la couverture de tests Go de manière formatée ✓ 2026-05-28
- [x] go-imports-check : détecte les imports inutilisés dans un projet Go ✓ 2026-05-28
- [x] json-to-sqlite : importe des données JSON dans une base SQLite ✓ 2026-05-28
- [x] csv-to-sqlite : importe des fichiers CSV dans une base SQLite ✓ 2026-05-28
- [x] git-worktree-manager : gestionnaire de worktrees git en CLI ✓ 2026-05-28

## Vague 53 — CLI Tools
- [x] http-range : vérifie le support des requêtes HTTP Range (contenu partiel) ✓ 2026-05-28
- [x] dns-overview : affiche un aperçu complet des enregistrements DNS d'un domaine ✓ 2026-05-28

## Vague 54 — CLI Tools
- [x] csv-dedup-cli : supprime les doublons dans un CSV par colonnes spécifiées ✓ 2026-05-28
- [x] git-reflog-analyzer : analyse le reflog git avec stats (actions, branches, timeline) ✓ 2026-05-28

## Vague 55 — CLI Tools
- [x] json-to-typescript : convertisseur JSON vers types TypeScript ✓ 2026-05-28
- [x] csv-to-protobuf : convertisseur CSV vers définitions Protocol Buffers ✓ 2026-05-28
- [x] yaml-to-protobuf : convertisseur YAML vers définitions Protocol Buffers ✓ 2026-05-28

## Vague 56 — CLI Tools
- [x] go-exported-api : liste la surface API exportée d'un package Go ✓ 2026-05-28
- [x] csv-to-html : convertisseur CSV vers tableau HTML ✓ 2026-05-28
- [x] git-first-commit : affiche le premier commit d'un repo git ✓ 2026-05-28
- [x] http-method-check : vérifie les méthodes HTTP supportées par une URL ✓ 2026-05-28
- [x] markdown-to-json : extrait des données structurées depuis du Markdown ✓ 2026-05-28

## Vague 57 — CLI Tools
- [x] go-staticcheck : lance staticcheck et formate le rapport ✓ 2026-05-28
- [x] json-to-graphql : génère un schema GraphQL depuis du JSON ✓ 2026-05-28
- [x] csv-to-sql-schema : génère un CREATE TABLE SQL depuis un CSV ✓ 2026-05-28
- [x] git-submodule-status : affiche le statut des submodules git ✓ 2026-05-28
- [x] env-to-graphql : convertit un fichier .env en schema GraphQL ✓ 2026-05-28

## Vague 58 — CLI Tools
- [x] go-deadcode-check : detecte le code mort dans un projet Go ✓ 2026-05-28
- [x] json-to-rust : genere des structs Rust depuis du JSON ✓ 2026-05-28
- [x] csv-to-graphviz : genere un graphe Graphviz depuis un CSV (nœuds/arêtes) ✓ 2026-05-28
- [x] git-lfs-status : affiche le statut des fichiers suivis par Git LFS ✓ 2026-05-28
- [x] env-to-openapi : convertit un fichier .env en spec OpenAPI/Swagger ✓ 2026-05-28

## Vague 59 — CLI Tools
- [x] go-cycle-check : detecte les import cycles dans un projet Go ✓ 2026-05-28
- [x] json-to-java : genere des classes Java depuis du JSON ✓ 2026-05-28
- [x] csv-to-mermaid : genere un diagramme Mermaid depuis un CSV ✓ 2026-05-28
- [x] git-squash-helper : assistant de squash de commits interactif ✓ 2026-05-28
- [x] env-to-jsonschema : convertit un fichier .env en JSON Schema ✓ 2026-05-28

## Vague 60 — CLI Tools
- [x] csv-to-latex : convertisseur CSV vers tableau LaTeX (tabular, longtable, tabularx) ✓ 2026-05-28
- [x] json-to-python : genere des dataclasses Python depuis du JSON ✓ 2026-05-28
- [x] git-cherry-pick-helper : assistant de cherry-pick interactif ✓ 2026-05-28
- [x] env-to-docker-compose : convertit un .env en docker-compose.yml ✓ 2026-05-28
- [x] go-golint-check : lance golint/staticcheck avec rapport formate ✓ 2026-05-28

## Vague 61 — CLI Tools
- [x] toml-diff : compare deux fichiers TOML et affiche les différences ✓ 2026-05-28
- [x] json-to-terraform : genere des blocs Terraform HCL depuis du JSON ✓ 2026-05-28
- [x] yaml-to-terraform : genere des blocs Terraform HCL depuis du YAML ✓ 2026-05-28
- [x] toml-to-terraform : genere des blocs Terraform HCL depuis du TOML ✓ 2026-05-28
- [x] git-merge-stats : statistiques sur les merge commits d'un repo git ✓ 2026-05-28
- [x] env-to-hcl : convertit un fichier .env en variables Terraform HCL ✓ 2026-05-28

## Vague 62 — CLI Tools
- [x] csv-to-chart : genere des graphiques en barres terminaux depuis un CSV ✓ 2026-05-28
- [x] env-to-kubernetes : convertit un .env en ConfigMap/Secret Kubernetes ✓ 2026-05-28

## Vague 70 — CLI Tools
- [x] http-response-diff : compare les réponses HTTP de deux URLs (status, headers, body) ✓ 2026-05-28

## Vague 71 — CLI Tools
- [x] http-websocket-cli : client WebSocket simple pour tester des endpoints WS/WSS ✓ 2026-05-28
- [x] json-to-avro : genere un schema Avro depuis du JSON ✓ 2026-05-28
- [x] csv-to-delta : exporte un CSV en format Delta Lake ✓ 2026-05-28
- [x] git-branch-coverage : affiche la couverture de tests par branche ✓ 2026-05-28
## Vague 63 — CLI Tools
- [x] robots-checker : fetch et parse les fichiers robots.txt de sites web ✓ 2026-05-28
- [x] cookie-parser : parse les headers Set-Cookie HTTP et strings de cookies ✓ 2026-05-28
- [x] git-churn : affiche le churn de fichiers (lignes ajoutées/supprimées) dans un repo git ✓ 2026-05-28
- [x] env-to-helm : convertit des fichiers .env en format Helm values.yaml ✓ 2026-05-28
- [x] http-retry : client HTTP avec retry automatique et exponential backoff ✓ 2026-05-28

## Vague 64 — CLI Tools
- [x] json-to-parquet : convertisseur JSON/JSONL vers Apache Parquet ✓ 2026-05-28
- [x] csv-to-parquet : convertisseur CSV vers Apache Parquet ✓ 2026-05-28
- [x] markdown-to-csv : extrait des tableaux Markdown et exporte en CSV ✓ 2026-05-28
- [x] yaml-to-sql-schema : génère des CREATE TABLE SQL depuis du YAML ✓ 2026-05-28
- [x] git-commit-template : gestionnaire de templates de messages git ✓ 2026-05-28

## Vague 65 — CLI Tools
- [x] parquet-to-json : convertisseur Parquet vers JSON/JSONL ✓ 2026-05-28
- [x] tsv-to-json : convertisseur TSV vers JSON ✓ 2026-05-28
- [x] ini-to-toml : convertisseur INI vers TOML ✓ 2026-05-28
- [x] xml-to-csv : convertisseur XML vers CSV ✓ 2026-05-28
- [x] git-merge-driver : configure et gère des merge drivers git ✓ 2026-05-28

## Vague 66 — CLI Tools
- [x] ssl-cert-info : affiche les infos d'un certificat SSL/TLS (host, URL, fichier PEM) ✓ 2026-05-28
- [x] dns-propagation : vérifie la propagation DNS sur 7 serveurs publics ✓ 2026-05-28
- [x] whois-lookup : requête WHOIS pour des domaines avec parsing structuré ✓ 2026-05-28
- [x] http-retry : client HTTP avec retry automatique et exponential backoff ✓ 2026-05-28

## Vague 67 — CLI Tools
- [x] conn-monitor : affiche les connexions TCP actives depuis /proc/net/tcp ✓ 2026-05-28
- [x] http-redirect-trace : trace les chaînes de redirection HTTP avec timing ✓ 2026-05-28
- [x] json-to-csv : convertisseur JSON arrays vers CSV ✓ 2026-05-28
- [x] csv-to-yaml : convertisseur CSV vers YAML ✓ 2026-05-28
- [x] file-du : affiche l'utilisation disque des fichiers triée par taille ✓ 2026-05-28

## Vague 68 — CLI Tools
- [x] html-to-markdown : convertit des fichiers HTML ou URLs en Markdown ✓ 2026-05-28
- [x] file-encoding-detect : detecte l'encodage de fichiers (UTF-8, Latin-1, etc.) ✓ 2026-05-28
- [x] markdown-link-extract : extrait tous les liens de fichiers Markdown ✓ 2026-05-28
- [x] git-blame-heatmap : affiche un heatmap de modifications par ligne ✓ 2026-05-28
- [x] http-compression-check : verifie le support de la compression HTTP (gzip, brotli, zstd) ✓ 2026-05-28
- [x] http-cache-check : analyse les headers de cache HTTP (Cache-Control, ETag, Last-Modified, Expires) ✓ 2026-05-28
- [x] http-security-headers : analyse les headers de sécurité HTTP (CSP, HSTS, X-Frame-Options, etc.) ✓ 2026-05-28
- [x] http-cookie-analyzer : analyse les cookies HTTP (SameSite, Secure, HttpOnly, expiration) ✓ 2026-05-28
- [x] http-protocol-check : detecte la version HTTP et les capabilities (HTTP/1.1, HTTP/2, HTTP/3) ✓ 2026-05-28

## Vague 69 — CLI Tools
- [x] http-content-type : verifie le Content-Type et l'encodage de réponses HTTP ✓ 2026-05-28
- [x] json-to-avro : genere un schema Avro depuis du JSON ✓ 2026-05-28
- [x] git-remote-branches : liste les branches remote avec dernier commit ✓ 2026-05-28
- [x] csv-to-parquet : convertisseur CSV vers Apache Parquet ✓ 2026-05-28
- [x] env-to-kubernetes : convertit un .env en ConfigMap/Secret Kubernetes ✓ 2026-05-28

## Vague 70 — CLI Tools
- [x] http-response-diff : compare les réponses HTTP de deux URLs (status, headers, body) ✓ 2026-05-28

## Vague 71 — CLI Tools
- [x] http-websocket-cli : client WebSocket simple pour tester des endpoints WS/WSS ✓ 2026-05-28

## Vague 72 — CLI Tools
- [x] markdown-word-count : compteur de mots/caractères/paragraphes pour fichiers Markdown ✓ 2026-05-28
- [x] yaml-to-sql : genere des requêtes SQL (CREATE TABLE + INSERT) depuis du YAML ✓ 2026-05-28
- [x] csv-to-protobuf : convertisseur CSV vers définitions Protocol Buffers ✓ 2026-05-28

## Vague 73 — CLI Tools
- [x] proc-kill : trouve et termine des processus par nom ou pattern ✓ 2026-05-28
- [x] env-to-sysctl : convertit .env en commandes sysctl Linux ✓ 2026-05-28

## Vague 74 — CLI Tools
- [x] file-sync : synchronise deux répertoires avec comparaison SHA-256 ✓ 2026-05-28
- [x] port-listener : affiche les ports TCP/UDP écoutants avec infos processus ✓ 2026-05-28
- [x] mem-usage : affiche l'utilisation mémoire des processus triée par RSS ✓ 2026-05-28
- [x] cpu-top : moniteur d'utilisation CPU en temps réel ✓ 2026-05-28
- [x] inode-list : liste l'utilisation d'inodes par répertoire ✓ 2026-05-28

## Vague 75 — CLI Tools
- [x] markdown-stats : analyse de statistiques de fichiers Markdown (mots, lignes, code, liens, temps de lecture) ✓ 2026-05-28
- [x] http-proxy-cli : proxy HTTP/HTTPS local configurable en CLI ✓ 2026-05-28
- [x] json-to-dhall : convertisseur JSON vers Dhall ✓ 2026-05-28
- [x] csv-to-sqlite-schema : génère un CREATE TABLE SQL depuis un CSV ✓ 2026-05-28

## Vague 76 — CLI Tools
- [x] yaml-to-env : convertisseur YAML vers .env ✓ 2026-05-28
- [x] git-merge-base : trouve le dernier ancêtre commun entre deux branches ✓ 2026-05-28
- [x] json-to-shell : exporte du JSON en variables shell ✓ 2026-05-28
- [x] csv-to-graph : génère un graphe terminal depuis un CSV ✓ 2026-05-28
- [x] http-timeout-test : teste les timeouts HTTP avec timing détaillé ✓ 2026-05-28

## Vague 77 — CLI Tools
- [x] robots-checker : fetch et parse les fichiers robots.txt de sites web ✓ 2026-05-29
- [x] cookie-parser : parse les headers Set-Cookie HTTP et strings de cookies ✓ 2026-05-29
- [x] git-churn : affiche le churn de fichiers (lignes ajoutées/supprimées) dans un repo git ✓ 2026-05-29
- [x] env-to-helm : convertit des fichiers .env en format Helm values.yaml ✓ 2026-05-29
- [x] http-retry : client HTTP avec retry automatique et exponential backoff ✓ 2026-05-29

## Vague 78 — CLI Tools
- [x] http-rate-limit : mesure le rate limiting d'une API (headers, 429 responses) ✓ 2026-05-29
- [x] json-to-toml : convertisseur JSON vers TOML ✓ 2026-05-29
- [x] csv-to-sql : genere des requetes SQL INSERT depuis un CSV ✓ 2026-05-29
- [x] git-commit-size : affiche la taille des commits (fichiers touches, LOC) ✓ 2026-05-29
- [x] env-to-varnish : convertit un .env en config Varnish VCL ✓ 2026-05-29

## Vague 79 — CLI Tools
- [x] sysctl-tune : gestionnaire de paramètres kernel Linux via sysctl (list, get, set, backup, restore) ✓ 2026-05-29
- [x] net-speed-cli : moniteur de vitesse réseau en temps réel via /proc/net/dev ✓ 2026-05-29
- [x] route-table : gestionnaire de tables de routage Linux (list, add, delete, flush) ✓ 2026-05-29

## Vague 80 — CLI Tools
- [x] access-log-parser : parse et affiche les logs d'accès Apache/Nginx avec stats ✓ 2026-05-29

## Vague 81 — CLI Tools
- [x] ssh-audit : audite la config SSH d'un serveur distant (ciphers, MACs, key exchange) ✓ 2026-05-29
- [x] json-to-dhall : convertisseur JSON vers Dhall ✓ 2026-05-29
- [x] csv-to-sqlite : importe des fichiers CSV dans une base SQLite ✓ 2026-05-29 (dup Vague 36)
- [x] git-commit-lint : linter de messages de commits (Conventional Commits) ✓ 2026-05-29 (dup Vague 49)
- [x] http-sni-check : vérifie le SNI (Server Name Indication) d'un serveur HTTPS ✓ 2026-05-29

## Vague 82 — CLI Tools
- [x] ipcalc-cli : calculateur de sous-réseaux IP (CIDR, broadcast, masque, hosts) ✓ 2026-05-29
- [x] dns-txt : récupère et affiche les enregistrements TXT DNS d'un domaine ✓ 2026-05-29
- [x] json-to-toml : convertisseur JSON vers TOML ✓ 2026-05-29 (dup Vague 78)
- [x] csv-to-sql : génère des requêtes SQL INSERT depuis un CSV ✓ 2026-05-29 (dup Vague 78)
- [x] git-commit-size : affiche la taille des commits (fichiers touchés, LOC) ✓ 2026-05-29 (dup Vague 78)

## Vague 83 — CLI Tools
- [x] netstat-filter : filtre les connexions réseau par état/hôte/port depuis /proc/net ✓ 2026-05-29
- [x] json-to-xml : convertisseur JSON vers XML ✓ 2026-05-29
- [x] csv-to-ini : convertisseur CSV vers INI avec regroupement par section ✓ 2026-05-29
- [x] git-branch-depth : affiche la profondeur d'historique par branche ✓ 2026-05-29
- [x] env-to-nginx : convertit un .env en variables Nginx config ✓ 2026-05-29

## Vague 84 — CLI Tools
- [x] proc-environ : affiche les variables d'environnement d'un processus par PID ✓ 2026-05-29
- [x] json-to-yaml : convertisseur JSON vers YAML ✓ 2026-05-29 (dup Vague 22)
- [x] csv-to-tsv : convertisseur CSV vers TSV ✓ 2026-05-29 (dup Vague 30)
- [x] git-tag-depth : affiche la profondeur d'historique par tag ✓ 2026-05-29
- [x] env-to-apache : convertit un .env en variables Apache config ✓ 2026-05-29

## Vague 85 — CLI Tools
- [x] http-h2-check : vérifie le support HTTP/2 d'un serveur et compare avec HTTP/1.1 ✓ 2026-05-29
- [x] json-to-csv-flat : convertisseur JSON profond vers CSV avec aplatissement configurable ✓ 2026-05-29
- [x] csv-to-jsonl : convertisseur CSV vers JSONL ✓ 2026-05-29
- [x] git-commit-distribution : histogramme des commits par heure du jour ✓ 2026-05-29
- [x] env-to-caddy : convertit un .env en config Caddy ✓ 2026-05-29

## Vague 86 — CLI Tools
- [x] json-to-sql : genere des requetes SQL INSERT depuis du JSON ✓ 2026-05-29
- [x] csv-to-markdown : convertisseur CSV vers tableau Markdown ✓ 2026-05-29
- [x] git-merge-conflict-finder : detecte les conflits de merge dans un repo git ✓ 2026-05-29
- [x] env-to-supabase : convertit un .env en config Supabase ✓ 2026-05-29
- [x] http-redirect-trace : trace les chaines de redirection HTTP avec timing ✓ 2026-05-29

## Vague 87 — CLI Tools
- [x] tsv-to-csv : convertisseur TSV vers CSV ✓ 2026-05-29
- [x] json-to-sql-schema : genere un schema SQL CREATE TABLE depuis du JSON ✓ 2026-05-29
- [x] git-merge-base : trouve le dernier ancetre commun entre deux branches ✓ 2026-05-29

## Vague 88 — CLI Tools
- [x] yaml-to-json : convertisseur YAML vers JSON ✓ 2026-05-29
- [x] csv-to-sql : genere des requetes SQL INSERT depuis un CSV ✓ 2026-05-29
- [x] git-log-oneline : affiche l'historique git en une ligne par commit ✓ 2026-05-29

## Vague 89 — CLI Tools
- [x] env-to-json : convertit un fichier .env en JSON ✓ 2026-05-29
- [x] json-to-env : convertisseur JSON vers .env ✓ 2026-05-29
- [x] git-branch-sort : trie les branches git par date de dernier commit ✓ 2026-05-29

## Vague 90 — CLI Tools
- [x] toml-to-json : convertisseur TOML vers JSON ✓ 2026-05-29
- [x] xml-to-yaml : convertisseur XML vers YAML ✓ 2026-05-29
- [x] git-tag-list : liste les tags git avec details ✓ 2026-05-29

## Vague 91 — CLI Tools
- [x] ini-to-toml : convertisseur INI vers TOML ✓ 2026-05-29
- [x] tsv-to-csv : convertisseur TSV vers CSV ✓ 2026-05-29 (dup Vague 87)
- [x] json-to-sql : genere des requetes SQL INSERT depuis du JSON ✓ 2026-05-29 (dup Vague 86)

## Vague 92 — CLI Tools
- [x] net-arp-table : affiche la table ARP du système depuis /proc/net/arp ✓ 2026-05-29
- [x] json-to-hcl : convertisseur JSON vers HCL (Terraform) ✓ 2026-05-29
- [x] csv-to-rust : génère des structs Rust depuis les headers CSV ✓ 2026-05-29
- [x] git-wip-manager : gestionnaire de commits WIP (save, restore, list) ✓ 2026-05-29
- [x] http-sse-cli : client Server-Sent Events pour tester des endpoints SSE ✓ 2026-05-29

## Vague 94 — CLI Tools
- [x] xml-to-toml : convertisseur XML vers TOML ✓ 2026-05-29
- [x] tsv-to-json : convertisseur TSV vers JSON ✓ 2026-05-29
- [x] csv-to-sql : genere des requetes SQL INSERT depuis un CSV ✓ 2026-05-29
- [x] git-branch-depth : affiche la profondeur d'historique par branche ✓ 2026-05-29
- [x] env-to-caddy : convertit un .env en config Caddy ✓ 2026-05-29

## Vague 93 — CLI Tools
- [x] csv-to-typescript : genere des interfaces TypeScript depuis un CSV avec inference de types ✓ 2026-05-29
- [x] yaml-to-typescript : genere des interfaces TypeScript depuis un fichier YAML ✓ 2026-05-29
- [x] toml-to-typescript : genere des interfaces TypeScript depuis un fichier TOML ✓ 2026-05-29
- [x] ini-to-typescript : genere des interfaces TypeScript depuis un fichier INI ✓ 2026-05-29

## Vague 95 — CLI Tools
- [x] http-benchmark-cli : outil de benchmark HTTP avec percentiles de latence, RPS, et taux d'erreurs ✓ 2026-05-29
- [x] log-anonymizer : anonymise les donnees sensibles dans les logs (emails, IPs, telephones, etc.) ✓ 2026-05-29
- [x] html-gallery-gen : genere une galerie HTML statique depuis un dossier d'images ✓ 2026-05-29

## Vague 96 — CLI Tools
- [x] markdown-checklist-cli : gestionnaire de checklists dans des fichiers Markdown (list, add, toggle, check, uncheck, delete, stats) ✓ 2026-05-29
- [x] csv-to-protobuf : convertisseur CSV vers définitions Protocol Buffers ✓ 2026-05-29
- [x] git-commit-lint : linter de messages de commits (Conventional Commits) ✓ 2026-05-29
- [x] env-to-kubernetes : convertit un .env en ConfigMap/Secret Kubernetes ✓ 2026-05-29
- [x] http-sse-cli : client Server-Sent Events pour tester des endpoints SSE ✓ 2026-05-29

## Vague 97 — CLI Tools
- [x] project-env-vars : scanne un projet et extrait toutes les références de variables d'environnement ✓ 2026-05-29
- [x] go-test-runner : wrapper pratique autour de go test (race, coverage, benchmarks) ✓ 2026-05-29
- [x] go-dep-tree : affiche l'arbre de dépendances d'un module Go avec versions ✓ 2026-05-29
- [x] go-mod-outdated-check : vérifie les dépendances Go obsolètes et signale les nouvelles versions ✓ 2026-05-29

## Vague 98 — CLI Tools
- [x] git-commit-calendar : calendrier de commits style GitHub en terminal ✓ 2026-05-29
- [x] http-redirect-trace : trace les chaînes de redirection HTTP avec timing ✓ 2026-05-29
- [x] json-to-csv : convertisseur JSON arrays vers CSV ✓ 2026-05-29
- [x] csv-to-yaml : convertisseur CSV vers YAML ✓ 2026-05-29
- [x] file-du : affiche l'utilisation disque des fichiers triée par taille ✓ 2026-05-29

## Vague 99 — CLI Tools
- [x] git-blame-summary : résumé de code ownership par fichier (auteurs, %) ✓ 2026-05-29
- [x] yaml-to-env : convertisseur YAML vers .env ✓ 2026-05-29
- [x] csv-to-tsv : convertisseur CSV vers TSV ✓ 2026-05-29
- [x] http-status-cli : vérifie le code HTTP d'une URL avec timing ✓ 2026-05-29
- [x] env-to-json : convertit un fichier .env en JSON ✓ 2026-05-29

## Vague 100 — CLI Tools
- [x] xml-to-ini : convertisseur XML vers INI ✓ 2026-05-29
- [x] tsv-to-yaml : convertisseur TSV vers YAML ✓ 2026-05-29
- [x] git-tag-annotate : annote des tags git en batch ✓ 2026-05-29
- [x] http-cookie-parser : parse et affiche les cookies HTTP d'une URL ✓ 2026-05-29
- [x] json-to-ini : convertisseur JSON vers INI ✓ 2026-05-29

## Vague 101 — CLI Tools
- [x] ini-to-xml : convertisseur INI vers XML ✓ 2026-05-29
- [x] yaml-to-tsv : convertisseur YAML vers TSV ✓ 2026-05-29
- [x] git-tag-delete : supprime des tags git en batch (local et remote) ✓ 2026-05-29
- [x] http-header-check : vérifie la présence de headers HTTP spécifiques ✓ 2026-05-29
- [x] json-to-env : convertisseur JSON vers .env ✓ 2026-05-29

## Vague 102 — CLI Tools
- [x] xml-to-sql : génère des requêtes SQL INSERT depuis un XML ✓ 2026-05-29
- [x] tsv-to-ini : convertisseur TSV vers INI ✓ 2026-05-29
- [x] git-branch-rename : renomme des branches git en batch ✓ 2026-05-29 (dup Vague 50)
- [x] http-h2-check : vérifie le support HTTP/2 d'un serveur ✓ 2026-05-29 (dup Vague 85)
- [x] ini-to-env : convertisseur INI vers .env ✓ 2026-05-29

## Vague 103 — CLI Tools
- [x] markdown-heading-list : liste tous les titres d'un fichier Markdown avec niveaux ✓ 2026-05-29
- [x] http-basic-auth : client HTTP avec authentification basique ✓ 2026-05-29
- [x] xml-to-env : convertisseur XML vers .env ✓ 2026-05-29
- [x] git-commit-range : affiche les commits entre deux refs avec stats ✓ 2026-05-29
- [x] file-word-count : compte les mots uniques dans des fichiers ✓ 2026-05-29

## Vague 112 — CLI Tools
- [x] json-to-sqlite-schema : genere du SQL SQLite depuis du JSON ✓ 2026-05-29
- [x] http-body-check : valide le corps HTTP contre des regles ✓ 2026-05-29

## Vague 113 — CLI Tools
- [x] file-concat : concatene des fichiers avec separateurs, verification SHA-256 et globs ✓ 2026-05-29
- [x] file-split-by-size : divise des fichiers en chunks par taille ou nombre avec verification SHA-256 ✓ 2026-05-29
- [x] http-pipeline-cli : enchainement de requetes HTTP ou chaque reponse alimente la suivante ✓ 2026-05-29

## Vague 114 — CLI Tools
- [x] markdown-toc-cli : genere et insere une table des matieres dans des fichiers Markdown ✓ 2026-05-29
- [x] csv-to-json-schema : genere un JSON Schema depuis un fichier CSV ✓ 2026-05-29

## Vague 115 — CLI Tools
- [x] word-frequency : compteur de frequence de mots dans des fichiers texte ✓ 2026-05-29
- [x] tsv-to-csv : convertisseur TSV vers CSV avec delimitateur configurable ✓ 2026-05-29
- [x] html-to-text : convertit du HTML en texte brut ✓ 2026-05-29
- [x] yaml-diff : compare deux fichiers YAML et affiche les differences ✓ 2026-05-29
- [x] markdown-link-list : extrait tous les liens d'un fichier Markdown ✓ 2026-05-29

## Vague 116 — CLI Tools
- [x] http-cors-check : vérifie les headers CORS d'un endpoint HTTP ✓ 2026-05-29
- [x] markdown-to-slides : convertit du markdown en présentation HTML ✓ 2026-05-29

## Vague 117 — CLI Tools
- [x] csv-to-pie-chart : convertisseur CSV vers graphique camembert SVG ✓ 2026-05-29
- [x] markdown-code-extract : extrait les blocs de code depuis des fichiers Markdown ✓ 2026-05-29
- [x] json-to-nix : convertisseur JSON vers expressions Nix ✓ 2026-05-29

## Vague 118 — CLI Tools
- [x] dockerfile-lint : analyse et signale les mauvaises pratiques dans les Dockerfiles ✓ 2026-05-29
- [x] json-to-terraform-var : genere des variables Terraform depuis du JSON ✓ 2026-05-29
- [x] csv-to-chart-svg : genere des graphiques SVG (barres, lignes) depuis un CSV ✓ 2026-05-29
- [x] git-commit-author : affiche les infos de l'auteur du dernier commit ✓ 2026-05-29
- [x] http-follow-redirect : fetch une URL en suivant toutes les redirections ✓ 2026-05-29

## Vague 119 — CLI Tools
- [x] http-batch-fetch : fetch multiple URLs concurrently with status, size, and timing ✓ 2026-05-29

## Vague 120 — CLI Tools
- [x] markdown-frontmatter : lit, écrit et gère le YAML frontmatter dans des fichiers Markdown ✓ 2026-05-29
- [x] http-stream-cli : stream HTTP responses en temps réel avec progress tracking ✓ 2026-05-29

## Vague 121 — CLI Tools
- [x] csv-to-chart-svg : genere des graphiques SVG (barres, lignes) depuis un CSV ✓ 2026-05-29
- [x] git-commit-lint : linter de messages de commits (Conventional Commits) ✓ 2026-05-29

## Vague 122 — CLI Tools
- [x] markdown-heading-validate : valide la hierarchie des titres Markdown (pas de niveaux sautes, H1 en premier) ✓ 2026-05-29
- [x] http-content-negotiation : teste la negociation de contenu HTTP (Accept headers vs reponses) ✓ 2026-05-29
- [x] file-compress : compresse des fichiers avec gzip, zip, tar.gz, zlib et compare les ratios ✓ 2026-05-29

## Vague 123 — CLI Tools
- [x] http-multipart-cli : envoie des requetes HTTP multipart/form-data depuis le terminal ✓ 2026-05-29
- [x] file-decompress : decompresse des fichiers gzip, zip, tar.gz, bzip2, xz, zlib ✓ 2026-05-29
- [x] file-encrypt : chiffre/dechiffre des fichiers avec AES-256-GCM ✓ 2026-05-29

## Vague 124 — CLI Tools
- [x] http-idempotency-check : vérifie l'idempotence HTTP en envoyant la même requête N fois et en comparant les réponses ✓ 2026-05-29

## Vague 125 — CLI Tools
- [x] rss-reader-cli : lecteur de flux RSS/Atom en terminal (RSS 2.0, Atom, sortie text/JSON) ✓ 2026-05-29
- [x] file-watch-exec : surveille des fichiers et lance des commandes sur changement (inotify + polling) ✓ 2026-05-29

## Vague 126 — CLI Tools
- [x] git-commit-lint : linter de messages de commits git (conventional commits, longueur, style) ✓ 2026-05-29
- [x] json-to-toml-strict : convertisseur JSON vers TOML avec validation stricte ✓ 2026-05-29
- [x] csv-sample : échantillonne aléatoirement N lignes d'un CSV ✓ 2026-05-29
- [x] http-method-check : vérifie les méthodes HTTP supportées par une URL ✓ 2026-05-29
- [x] yaml-to-env : convertisseur YAML vers .env ✓ 2026-05-29

## Vague 127 — CLI Tools
- [x] text-readability : analyse la lisibilité de texte (Flesch, Gunning Fog, SMOG, Coleman-Liau) ✓ 2026-05-29
- [x] file-hardlink : crée et gère des hardlinks entre fichiers ✓ 2026-05-29
- [x] http-chunked-cli : client HTTP pour tester le transfer-encoding chunked ✓ 2026-05-29
- [x] json-to-yaml-strict : convertisseur JSON vers YAML avec validation de types ✓ 2026-05-29
- [x] csv-to-sqlite-bulk : import bulk CSV vers SQLite avec transactions ✓ 2026-05-29

## Vague 128 — CLI Tools
- [x] json-patch : outil JSON Patch CLI (RFC 6902) avec apply, diff, inverse ✓ 2026-05-29
- [x] yaml-merge : fusionne plusieurs fichiers YAML avec priorité configurable ✓ 2026-05-29
- [x] csv-aggregate : agrège des données CSV par groupe (sum, avg, count, min, max) ✓ 2026-05-29
- [x] git-commit-daily : compteur de commits par jour sur une période ✓ 2026-05-29
- [x] http-digest-auth : client HTTP avec authentification Digest ✓ 2026-05-29

## Vague 129 — CLI Tools
- [x] csv-to-markdown-table : convertit un CSV en tableau Markdown formaté ✓ 2026-05-29
- [x] git-branch-list : liste les branches git avec dernier commit ✓ 2026-05-29
- [x] env-diff-cli : compare deux fichiers .env et affiche les différences ✓ 2026-05-29

## Vague 130 — CLI Tools
- [x] proc-open-files : affiche les fichiers ouverts par un processus depuis /proc/PID/fd ✓ 2026-05-29
- [x] net-interface-stats : affiche les statistiques des interfaces reseau ✓ 2026-05-29
- [x] json-to-go-interfaces : genere des interfaces Go depuis du JSON ✓ 2026-05-29
- [x] csv-to-mermaid-flow : genere un diagramme Mermaid flowchart depuis un CSV ✓ 2026-05-29
- [x] http-connection-pool-test : teste le connection pooling HTTP ✓ 2026-05-29

## Vague 131 — CLI Tools
- [x] markdown-image-list : liste toutes les images referencees dans des fichiers Markdown ✓ 2026-05-29
- [x] http-response-size : mesure la taille des reponses HTTP (headers + body) ✓ 2026-05-29
- [x] json-to-csv-schema : genere un schema de colonnes CSV depuis du JSON ✓ 2026-05-29
- [x] git-commit-merge : liste les commits de merge dans un repo git ✓ 2026-05-29
- [x] env-to-script : genere un script shell depuis un fichier .env ✓ 2026-05-29

## Vague 132 — CLI Tools
- [x] docker-volume-list : liste les volumes Docker avec tailles et containers connectes ✓ 2026-05-29

## Vague 133 — CLI Tools
- [x] file-age-cli : affiche l'âge des fichiers (création, modification, accès) avec format humain ✓ 2026-05-29

## Vague 134 — CLI Tools
- [x] markdown-checklist-stats : affiche les stats de checklists dans des fichiers Markdown ✓ 2026-05-29

## Vague 135 — CLI Tools
- [x] json-to-toml : convertisseur JSON vers TOML ✓ 2026-05-29
- [x] csv-to-sql : génère des requêtes SQL (CREATE + INSERT) depuis un CSV ✓ 2026-05-29
- [x] git-commit-size : affiche la taille des commits (fichiers ajoutés/modifiés/supprimés) ✓ 2026-05-29
- [x] env-to-markdown : génère un tableau Markdown depuis un .env ✓ 2026-05-29
- [x] xml-to-toml : convertisseur XML vers TOML ✓ 2026-05-29

## Vague 136 — CLI Tools
- [x] csv-to-ini : convertisseur CSV vers INI avec regroupement par clé ✓ 2026-05-29
- [x] json-to-yaml : convertisseur JSON vers YAML avec indentation configurable ✓ 2026-05-29
- [x] git-tag-create : crée des tags git annotés en batch depuis un fichier ✓ 2026-05-29
- [x] env-to-csv : convertit un fichier .env en CSV ✓ 2026-05-29
- [x] xml-to-json : convertisseur XML vers JSON ✓ 2026-05-29

## Vague 137 — CLI Tools
- [x] csv-to-protobuf : convertisseur CSV vers définitions Protocol Buffers ✓ 2026-05-29
- [x] json-to-go : génère des structs Go depuis du JSON ✓ 2026-05-29
- [x] git-branch-age : affiche l'âge de chaque branche git ✓ 2026-05-29
- [x] env-to-yaml : convertit un fichier .env en YAML ✓ 2026-05-29
- [x] yaml-to-csv : convertisseur YAML vers CSV ✓ 2026-05-29

## Vague 138 — CLI Tools
- [x] json-to-rust : génère des structs Rust depuis du JSON ✓ 2026-05-29
- [x] csv-to-sql : génère des requêtes SQL INSERT depuis un CSV ✓ 2026-05-29
- [x] git-commit-lint : linter de messages de commits (Conventional Commits) ✓ 2026-05-29
- [x] env-to-toml : convertit un fichier .env en TOML ✓ 2026-05-29
- [x] xml-to-yaml : convertisseur XML vers YAML ✓ 2026-05-29

## Vague 139 — CLI Tools
- [x] toml-to-csv : convertisseur TOML vers CSV ✓ 2026-05-29
- [x] json-to-ini : convertisseur JSON vers INI ✓ 2026-05-29
- [x] csv-to-markdown : convertisseur CSV vers tableau Markdown ✓ 2026-05-29
- [x] git-tag-list : liste les tags git avec détails ✓ 2026-05-29
- [x] env-validate : valide un .env contre une liste de clés requises ✓ 2026-05-29

## Vague 140 — CLI Tools
- [x] ini-to-csv : convertisseur INI vers CSV ✓ 2026-05-29
- [x] yaml-flatten : aplatit du YAML imbriqué en paires clé-valeur ✓ 2026-05-29
- [x] csv-to-typescript : génère des interfaces TypeScript depuis un CSV ✓ 2026-05-29
- [x] git-diff-summary : résumé des changements entre deux commits ✓ 2026-05-29
- [x] http-redirect-trace : trace les chaînes de redirection HTTP ✓ 2026-05-29

## Vague 141 — CLI Tools
- [x] json-to-ruby : génère des structs Ruby depuis du JSON ✓ 2026-05-29
- [x] csv-to-latex : convertisseur CSV vers tableau LaTeX ✓ 2026-05-29
- [x] git-commit-graph : affiche un graphique ASCII du historique de commits ✓ 2026-05-29
- [x] env-to-dhall : convertit un fichier .env en expression Dhall ✓ 2026-05-29
- [x] http-chunked-cli : client HTTP pour tester le transfer-encoding chunked ✓ 2026-05-29 (dup Vague 127)

## Vague 142 — CLI Tools
- [x] json-to-xlsx : convertisseur JSON arrays vers fichiers Excel (.xlsx) ✓ 2026-05-29
- [x] csv-to-pdf : genere des tableaux PDF depuis des fichiers CSV ✓ 2026-05-29
- [x] git-commit-calendar : calendrier de commits style GitHub en terminal ✓ 2026-05-29
- [x] env-to-nginx : convertit un .env en variables Nginx config ✓ 2026-05-29
- [x] http-cors-check : vérifie les headers CORS d'un endpoint HTTP ✓ 2026-05-29

## Vague 144 — CLI Tools
- [x] markdown-to-text : supprime le formatage markdown et affiche du texte brut ✓ 2026-05-29
- [x] json-to-sql : genere des requetes SQL INSERT depuis du JSON ✓ 2026-05-29
- [x] csv-to-yaml : convertisseur CSV vers YAML ✓ 2026-05-29
- [x] git-branch-list : liste les branches git avec dernier commit ✓ 2026-05-29
- [x] env-diff-cli : compare deux fichiers .env et affiche les différences ✓ 2026-05-29

## Vague 145 — CLI Tools
- [x] json-to-graphql : genere un schema GraphQL depuis du JSON ✓ 2026-05-29
- [x] csv-to-latex : convertisseur CSV vers tableau LaTeX ✓ 2026-05-29
- [x] git-commit-lint : linter de messages de commits (Conventional Commits) ✓ 2026-05-29
- [x] env-to-docker-compose : convertit un .env en docker-compose.yml ✓ 2026-05-29
- [x] http-cache-analyzer : analyse les headers de cache HTTP (Cache-Control, ETag, Expires) ✓ 2026-05-29

## Vague 146 — CLI Tools
- [x] tcp-ping : teste la connectivité TCP vers un hôte/port avec timing (alternative à ping sans ICMP) ✓ 2026-05-29
- [x] json-to-go : génère des structs Go depuis du JSON avec type inference ✓ 2026-05-29
- [x] csv-to-xlsx : convertisseur CSV vers Excel (.xlsx) ✓ 2026-05-29
- [x] git-commit-template : gestionnaire de templates de messages git ✓ 2026-05-29
- [x] http-headers-cli : affiche tous les headers HTTP d'une URL formatés ✓ 2026-05-29

## Vague 147 — CLI Tools
- [x] json-to-protobuf : génère des fichiers .proto depuis du JSON ✓ 2026-05-29
- [x] csv-to-sql : génère des requêtes SQL (CREATE + INSERT) depuis un CSV ✓ 2026-05-29
- [x] git-log-filter : filtre l'historique git par auteur, date, fichier ✓ 2026-05-29
- [x] yaml-to-env : convertisseur YAML vers .env ✓ 2026-05-29
- [x] http-status-cli : vérifie le code HTTP d'une URL et affiche le résultat ✓ 2026-05-29

## Vague 148 — CLI Tools
- [x] file-perm-check : vérifie les permissions de fichiers et signale les problèmes de sécurité ✓ 2026-05-29
- [x] json-to-avro : génère un schema Avro depuis du JSON ✓ 2026-05-29
- [x] csv-to-xlsx : convertisseur CSV vers Excel (.xlsx) ✓ 2026-05-29
- [x] git-commit-sign : signe des commits git avec GPG ✓ 2026-05-29
- [x] http-etag-check : vérifie et compare les ETags HTTP d'une URL ✓ 2026-05-29

## Vague 149 — CLI Tools
- [x] xlsx-to-csv : convertisseur Excel (.xlsx) vers CSV ✓ 2026-05-29
- [x] git-prune-refs : nettoie les références distantes obsolètes ✓ 2026-05-29
- [x] json-to-protobuf : génère un fichier .proto depuis du JSON ✓ 2026-05-29

## Vague 150 — CLI Tools
- [x] xml-lint : vérifie la syntaxe et le style XML ✓ 2026-05-29
- [x] json-to-toml-nested : convertisseur JSON imbriqué vers TOML ✓ 2026-05-29
- [x] csv-unique : extrait les lignes uniques d'un CSV par colonne ✓ 2026-05-29
- [x] git-commit-diff : affiche le diff d'un commit spécifique ✓ 2026-05-29
- [x] env-to-sql : génère des requêtes SQL INSERT depuis un .env ✓ 2026-05-29

## Vague 152 — CLI Tools
- [x] toml-merge : fusionne plusieurs fichiers TOML avec priorité configurable ✓ 2026-05-29
- [x] markdown-outline : génère un outline de titres depuis du Markdown ✓ 2026-05-29
- [x] git-commit-revert : assistant interactif de revert de commits ✓ 2026-05-29
- [x] csv-to-bar-chart : génère des graphiques en barres terminaux depuis CSV ✓ 2026-05-29
- [x] http-response-time-histogram : mesure et affiche un histogramme de temps de réponse HTTP ✓ 2026-05-29

## Vague 153 — CLI Tools
- [x] regex-tester : teste des patterns regex contre des chaînes avec détails des captures ✓ 2026-05-29
- [x] json-to-typescript : génère des interfaces TypeScript depuis du JSON ✓ 2026-05-29

## Vague 154 — CLI Tools
- [x] cron-next-run : calcule les prochaines exécutions d'expressions cron (timezone, JSON) ✓ 2026-07-28
- [x] git-standup : résumé style standup des commits récents multi-repos (JSON, stats) ✓ 2026-07-28
- [x] totp-cli : générateur TOTP RFC 6238 (sha1/256/512, watch mode, JSON) ✓ 2026-07-28
- [x] hmac-cli : calcule des signatures HMAC (md5/sha1/sha2) depuis le terminal ✓ 2026-07-28
- [x] url-query-builder : parse, build et manipulation de query strings URL (JSON, stdin) ✓ 2026-07-28
- [x] json-lines : flatten JSON vers lignes `path = value` grep-friendly et inverse ✓ 2026-07-28
## Vague 155 — CLI Tools
- [x] epoch-cli : convertit entre timestamps Unix (s/ms/us/ns) et dates humaines (multi-formats, timezone, relatif, JSON) ✓ 2026-07-28
- [x] text-case : convertit entre camelCase, snake_case, kebab-case, PascalCase, CONSTANT_CASE, Title Case (batch, stdin) ✓ 2026-07-28
- [x] fuzzy-cli : fuzzy matching de chaînes avec scores (Levenshtein, Jaro-Winkler, sous-séquence, tri des résultats) ✓ 2026-07-28
- [x] lorem-gen : génère du texte lorem ipsum (mots, phrases, paragraphes, seed reproductible) ✓ 2026-07-28
- [x] base62-cli : encode/decode base62 (alphabet alphanumérique) avec support bigint et stdin ✓ 2026-07-28
- [x] ulid-cli : génère et décode des ULID (Crockford Base32, sortable, JSON) ✓ 2026-07-28

## Vague 156 — CLI Tools
- [x] roman-cli : convertit entre chiffres romains et arabes (validation, batch, JSON) ✓ 2026-07-28
- [x] levenshtein-cli : distance d'édition entre chaînes (matrice, ops, batch) ✓ 2026-07-28
- [x] iban-validate : valide des IBAN selon ISO 13616 (checksum, format pays, JSON) ✓ 2026-07-28
- [x] punycode-cli : encode/decode punycode (RFC 3492, IDN) ✓ 2026-07-28
- [x] hashid-cli : génère des hashids YouTube-style depuis des entiers (réversible, salt) ✓ 2026-07-28
## Vague 157 — CLI Tools
- [x] base85-cli : encode/decode Ascii85 (Adobe) et Z85 (ZeroMQ) avec stdin/batch/JSON ✓ 2026-07-28
- [x] morse-cli : encode/decode code Morse international (audio opt, batch, JSON) ✓ 2026-07-28
- [x] soundex-cli : algorithmes phonétiques Soundex, Metaphone, NYSIIS (similarité, batch, JSON) ✓ 2026-07-28
- [x] damerau-cli : distance de Damerau-Levenshtein avec détail des opérations ✓ 2026-07-28
- [x] caesar-cli : chiffre/déchiffre Caesar et ROT13 avec brute-force intégré ✓ 2026-07-28

## Vague 158 — CLI Tools
- [x] rot5-cli : chiffre/déchiffre ROT5 sur les chiffres uniquement (batch, stdin, JSON) ✓ 2026-07-28
- [x] atbash-cli : chiffre/déchiffre Atbash (alphabet inversé, batch, stdin, JSON) ✓ 2026-07-28
- [x] vigenere-cli : chiffre/déchiffre Vigenère avec clé (analyze, stdin, JSON) ✓ 2026-07-28
- [x] xor-cli : chiffre/déchiffre XOR hex ou base64 (single-byte, key, bruteforce) ✓ 2026-07-28
- [x] pig-latin-cli : traduit en/fr pig latin (batch, stdin, JSON) ✓ 2026-07-28

## Vague 159 — CLI Tools (chiffrement, suite)
- [x] caesar-cli : chiffre/déchiffre César avec décalage configurable (analyze, bruteforce, batch, stdin, JSON) ✓ 2026-07-28
- [x] morse-cli : encode/décode le code Morse (batch, stdin, JSON) ✓ 2026-07-28
- [x] tap-code-cli : encode/décode le tap code carré Polybe 5x5 (batch, stdin, JSON) ✓ 2026-07-28
- [x] bacon-cipher-cli : chiffre de Bacon bilitère A/B (encode/decode, batch, stdin, JSON) ✓ 2026-07-28
- [x] rail-fence-cli : chiffre des rails (zigzag) encode/decode avec N rails (batch, stdin, JSON) ✓ 2026-07-28
- [x] scytale-cli : chiffre scytale cylindrique (transposition grecque, encode/decode, batch, stdin, JSON) ✓ 2026-07-28
- [x] route-cipher-cli : chiffre de transposition par route (rows/columns/boustrophedon/diagonal/spiral, batch, stdin, JSON) ✓ 2026-07-28
- [x] columnar-transposition-cli : transposition columnar avec clé (encode/decode, batch, stdin, JSON) ✓ 2026-07-28
- [x] gronsfeld-cli : chiffre Gronsfeld (Vigenère à clé numérique 0-9, analyze, crack, batch, stdin, JSON) ✓ 2026-07-28
- [x] affine-cli : chiffre affine E(x)=(ax+b) mod 26 (bruteforce, inverse modulaire, batch, stdin, JSON) ✓ 2026-07-28
- [x] rot47-cli : rotation sur 94 caractères ASCII imprimables (auto-inverse, custom -n, batch, stdin, JSON) ✓ 2026-07-28

## Vague 160 — CLI Tools (chiffrement, suite)
- [x] base32-cli : encode/decode Base32 RFC 4648 (std/hex/nopad/crockford, batch, stdin, JSON) ✓ 2026-07-28
- [x] crc32-cli : checksums CRC-32 (IEEE, Castagnoli, Koopman) et CRC-64 (ISO, ECMA), verify, batch, JSON ✓ 2026-07-28
- [x] nato-cli : alphabet phonétique NATO/ICAO (encode, decode, batch, JSON) ✓ 2026-07-28
- [x] beaufort-cli : chiffre Beaufort réciproque C=(K-P) mod 26 (encode=decode, batch, stdin, JSON) ✓ 2026-07-28
- [x] keyword-cipher-cli : substitution monoalphabétique par mot-clé (alphabet, encode/decode, batch, JSON) ✓ 2026-07-28
- [x] playfair-cli : chiffre Playfair 5x5 avec I/J fusionnés (square, digraphes, padding X, batch, JSON) ✓ 2026-07-28

## Vague 161 — CLI Tools (chiffrement, suite)
- [x] adfgvx-cli : chiffre allemand ADFGVX 1918 (carré 6x6 A-Z0-9 + transposition columnar, encode/decode, batch, JSON) ✓ 2026-07-28
- [x] hill-cli : chiffre de Hill matriciel 2x2/3x3 mod 26 (inverse modulaire, padding X, batch, JSON) ✓ 2026-07-28
- [x] trifid-cli : chiffre Trifid de Delastelle (cube 3x3x3 + fractionation horizontale, période, batch, JSON) ✓ 2026-07-28
- [x] four-square-cli : chiffre Four-Square de Delastelle (2 carrés 5x5, digraphes, X-injection, batch, JSON) ✓ 2026-07-28
- [x] two-square-cli : chiffre Two-Square de Wheatstone (Playfair horizontal, involution, batch, JSON) ✓ 2026-07-28
- [x] nihilist-cli : chiffre Nihiliste russe (Polybe 5x5 + Vigenère numérique, batch, JSON) ✓ 2026-07-28

## Vague 162 — CLI Tools (chiffrement, suite)
- [x] amsco-cli : chiffre AMSCO (transposition incomplète colonar avec séquences 2-1, encode/decode, batch, stdin, JSON) ✓ 2026-07-28
- [x] bifid-cli : chiffre Bifid de Delastelle (carré Polybe 5x5 + fractionation, période, batch, stdin, JSON) ✓ 2026-07-28
- [x] beaufort-variant-cli : chiffre Beaufort variant C=(P+K) mod 26 (encode/decode, batch, JSON) ✓ 2026-07-28
- [x] autokey-cli : chiffre Autokey (Vigenère avec clé étendue par le plaintext, batch, stdin, JSON) ✓ 2026-07-28
- [x] running-key-cli : chiffre Running Key (texte long comme clé, batch, stdin, JSON) ✓ 2026-07-28

## Vague 163 — CLI Tools (chiffrement, suite)
- [x] porta-cli : chiffre Porta (tableau réciproque 13 paires, batch, stdin, JSON) ✓ 2026-07-28
- [x] quagmire-cli : chiffres Quagmire I-II-III-IV (clé K1/K2 with alphabet indicateur, batch, JSON) ✓ 2026-07-28
- [x] checkerboard-cli : chiffre à damier straddling checkerboard (homophone, batch, JSON) ✓ 2026-07-28
- [x] progressive-key-cli : chiffre à clé progressive (Vigenère avec shift +N par période, batch, JSON) ✓ 2026-07-28
- [x] trithemius-cli : chiffre de Trithémius (tabula recta progressive, shift 0,1,2,..., batch, JSON) ✓ 2026-07-28

## Vague 164 — CLI Tools (chiffrement, suite)
- [x] homophonic-cli : substitution homophonique (codes proportionnels à la fréquence, clé JSON, batch, JSON) ✓ 2026-07-28
- [x] nicodemus-cli : chiffre Nicodemus (Vigenère + transposition columnar à même clé, batch, JSON) ✓ 2026-07-28
- [x] phillips-cli : chiffre Phillips (8 carrés 5x5 shiftés, clé, batch, JSON) ✓ 2026-07-28
- [x] polybius-square-cli : carré de Polybe classique (coordonnées 11-55, keyed square, batch, JSON) ✓ 2026-07-28
- [x] nomenclator-cli : chiffre par nomenclateur (codes mots 90xx + table lettres shiftée, nulls, batch, JSON) ✓ 2026-07-28

## Vague 165 — CLI Tools (chiffrement, suite)
- [x] vic-cli : chiffre VIC simplifié (straddling checkerboard ATONESIR + double transposition à mot-clé, batch, JSON) ✓ 2026-07-28

## Vague 166 — CLI Tools (utilitaires)
- [x] adfgx-cli : chiffre ADFGX WWI (carré Polybius 5x5 + transposition columnar, batch, JSON) ✓ 2026-07-28
- [x] newsletter-template-cli : génère un squelette newsletter Markdown/HTML avec sections et placeholders ✓ 2026-07-28
- [x] csv-normalize : normalise CSV (détection délimiteur, conversion, quoting, EOL) ✓ 2026-07-28
- [x] json-escape : échappe/déséchappe du contenu de chaînes JSON (stdin/arg, JSON out) ✓ 2026-07-28
- [x] url-canonicalize : normalise les URLs (lower host, sans fragment, query triée, strip tracking, dedup) ✓ 2026-07-28
- [x] env-var-export : convertit un .env en export sh/fish/csh/PowerShell/cmd/JSON ✓ 2026-07-28

## Vague 167 — CLI Tools (utilitaires)
- [x] markdown-check-links : vérifie les ancres internes d'un Markdown (règles de slug GitHub, JSON, exit codes) ✓ 2026-07-28
- [x] diff-lines : diff ligne par ligne LCS minimal (unified context, JSON, couleurs) ✓ 2026-07-28
- [x] file-mime-sniff : détecte le type MIME par magic bytes (zip-subtypes docx/xlsx/apk/jar, text heuristics) ✓ 2026-07-28
- [x] cli-wc : compte lignes/mots/caractères/bytes, JSON, paragraphs ✓ 2026-07-28

## Vague 168 — CLI Tools (utilitaires)
- [x] string-case-convert : conversion camelCase/snake_case/kebab/Pascal/CONSTANT ✓ 2026-07-28
- [x] shuffle-lines : mélange aléatoirement les lignes (Fisher–Yates, seedable) ✓ 2026-07-28
- [x] truncate-lines : tronque chaque ligne à N colonnes avec ellipsis ✓ 2026-07-28
- [x] indent-convert : convertit entre espaces et tabs en début de ligne ✓ 2026-07-28
- [x] timestamp-diff : calcule la différence entre deux timestamps ✓ 2026-07-28

## Vague 169 — CLI Tools (utilitaires)
- [x] rot13-cli : applique ROT13/ROT47 sur des fichiers ou stdin (batch, JSON) ✓ 2026-07-28
- [x] reverse-lines : inverse l'ordre des lignes (stdin/fichier, in-place option) ✓ 2026-07-28
- [x] csv-sum : somme/moyenne/min/max rapide sur une colonne CSV ✓ 2026-07-28
- [x] dir-size-tree : affiche la taille des sous-dossiers en arbre ✓ 2026-07-28
- [x] uppercase-files : renomme des fichiers en majuscules/minuscules (batch) ✓ 2026-07-28

## Vague 170 — CLI Tools (utilitaires)
- [x] trim-lines : supprime les espaces/tabulations en début/fin de ligne ✓ 2026-07-28
- [x] uniq-lines : détecte et supprime les lignes dupliquées (count, sort) ✓ 2026-07-28
- [x] join-lines : joint des lignes avec un délimiteur personnalisé ✓ 2026-07-28
- [x] split-lines : découpe les lignes longues en chunks de N caractères ✓ 2026-07-28
- [x] column-align : aligne des colonnes de texte sur un délimiteur ✓ 2026-07-28

## Vague 171 — CLI Tools (texte)
- [x] word-frequency : compte la fréquence des mots dans un texte ✓ 2026-07-28
- [x] ngram-count : génère et compte des n-grams de mots ou caractères ✓ 2026-07-28
- [x] case-convert : convertit casse camel/snake/kebab/title par ligne ✓ (déjà fait: string-case-convert)
- [x] number-lines : numérote les lignes avec format configurable ✓ 2026-07-28
- [x] squeeze-blank : compresse plusieurs lignes vides en une seule ✓ 2026-07-28

## Vague 172 — CLI Tools (texte)
- [x] text-wrap : rewrap un texte à une largeur donnée (soft/hard) ✓ 2026-07-28
- [x] ansi-strip : supprime les codes ANSI d'un flux ✓ 2026-07-28
- [x] line-filter : garde/supprime les lignes contenant/matchant un pattern ✓ 2026-07-28
- [x] text-rotate : décale les caractères (rot47 custom, caesar) ✓ 2026-07-28
- [x] grep-context : grep-like avec lignes de contexte avant/après ✓ 2026-07-28

## Vague 173 — CLI Tools (texte)
- [x] tail-head : affiche les N premières et N dernières lignes d'un fichier (milieu élidé) ✓ 2026-07-28
- [x] col-swap : échange l'ordre de colonnes délimitées (CSV/TSV/espace) ✓ 2026-07-28
- [x] line-shuffle : mélange aléatoirement les lignes (Fisher-Yates, seed optionnel) ✓ 2026-07-28
- [x] text-pad : complète/tronque chaque ligne à une largeur fixe ✓ 2026-07-28
- [x] char-decode : affiche les points de code Unicode de chaque caractère ✓ 2026-07-28

## Vague 174 — CLI Tools (texte)
- [x] word-reverse : inverse l'ordre des mots de chaque ligne ✓ 2026-07-28
- [x] line-reverse : inverse l'ordre des lignes d'un fichier (tac-like) ✓ 2026-07-28
- [x] text-columns : reformate un texte en N colonnes équilibrées ✓ 2026-07-28
- [x] char-count : compte les occurrences de chaque caractère (histogramme) ✓ 2026-07-28
- [x] text-justify : justifie un texte à une largeur donnée ✓ 2026-07-28

## Vague 175 — CLI Tools (texte)
- [x] text-dedupe : supprime les lignes dupliquées adjacentes (uniq-like avec options) ✓ 2026-07-28
- [x] tab-expand : convertit tabs en espaces (et inversement) avec tabstop configurable ✓ 2026-07-28
- [x] csv-fill-down : remplit les cellules vides d'un CSV avec la valeur précédente ✓ 2026-07-28
- [x] text-prefix : ajoute un préfixe/suffixe à chaque ligne ✓ 2026-07-28
- [x] random-lines : sélectionne N lignes aléatoires d'un fichier (échantillonnage) ✓ 2026-07-28

## Vague 179 — CLI Tools (texte/CSV)
- [x] text-wrap : ré-enveloppe un texte à largeur donnée (fmt-like) ✓ 2026-07-28
- [x] csv-drop-columns : supprime des colonnes CSV par nom ou index ✓ 2026-07-28
- [x] line-number : préfixe les lignes avec numéros (nl-like) ✓ 2026-07-28
- [x] char-reverse : inverse les caractères de chaque ligne (rev-like) ✓ 2026-07-28
- [x] csv-transpose : transpose lignes/colonnes d'un CSV ✓ 2026-07-28
- [x] csv-swap-rows : échange/déplace/inverse des lignes de CSV par index ✓ 2026-07-28

## Vague 176 — CLI Tools (texte)
- [x] text-replace : rechercher/remplacer multi-pattern (literal + regex) en fichier/stdin ✓ 2026-07-28
- [x] line-slice : extrait une tranche de lignes (start:end style) avec pas ✓ 2026-07-28
- [x] text-fold : replie des lignes longues en lignes multiples à largeur fixe ✓ 2026-07-28
- [x] strip-html : enlève les balises HTML d'un texte (texte brut) ✓ 2026-07-28
- [x] csv-filter-rows : filtre les lignes d'un CSV par expression sur colonnes ✓ 2026-07-28

## Vague 177 — CLI Tools (texte)
- [x] trim-lines : supprime espaces début/fin de chaque ligne (left/right/both) ✓ 2026-07-28
- [x] text-uniq-chars : compte les caractères uniques d'un texte (alphabet, set) ✓ 2026-07-28
- [x] csv-sample : échantillonne N lignes aléatoires d'un CSV (seed optionnel) ✓ 2026-07-28
- [x] indent-lines : indente/désindente des lignes de N espaces ou tabs ✓ 2026-07-28
- [x] text-between : extrait le texte entre deux marqueurs/délimiteurs ✓ 2026-07-28

## Vague 178 — CLI Tools (texte)
- [x] line-join : joint toutes les lignes d'un fichier avec un délimiteur ✓ 2026-07-28
- [x] text-split : divise un texte en lignes par délimiteur (inverse de line-join) ✓ 2026-07-28
- [x] csv-row-filter : garde/supprime les lignes selon une condition sur une colonne ✓ 2026-07-28
- [x] text-prefix-lines : préfixe chaque ligne avec numéro ou texte ✓ 2026-07-28 (déjà fait: text-prefix)
- [x] csv-column-swap : échange deux colonnes dans un CSV ✓ 2026-07-28

## Vague 179 — CLI Tools (fichiers & texte)
- [x] text-wrap : formate un texte à largeur fixe (mots entiers, pas de coupure) ✓ 2026-07-28
- [x] csv-drop-columns : supprime une ou plusieurs colonnes d'un CSV ✓ 2026-07-28
- [x] line-number : ajoute/supprime les numéros de ligne ✓ 2026-07-28
- [x] text-reverse-lines : inverse l'ordre des lignes ✓ 2026-07-28
- [x] csv-swap-rows : échange deux lignes dans un CSV ✓ 2026-07-28
- [x] text-dedup-words : supprime les mots dupliqués consécutifs ✓ 2026-07-28

## Vague 180 — CLI Tools (utilitaires)
- [x] json-escape : échappe/ déséchappe les caractères JSON dans un texte ✓ 2026-07-28
- [x] csv-to-markdown-table : convertit un CSV en tableau Markdown ✓ 2026-07-28
- [x] text-troncat : tronque un texte avec ajout de ... si dépassement ✓ 2026-07-28

## Vague 181 — CLI Tools (utilitaires)
- [x] base32-cli : encode/decode Base32 et Base32hex (RFC 4648) en terminal ✓ 2026-07-28
- [x] csv-to-toml : convertit un CSV en tableau TOML (array-of-tables, types, JSON wrapper) ✓ 2026-07-28
- [x] unix-timestamp-cli : conversion Unix timestamp ⇆ dates humaines (seconds/millis, timezone, JSON) ✓ 2026-07-28
- [x] jwt-decode-cli : décode et inspecte des JSON Web Tokens (header, payload, expiry, claims) ✓ 2026-07-28
- [x] yaml-to-go-struct : génère des structs Go depuis du YAML (nested, tags yaml/json, omitempty) ✓ 2026-07-28
- [x] toml-to-env : convertit TOML en .env (flatten, prefix, export, JSON) ✓ 2026-07-28
- [x] uuid-batch-gen : génère des UUIDs v4 en batch (count, formats, JSON) ✓ 2026-07-28
- [x] json-to-csv : convertit JSON array en CSV (flatten dot-notation, sélecteur de colonnes) ✓ 2026-07-28
- [x] sql-from-json : génère CREATE TABLE + INSERT SQL depuis JSON (types inférés, dialectes) ✓ 2026-07-28
- [x] ansi-to-html : convertit sortie terminal ANSI colorée en page HTML (dark/light themes) ✓ 2026-07-28

## Vague 182 — CLI Tools (git & data)
- [x] git-author-activity : analyse l'activité par auteur (commits, lignes, fichiers, jours actifs, histogramme) ✓ 2026-07-28
- [x] yaml-deep-merge : fusion profonde multi-fichiers YAML avec reporting des overrides et concat-arrays ✓ 2026-07-28
- [x] http-mock-server : serveur HTTP de mock piloté par spec JSON/YAML (status, headers, delay, hot-reload) ✓ 2026-07-28
- [x] ip-subnet-calc : calculateur sous-réseaux IPv4/IPv6 (broadcast, wildcard, hosts, split N subnets, JSON) ✓ 2026-07-28
- [x] openapi-validate : validateur statique de specs OpenAPI 3.x (erreurs, warnings, refs, $refs, JSON/strict) ✓ 2026-07-28
- [x] cron-describe : explique une expression cron en langage humain et liste les prochaines exécutions ✓ 2026-07-28
- [x] git-repo-summary : résumé haut niveau d'un repo git (branches, tags, activity, top authors, extensions) ✓ 2026-07-28

## Vague 183 — CLI Tools (utilitaires)
- [x] wait-for-it : attend qu'un port TCP ou une URL HTTP soit prêt (CI, entrypoints) ✓ 2026-07-28
- [x] env-set : édite des fichiers .env (set/get/unset/list) en préservant commentaires ✓ 2026-07-28
- [x] csv-stack : concaténation verticale de CSV à headers différents (union/intersect, source) ✓ 2026-07-28
- [x] json-unflatten : inverse de json-flatten (keys dotted -> objets imbriqués, arrays) ✓ 2026-07-28
- [x] text-template : rend des templates {{ keys }} depuis env/.env/JSON ✓ 2026-07-28
- [x] git-quick-stats : snapshot repo (auteurs, activity jours/heures, hotspots) ✓ 2026-07-28

## Vague 184 — CLI Tools (utilitaires)
- [x] percent-calc : calculs de pourcentages (of, change, part, vat, increase, decrease) ✓ 2026-07-28
- [x] duration-parse : parse des durées humaines (2h30m, 90min, 1d4h) vers secondes/ISO 8601/JSON ✓ 2026-07-28
- [x] number-format : formate des nombres (séparateurs milliers, décimales, notation scientifique, arrondi) ✓ 2026-07-28
- [x] roman-numerals : convertisseur nombres ⇆ chiffres romains (1..3999, batch, JSON) ✓ 2026-07-28
- [x] csv-max-min : valeurs min/max des colonnes numériques d'un CSV (fichier/stdin, JSON) ✓ 2026-07-28

## Vague 185 — CLI Tools (utilitaires)
- [x] csv-median : calcule la médiane des colonnes numériques d'un CSV ✓ 2026-08-01
- [x] text-strip-html : retire les balises HTML d'un texte (entités décodées) ✓ 2026-08-01
- [x] json-sort-keys : trie récursivement les clés d'un document JSON ✓ 2026-08-01
- [x] file-line-diff-count : compte les lignes ajoutées/supprimées entre deux fichiers ✓ 2026-08-01
- [x] text-capitalize-words : met en majuscule la première lettre de chaque mot ✓ 2026-08-01

## Vague 186 — CLI Tools (utilitaires)
- [x] csv-average : calcule la moyenne des colonnes numériques d'un CSV ✓ 2026-08-01
- [x] text-remove-diacritics : retire les accents et diacritiques d'un texte ✓ 2026-08-01
- [x] json-keys-list : liste toutes les clés d'un document JSON (chemins dot) ✓ 2026-08-01
- [x] text-nato-alphabet : épelle un texte avec l'alphabet phonétique OTAN ✓ 2026-08-01
- [x] csv-sum-columns : somme des colonnes numériques d'un CSV ✓ 2026-08-01

## Vague 187 — CLI Tools (texte & CSV)
- [x] text-collapse-spaces : réduit les espaces multiples à un seul ✓ 2026-08-01
- [x] csv-column-rename : renomme des colonnes dans un CSV (OLD=NEW) ✓ 2026-08-01
- [x] text-rot13 : applique le chiffrement ROT13 (ou ROT-n) ✓ 2026-08-01
- [x] line-sort-unique : trie les lignes et supprime les doublons (sort -u) ✓ 2026-08-01
- [x] text-strip-punctuation : retire la ponctuation d'un texte ✓ 2026-08-01

## Vague 190 — CLI Tools (CSV & texte)
- [x] csv-column-count : compte les colonnes d'un CSV (détection lignes irrégulières, --strict, JSON) ✓ 2026-08-01
- [x] path-base-dir : découpe des chemins en dirname/basename (+ extension, stdin, JSON) ✓ 2026-08-01
- [x] json-depth : profondeur maximale d'imbrication d'un document JSON (--max lint, exit 2) ✓ 2026-08-01
- [x] text-squeeze-blank-lines : réduit les suites de lignes vides à une seule (cat -s, --remove-all) ✓ 2026-08-01
- [x] csv-row-numbers : préfixe un CSV avec une colonne de numéros de ligne (id, --start, JSON) ✓ 2026-08-01

## Vague 189 — CLI Tools (CSV & texte)
- [x] csv-last-rows : garde les N dernières lignes de données d'un CSV (tail, header conservé, JSON) ✓ 2026-08-01
- [x] text-snake-case : convertit du texte/identifiants en snake_case (camelCase, kebab, espaces) ✓ 2026-08-01
- [x] text-camel-case : convertit du texte/identifiants en camelCase ou PascalCase ✓ 2026-08-01
- [x] csv-delimiter-detect : détecte le délimiteur d'un CSV (virgule, point-virgule, tab, pipe) ✓ 2026-08-01
- [x] text-wrap-width : re-wrap les paragraphes d'un texte à une largeur fixe ✓ 2026-08-01

## Vague 191 — CLI Tools (texte & env)
- [x] text-line-length : rapport de longueur par ligne + stats (min/max/avg), lint --max/--min exit 2, JSON ✓ 2026-08-01
- [x] text-extract-emails : extrait les adresses email d'un texte (unique, sort, count, domains, check) ✓ 2026-08-01
- [x] text-extract-urls : extrait les URLs http/https/ftp (filtre scheme/host, hosts/schemes, check, JSON) ✓ 2026-08-01
- [x] csv-row-sample : échantillonnage aléatoire de N lignes d'un CSV (seed reproductible, pourcentage, JSON) ✓ 2026-08-01
- [x] text-hash-lines : hash de chaque ligne (md5/sha1/sha256...), dedup, dup-check exit 2, JSON ✓ 2026-08-01
- [x] env-sort : trie les clés d'un .env en préservant commentaires et groupes (--check, --in-place) ✓ 2026-08-01

## Vague 188 — CLI Tools (CSV & texte)
- [x] csv-count-values : fréquence des valeurs d'une colonne CSV ✓ 2026-08-01
- [x] csv-first-rows : garde les N premières lignes de données d'un CSV (head) ✓ 2026-08-01
- [x] text-trail-space : supprime ou rapporte les espaces en fin de ligne ✓ 2026-08-01
- [x] json-array-wrap : encapsule/normalise/unwrap une valeur JSON en array ✓ 2026-08-01
- [x] text-double-space : convertit entre interligne simple et double ✓ 2026-08-01
