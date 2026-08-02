# OpenClaw — File d'attente des outils

## Vague 471 — CLI Tools (validation ISBN, IPv6 RFC 5952, barcodes GS1, NanoID, portabilité fichiers)
- [x] isbn-validate : valide les check digits ISBN-10 (mod-11) / ISBN-13 (mod-10 GS1), tolère tirets/espaces, --convert-10-to-13 (préfixe 978), --check/--require-valid/--require-min-valid/--require-all-isbn13 exit 2 CI, --invalid-only, --json, batch args/file/stdin ✓ 2026-08-02
- [x] ipv6-compress : normalise IPv6 en forme compressée RFC 5952 / expand 8 groupes / canonical (IPv4 embarqué ::ffff:192.0.2.1, /prefix réseau host-bits reset, mixte-case), --check exit 2 lint, --require-all-valid/--require-min-changes, --json ✓ 2026-08-02
- [x] ean13-check : valide EAN-13 / UPC-A / EAN-8 (mod-10 poids 1/3 alternés), --complete calcule le check digit manquant 7/11/12 digits, --check/--require-valid/--require-min-valid exit 2 CI, --invalid-only, --json ✓ 2026-08-02
- [x] nanoid-gen : génère des IDs NanoID (secrets.SystemRandom, alphabet A-Za-z0-9_- taille 21 default, presets url/numeric/hex/lowercase/uppercase/no-lookalike, custom alphabet, --size/--count, --seed reproductible test only, sampling rejection pour uniformité, --collision-check/--require-min-unique/--require-length exit 2, --json) ✓ 2026-08-02
- [x] filename-portability-check : audit portabilité fichiers Windows/macOS/Linux (chars illégaux <>:"/\|?*, noms réservés CON/PRN/AUX/NUL/COM1-9/LPT1-9 + extension, trailing dot/space, leading space, ctrl chars U+0000-1F/7F, >255 octets, NFC/NFD mismatch, collisions case-fold + NFC entre noms, --no-unicode/--no-collisions, --max-issues N, --json) ✓ 2026-08-02

## Vague 470 — CLI Tools (fan-out JSONL, stripes ANSI, justification CSV, intersection lignes, translittération)
- [x] json-lines-burst : partitionne un flux JSONL par clé (chemins pointillés, sanitise les noms de fichiers, seau null, --max-files 512 sécurité, --dry-run, --skip-invalid, gates require-min-lines/files exit 2 CI, --json) ✓ 2026-08-02
- [x] text-ansi-stripe : alterne les couleurs ANSI par ligne/groupes (16 couleurs fg/bg 90-97/100-107, --cycle N, --skip-empty, --no-reset, --check codes existants exit 2, --json) ✓ 2026-08-02
- [x] csv-row-justify : justifie les champs d'un CSV (auto numériques à droite/texte à gauche, sniff délimiteur, display width CJK, --columns nom/index, --char custom, --check idempotent exit 2, gates require-min-rows, --json) ✓ 2026-08-02
- [x] file-line-intersection : opérations ensemblistes sur lignes (intersection défaut/union/difference A-B, 2+ fichiers, casefold Unicode + first-seen order, --strip/--ignore-case/--count-only, gates require-min/max exit 2 CI, --json) ✓ 2026-08-02
- [x] text-unidecode-lite : translittération Unicode → ASCII locale (NFKD + overrides FR/DE/ES/PL/CZ/Nordic/GR : ß→ss, ł→l, ø→o, œ→oe, €→EUR, --strict remplace ? --replacement, --char-map custom, --check exit 2 non-ASCII, --json) ✓ 2026-08-02

## Vague 469 — CLI Tools (base58, versions UUID, temps de lecture, tables Markdown, résumé réseaux)
- [x] base58-encode : encode/décode base58 alphabet Bitcoin en batch (bytes ou entiers --integer, --per-line multi-entrées, préfixe zéro -> '1', --check round-trip, gates require-min/failmax exit 2 CI, --json) ✓ 2026-08-02
- [x] uuid-version-report : classifie et valide les UUIDs (versions 1-8, variantes RFC 4122/NCS/Microsoft/réservée, nil/max, --decode extrait timestamp v1/v7, batch args/file/stdin, --check exit 2 sur invalide, gates require-all-valid/min-valid/forbid-version exit 2 CI, --json) ✓ 2026-08-02
- [x] text-reading-time : estime temps de lecture d'un texte (WPM configurable 200 défaut, comptage mots Unicode/apostrophes, poids --figure/--code, plafond images --max-figure, --raw sans images/code, stats chars/words/sentences, gates require-max-minutes/min exit 2 CI, --json) ✓ 2026-08-02
- [x] markdown-table-format : normalise et aligne les tables pipe Markdown (largeurs display Unicode, padding, séparateur régénéré avec marqueurs :-- :-: --:, --colon-style, exit 2 si déjà normalisée en mode --check, fences ignorées, --in-place, gates require-min-tables exit 2 CI, --json) ✓ 2026-08-02
- [x] ip-network-summary : résumé d'un réseau IPv4/IPv6 (wildcard inverse, ip_class A/B/C/D/E, scopes is_private/global/loopback/link_local/multicast/reserved selon stdlib, hostmask, sous-réseaux --split N, agrégation --aggregate de réseaux adjacents, gates require-private/public/none-conflicting exit 2 CI, --json) ✓ 2026-08-02

## Vague 468 — CLI Tools (morse, JSON Pointer, diff JSONL, lignes vides, MinHash)
- [x] text-to-morse-code : encode/décode morse international (A-Z 0-9 ponctuation, séparateur mots /, --decode/--lower, --ignore-unknown, batch args/file/stdin, gates require-min-words/require-all-decoded exit 2 CI, --json) ✓ 2026-08-02
- [x] json-pointer-get : résout des pointers RFC 6901 sur JSON/JSONL (échappements ~0 ~1, indices tableaux, --jsonl multi-lignes, --values-only, échec pointer = exit 2, gates require-resolved/require-all-resolved/require-eq PTR VALUE exit 2 CI, --json) ✓ 2026-08-02
- [x] jsonl-key-diff : diff structure clés entre 2 streams JSONL (only-A/only-B/shared, --values join sur --key champ ou n° ligne, exit 2 sur différence, --exit-zero, gate require-identical exit 2 CI, --json) ✓ 2026-08-02
- [x] file-empty-line-report : localise lignes vides/whitespace dans fichiers (comptes, ratio, runs consécutifs, walk arborescence --ext, --list/--runs, blank-only strict, gates require-no-blanks/max-ratio/max-run exit 2 CI, --json) ✓ 2026-08-02
- [x] text-min-hash-sim : similarité approximative MinHash entre lignes/documents (shingles mots/caractères --n, k permutations --perms, --compare A vs B, --threshold/--top, gates require-max-sim/min-pairs exit 2 CI, --json) ✓ 2026-08-02

## Vague 467 — CLI Tools (drop colonnes CSV, table fixed-width, comptage types JSON, ordinaux anglais, profondeur chemins)
- [x] csv-except-columns : garde toutes les colonnes SAUF celles nommées/indexées (1-based/négatifs, globs --glob, --ignore-case, --keep-first N raccourci, délimiteur sniffé, gates check/require-remaining/require-dropped exit 2 CI, --json) ✓ 2026-08-02
- [x] csv-to-fixed-width : rend un CSV en table texte alignée (largeurs display-width Unicode CJK×2/combinaison 0, numériques à droite par défaut/--all-left/--all-right, --separator-rule, --max-col-width ellipse …, gates --max-width/--require-rows exit 2 lint terminal, --json) ✓ 2026-08-02
- [x] json-type-count : histogramme des types de valeurs JSON (object/array/string/integer/float/boolean/null + clés totales, JSONL auto-détecté/--jsonl, --paths profondeur max, gates require-type/require-min/require-max T N exit 2 CI, ex. interdire null, --json) ✓ 2026-08-02
- [x] text-ordinal-numbers : cardinaux ⇆ ordinaux anglais (forms courtes 1st/21st/112th avec règle 11-13, forms longues first/twenty-first/one millionth, --parse inverse y compris hundred/thousand, batch args/file/stdin, gate --require-min exit 2 CI, --json) ✓ 2026-08-02
- [x] file-path-depth-report : profondeur d'imbrication des chemins (walk arborescence base auto, paths args/--file/stdin, --base explicite, --files-only/--dirs-only, --histogram barres, --list/--top N, paths hors base skippés comptés, gates require-max-depth/min-depth/min-paths exit 2 CI, --json) ✓ 2026-08-02

## Vague 466 — CLI Tools (filtre lignes CSV, chart braille, shuffle clés JSON, case folding, inventaire dossiers)
- [x] csv-row-matcher : filtre les lignes CSV par conditions sur colonnes (eq/ne/contains/regex/gt/ge/lt/le/empty/not-empty/in-set, colonnes nom/index, --match all/any, --invert, --ignore-case/--stripws, délimiteur sniffé, --check/--require-min/max/none exit 2 CI, --json) ✓ 2026-08-02
- [x] text-braille-chart : rend des séries numériques en charts braille 2x4 dots (32 lignes de résolution par défaut, resample moyennant par colonne --width max 200, --extract depuis texte libre, échelle --min/--max, --stats, gates require-min-points/range/max exit 2 CI, --json) ✓ 2026-08-02
- [x] json-keys-shuffle : mélange/trie/inverse l'ordre des clés JSON (seed reproductible, --top-level, --reverse/--sort, JSONL --jsonl, --check exit 2 si ordre modifié, gate --require-key exit 2, --compact, --json) ✓ 2026-08-02
- [x] text-fold-case : Unicode case folding pour matching insensible à la casse (str.casefold: Straße↔STRASSE, sigma final, ligatures, --compare A B exit 0/2, --normalize NFC..NFKD, --lower, --unique/--sort, --in-place, gates check/require-changes/require-none-changed exit 2 CI, --json) ✓ 2026-08-02
- [x] dir-entry-filter : inventaire et filtre d'entrées de répertoires (types f/d/l, globs --pattern/--exclude, --ext, tailles min/max-unités, âge --days/--older-than, --depth, hidden, tris name/size/mtime/path, --relative/--limit, gates require-min/max/none exit 2 CI, --json) ✓ 2026-08-02

## Vague 465 — CLI Tools (stats CSV, entropie texte, flatten JSON, anagrammes, expansion CIDR)
- [x] csv-column-stats : stats descriptives des colonnes numériques CSV (count/sum/min/max/mean/médiane, quartiles Q1/Q3 interp. linéaire, variance+stdev pop/échantillon, délimiteur sniffé, colonnes nom/index, --no-header, gates --min/--max/--require-mean exit 2, --json) ✓ 2026-08-02
- [x] text-entropy-meter : entropie de Shannon d'un texte (global ou --by-line avec min/mean/max, modes chars/--bytes UTF-8, --max-theoretical alphabet observé, redondance, args positionnels multi-strings, gates --require-min/--require-max exit 2, --json) ✓ 2026-08-02
- [x] json-flatten-paths : aplatit JSON en paires path=value (styles dot/bracket, tableaux [N], objets/tableaux vides préservés, --types, --separator, JSONL auto/--jsonl, null exclus par défaut/--include-null, gates --require-path/--require-min/--require-max exit 2, --json) ✓ 2026-08-02
- [x] text-anagram-find : détecte et groupe les anagrammes (normalisation case/NFD accents/non-lettres, args positionnels mots si fichier inexistant, --pairs A B exit 2 sur non-anagramme, --min-length/--min-size/--all, tri size/key/alpha, gates require-min/max-groups/require-group exit 2, --json) ✓ 2026-08-02
- [x] ip-cidr-expand : développe CIDR IPv4/IPv6 en listes d'adresses (strict=False host-bits, --exclude-hosts réseau+broadcast IPv4 hors /31 //32, --max cap sécurité 65536 défaut, --limit tronqué, --count-only/--sum, --file avec # commentaires, gates require-total/min-total/max-total exit 2, --json) ✓ 2026-08-02

## Vague 464 — CLI Tools (base32 batch, JSONL pretty, comptage voyelles, positions tokens, normalisation EOL)
- [x] text-base32-decode : decode/encode base32 RFC 4648 en batch (args/--file/--stdin, padding toléré ou --require-padding, --allow-lower, UTF-8 ou fallback hex, --unique, gates --require-min/--require-none-failed exit 2/1, --json) ✓ 2026-08-02
- [x] json-lines-pretty : pretty-print/compact/lint de flux JSONL (indent espaces/tabs, --compact, --sort-keys, lignes vides tolérées, --skip-invalid, lint --check exit 2, gates require-valid-min/none-invalid/line exit 2 CI, --json) ✓ 2026-08-02
- [x] text-vowel-consonant-count : compte voyelles/consonnes/chiffres/espaces/autres par ligne + totaux (mode --extended voyelles accentuées via NFD, ratio voyelles/lettres, --no-list, gates require-min-vowels/consonants/words/vowel-ratio exit 2 CI, --json) ✓ 2026-08-02
- [x] text-token-position : liste les tokens avec leur position ligne/colonne (regex custom /--words/--numbers, 1-based par défaut/--zero-based, --unique première occurrence, --values-only, gates require-min-tokens/require-token répétable exit 2 CI, --json) ✓ 2026-08-02
- [x] csv-eol-normalize : normalise les fins de ligne CSV/texte vers LF ou CRLF (détection lf/crlf/cr/mixed/none, --in-place multi-fichiers ou stdout, --ensure-final-newline, binaires rejetés/--binary-safe, lint --check exit 2, gates require-none-changed/changed-min exit 2 CI, --json) ✓ 2026-08-02

## Vague 463 — CLI Tools (remplissage CSV, classes de caractères, histo tailles fichiers, format JSON, wrap colonnes)
- [x] csv-column-fill : remplit les cellules vides d'un CSV (mode down forward-fill / constant / static, --columns par nom ou index, --keep-leading-empty avec --value, délimiteur sniffé, --no-header, gates check/require-filled-min/require-empty exit 2 CI, --json) ✓ 2026-08-02
- [x] text-char-classify : compte les caractères par classe par ligne (alpha/digit/space/punct/other via str.is* et catégories Unicode P*, flag ascii par ligne, totaux fusionnés, --keep-newlines, gates require-ascii/class min/forbid-class/require-min-total exit 2 CI, --json) ✓ 2026-08-02
- [x] file-size-histogram : histogramme tailles de fichiers sous une arborescence (modes exponential puissances de 2 / log10 / linear --bucket-size, barres ASCII, filtre min/max-size, exclude-dir, symlinks non suivis, gates require-min/max-files/no-empty/total-gte/lte exit 2 CI, --json) ✓ 2026-08-02
- [x] json-format-cli : ré-indente/canonicalise documents JSON et JSONL (indent espaces ou tabs, --compact, --sort-keys récursif, --canonical RFC-style compact+sorted+UTF-8, --no-ensure-ascii, validation améliorée ligne/col, gates require-min/max-depth/min-keys/forbid-key exit 2 CI, --json) ✓ 2026-08-02
- [x] text-wrap-columns : coupe les lignes en dur à une largeur fixe (modes word (défaut) / any strict, expansion tabs avant wrap, préserve lignes vides et trailing-newline, --no-split-long, --check lint exit 2, gates require-max-line/lines-changed exit 2 CI, --json) ✓ 2026-08-02

## Vague 462 — CLI Tools (pivot CSV, schéma JSON inféré, fréquences mots, ancres MD, hex batch, texte XML)
- [x] csv-column-pivot : pivot long→large d'un CSV (groupe/colonne/valeur par nom ou index, agrégations sum/mean/min/max/count/first/last/list, --fill, sniff délimiteur, gates require-rows/max-pivot-cols exit 2, --json) ✓ 2026-08-02
- [x] json-schema-guess : infère un JSON Schema draft 2020-12 depuis un document ou flux JSONL (types, min/max numeric, formats string date-time/date/uuid/email/uri/ipv4, required = clés présentes partout, types mixtes en array, --no-format, --compact, gates require-property/type exit 2) ✓ 2026-08-02
- [x] text-word-frequency-csv : histogramme de fréquence des mots en CSV (tokenization Unicode, stop words EN embarqués + fichier custom, --min-len/--keep-case/--top/--sort, rank/word/count/share, gates require-min-words/max-distinct/top exit 2, --json) ✓ 2026-08-02
- [x] md-anchor-check : vérifie que les liens #anchor Markdown pointent vers un titre réel (anchors style GitHub, doublons -1/-2, liens cross-fichiers, fences ignorées, ref-defs, %20, exit 2 sur anchor cassé, --json) ✓ 2026-08-02
- [x] text-url-decode-batch : percent-decode/encode en batch URLs/strings (args/--file/stdin, --plus, --components scheme/host/path/query, --change-only, --safe encoder, gates check/require-decoded-min exit 2, --json) ✓ 2026-08-02
- [x] xml-plaintext-extract : extrait le texte brut de documents XML/HTML (script/style skip, entités résolues, --one-line/--separator, --lenient tag-stripping fallback, --check well-formed, --require-min-text exit 2, --json) ✓ 2026-08-02
- [x] text-hex-decode : decode/encode hex en batch (styles spaced/compact/0x/\xNN tolérés, UTF-8 ou escapes \xNN, --raw, --encode formats spaced/compact/c/python/0x, --lenient, gates require-text/min-decoded exit 2, --json) ✓ 2026-08-02

## Vague 461 — CLI Tools (comparaison hash, surlignage ANSI, blocs entropie, collisions noms, escape CSV, extensions manquantes)
- [x] file-hash-compare : compare deux fichiers ou arborescences par hash de contenu (md5/sha1/sha256/blake2b, symlinks ignorés, rapport same/differ/only-A/only-B, exit 2 sur diff, --json, --exit-zero) ✓ 2026-08-02
- [x] text-highlight : surligne les matchs regex/littéraux dans du texte avec couleurs ANSI (16 noms + code 0-255, --fixed, -i, --max N, --count, --check exit 2 lint) ✓ 2026-08-02
- [x] file-block-report : rapport par blocs de taille fixe (fill %, bytes distincts, entropie de Shannon) pour détecter régions sparse/padding/chiffré (--block-size, --limit, gates --max-blocks/--require-entropy-max/--min-fill-avg exit 2, --json) ✓ 2026-08-02
- [x] filename-collision : détecte collisions de noms sous normalisation (case fold, NFC/NFD/NFKC, cmf=NFC+fold) par dossier (--include-dirs, exit 2 sur collision, --json, --exit-zero) ✓ 2026-08-02
- [x] csv-row-escape : escape/unescape des lignes comme champ CSV unique (auto-quote delim/quote/newline, round-trip, --delimiter/--quote custom, --check exit 2, --json) ✓ 2026-08-02
- [x] file-extension-missing-report : liste les fichiers sans extension dans une arborescence (dotfiles comptés comme sans extension, symlinks ignorés, gates --require-none/--max/--min exit 2, --json) ✓ 2026-08-02

## Vague 460 — CLI Tools (vérification gitignore, retrait de quotes, normalisation quoting CSV, glob→regex, somme des tailles disque)
- [x] gitignore-check : évalue si des chemins seraient ignorés par un fichier .gitignore (sémantique git : glob * ? [cls], **/ segments, dir-only trailed /, ! negation règle tardive gagne, / anchored, # commentaires, args/--file/stdin, --list/--ignored-only/--not-ignored-only, --json, gates require-min/max/none exit 2 CI) ✓ 2026-08-02
- [x] text-unquote : retire les paires de quotes simples/doubles/backticks autour des lignes (mismatched/unterminated conservés, kind filtrable, --allow-escaped, --delimiter mode champs, --stripws, --check exit 2, --in-place, --json, gates require-stripped/none exit 2 CI) ✓ 2026-08-02
- [x] csv-escape : normalise le quoting des champs CSV (styles minimal RFC 4180/all/non-numeric coercit nombre/none avec --unsafe-replacement, délimiteur sniffé, --check exit 2, --in-place, --json, gates require-unsafe-max/unchanged exit 2 CI) ✓ 2026-08-02
- [x] path-to-regex : convertit un glob path-style en regex équivalente (* non-séparateur, ** segment multi-niveaux incluant zéro, ? un char, [classes], {a,b} alternance, \ escape, anchoré par défaut/--no-anchor, --separator, --test MATCH/NO-MATCH, --json, gates require-match-min/no-match exit 2 CI) ✓ 2026-08-02
- [x] file-size-total : somme les tailles de fichiers et dossiers récursivement (symlinks non suivis, args/--file/stdin, --human, --no-list/--summary, --json, gates require-min/max-bytes exit 2 CI, --ignore-missing) ✓ 2026-08-02

## Vague 459 — CLI Tools (cron expressions, versions semver, fins de ligne, slugs URL, subnets IPv4)
- [x] cron-expression-parse : parse et valide les cron 5 champs (*, N, A-B, A-B/S, */S, N/S, listes, noms jan-dec/sun-sat, 0/7=dimanche, OR sémantique dom+dow) + macros @hourly/@daily/@weekly/@monthly/@yearly/@annually/@midnight, expansion compacte par champ, description texte, --next N occurrences futures (--start ISO), batch args/--file/stdin avec # commentaires, gates require-min/valid/next exit 2 CI, --strict, JSON ✓ 2026-08-02
- [x] text-extract-semantic-versions : extrait versions semver 2.0.0 strictes (préfixe v, prerelease/build, leading-zeros rejetés, pas d'embed alphanum), --allow-partial 1.2→1.2.0, tri precedence semver complet (numeric<alpha, release>pre), --range '>=a,<b' multi-tokens, --unique/--counts/--sort/--latest-only/--stats, gates require-min/max/none/version/latest exit 2 CI, JSON ✓ 2026-08-02
- [x] file-line-ending-report : audit fins de ligne LF/CRLF/CR par fichier (styles lf/crlf/cr/mixed/empty/no-newlines/binary-NUL, flag no-final-newline, lignes partielles comptées), walk arborescence sans symlinks, filtres --ext/--pattern/--skip-dir, --only-style, gates require-style/no-mixed/final-newline/min/max-binary exit 2 CI, JSON ✓ 2026-08-02
- [x] text-slugify-convert : convertit textes en slugs URL (translitération NFKD + ß/æ/ø/ł, lowercase, --separator, stop words EN + customs, --max-length sur frontière mot, --unique suffixe -2/-3, --keep-case/--keep-unicode, args/--file/stdin, gates require-min/max-length/no-empty/unique/changed exit 2 CI, JSON) ✓ 2026-08-02
- [x] ipv4-subnet-containment : containment IP/subnet dans réseaux IPv4 (CIDR, ranges a-b mono-CIDR alignés, bare IP→/32, host-bits normalisés/--strict) modes contain/cover/overlaps (kinds duplicate/contains/partial), cibles args/--file/stdin, --exit-nonzero-out firewall-style, gates require-all/none/min-contained/no-overlaps/min-overlaps exit 2 CI, JSON ✓ 2026-08-02

## Vague 458 — CLI Tools (arborescence disque par groupe, explication de regex, diff/merge INI, DNS forward/reverse, distance de Levenshtein)
- [x] file-tree-size-report : audit espace disque d'une arborescence (groupes ext/dir, compte/total/mean/min/max human, filtres --ext/--no-ext/--skip-dir/--hidden/--min-size/--max-size, --largest N/--empty, symlinks ignorés, gates require-files-min/max/under-bytes/empty-zero/no-empty-sets/groups-min exit 2 CI, JSON) ✓ 2026-08-02
- [x] regex-explain-parse : explique une regex en langage clair (AST own-parser: littéraux, classes, ancres, quantifieurs {m,n}, groupes nommés/lookahead/flags inline, alternance) + matching de strings de test (spans, groupes capturés, -i/-m, --strings-file -, --max-matches), gates require-match-min/max/all/none/groups-min exit 2 CI, JSON ✓ 2026-08-02
- [x] ini-diff-merge : diff et merge INI section→clé (configparser sans interpolation, +/-/~ champ-level, --keep-case, modes --merge theirs/ours/add-missing, stdin '-', gates check/require-identical/min/max-changes/no-removed/section exit 2 CI, JSON) ✓ 2026-08-02
- [x] ip-hostname-resolve : résolution DNS forward A/AAAA et reverse PTR (auto-détection ipaddress, cibles CLI/--file/-, --family any/ipv4/ipv6, --timeout, --unique, elapsed ms, gates require-all/min/max-failed/none-failed exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-levenshtein-distance : distance de Levenshtein strings/fichiers/paires TSV (DP O(min) mémoire, ratio similarité, préfixe/suffixe communs, --trace opérations difflib, gates require-max-distance/min-similarity/identical/different exit 2 CI, JSON) ✓ 2026-08-02

## Vague 457 — CLI Tools (audit permissions POSIX, vendor OUI MAC, CSV diff, iTunes playlists, montants de devises)
- [x] file-perms-report : audit des bits de permission POSIX (symboliques+octal, setuid/setgid/sticky, world-writable WW, filtres --world-writable/--special-only/--mode/--pattern, --counts, gates require-none-world-writable/special/max/min exit 2 CI, JSON) ✓ 2026-08-02
- [x] mac-vendor-lookup : extrait MAC-48 (colon/hyphen/dotted Cisco, --bare optionnel) et résout le vendor via table OUI embarquée ~1200 entrées (Apple, Cisco, Dell, Huawei, Intel, Raspberry, VMware, Espressif, Netgear, TP-Link, Belkin, Xiaomi, Samsung, Google, Amazon, Sony, LG...), --lookup unique, flag locally-administered, --unknown, gates require-vendor/all-known/min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-diff : compare deux CSV clé→lignes (délimiteur sniffé, --key répétable, --ignore-column/--ignore-case, --no-header, --kind added/removed/changed, --summary, sortie -/+/~ champ-level, exit 2 sur diff par défaut, --exit-zero, gates require-identical/min/max-changes, JSON) ✓ 2026-08-02
- [x] itunes-playlist-parse : parse Library.xml iTunes/Music.app (plistlib stdlib, --list-playlists, --playlist NAME ou plus grande réelle, --all-tracks, formats table/csv/json/m3u, --fields sélecteur, --sort name/artist/play_count/rating/date_added/duration_ms, file:// décodé URL, ratings ★, durations h:mm:ss, --stats, gates require-min/max-tracks/playlist/none-missing exit 2 CI) ✓ 2026-08-02
- [x] currency-amount-extract : extrait montants de devises ($ pré/suffixe, codes ISO, 60+ monnaies dont BTC/ETH, séparateurs US 1,234.56 / EU 1.234,56 / FR 1 234,56 / CH 1'234.56, multiplicateurs k/M/mm/bn/b/t, --bare opt-in, --assume, --sum par devise, --amounts-only/--codes-only, --table, gates require-min/max/none/currency/sum-gte/single-currency exit 2 CI, JSON) ✓ 2026-08-02

## Vague 456 — CLI Tools (extrait checkboxes Markdown, parse ssh_config, shift headings Markdown, longueur octets UTF-8, audit longueur clés JSON)
- [x] markdown-checkbox-extract : extrait items checkbox Markdown (- [ ] / - [x], bullets - * + et numérotés, niveaux d'indentation, ratio done, listes --pending/--done/--all, gates require-done-ratio/min/no-pending exit 2 CI, JSON) ✓ 2026-08-02
- [x] ssh-config-parse : parse ssh_config (options globales, blocks Host, Include, shlex) et résout les options effectives par alias avec glob, --names-only, gates require-min-hosts/host exit 2 CI, JSON) ✓ 2026-08-02
- [x] markdown-heading-shift : décale les niveaux de titres ATX (# ⇆ ######, --delta ±N, clamp par défaut/--no-clamp/--strict, fences code préservées, --in-place, gates require-shifted/no-clamped exit 2 CI, JSON) ✓ 2026-08-02
- [x] utf8-byte-length : rapporte octets vs caractères UTF-8 par ligne (classes 1-4 octets/char, flag multibyte, --summary/--strings, gates require-max-bytes/ascii exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-key-length-audit : audit des longueurs de clés JSON (récursif objets+arrays, stats min/max/moyenne chars+bytes, --sort/--top/--paths, gates require-max/min/count exit 2 CI, JSON) ✓ 2026-08-02

## Vague 455 — CLI Tools (similarité Jaccard/Dice, sparklines Unicode, tri IPv4, nearest xterm-256, valeurs lettres mots)
- [x] text-jaccard-similarity : coefficients Jaccard / Dice-Sorensen / overlap entre textes (tokens mots/n-grammes mots ou caractères --chars --n, mode paire/--pairs toutes lignes triées, --min-similarity filtrage, gates require-min/max-similarity exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-sparkline : rend des séries numériques en sparklines Unicode ▁▂▃▄▅▆▇█ (--min/--max échelle, --width sous-échantillonnage moyennes, --extract numéraux dans texte libre, --stats, gates require-range/max-under/min-over exit 2 CI, JSON) ✓ 2026-08-02
- [x] ipv4-sort : trie les lignes par IPv4 octet-aware (9<10 contrairement à sort lexico, CIDR triés réseau puis préfixe croissant, lignes sans IP droppables/défaut début, --unique/--strict/--check exit 2 CI, JSON) ✓ 2026-08-02
- [x] ansi-nearest-color : mappe couleurs RGB/hex vers index xterm-256 le plus proche (palette 256 complète, distance redmean pondérée, fragments SGR 38;5;N/48;5;N, gates require-max-delta/index exit 2 CI, JSON) ✓ 2026-08-02
- [x] word-letter-value : score les mots par positions alphabétiques A=1..Z=26 (somme/produit/racine numérique, --breakdown, --anagram-value groupes même multiset, --pair-sum S, gates require-sum/min/max/root exit 2 CI, JSON) ✓ 2026-08-02

## Vague 454 — CLI Tools (factorisation entiers, interleave lignes, notation A1 Excel, contraste WCAG, semaines ISO)
- [x] int-factor-tools : factorisation premiers + gcd/lcm/diviseurs/totient (Miller-Rabin déterministe <2^64, batch stdin/--numbers, gates --require-prime-min/coprime/no-composites exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-interleave-lines : entrelace N fichiers ligne-à-ligne round-robin (--group K par bloc, --skip-empty) et inverse --split N vers stdout ou --out-dir/lane-N.txt (gates require-equal-lanes/lines exit 2 CI, JSON) ✓ 2026-08-02
- [x] excel-a1-notation : convertit références A1 ⇆ indices (lettres↔index XXE=XFD, ranges A1:C10 normalisés width/height/cells, markers absolus $, gates require-max-col/min exit 2 CI, JSON) ✓ 2026-08-02
- [x] color-contrast-ratio : ratio de contraste WCAG 2.x entre couleurs (hex 3/4/6/8, rgb()/rgba(), 16 nommées CSS, verdicts AA/AAA normal+large, multi-paires, gates --require/aa/aaa exit 2 CI, JSON) ✓ 2026-08-02
- [x] iso-week-date : dates calendaires ⇆ semaines ISO 8601 (to-week avec métadonnées, to-date YYYY-Www[-D] expansion 7 jours, forme compacte, années 53 semaines, gates require-week/min exit 2 CI, JSON) ✓ 2026-08-02

## Vague 453 — CLI Tools (fréquence d'octets, expansion ranges IPv4, CSV→dotenv, jours ouvrés, valeurs JSON par chemin)
- [x] text-byte-frequency-stats : histogramme des occurrences de chaque octet (classes control/printable/extended, barres ASCII, --top/--ascii-only/--no-list, gates --require-ascii/min-distinct/max-distinct exit 2 CI, JSON) ✓ 2026-08-02
- [x] ipv4-range-expand : expand ranges IPv4 en adresses (CIDR, a-b forme courte, wildcard suffixe, /32 implicite, --summary/--all/--count-only/--strict, gates --require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-columns-from-csv : convertit les colonnes d'un CSV en variables dotenv (header normalisé UPPER_SNAKE, prefixe row ROWn_, quoting shell POSIX, --no-row-prefix/--row N/--skip-empty, gates --require-min-rows/column exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-business-days-count : compte les jours ouvrés (lun-ven) entre les dates extraites d'un texte (ISO + slash ambigu DD/MM fall-back MM/DD, fichier --holidays ISO, --exclusive [start,end[ , --all-pairs, --sum, gates --require-min-days/min-pairs exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-extract-path-values : extrait les valeurs d'un JSON/JSONL à des dot-paths (indices numériques, wildcard `*` sur objets/listes, filtres --type, --unique/--counts/--paths, --strict JSONL, gates --require-min/max/count exit 2 CI, JSON) ✓ 2026-08-02

## Vague 452 — CLI Tools (extrait base64, doublons par hash, liens Markdown, conversion de casse, profil statistique)
- [x] text-extract-base64 : extrait et décode les blobs base64/base64url d'un texte (preview texte/flag binaire, --min-len 8, --decode/--raw-out DIR, --unique/--counts, gates --require-min/max/none/text exit 2 CI, JSON) ✓ 2026-08-02
- [x] file-duplicate-report : groupe les fichiers dupliqués par hash de contenu (sha256/sha1/md5/blake2b, pré-bucketing par taille, bytes gaspillés triés, --delete-extras/--dry-run, --min-size/--skip-dir, gates --require-none/max-wasted/min-groups exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-markdown-links : extrait liens/images/références Markdown (inline avec titres, images, refs résolues via définitions, autolinks, mailto, refs non résolues flaggées, --kind/--targets-only, gates --require-no-unresolved/--require-host exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-case-convert : convertit les identifiants entre kebab/snake/camel/pascal/screaming/dot (split frontières camel/Pascal, modes ligne/--tokens, --detect, --in-place, --check/--require-change exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-stat-profile : statistiques complètes d'un texte (lignes/mots/unicité/richesse vocab, classes de caractères Unicode, moyennes lignes/mots, top-N mots, --keys, gates --require-min/max-words/max-line-length/min-richness exit 2 CI, JSON) ✓ 2026-08-02

## Vague 451 — CLI Tools (extrait CIDR IPv4/IPv6, chiffres romains, morse, décode JWT, inventaire symlinks)
- [x] text-extract-cidr : extrait et valide les blocs CIDR IPv4/IPv6 d'un texte (réseau/broadcast/taille via ipaddress, host-bits flaggés, --version 4|6, --unique/--counts/--sort value, --strict-network/--require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-roman-numerals : extrait les chiffres romains d'un texte et les convertit (paires soustractives validées, formes malformées rejetées, --convert N vers romain, --min/--max, --unique/--counts/--sort, gates --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] morse-tool : encode/décode le morse international (lettres/chiffres/ponctuation, séparateurs --char-sep/--word-sep, glyphes Unicode alternatifs décodés, --lowercase, --check round-trip, --strict/--require-encoded/--require-decoded exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-jwt : extrait et décode les JWT d'un texte sans vérification de signature (header/payload claims, statut expiration --now, --claims/--header-only, --unique/--counts, gates --require-min/max/none/valid/--fail-on-expired exit 2 CI, JSON) ✓ 2026-08-02
- [x] file-symlink-report : inventorie les symlinks d'une arborescence sans les suivre (flags BROKEN/ABS/CHAIN, filtres --broken-only/--abs-only, --unlink-broken, gates --require-none/no-broken/no-absolute/max exit 2 CI, JSON) ✓ 2026-08-02

## Vague 450 — CLI Tools (fichiers vides, renommage séquentiel, hash multi-algos CSV, inventaire Unicode, bench lecture disque)
- [x] file-zero-byte-report : rapporte les fichiers de 0 octet dans une arborescence (filtres regex, --skip-dir, --delete/--dry-run, gates --require-none/min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] file-rename-sequential : renomme des fichiers en séquence numérotée (prefix/suffix/digits/start/step, tri name/mtime, dry-run par défaut, renommage 2 phases anti-collision, --check/--require-renames exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-multihash : hash les lignes/colonnes d'un CSV en multi-algorithmes (md5/sha1/sha256/blake2b, --columns, --key-column, --whole-file, --check-duplicate-rows exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-unicode-name-report : inventaire Unicode des caractères d'un texte (codepoint/nom/catégorie/combining, --summary, tris, gates --require-ascii/--require-category/--forbid-category exit 2 CI, JSON) ✓ 2026-08-02
- [x] file-read-speed : mesure le débit de lecture de fichiers (sequential/lines/random, --chunk-size/--runs/--seed, cache-drop hint, --min-mbps exit 2 CI, JSON) ✓ 2026-08-02

## Vague 449 — CLI Tools (JSON union, indentation audit, dotenv→TypeScript, CSV⇄JSONL round-trip, comptage commentaires)
- [x] json-array-merge-union : fusionne des arrays JSON en union stable dédupliquée (par identité ou --key dot-path, JSONL, gates CI, --stats/--json) ✓ 2026-08-02
- [x] text-space-indent-check : audite l'indentation d'un texte (tabs / mix / largeur impaire, --require-uniform, --json, exit 2 CI) ✓ 2026-08-02
- [x] env-generate-ts : génère un fichier .d.ts typé depuis un .env (quotes, export, gates --forbid-empty / --forbid-duplicates / --require-min, --json) ✓ 2026-08-02
- [x] csv-jsonl-roundtrip : convertit CSV ⇄ JSONL avec map de types réversible (inférence int/float/bool, --map-out / --map pour round-trip identique, gates CI) ✓ 2026-08-02
- [x] file-comment-lines-count : compte les lignes de commentaires par extension dans un arbre (#, //, /* */, --, REM, <!-- -->, --per-file, --top, --ext, --require-comments, JSON) ✓ 2026-08-02

## Vague 448 — CLI Tools (extrait coordonnées GPS, extrait numéros téléphone, extrait expressions temporelles, extrait couleurs CSS, extrait références packages)
- [x] text-extract-coordinates : extrait les coordonnées GPS (paires décimales, DMS compact, latitudes seules avec hémisphère, validation plages, --pairs-only/--lat-only/--decimal-only/--dms-only, --precision, --unique/--counts/--sort, gates exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-phone-numbers : extrait les numéros de téléphone (E.164, 00CC international, US groupé, FR groupé, extensions xNNN, normalisation E.164, --unique/--counts/--with-extension, gates exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-time-expressions : extrait les expressions temporelles (horloge 24h/12h validées, durées composées 1h30m sommées en secondes, périodes ISO 8601 PT...P, intervalles, filtre --kind répétable, --unique/--counts, gates exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-css-colors : extrait les couleurs CSS (hex 3/4/6/8 chiffres normalisés, rgb/hsl/hwb/lab/lch/oklab/oklch/color(), 148 nommées CSS4, --kind répétable, --normalize, --no-named, gates exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-package-refs : extrait les références de packages (npm@scope, pip=extras>=, cargo="", gem~>, archives .tar.gz, Docker :tag/@sha256, --kind répétable, --names-only/--format, gates exit 2 CI, JSON) ✓ 2026-08-02

## Vague 447 — CLI Tools (extrait params query URL, extrait hashes git, extrait tailles fichiers, mots-nombres→digits, répète première ligne)
- [x] text-extract-urls-params : extrait les paramètres query des URLs http(s) (clé=valeur dupliquées, percent-decode --no-decode, --key/--exclude-key répétables, --keys-only/--values-only/--counts, gates exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-git-hashes : extrait les SHAs git (full 40 / abrégés 7-39 hex, filtre pure-digit, --full-only/--short-only, --prefix-match, --unique/--counts, gates exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-file-sizes : extrait les tailles humaines (1.5 MB SI / 2 GiB IEC / 512K shell) en bytes (--human, --raw, --total, filtres --min/--max-bytes, gates exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-number-words-to-digits : convertit les nombres en lettres anglaises en chiffres (twenty-one, a hundred, two million..., runs stoppés à la ponctuation, --list, --in-place, --check/--require-conversions exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-repeat-first-line : répète la première ligne (ou --line N / last) N fois (prepend/append/replace, --in-place, --check exit 2 CI, ligne hors-plage exit 2, JSON) ✓ 2026-08-02

## Vague 446 — CLI Tools (extrait MAC-48, extrait ports TCP/UDP, retire séquences ANSI, extrait entités HTML, inverse la casse)
- [x] text-extract-mac48 : extrait les adresses MAC-48 (colon/hyphen/dotted Cisco, --normalize/--lowercase, --vendor-oui, filtres multicast/broadcast, --unique/--counts, gates exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-port-numbers : extrait les ports TCP/UDP (8080/tcp, tcp/443, host:port, keyword port/listen/dport, filtres --proto/--well-known-only, --unique/--counts/--sort/--sum, gates exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-strip-ansi-sequences : retire les séquences ANSI (CSI/OSC/DCS/SGR/C1, --in-place, --check/--require-none/--require-removed exit 2 CI, --stats, JSON) ✓ 2026-08-02
- [x] text-extract-html-entities : extrait les entités HTML (&amp;/&#233;/&#xE9;, table HTML5, --kind répétable, --decode, --allow-unknown, --unique/--counts, gates exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-swapcase : inverse la casse de chaque lettre (swap Unicode, alternating spongebob, random --seed, --ascii, --in-place, --check/--require-changes exit 2 CI, JSON) ✓ 2026-08-02

## Vague 445 — CLI Tools (extrait flottants, run-length encode, arrondit floats JSON, extrait hash hex, retire commentaires C)
- [x] text-extract-floats : extrait les littéraux flottants d'un texte (décimaux/exposants/inf/nan, --precision, --unique/--counts/--sort, filtres min/max, gates exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-run-length : run-length encode/décode un texte (tokens Nx{char} par ligne, --min-run/--all, --decode, gates exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-round-floats : arrondit récursivement les floats d'un JSON à N décimales (--drop-float vers int, --in-place, gates exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-hashes : extrait les empreintes hex d'un texte (md5/sha1/sha224/sha256/sha384/sha512, kind long-hex, --kind filtre répétable, --lowercase/--unique/--counts, gates exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-strip-c-comments : retire les commentaires C/C++ // et /* */ (chaînes préservées, --keep-lines, --in-place, --check/--require-removed exit 2 CI, JSON) ✓ 2026-08-02

## Vague 444 — CLI Tools (csv supprime colonne par index, wrap conserve indentation, compte phrases, quoted-printable, extrait entiers)
- [x] csv-remove-column-by-index : supprime une colonne CSV par index (1-based/--zero-based, négatifs, --dry-run, --require-name/min-cols exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-wrap-indent : wrap chaque ligne à une largeur cible en préservant l'indentation (continuation indentée, --break-long, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-count-sentences : compte les phrases d'un texte (abréviations Mr./e.g., décimales, initiales, --list, --min-words, --require-min/max/exact exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-quoted-printable : encode/décode les lignes en quoted-printable (=XX hex UTF-8, espaces trailing, --encode-wsp, --require-encoded/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-integers : extrait les entiers décimaux/hex/octal/binaire d'un texte (0x/0o/0b, signes, --hex/--oct/--bin, --unique/--counts/--sort, filtres --min/--max, --sum, gates exit 2 CI, JSON) ✓ 2026-08-02

## Vague 443 — CLI Tools (tabs→espaces indentation, escape strings JSON, extrait key=value, ajoute colonne numéros, profondeur d'indentation)
- [x] text-space-indent-migrate : convertit les tabs d'indentation en espaces (--width, --inner, --in-place, --check exit 2 CI, gates require-conversions/none, JSON) ✓ 2026-08-02
- [x] json-string-escape : échappe/déséchappe les caractères spéciaux des strings d'un JSON (récursif, \n \t \uXXXX décodeur manuel, --require-changes/unchanged exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-kv-pairs : extrait les paires key=value d'un texte (séparateur custom, valeurs quotes, --pairs à plat, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-add-line-numbers : ajoute une colonne numéro de ligne aux rows CSV (--start/--step, front/--end, --skip-empty, --no-header, --check/--require-rows exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-indent-depth : rapporte la profondeur d'indentation par ligne (unités tab/--width espaces, remainder, --stats, --strict/--require-max-depth exit 2 CI, JSON) ✓ 2026-08-02

## Vague 442 — CLI Tools (csv mapping header→index, supprime lignes congruence, extrait accolades équilibrées, profil whitespace, fenêtres contexte grep)
- [x] csv-header-index-map : mappe les en-têtes CSV vers leurs indices colonnes (1-based/--zero-based, --invert, --delimiter, --only/--require-column répétables exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-delete-nth-line : supprime les lignes par congruence (i-offset) mod n (--every/--offset, --list, --in-place, --keep-blank, --check/--require-drops exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-braces : extrait les segments entre accolades équilibrées (pile réelle, --innermost/--outermost, --strip/--with-delim/--lines, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-whitespace-profile : profile le whitespace d'un texte (indent espaces/tabs/mixte, trailing, vides, stats largeurs, --require-indent-unit/no-tabs/no-trailing/max-mixed exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-window-lines : fenêtres de contexte glissantes autour de matchs regex (-B/-A/-C, fusion fenêtres, -n/-i/-v, --require-matches/none exit 2 CI, JSON) ✓ 2026-08-02

## Vague 441 — CLI Tools (csv garde colonnes par plages, garde lignes congruence, extrait parenthèses équilibrées, expansion ~ utilisateurs, compte caractères par ligne)
- [x] csv-slice-columns : garde un sous-ensemble de colonnes CSV par noms et/ou plages d'indices 1-based (négatifs, A-B / A:B, --invert, --require-columns/--check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-keep-nth-line : garde les lignes par congruence (i-offset) mod n (--every/--offset, --sample, --invert, --require-matches/--check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-parens : extrait les segments entre parenthèses équilibrées d'un texte (imbrication pile réelle, paires custom, --innermost/--outermost, --lines/--with-delim, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-expand-user-home : expanse les préfixes ~/ et ~user ligne par ligne (--separator champs, --strict compte inconnu, --check/~ restant, --require-expansions exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-lines-total-chars : rapport caractères par ligne + totaux (min/max/moyenne, --delimiter, --no-list, --require-max-total/max-line/lines exit 2 CI, JSON) ✓ 2026-08-02

## Vague 440 — CLI Tools (csv trie lignes par clés multi, extrait %, extrait ranges SemVer, json insère item array, ini liste sections)
- [x] csv-reorder-rows-by-key : trie les lignes CSV par une ou plusieurs colonnes clés (stable, --numeric/--desc/--ignore-case, --check exit 2 CI, --in-place, JSON) ✓ 2026-08-02
- [x] text-extract-percentages : extrait les littéraux pourcentage d'un texte (12.5%, virgule décimale, filtres --min/--max, --unique/--counts/--sort, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-semver-range : extrait les ranges SemVer npm-style (caret/tilde/comparator-set/hyphen/exact, classification --kind, --versions-only, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-insert-array-item : insère un item typé dans un array JSON à un dot-path (--index négatif ok, --dedupe-skip, --require-inserted exit 2 CI, --in-place, JSON) ✓ 2026-08-02
- [x] ini-list-sections : liste les sections INI avec comptes de clés (--with-counts, --include-default, --filter/--sort, --require-min/--require-section/--require-none exit 2 CI, JSON) ✓ 2026-08-02

## Vague 439 — CLI Tools (texte reformate paragraphes, extrait E.164, csv exige valeurs, json compacte nulls arrays, texte reformate ISO 8601)
- [x] text-reflow-paragraphs : reformate les paragraphes à une largeur cible (wrap/unwrap, --single-blank/--trim, --whitespace-blank, --in-place, --require-max-width/--require-paragraphs exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-e164 : extrait les numéros E.164 d'un texte (formes groupées, préfixe 00→+, strict 15 chiffres/--lenient, --unique/--counts/--sort, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-require-values : exige des valeurs non vides dans des colonnes CSV (gate avec liste violations colonne+ligne, --drop nettoie, --stripws, --require-max-violations exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-array-compact : supprime les éléments null des arrays JSON (--recursive, --empty-strings/--empty-containers, --drop-empty-arrays, --compact, --in-place, --require-removed/--require-none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-format-iso8601 : parse et reformate dates/heures/datetimes ISO 8601 ligne par ligne (--tz-utc suffixe Z, --format strftime, --strict, --require-min/--require-none-invalid exit 2 CI, JSON) ✓ 2026-08-02

## Vague 438 — CLI Tools (texte dédup mots par ligne, parse URI/URL, compte paragraphes, extraction e-mails, rapport lignes vides)
- [x] text-dedupe-words : supprime les mots en double de chaque ligne (première occurrence conservée, --ignore-case, --in-place, --check/--require-collapses exit 2 CI, JSON) ✓ 2026-08-02
- [x] uri-parse : parse des URIs/URLs en composants (scheme/host/port/user/path/query/fragment, --decode, --query éclaté, --require-scheme/--require-absolute exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-count-paragraphs : compte les paragraphes d'un texte (runs de lignes non vides, --whitespace-blank, --list/--words, --require-min/max/exact exit 2 CI, JSON) ✓ 2026-08-02 (repo existant, réécriture CLI propre)
- [x] text-extract-emails : extrait les adresses e-mail d'un texte (--unique/--counts/--sort, --domain/--exclude-domain, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02 (repo existant, réécriture CLI propre)
- [x] text-blank-line-report : rapport des lignes vides et runs de lignes vides (positions 1-based, --whitespace-blank, --list, --require-max-run/max-total/none exit 2 CI, JSON) ✓ 2026-08-02

## Vague 437 — CLI Tools (texte extrait SemVer, texte extrait IBAN mod-97, validateur cron 5 champs, fusion profonde JSON, score force mot de passe)
- [x] text-extract-semver : extrait les versions SemVer 2.0.0 d'un texte (pré-release/build, --unique/--counts/--sort précédence, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-iban : extrait et valide les IBAN ISO 13616 (mod-97, longueurs pays registre, formes espacées, --country, --with-invalid, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] cron-validate : valide des expressions cron 5 champs ligne par ligne (plages/listes/pas, noms JAN..DEC MON..SUN, --describe, --require-min/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-deep-merge : fusionne profondément plusieurs JSON (objets récursifs, stratégies arrays replace/concat/union, --check exit 2 CI, --compact, -o) ✓ 2026-08-02
- [x] password-strength : score la force de mots de passe 0-100 (heuristique locale, --min-score exit 2 CI, --feedback, --file/stdin, JSON) ✓ 2026-08-02

## Vague 436 — CLI Tools (texte extrait hostnames RFC1034, data URIs décodées, Accept-Language q-values, templates URI RFC6570, codes SWIFT/BIC ISO9362)
- [x] text-extract-hostnames : extrait les hostnames RFC 1034/1035 d'un texte (filtre --tld, --lowercase, --apex, --unique/--counts, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-data-uri : extrait et décode les data: URIs RFC 2397 (base64/percent, --mime-filter, --decode brut, --require-max-bytes/min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] accept-language-parse : parse les en-têtes Accept-Language (q-values triées, best-match RFC 4647 avec fallback préfixe, --supported/--default, --strict, --require-match exit 2 CI, JSON) ✓ 2026-08-02
- [x] uri-template-expand : expanse les templates URI RFC 6570 ({var}/{+var}/{#}/{.}/{/}/{;}/{?}/{&}, percent-encoding, --var/--vars-json, --list-vars, --require-vars exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-swift-bic : extrait et valide les codes SWIFT/BIC ISO 9362 (structure banque/pays/lieu/agence, pays ISO 3166-1 vérifié, --country/--parts, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02

## Vague 435 — CLI Tools (texte extrait UUIDs, csv remplissage vers le haut, résumé codes ANSI, texte extrait ISBN validés, texte extrait IMEI Luhn)
- [x] text-extract-uuids : extrait les UUID d'un texte (filtre --version-filter 1-5, --normalize, --unique/--counts, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-fill-up : remplit les cellules CSV vides avec la valeur de la ligne du dessous (--columns nom/index, --in-place, --require-fills/--require-none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-ansi-summary : résume les séquences d'échappement ANSI d'un texte (comptes par séquence, décodage SGR en noms lisibles, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-isbn : extrait et valide les ISBN-10/13 d'un texte (checksum mod-11/mod-10, --type, --invalid, --counts, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-imei : extrait et valide les IMEI 15 chiffres d'un texte (Luhn, formes groupées, --tac, --invalid, --counts, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02

## Vague 434 — CLI Tools (texte couleurs hex CSS, json chemins booléens, années, masquage IPv4, grep CSV regex)
- [x] text-extract-hex-colors : extrait les couleurs hex CSS d'un texte (#rgb/#rgba/#rrggbb/#rrggbbaa, expansion, --unique/--counts/--top, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-list-bool-paths : liste les chemins dot-path des valeurs booléennes d'un JSON (--true-only/--false-only, --require-true-min/--require-false-max/--require-none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-years : extrait les années YYYY d'un texte (frontières de mots, plage --min-year/--max-year, --unique/--counts/--sort/--top, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-mask-ipv4 : anonymise les IPv4 d'un texte (octets conservés, token/zero, skip private/loopback, --in-place, --check/--require-masks/--require-none exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-keep-matching-rows : garde les lignes CSV dont une colonne matche des regex (OR/AND, --ignore-case/--full-match/--invert, --in-place, --check/--require-rows/--require-none exit 2 CI, JSON) ✓ 2026-08-02

## Vague 433 — CLI Tools (csv fusion colonnes fallback, json chemins iso8601, texte extrait FQDN, ROT5 digits, csv fusion clé première valeur non vide)
- [x] csv-coalesce-columns : fusionne des colonnes CSV avec fallback première non vide (--merge NEW=COL1,COL2, --in-place, --check, --require-min-fills exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-list-iso8601-paths : liste les chemins dot-path des strings ISO 8601 (date/time/datetime, Z normalisé, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-fqdn : extrait les FQDN d'un texte (--unique, --top, --public-suffix-verify allowlist TLD, --require-min/domain/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-rot5 : applique ROT5 aux chiffres d'un texte (auto-inverse, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-first-nonempty-per-key : fusionne les lignes CSV par clé gardant la première valeur non vide par colonne (--in-place, --check/--require-min-groups/--require-all-filled exit 2 CI, JSON) ✓ 2026-08-02

## Vague 432 — CLI Tools (csv drop colonnes par plages, json garde clés listées, texte extrait IPv6 validées, collapse lignes vides à N max, ROT18)
- [x] csv-drop-column-ranges : supprime des colonnes CSV par nom, index ou plages inclusives (A-B / A:B, négatifs, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-only-keys : ne garde que les clés racine listées d'un objet JSON (--require-all/--require-min exit 2 CI, --compact, JSON) ✓ 2026-08-02
- [x] text-extract-ipv6 : extrait les adresses IPv6 d'un texte validées via ipaddress (compressed/expanded, fe80::/global filters, zone ids, --require-min exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-collapse-blank-runs-max : collapse les runs de lignes vides à N max (--n, --whitespace-blank, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] rot18-tool : applique ROT18 (ROT13 lettres + ROT5 digits) à un texte (auto-inverse, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 431 — CLI Tools (csv échange colonnes clé/valeur, json trie clés racine seulement, texte extrait EUI-64, compte points de code Unicode, base64url batch fichiers)
- [x] csv-swap-key-value-columns : échange les colonnes clé et valeur d'un CSV clé→valeur (noms/indices, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-sort-keys-shallow : trie uniquement les clés du niveau racine d'un objet JSON (objets imbriqués intacts, --reverse/--ignore-case, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-mac-eui64 : extrait les identifiants EUI-64 d'un texte (--also-mac48, --normalize, --require-min exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-count-codepoints : compte les points de code Unicode d'un texte (--graphemes, --categories, --per-line, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] base64url-file-tool : encode/décode base64url en batch fichiers/arborescences (--out-dir miroir, --recursive, --check round-trip exit 2 CI, --dry-run, JSON) ✓ 2026-08-02

## Vague 430 — CLI Tools (csv trie lignes par colonne, json supprime clés vides, texte extrait adresses IP, détecte CRLF/LF mixte, texte lignes avec exactement N mots)
- [x] csv-sort-rows : trie les lignes d'un CSV par colonne (--col, --numeric, --desc, --in-place, --check exit 2 CI si pas trié, JSON) ✓ 2026-08-02
- [x] json-prune-empty : supprime les clés dont la valeur est null, "", [] ou {} d'un JSON (--recursive, --require-none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-ipv4 : extrait les adresses IPv4 d'un texte (--unique, --private-only, --require-min exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-line-ending-check : détecte les fins de ligne LF/CRLF/mixtes d'un fichier (--require-lf/--require-crlf exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-lines-word-count-exact : garde les lignes ayant exactement N mots (--n, --require-min exit 2 CI, JSON) ✓ 2026-08-02

## Vague 429 — CLI Tools (csv supprime colonnes dupliquées, json compte clés par chemin, texte lignes les plus longues, base58check décode, texte mots en minuscules seulement)
- [x] csv-drop-duplicate-columns : supprime les colonnes au contenu identique d'un CSV (toutes lignes, --keep-first, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-count-keys : compte le nombre de clés objets par chemin d'un JSON (dot-path, --top, --require-max exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-longest-lines : affiche les lignes les plus longues d'un texte (--top, --min-length, --require-max exit 2 CI, JSON) ✓ 2026-08-02
- [x] base58check-tool : encode/décode base58check (version byte + checksum double-SHA256, --decode, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-lowercase-only-words : garde les lignes ne contenant que des mots en minuscules (--invert, --require-min exit 2 CI, JSON) ✓ 2026-08-02

## Vague 428 — CLI Tools (csv compte valeurs par colonne, json chemins strings, texte mots les plus longs, détecte encodage BOM, texte lignes contenant nombre)
- [x] csv-value-counts : compte les occurrences de chaque valeur d'une colonne CSV (--col, --top, --require-max exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-list-string-paths : liste les chemins vers toutes les valeurs string d'un JSON (dot-path, --min-length, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-longest-words : liste les mots les plus longs d'un texte (--top, --min-length, --require-min exit 2 CI, JSON) ✓ 2026-08-02
- [x] detect-bom : détecte la présence d'un BOM (UTF-8/UTF-16LE/BE) dans un fichier (--require-none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-lines-containing-number : garde les lignes contenant au moins un nombre (--invert, --require-min exit 2 CI, JSON) ✓ 2026-08-02

## Vague 427 — CLI Tools (csv compte distincts, json chemins arrays, texte extraire URLs, hexdump octets, texte mots uniques)
- [x] csv-distinct-count : compte les valeurs distinctes par colonne d'un CSV (--col, --top, --require-max exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-list-array-paths : liste les chemins vers toutes les valeurs array d'un JSON (dot-path, --min-length, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-urls : extrait les URLs http(s) d'un texte (--unique, --domain, --require-min exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-to-hexdump : affiche un texte/fichier en hexdump classique (offsets, ascii, --width, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-unique-words : liste les mots n'apparaissant qu'une seule fois (hapax, --top, --require-min exit 2 CI, JSON) ✓ 2026-08-02

## Vague 426 — CLI Tools (texte compte underscores, csv garde colonnes non constantes, json chemins objets, uuencode, texte lignes sans ponctuation finale)
- [x] text-count-underscores : compte les underscores _ d'un texte (--per-line, --require-min/max, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-drop-constant-columns : supprime les colonnes ayant une seule valeur distincte (--constant-seuil N valeurs distinctes, --check exit 2 CI, JSON) ✓ 2026-08-02 (repo existant, ajout alias --max-distinct et --stripws pour compat)
- [x] json-list-object-paths : liste les chemins vers toutes les valeurs objet d'un JSON (dot-path, --min-keys, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] uuencode-tool : encode/décode en uuencode (format historique usenet, --decode, --strict, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-find-no-final-punct : liste les lignes ne finissant pas par un signe de ponctuation (.!?:;…), (--per-paragraph, --require-none exit 2 CI, JSON) ✓ 2026-08-02

## Vague 425 — CLI Tools (csv wide en long, json supprime chemin, texte extraits mentions, csv garde colonnes non vides, texte lignes avec longueur)
- [x] csv-to-long : dépivote un CSV large en long (valeurs pivots en lignes, --key --pivot-cols, --drop-empty, --require-rows exit 2 CI, JSON) ✓ 2026-08-02 (repo existant, support --key multi-colonnes)
- [x] json-delete-path : supprime la valeur à un chemin dot-path d'un JSON (--path, clé ou index, --require-exists exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-mentions : extrait les @mentions d'un texte (--unique, comptes, --require-min exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-nonempty-columns : garde uniquement les colonnes ayant au moins une valeur non vide (--min-nonempty, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-line-lengths : affiche la longueur de chaque ligne d'un texte (--top, min/max/avg, --require-max exit 2 CI, JSON) ✓ 2026-08-02

## Vague 424 — CLI Tools (diff deux textes, csv compte cellules vides, json remplace valeur, texte extraits hashtags, csv pivot long en large)
- [x] text-diff-lines : diff unifié ligne à ligne entre deux textes (contexte N, --check exit 2 CI si différent, JSON) ✓ 2026-08-02
- [x] csv-count-empty-cells : compte les cellules vides par colonne d'un CSV (--total, --require-max exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-set-value : remplace la valeur à un chemin dot-path d'un JSON (--path --value typé, --require-exists exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-hashtags : extrait les #hashtags d'un texte (--unique, comptes, --require-min exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-to-wide : pivote un CSV long en large (clé/pivot/valeur, --agg first/sum/count, --require-keys exit 2 CI, JSON) ✓ 2026-08-02

## Vague 423 — CLI Tools (texte compte voyelles par ligne, csv supprime lignes selon pattern, json chemins strings max, rot47, texte titres capitalisés)
- [x] text-vowels-per-line : compte les voyelles par ligne d'un texte (min/max/avg, --require-max exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-drop-matching-rows : supprime les lignes CSV dont une colonne matche un pattern (--col --pattern regex, --invert, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-list-long-string-paths : liste les chemins des strings de longueur >= N (dot-path, --min-length, --require-none exit 2 CI, JSON) ✓ 2026-08-02
- [x] rot47-tool : applique ROT47 à un texte (printables ASCII, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-title-case : transforme un texte en Title Case (small words list, --sentence-style, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 422 — CLI Tools (compte mots fréquents, csv header add, json empty objects, anonymize emails, csv row to lines)
- [x] text-word-frequency : fréquence des mots d'un texte (top N, --min-count, --check mot absent exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-add-column : ajoute une colonne constante/calculée à un CSV (--name --value, --in-place, --check existe exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-list-empty-object-paths : liste les chemins vers les objets vides {} d'un JSON (dot-path, --require-none exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-mask-emails : masque les emails d'un texte (user@domain -> u***@d***.tld, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-rows-to-files : éclate chaque ligne d'un CSV en fichier séparé (template, --dry-run, --count exit 2 CI, JSON) ✓ 2026-08-02

## Vague 421 — CLI Tools (lines strip prefix/suffix, csv header case, json paths empty arrays, slugify, csv transpose)
- [x] text-strip-prefix : retire un préfixe donné (ou auto-commun) de chaque ligne (--boundary, --check exit 2 CI, JSON) ✓ 2026-08-02 (repo existant, ajout __main__ + support '-' stdin)
- [x] csv-header-case : normalise les en-têtes d'un CSV (snake/camel/pascal/kebab/upper/title/lower, --only, --dry-run, --check exit 2 CI, JSON) ✓ 2026-08-02 (repo existant, fix args intermixed)
- [x] json-list-empty-array-paths : liste les chemins vers les arrays vides d'un JSON (dot-path, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] slugify-text : convertit un texte en slug URL-safe (lower, ascii, tirets, --max-len, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-transpose : transpose un CSV (lignes ↔ colonnes, --strict, --require-square exit 2 CI, JSON) ✓ 2026-08-02 (repo existant, retrait binaire commité)

## Vague 420 — CLI Tools (texte compte mots par paragraphe, csv garde premières lignes par clé, json chemins floats, base64 décode strict, texte joindre lignes avec séparateur)
- [x] text-words-per-paragraph : compte les mots par paragraphe d'un texte (blank-line separated, stats min/max/moyenne, --require-max exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-first-row-per-key : garde la première ligne de chaque valeur de clé dans un CSV (group-by first, --require-min-groups exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-list-float-paths : liste les chemins vers les valeurs float d'un JSON (dot-path, ints exclus, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] base64-decode-strict : décode base64 avec validation stricte (alphabet, padding, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-join-lines : joint toutes les lignes d'un texte avec un séparateur (--sep, --strip, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 419 — CLI Tools (texte compte parenthèses ouvrantes, csv supprime lignes dupliquées consécutives, json chemins entiers, base32hex encode, texte swap première/dernière ligne)
- [x] text-count-open-parens : compte les parenthèses ouvrantes ( dans un texte (--close, --balance, --require-min/max, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-dedupe-adjacent-rows : supprime les lignes dupliquées consécutives d'un CSV (uniq, --count, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-list-integer-paths : liste les chemins vers les valeurs entières d'un JSON (dot-path, bool exclus, --require-min/max/none exit 2 CI, JSON) ✓ 2026-08-02
- [x] base32hex-tool : encode/décode base32hex (RFC 4648 extended hex alphabet, --no-pad, --decode, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-swap-first-last-line : permute la première et la dernière ligne d'un texte (--in-place, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 418 — CLI Tools (texte compte chevrons, csv drop lignes vides, json chemins number, base85, texte reverse ordre lignes)
- [x] text-count-angles : compte les chevrons < et > dans un texte (--pair-check, --require-min/max, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-drop-empty-rows : supprime les lignes entièrement vides d'un CSV (toutes cellules vides, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-list-number-paths : liste les chemins vers les valeurs number d'un JSON (dot-path, ints vs floats, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] base85-tool : encode/décode base85 (RFC 1924 alphabet, --decode, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-reverse-lines : inverse l'ordre des lignes d'un texte (tac, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 417 — CLI Tools (texte compte dièses, csv strip colonnes vides, json chemins false, base16, texte dédup lignes gardant ordre)
- [x] text-count-hashes : compte les dièses # dans un texte (--per-line, --require-min/max, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-strip-empty-columns : supprime les colonnes entièrement vides d'un CSV (toutes lignes vides, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-list-false-paths : liste les chemins vers les valeurs false d'un JSON (dot-path, --require-none exit 2 CI, JSON) ✓ 2026-08-02
- [x] base16-tool : encode/décode base16 hex (RFC 4648, --upper, --no-0x, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-dedupe-lines : supprime les lignes dupliquées en gardant la première occurrence (ordre stable, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 416 — CLI Tools (texte compte astérisques, csv trim colonnes vides droite, json chemins booléens true, url base64, texte squeeze lignes consécutives identiques)
- [x] text-count-asterisks : compte les astérisques * dans un texte (--per-line, --require-min/max, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-strip-right-empty : supprime les colonnes vides en queue d'un CSV (trailing empty columns, --all-empty exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-list-true-paths : liste les chemins vers les valeurs true d'un JSON (dot-path, --require-none exit 2 CI, JSON) ✓ 2026-08-02
- [x] base64-tool : encode/décode base64 standard (RFC 4648, --wrap, --no-pad, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-squeeze-repeat-lines : supprime les lignes consécutives identiques en gardant la 1re (--count, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 415 — CLI Tools (texte compte backticks, csv strip colonne vide gauche, json chemins strings vides, rot13, texte trim chaque ligne)
- [x] text-count-backticks : compte les backticks ` dans un texte (--require-min/max, --require-balanced, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-strip-left-empty : supprime les colonnes vides en tête d'un CSV (leading empty columns, --max-strip, --check/--all-empty exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-list-empty-string-paths : liste les chemins vers les strings vides d'un JSON (dot-path, --whitespace, --require-none exit 2 CI, JSON) ✓ 2026-08-02
- [x] rot13-tool : applique ROT13 à un texte (lettres seules, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-rstrip-lines : retire l'espace/tab de fin de chaque ligne (LF/CRLF préservés, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 414 — CLI Tools (texte compte tildes, csv supprime colonne, json liste valeurs null, base58 encode, texte lignes indentées)
- [x] text-count-tildes : compte les tildes ~ dans un texte (--require-min/max, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-drop-column : supprime une colonne d'un CSV (par nom ou index, --keep autres, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-list-null-paths : liste les chemins vers les valeurs null d'un JSON (dot-path, --require-none exit 2 CI, JSON) ✓ 2026-08-02
- [x] base58-tool : encode/décode base58 (Bitcoin alphabet, --decode, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-indent-lines : indente les lignes d'un texte (N espaces, --tabs, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02 (repo existant, ajout --in-place)

## Vague 413 — CLI Tools (texte lignes TRI, csv compte par clé, json clés par type, hexdump inverse, texte phrase la plus longue)
- [x] text-sort-lines : trie les lignes d'un texte (--numeric, --reverse, --unique, --locale-like casefold, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-count-by-key : compte les lignes par valeur de clé (group-by, --top, --min-count, --require-max-groups exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-keys-by-type : liste les clés menant à chaque type JSON (string/number/bool/null/array/object, --require-only-types exit 2 CI, JSON) ✓ 2026-08-02
- [x] hex-to-binary-file : décode un texte hex en fichier binaire (--strict, --offset, --check round-trip exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-longest-sentence : extrait la phrase la plus longue d'un texte (split .!?, --top, --require-max-words exit 2 CI, JSON) ✓ 2026-08-02

## Vague 412 — CLI Tools (texte mots par ligne, csv ligne aléatoire, json liste strings, base64 url encode, texte compte mots répétés)
- [x] text-words-per-line : compte les mots par ligne d'un texte (min/max/avg, --require-max exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-random-row : tire une ligne aléatoire d'un CSV (--count, --seed reproductible, --fields colonnes, --require-min-rows exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-list-strings : extrait toutes les strings d'un JSON récursivement (--path dot-path wildcard, --unique, --min-length, --require-max exit 2 CI, JSON) ✓ 2026-08-02
- [x] base64url-tool : encode/décode base64url RFC 4648 (sans padding --no-pad, --decode, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-repeated-words : détecte les mots répétés consécutifs ("the the", --ignore-case, --require-none exit 2 CI, JSON) ✓ 2026-08-02

## Vague 411 — CLI Tools (texte squeeze lignes vides, csv colonnes select, json longueur chemin, texte swap casse, hex encode)
- [x] text-squeeze-blank-lines : compresse les séries de lignes vides en une seule (--max, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-select-columns : sélectionne/réordonne des colonnes d'un CSV (--columns, --exclude, --require-cols exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-path-value-length : longueur des strings à un chemin dot-path JSON (wildcard *, min/max/avg, --require-max exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-swap-case : inverse la casse d'un texte (tout ou --words-only, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] hex-encode-binary : encode binaire/stdin en hex (--upper, --0x-prefix, alignement --group, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 410 — CLI Tools (csv dédup lignes, texte acronymes, json profondeur moyenne, uuid génère/valide, texte censures e-mail)
- [x] csv-dedupe-rows : supprime les lignes dupliquées d'un CSV (--key cols partielles, --keep first/last, --require-max-dupes exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-extract-acronyms : extrait les acronymes ALL-CAPS d'un texte (--min-length, avec comptes, --require-min exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-avg-depth : profondeur moyenne/max des feuilles d'un JSON (par branche, --require-max-depth exit 2 CI, JSON) ✓ 2026-08-02
- [x] uuid-tool : génère/valide des UUID (v4 aléatoire, v5 nommé, --count, --validate, --namespace, exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-redact-emails : masque les adresses e-mail dans un texte (--domain-only, --keep-first, --require-redactions exit 2 CI, JSON) ✓ 2026-08-02

## Vague 409 — CLI Tools (texte lignes plus longues, csv lignes jointes, json trie clés, base32 encode, texte compte ponctuation)
- [x] text-longest-lines : extrait les N lignes les plus longues d'un texte (--top, --min-length, --require-max-length exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-join-rows : joint les lignes CSV consécutives partageant une clé (--key, --agg concat/first/last par colonne, --require-joins exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-sort-keys : trie récursivement les clés d'un JSON (--reverse, --locale, --check exit 2 CI si déjà différent, JSON) ✓ 2026-08-02
- [x] base32-tool : encode/décode base32/base32hex (RFC 4648, --decode, --no-padding, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-count-punctuation : compte les signes de ponctuation par type (.,;:!?…, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02

## Vague 408 — CLI Tools (csv somme colonnes, text compte majuscules, json extraire tableau, yaml compter docs, hex decode stdin)
- [x] csv-sum-column : somme numérique des colonnes choisies (count/sum/avg, --columns, --require-sum-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-capitalized-count : compte les mots capitalisés d'un texte (--min-length, --ignore-sentence-start, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-extract-array : extrait un tableau JSON par chemin pointé (a.b.c, wildcard *, --require-length-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] yaml-count-documents : compte les documents d'un YAML multi-doc (---/... markers, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] hex-to-binary : décode du hex (espaces/0x tolérés) en binaire sur stdout (decode strict --strict, --require-bytes exit 2 CI, JSON report) ✓ 2026-08-02

## Vague 407 — CLI Tools (csv filtre lignes, text compte chiffres, json compact jsonl, env export shell, md titres)
- [x] csv-filter-rows : filtre les lignes CSV par expression sur colonne (--where col=val/col>val/regex, --invert, --require-min-rows exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-count-digits : compte les chiffres d'un texte (par caractère 0-9, total/par ligne, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-lines-compact : compacte un JSON multi-lignes en JSONL ou inverse (--expand, --validate, --require-count exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-to-shell-export : convertit un dotenv en commandes export shell (quoting POSIX, --unset-empty, --prefix, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] md-list-headings : liste les titres d'un markdown avec niveaux (atx + setext, --max-depth, --toc, --require-min exit 2 CI, JSON) ✓ 2026-08-02

## Vague 406 — CLI Tools (rot13 cipher, csv renomme colonnes, json liste tailles, text palindromes, file mime type)
- [x] rot-cipher : chiffre/déchiffre par rotation César (--shift N, lettres seules, --brute-force les 25 clés, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-rename-columns : renomme des colonnes CSV selon un mapping (--map old=new,..., --require-mapped exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-key-lengths : liste les clés JSON avec leur profondeur et longueur (--max-depth, --require-max-key-length exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-find-palindromes : trouve les mots/phrases palindromes dans un texte (--min-length, --phrases, --require-min-found exit 2 CI, JSON) ✓ 2026-08-02
- [x] file-detect-type : devine le type de fichier par signatures magiques (png/jpg/pdf/zip/gzip/..., multi-fichiers, --require-type exit 2 CI, JSON) ✓ 2026-08-02

## Vague 405 — CLI Tools (hex dump, csv transpose, markdown liens, json aplati, texte mots uniques)
- [x] hex-dump-view : hexdump canonique d'un fichier ou stdin (offset/hex/ascii, --length, --width, --skip, JSON) ✓ 2026-08-02
- [x] csv-transpose : transpose un CSV (lignes <-> colonnes, --max-cells garde-fou, --check carré exit 2 CI, JSON) ✓ 2026-08-02
- [x] md-extract-links : extrait les liens d'un markdown (inline, référence, autolinks, --unique, --require-no-broken-anchor exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-flatten-keys : aplati un JSON imbriqué en clés pointées (a.b.c, tableaux [i], --unflatten inverse, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-unique-words : liste les mots uniques d'un texte avec fréquences (--min-length, --top, --require-min-unique exit 2 CI, JSON) ✓ 2026-08-02

## Vague 404 — CLI Tools (compte lignes code, base64 batch, csv stat colonne, env diff, text voyelles)
- [x] code-line-count : compte lignes totales/code/commentaires/blank par langage (table intégrée py/js/go/c/..., multi-fichiers, --require-max-lines exit 2 CI, JSON) ✓ 2026-08-02
- [x] base64-batch : encode/décode base64/base64url de fichiers ou stdin (--decode, --url-safe, --wrap N, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-column-stats : stats numériques d'une colonne CSV (count/min/max/sum/mean/median/stdev, --column, --require-min-mean exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-diff-keys : compare deux fichiers .env (clés ajoutées/supprimées/modifiées, valeurs masquées, --require-identical-keys exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-vowel-ratio : ratio voyelles/consonnes d'un texte (par ligne et global, --require-min-ratio/--require-max-ratio exit 2 CI, JSON) ✓ 2026-08-02

## Vague 403 — CLI Tools (comptage phrases texte, extraction URL, validation colonnes CSV, forme JSON canonique, trim lignes)
- [x] text-sentence-count : compte les phrases d'un texte (découpage heuristique . ! ?, stats mots/par phrase, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] url-extract-domain : extrait host + domaine effectif (édicule TLD simple) des URLs en flux (--unique, --require-domain/--forbid-domain exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-validate-columns : valide l'header CSV contre un schéma attendu (--require-cols, --forbid-cols, --order, --allow-extra, exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-canonical-form : écrit JSON en forme canonique (clés triées récursivement, floats normalisés, --compact, RFC 8785-like lite, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-trim-lines : trim chaque ligne (gauche/droite/--both, chars custom, --collapse, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 402 — CLI Tools (comptage modifiés git, slug texte, top valeurs CSV, validation JSON, wrap texte)
- [x] git-diff-stat-names : liste les fichiers modifiés entre deux refs git (+/- par fichier depuis git diff --numstat, --require-max-files/--require-max-changes exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-slugify : convertit du texte en slug URL-safe (translittération ASCII, lowercase, --separator, --max-length, --require-slug exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-top-values : top-N des valeurs d'une colonne CSV avec fréquences (--column, --top, --min-count, --require-min-frequency exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-check-valid : valide la syntaxe JSON/JSONL de fichiers (multi-fichiers, --require-valid exit 2 CI si un fichier invalide, rapport ligne/colonne erreur, JSON) ✓ 2026-08-02
- [x] text-wrap-width : formate du texte à une largeur cible (word wrap, --width, --no-break-long-words, --require-max-line-length exit 2 CI, JSON) ✓ 2026-08-02

## Vague 401 — CLI Tools (valideur JSON Schema complet, diff YAML, stats mots parlés, template env, format durée)
- [x] json-schema-validator : valide un JSON contre un JSON Schema complet (draft 7+, types, formats, $ref locales, additionalProperties, patternProperties, oneOf/anyOf/allOf, exit 2 CI, JSON) ✓ 2026-08-02
- [x] yaml-diff-tool : compare deux fichiers YAML et affiche les différences (structure-aware, --ignore-order lists, --show-values, exit 2 CI si différent, JSON) ✓ 2026-08-02
- [x] text-speaking-time : estime le temps de lecture/parole d'un texte (mots par minute configurable, stats pauses/ponctuation, --require-min/max-time exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-template-generator : génère un fichier .env.example depuis un .env (valeurs masquées, --keep-comments, --sort, --require-keys, JSON) ✓ 2026-08-02
- [x] text-format-duration : convertit des durées entre formats (secondes ↔ "1h 23m 45s" ↔ "83:45", parse/format, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02

## Vague 400 — CLI Tools (anagrammes texte, validation INI, profondeur YAML, stats lignes shell, merge CSV)
- [x] text-anagram-finder : trouve les anagrammes dans un texte (normalisation Unicode casefold, groupes de mots anagrammes, --min-length/--max-group-size, --require-anagrams exit 2 CI, JSON) ✓ 2026-08-02
- [x] ini-validate-syntax : valide la syntaxe d'un fichier INI (sections, clés=valeurs, commentaires, doublons section/clé, --allow-no-value, --require-sections, exit 2 CI, JSON) ✓ 2026-08-02
- [x] yaml-max-depth : profondeur maximale d'imbrication d'un YAML (mapping/sequences, --require-max-depth exit 2 CI, --show-path, JSON) ✓ 2026-08-02
- [x] shell-line-length : lint la longueur des lignes de scripts shell (--max-len, --ignore-comments, --ignore-heredoc, exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-merge-files : fusionne plusieurs CSV avec header identique (--check-headers, dédupe --unique, --sort colonne, JSON) ✓ 2026-08-02

## Vague 399 — CLI Tools (palindromes, audit quoting shell, catégories Unicode, split CSV, paragraphes)
- [x] text-palindrome-checker : détecte lignes palindromes et mots palindromes (normalisation Unicode casefold, --min-length/--max-results, --require-palindromes exit 2 CI, JSON) ✓ 2026-08-02
- [x] shell-quote-audit : lint les dangers de quoting shell ($VAR, $(...), backticks, $@ et $*, heredoc/comment-aware, --only/--skip, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-unicode-category : compte les catégories générales Unicode (par cat Lu/Ll/Nd/... + groupes letters/numbers/...), --show-category avec codepoints, --require-ascii/--max-non-ascii-pct/--min-letters-pct exit 2 CI, JSON ✓ 2026-08-02
- [x] csv-split-file : découpe un CSV en N chunks en préservant le header (--rows-per-chunk ou --num-chunks, stdin/out-dir/prefix/digits, JSON) ✓ 2026-08-02
- [x] text-paragraph-count : compte les paragraphes (runs de lignes non-blanc, --blank-run, stats mots/chars par paragraphe, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02

## Vague 398 — CLI Tools (score lisibilité texte, tri imports Python, pièges shell, validation JSON)
- [x] text-readability-checker : score de lisibilité heuristique (Flesch Reading Ease, Flesch-Kincaid Grade, SMOG, stats phrases/mots/syllabes, --require-readability/--max-grade/--max-sentence-length/--complex-pct exit 2 CI, JSON) ✓ 2026-08-02
- [x] py-import-sort-checker : audit/fix l'ordre des blocs d'imports Python top-level (__future__ d'abord, groupes stdlib < third-party < first-party < local, tri alpha insensible à la casse, --fix in-place, --first-party, ligne commentée = ancre, exit 2 CI, JSON) ✓ 2026-08-02
- [x] shell-pitfall-checker : lint heuristique des pièges bash (vars non quotées, for-in-$(ls), cd sans garde, sudo > file, ps|grep, read sans -r, for in $@, heredocs ignorés, --only/--skip/--require-pipefail, exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-schema-lite : validation JSON/JSONL contre un schéma minimal sans dépendance (type/required/properties/items/enum/bornes/pattern/additionalProperties, --max-errors, -q, exit 2 CI, JSON) ✓ 2026-08-02

## Vague 397 — CLI Tools (audit quoting CSV, stats longueur mots, tri dotenv, valeurs vides YAML, escapes JSON, ENV Dockerfile, messages commit)
- [x] csv-detect-quoted-fields : audit le quoting des champs CSV (re-lexing RFC 4180, cellules quoted/needs_quotes/unquoted_special, --require-quoted-columns, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-median-word-length : distribution des longueurs de mots (min/max/mean/median/mode, histogramme ASCII, Unicode, --require-median/mean-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-lex-sort : vérifie/applique l'ordre alphabétique des entrées dotenv (commentaires attachés aux entrées, --ignore-case, --fix/--in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] yaml-empty-values : lint les valeurs de mapping vides (key: sans valeur, block-aware, --include-explicit-null, --allow, --check/--require-max exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-string-escape-audit : audit les séquences d'échappement des strings JSON (unicode/simple/invalid, JSONL, --verbose, --require-max-unicode/--require-no-invalid exit 2 CI, JSON) ✓ 2026-08-02
- [x] dockerfile-lint-env-spaces : lint les espaces vs '=' dans ENV Dockerfile (forme legacy dépréciée, continuations, --fix/--in-place, exit 2 CI, JSON) ✓ 2026-08-02
- [x] git-commit-msg-lint : lint les messages de commit (largeur sujet/corps, conventional commits --require-conventional/--types, impératif, ligne 2 vide, hook commit-msg, exit 2 CI, JSON) ✓ 2026-08-02

## Vague 396 — CLI Tools (comptage mots texte, header CSV, taille JSON, commentaire licence dotenv, quotes YAML)
- [x] text-count-distinct-words : compte les mots distincts d'un texte (fréquences, --ignore-case, --top, --min-length, lexical diversity, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-check-header : exige/interdit des colonnes dans un header CSV (--require/--forbid, ordre --require-order, doublons, count exact, exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-byte-size : mesure la taille en octets d'un JSON sérialisé (compact/indent, JSONL per-document, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-license-header : exige un commentaire SPDX en tête d'un dotenv (--require SPDXID, --set insertion, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] yaml-quote-check : lint les plain scalars YAML à résolution implicite risquée (boolean/null/number YAML 1.1, --fix quote, --allow, --kinds, exit 2 CI, JSON) ✓ 2026-08-02

## Vague 395 — CLI Tools (whitespace texte, casse valeurs CSV, profondeur clés JSON, trailing comment dotenv, block scalar YAML)
- [x] text-collapse-whitespace : compresse les espaces multiples des lignes (indent préservée par défaut, --collapse-leading, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-value-casing : lint la casse des valeurs d'une colonne CSV (7 styles, distribution, --require-style exit 2 CI, --list-violations, JSON) ✓ 2026-08-02
- [x] json-deep-keys : liste les clés JSON à profondeur >= N (--min-depth, --max-results, --require-none exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-trailing-comment : lint les commentaires de fin de ligne dotenv (espace avant # requis, --fix, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] yaml-block-scalar-check : vérifie la cohérence des block scalars (|/-, +/- chomping, --require-style literal|folded exit 2 CI, JSON) ✓ 2026-08-02

## Vague 394 — CLI Tools (largeur texte, champs CSV, entropie JSON, commentaires dotenv, retours chariot YAML)
- [x] text-count-grapheme-width : largeur d'affichage par ligne (ASCII=1, large chars CJK/emoji=2, --require-max-width exit 2, JSON) ✓ 2026-08-02
- [x] csv-extract-column : extrait une colonne CSV par nom/index (valeur brute, --unique, --sort, --count, JSON) ✓ 2026-08-02
- [x] json-entropy-shannon : entropie de Shannon des strings/bytes d'un JSON (--require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-comment-style : lint le style des commentaires dotenv (require space après #, --fix, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] yaml-normalize-eol : détecte et normalise les fins de ligne d'un YAML (CRLF/LF/CR mixtes, --to lf|crlf, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 393 — CLI Tools (contrôle texte, profondeur JSON, longueur valeurs CSV, trailing whitespace dotenv, directives YAML)
- [x] text-strip-control-chars : retire les caractères de contrôle ASCII (C0 sauf tab/lf/cr configurable via --keep, DEL+C1, --replace, --list codepoints, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-max-depth : profondeur maximale d'imbrication d'un JSON (--require-max-depth/--require-min-depth exit 2 CI, JSONL, --show-path pointe la feuille la plus profonde, JSON) ✓ 2026-08-02
- [x] csv-value-max-length : longueur max des valeurs par colonne CSV (--columns nom/index, --require-max-len exit 2 CI, --truncate + --ellipsis, JSON) ✓ 2026-08-02
- [x] env-strip-trailing-space : retire l'espace/tab de fin de ligne d'un .env (quotés préservés, --check exit 2 CI, --in-place, JSON) ✓ 2026-08-02
- [x] yaml-lint-directives : valide les directives d'un YAML (%YAML version, %TAG, document markers ---/..., --require-version exit 2 CI, JSON) ✓ 2026-08-02

## Vague 392 — CLI Tools (rotation CSV, nulls JSON, tabs→espaces, casse clés dotenv, commentaires YAML)
- [x] csv-rotate-columns : décale les colonnes d'un CSV vers la gauche/droite (header compris ou --keep-header, --left N/--right N, --delimiter, --check --reference exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-count-null : compte les valeurs null récursivement (--path dot-path avec wildcard *, JSONL, --list-paths, --require-max/--require-zero exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-tabs-to-spaces : convertit les tabulations en espaces avec vrais tab stops (--width N, --leading-only, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-key-casing : lint la casse des clés dotenv (7 styles, whitelist exacte/préfixe/regex, --require-max-violations tolérance CI, --list-styles, valeurs jamais affichées, exit 2 CI, JSON) ✓ 2026-08-02
- [x] yaml-strip-comments : retire les commentaires d'un YAML en préservant les # dans les strings (quoting simple/double avec échappements '' et \", séparation whitespace requise, --keep-blank, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 391 — CLI Tools (BOM texte, checksum lignes CSV, casse clés JSON, ARG avant FROM, trim valeurs dotenv)
- [x] text-detect-bom : détecte et strip les BOM (UTF-8/16/32 LE/BE, UTF-7), multi-fichiers/stdin, --require-none/--require ENCODING exit 2 CI, --strip in-place, JSON ✓ 2026-08-02
- [x] csv-row-checksum : calcule et vérifie un checksum par ligne CSV (tous algos hashlib, --exclude colonnes volatiles, --standalone, --check-column anti-tampering exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-key-casing : lint la convention de casse des clés JSON (snake/camel/kebab/pascal/upper/lower, whitelist exacte/préfixe/regex, --require-max-violations tolérance CI, --list-styles, JSON) ✓ 2026-08-02
- [x] dockerfile-arg-before-from : exige que tout ARG référencé dans un FROM soit déclaré avant le premier FROM (règle BuildKit, $VAR/${VAR}, --allow, multi-stages, exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-trim-values : trimme l'espace autour des valeurs dotenv brutes (quotés préservés sauf --include-quoted, --collapse-inner, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 390 — CLI Tools (valeurs uniques CSV, clés dupliquées JSON, comptage emojis, volumes Dockerfile, délimiteur CSV sniffé)
- [x] csv-unique-values : liste les valeurs uniques d'une colonne CSV avec occurrences (--ignore-case, --trim, --require-min/max-unique, --require-max-frequency, --check-dominated PCT exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-lint-duplicate-keys : détecte les clés d'objet JSON dupliquées que json.loads masque silencieusement (JSON+JSONL, --skip-invalid, --require-max tolérance CI, exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-emoji-count : compte les caractères emoji d'un texte (emoji/ZWJ/keycap/VS16, by-kind et by-category Unicode, --list codepoints+noms, --unique, --require-max/--require-zero exit 2 CI, JSON) ✓ 2026-08-02
- [x] dockerfile-check-volume : lint VOLUME Dockerfile (forms shell+JSON, --forbid-volume/--require-volume, --allow-path/--forbid-path/--forbid-prefix, --forbid-variable, --require-absolute, --final-only, exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-sniff-delimiter : détecte le délimiteur le plus probable d'un CSV (scoring consistency×columns×rows, --candidates, --check exit 2 CI, --require-min-columns/consistency, JSON) ✓ 2026-08-02

## Vague 389 — CLI Tools (index colonnes CSV, types JSON, lignes à tirets, couches Dockerfile, valeurs dotenv masquées)
- [x] csv-column-index : mapping nom→index des colonnes CSV (0 et 1-based, --structs, --strict doublons exit 2, JSON) ✓ 2026-08-02
- [x] json-value-types : distribution des types JSON récursivement (object/array/string/integer/float/boolean/null, JSONL, --require-type, --require-null-max-pct exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-hyphen-heavy-lines : lint typographique des lignes surchargées en tirets (--max-per-line, --unicode, --ignore-comments/separators, exit 2 CI, JSON) ✓ 2026-08-02
- [x] dockerfile-layer-count : estime les couches d'image par stage Dockerfile (FROM/RUN/COPY/ADD, --max-layers/--require-min exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-redact-values : masque les valeurs d'un .env en préservant clés/commentaires/quoting (--only-secrets, --keep-prefix, --hash sha256, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 388 — CLI Tools (cellules vides CSV, min/max JSON par clé, voyelles texte, tags base Dockerfile, secrets dotenv)
- [x] csv-count-null-cells : compte les cellules vides d'un CSV par colonne + total (--no-header, --delimiter, --require-max/--forbid-empty exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-min-max-key : min/max/somme/moyenne des valeurs numériques d'un chemin dot-path JSON (wildcard *, JSONL, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-count-vowels : compte voyelles/consonnes d'un texte (accents via NFD, --y-is-vowel, --ascii-only, ratio, --require-ratio-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] dockerfile-base-tag-check : exige des base images épinglées (pas de tag manquant/:latest, --require-digest sha256, --forbid-pattern, multi-stage, registry ports, exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-detect-secrets : détecte les secrets probables dans les .env sans jamais afficher les valeurs (formats AWS/GitHub/Slack/JWT/PEM, entropie Shannon, placeholders ignorés, exit 2 CI, JSON) ✓ 2026-08-02

## Vague 387 — CLI Tools (transposition CSV, strings vides JSON, acronymes texte, dotenv→ConfigMap, EXPOSE interdit)
- [x] csv-transpose-rows : transpose lignes/colonnes d'un CSV (pad rows courtes, --delimiter, --check square exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-count-empty-strings : compte les strings vides récursivement (--whitespace, --list-paths dot-paths, JSONL, --require-max/--require-zero exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-count-uppercase-words : compte les mots tout-majuscules/acronymes (--min-length, --include-single, --list fréquences, --unique, --require-min/max exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-to-k8s-configmap : génère un manifest ConfigMap/Secret depuis dotenv (quoting YAML, base64 --as-secret, labels/annotations, validation noms exit 2, JSON) ✓ 2026-08-02
- [x] dockerfile-verify-no-expose : interdit/restreint les EXPOSE d'un Dockerfile (--allow PORT/PROTO, --allow-env, --final-only, continuations, exit 2 CI, JSON) ✓ 2026-08-02

## Vague 386 — CLI Tools (remplissage CSV vide, somme récursive JSON, strip chiffres, commentaires Dockerfile, dotenv→properties)
- [x] csv-fill-missing : remplit les cellules vides d'un CSV (--fill valeur fixe, --ffill forward-fill, --empty-as-zero, colonnes nom/index, --skip-rows regex, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-sum-numbers : somme récursive des nombres d'un JSON/JSONL (bool exclus, --path dot-path avec *, --require-min/max exit 2 CI, --count, --int, JSON) ✓ 2026-08-02
- [x] text-strip-digits : retire les chiffres d'un texte (ASCII + Unicode Nd, --replace, --skip-first, --only-matching, --in-place, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] dockerfile-check-comment : lint des commentaires Dockerfile (--require-header, --forbid-todo TODO/FIXME/HACK, --style, --allow-marker, exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-to-properties : convertit dotenv en .properties Java (échappement = : # !, comments, --sort, --prefix, --lowercase-keys, export-aware, JSON) ✓ 2026-08-02

## Vague 385 — CLI Tools (lignes CSV vides, valeurs non vides dotenv, ENTRYPOINTs Dockerfile, longueur max strings JSON, lignes non vides texte)
- [x] csv-remove-comma-only-rows : retire les lignes CSV dont toutes les cellules sont vides (header préservé, --no-header, --delimiter, --keep-empty-rows, --in-place, --check exit 2 CI, --require-max, JSON)
- [x] env-count-non-empty : compte les entrées dotenv à valeur non vide (export-aware, quotes, --names sans révéler les valeurs, --require-min/max, --forbid-empty exit 2 CI, JSON)
- [x] dockerfile-entrypoint-count : compte les ENTRYPOINT par stage d'un Dockerfile (--final-only, form exec/shell, --require-entrypoint, --forbid-multiple, --require-form, continuations, exit 2 CI, JSON)
- [x] json-string-length-max : min/max des longueurs de strings JSON récursivement (--path dot-path, JSONL, --include-keys, --show-longest, --ignore-empty, --require-min/max-len exit 2 CI, JSON)
- [x] text-count-non-empty-lines : compte les lignes non vides de fichiers texte (modes non-empty/non-blank, multi-fichiers + TOTAL, --percent, --sort, --require-min/max, --aggregate, JSON)

## RÈGLE
Chaque outil = son propre repo Git sur github.com/TataneSan.
JAMAIS mentionner IA/agent dans le code ou les commits.
Push automatique après chaque outil.

## Outil publié — GPU Research
- [x] modal-research-arena : orchestration d'expériences reproductibles avec planification de coût, budgets verrouillés, exécution locale déterministe et backend Modal distant opt-in ✓ 2026-08-01

## Vague 384 — CLI Tools (comptage caractères texte, aplatissement JSON, labels Dockerfile, CMDs Dockerfile, suppression lignes vides)
- [x] text-count-chars : compte caractères/mots/lignes d'un texte (multi-fichiers, agrégat, --mode total/no-whitespace/no-newlines, --only, --require-min/max CI, JSON) ✓ 2026-08-02
- [x] json-flatten-array : aplatit récursivement les tableaux JSON imbriqués (--max-depth, --require-array, --require-length-min/max, JSONL, --compact, JSON) ✓ 2026-08-02
- [x] dockerfile-check-label : exige/interdit des LABEL dans un Dockerfile (multi-stages, --final-only, forms key=value et legacy key value, continuations, --list, exit 2 CI, JSON) ✓ 2026-08-02
- [x] dockerfile-cmd-count : compte les CMD par stage d'un Dockerfile (form exec/shell, --final-only, --require-cmd, --forbid-multiple, exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-remove-empty-lines : retire les lignes vides d'un texte (whitespace-only, --keep-one, --in-place, --check exit 2 CI, --require-max, JSON) ✓ 2026-08-02

## Vague 383 — CLI Tools (comptage types JSON, doublons CSV, HEALTHCHECK requis, strip markdown, validation clés dotenv)
- [x] json-count-type-values : compte les valeurs d'un type JSON récursivement (mode all, JSONL, --require-min/max CI, JSON) ✓ 2026-08-02
- [x] csv-detect-duplicate-rows : détecte les lignes dupliquées d'un CSV (--columns nom/index, --ignore-case/--trim, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] dockerfile-check-healthcheck : exige un HEALTHCHECK par stage (--all-stages, --forbid-none, --require-cmd, --max-interval/timeout, --min-retries, exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-strip-markdown : retire la syntaxe Markdown (titres, liens, emphase, listes, code, blockquote, HR, --keep-urls, --strip-html, --in-place, --check CI, JSON) ✓ 2026-08-02
- [x] env-validate-keys : valide les noms de clés dotenv (format, conventions upper/lower-snake/camel, préfixes, doublons, valeurs jamais affichées, exit 2 CI, JSON) ✓ 2026-08-02

## Vague 382 — CLI Tools (dédoublonnage JSON, tri JSON, inversion CSV, ADD interdit Dockerfile, COPY best practices Dockerfile)
- [x] json-dedupe-array : déduplique un tableau JSON en préservant l'ordre d'apparition (comparaison par forme canonique, JSONL, --sort, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-sort-array : trie les items d'un tableau JSON (ordre naturel, --key dot-path, --reverse, JSONL, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-reverse-rows : inverse l'ordre des lignes d'un CSV en gardant le header (délimiteur custom, --no-header, --check + --reference CI, JSON) ✓ 2026-08-02
- [x] dockerfile-add-check : interdit ADD en faveur de COPY (URL/archive/local, --allow-url/--allow-archive, continuations, exit 2 CI, JSON) ✓ 2026-08-02
- [x] dockerfile-copy-check : règles COPY Dockerfile (--require-copy, --forbid-copy-dot, --forbid-wildcard, --require-trailing-slash, exit 2 CI, JSON) ✓ 2026-08-02

## Vague 381 — CLI Tools (padding lignes CSV, strip ANSI, WORKDIR non-root Dockerfile, valeurs longues dotenv, USER non-root Dockerfile)
- [x] csv-pad-rows : normalise les lignes courtes d'un CSV à la largeur attendue (fill custom, --drop, --pad-to, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] text-strip-ansi : retire les séquences d'échappement ANSI (CSI/OSC/ESC) d'un texte (--check exit 2 CI, JSON) ✓ 2026-08-02
- [x] dockerfile-workdir-check : exige un WORKDIR absolu non-root dans chaque stage (variables ok, --all-stages, exit 2 CI, JSON) ✓ 2026-08-02
- [x] env-find-long-lines : détecte les valeurs dotenv dépassant un seuil de longueur (valeurs jamais affichées, --names-only, exit 2 CI, JSON) ✓ 2026-08-02
- [x] dockerfile-check-user : exige un USER non-root dans chaque stage Dockerfile (root/0/0:0 interdits, uid>=1000 ok, --all-stages, exit 2 CI, JSON) ✓ 2026-08-02

## Vague 380 — CLI Tools (underscore→espace, chaînage RUN Dockerfile, comptage booléens JSON, dotenv→Dockerfile, colonnes vides CSV)
- [x] text-underscore-to-space : remplace _ par des espaces (--in-place, --check CI, --json) ✓ 2026-08-02
- [x] dockerfile-run-one-layer : détecte les RUN consécutifs à chaîner (3+ par défaut, --strict 2+, exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-boolean-count : compte les true/false récursivement (--require-true-min/--require-false-max CI, JSON) ✓ 2026-08-02
- [x] env-to-dockerfile : convertit dotenv en instructions ENV Dockerfile (échappement backslash/quotes, séparateur =/espace, JSON) ✓ 2026-08-02
- [x] csv-detect-empty-columns : détecte les colonnes entièrement vides d'un CSV (header, --no-header, --check exit 2 CI, JSON) ✓ 2026-08-02

## Vague 379 — CLI Tools (extraction emails, contraintes EXPOSE Dockerfile, merge unique tableaux JSON, valeurs vides dotenv, types colonnes CSV)
- [x] text-email-extract : extrait les adresses email d'un texte (dédoublonnage case-insensitive, --count fréquences, exit 2 sans match, JSON) ✓ 2026-08-02
- [x] dockerfile-expose-check : vérifie les ports EXPOSE d'un Dockerfile (--require-port/--forbid-port/--forbid-udp, continuations, exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-merge-arrays-unique : concatène des tableaux JSON et déduplique (1re occurrence gardée, --sort, stdin via -) ✓ 2026-08-02
- [x] env-list-empty-values : liste les clés vides d'un dotenv (KEY= et KEY="", export-aware, --check exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-column-types : infère le type de chaque colonne CSV (integer/float/boolean/string/mixed, null-aware, JSON) ✓ 2026-08-02

## Vague 378 — CLI Tools (tri mots par ligne, comptage stages Dockerfile, filtre tableau JSON, dotenv→JSON, suppression dernières lignes CSV)
- [x] text-sort-words-line : trie les mots par ordre alphabétique dans chaque ligne (--reverse, indent préservée, --in-place, --check CI, JSON) ✓ 2026-08-02
- [x] dockerfile-count-stages : compte les stages FROM d'un Dockerfile (alias, --names, --require-min/max CI, JSON) ✓ 2026-08-02
- [x] json-filter-array-by-value : filtre les objets d'un tableau JSON par valeur de champ (dot-path, --number, exit 2 sans match) ✓ 2026-08-02
- [x] env-to-json : convertit dotenv en objet JSON (types inférés int/float/bool/null, --raw, --compact) ✓ 2026-08-02 (repo pré-existant, hybride Go+Python)
- [x] csv-remove-last-rows : retire les N dernières lignes d'un CSV (header préservé, --delimiter, --no-header, JSON) ✓ 2026-08-02

## Vague 377 — CLI Tools (capitalisation phrases, apt-get cleanup Dockerfile, clés par préfixe JSON, dotenv→shell, swap 1re/dernière colonne CSV)
- [x] text-capitalize-sentences : majuscule en début de phrase (. ! ? préservés, --in-place, --check CI, --json) ✓ 2026-08-02
- [x] dockerfile-check-run-apt-get-clean : exige cleanup après apt-get install (rm lists ou apt-get clean, exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-count-keys-by-prefix : compte les clés JSON par préfixe récursivement (objets+arrays, --require-min/max CI, JSON) ✓ 2026-08-02
- [x] env-to-shell : convertit dotenv en lignes export shell (échappement valeurs, --prefix, JSON) ✓ 2026-08-02
- [x] csv-swap-first-last : permute 1re et dernière colonnes d'un CSV (header inclus, --delimiter, --no-header, JSON) ✓ 2026-08-02

## Vague 376 — CLI Tools (duplication mots, CMD exec Dockerfile, min/max tableau JSON, strip commentaires dotenv, ligne min CSV)
- [x] text-duplicate-words : duplique chaque mot par ligne (indent préservée, --in-place, --check CI, --json) ✓ 2026-08-02
- [x] dockerfile-verify-no-shell-cmd : exige CMD forme exec JSON (multi-stages, --final-only, --require-cmd, exit 2 shell form, JSON) ✓ 2026-08-02
- [x] json-array-max-min : min/max d'un tableau JSON de nombres (bool exclus, --require-min/--require-max CI, JSON) ✓ 2026-08-02
- [x] env-strip-comments : retire les lignes # d'un dotenv (pairs et blancs préservés, --in-place, --check CI, JSON) ✓ 2026-08-02
- [x] csv-min-row : affiche la ligne CSV au min d'une colonne numérique (nom/index 1-based, --no-header, --delimiter, --require-min/max CI, JSON) ✓ 2026-08-02

## Vague 375 — CLI Tools (répétition mots, ENTRYPOINT exec Dockerfile, profondeur JSON, doublons dotenv, ligne max CSV)
- [x] text-repeat-words : répète chaque mot N fois par ligne (--sep, indentation préservée, --in-place, --check CI, --json) ✓ 2026-08-02
- [x] dockerfile-verify-entrypoint-json : exige ENTRYPOINT forme exec JSON (multi-stages, --final-only, --require-entrypoint, exit 2 shell form, JSON) ✓ 2026-08-02
- [x] json-count-object-depth : profondeur max d'imbrication JSON/JSONL (objets+arrays, --require-min/--require-max CI, JSON) ✓ 2026-08-02
- [x] env-check-duplicates : détecte les clés KEY= dupliquées d'un dotenv (lignes rapportées, export-aware, exit 2 CI, JSON) ✓ 2026-08-02
- [x] csv-max-row : affiche la ligne CSV au max d'une colonne numérique (nom/index 1-based, --no-header, --delimiter, --require-min/max CI, JSON) ✓ 2026-08-02

## Vague 374 — CLI Tools (swap case mots, préfixes dotenv, FROM latest Dockerfile, pretty JSONL, longueurs clés dotenv)
- [x] text-swap-case-words : inverse la casse de chaque mot par ligne (swapcase Unicode, espaces préservés, --skip-first, --check CI, --json) ✓ 2026-08-02
- [x] env-prefix-check : vérifie que les clés d'un dotenv commencent par un préfixe autorisé (multi-préfixes, export-aware, violations listées, exit 2 CI, JSON) ✓ 2026-08-02
- [x] dockerfile-verify-no-latest : interdit FROM latest explicite/implicite (digest ok, multi-stages, continuations, --allow-arg pour ${BASE}, exit 2 CI, JSON) ✓ 2026-08-02
- [x] json-pretty-lines : pretty-print chaque ligne JSONL en document indenté (--indent/--sort-keys/--ascii, strict ou --skip-invalid) ✓ 2026-08-02
- [x] env-key-lengths : rapporte la longueur de chaque clé d'un dotenv (--sort key/length, --require-max CI, longest/shortest, JSON) ✓ 2026-08-02

## Vague 373 — CLI Tools (reverse word order, vérification CMD exec JSON, longueur tableau JSON, présence clés dotenv, colonne TOTAL CSV)
- [x] text-reverse-word-order : inverse l'ordre des mots de chaque ligne (indent préservée, --sep, --skip-first, --check CI, --json) ✓ 2026-08-02
- [x] dockerfile-verify-cmd-json : vérifie que les CMD d'un Dockerfile utilisent la forme exec JSON (multi-stages, --final-only, --require-cmd, exit 2 shell form, JSON) ✓ 2026-08-02
- [x] json-count-array-length : compte les items d'un tableau JSON/JSONL (dot-path avec indices, --require-min/--require-max CI, JSON) ✓ 2026-08-02
- [x] env-check-key : vérifie la présence de clés dans un dotenv (--require-non-empty, --match REGEX, jamais d'impression des valeurs, doublons, JSON) ✓ 2026-08-02
- [x] csv-row-total : ajoute une colonne TOTAL sommant les colonnes numériques par ligne (--columns, --delimiter, --no-header, --name, --empty-as-zero, JSON) ✓ 2026-08-02

## Vague 372 — CLI Tools (suppression dernier mot, ENTRYPOINT Dockerfile, minification canonique JSON, masquage valeurs dotenv, colonne N CSV)
- [x] text-delete-last-word : supprime le dernier mot de chaque ligne (indent/trail préservés, --skip-first/--only-matching/--skip-blank, --in-place, --check CI, JSON) ✓ 2026-08-02
- [x] dockerfile-extract-entrypoint-json : extrait l'ENTRYPOINT per stage d'un Dockerfile (exec/JSON vs shell, --final-only, --command-only, --require-entrypoint/--require-exec-form/--forbid-shell-form CI, JSON) ✓ 2026-08-02
- [x] json-minify-clone : réécrit JSON/JSONL en forme minifiée canonique (--sort-keys, --ascii, --in-place/-o, --check CI, --require-object/--require-array, rapport octets JSON) ✓ 2026-08-02
- [x] env-mask-values : masque les valeurs de clés données d'un dotenv (quotes + commentaires préservés, --unmask, --check CI, --require-present/--require-masked, JSON) ✓ 2026-08-02
- [x] csv-nth-column : imprime les valeurs de la colonne N d'un CSV (nom/1-based index, --no-header, --delimiter, --unique/--count, --require-min/max/value/non-empty CI, JSON) ✓ 2026-08-02

## Vague 371 — CLI Tools (swap 1er/dernier mot, RUN longs Dockerfile, delete nulls JSON, tri clés dotenv, comptage remplissage colonne CSV)
- [x] text-swap-first-last-word : permute le premier et dernier mot de chaque ligne (indent/trail préservés, --min-words, --skip-first, --check CI, JSON) ✓ 2026-08-02
- [x] dockerfile-extract-run-long : liste les RUN d'un Dockerfile dépassant un seuil (chars/lignes physiques, comptage &&, --all, --require-none-long/max-len/max-chained CI, JSON) ✓ 2026-08-02
- [x] json-delete-nulls : supprime les clés null des objets JSON/JSONL récursivement (--no-recursive, --prune-empty, tableaux préservés, --require-removed/--require-none CI, JSON) ✓ 2026-08-02
- [x] env-sort-keys : trie les pairs KEY=VALUE d'un dotenv (blocs par commentaires/blancs préservés, --group-comments, --reverse, --in-place, --check CI, JSON) ✓ 2026-08-02
- [x] csv-count-filled-column : compte les cellules non vides d'une colonne CSV (nom/index 1-based, --print-empty, --require-min/max/all-filled CI, JSON) ✓ 2026-08-02

## Vague 370 — CLI Tools (suppression premier mot, COPY --from Dockerfile, insert champ JSON à un path, delete clé dotenv, 3e colonne CSV)
- [x] text-delete-first-word : supprime le premier mot de chaque ligne (indent préservée, --skip-first, --only-matching, --check CI, JSON) ✓ 2026-08-02
- [x] dockerfile-extract-copy-from : liste les COPY --from d'un Dockerfile par stage (refs stage/image/numeric, --refs-only, --require/from/forbid-from/--require-stage-ref/--forbid-image-ref/--require-resolved CI, JSON) ✓ 2026-08-02
- [x] json-insert-field : insère une clé=valeur à un dot-path JSON/JSONL (wildcards *, indices, création de parents, types inférés, --overwrite, --require-inserted/--require-absent CI, JSON) ✓ 2026-08-02
- [x] env-delete-key : supprime des clés KEY=VALUE d'un dotenv (batch, --keys-file, --invert keep-only, export-aware, stdin/out, --dry-run, --require-absent CI, JSON) ✓ 2026-08-02
- [x] csv-third-column : imprime les valeurs de la 3e colonne CSV (header-aware, délimiteur custom, short rows skippés, --unique/--count, --require-min/max/value CI, JSON) ✓ 2026-08-02

## Vague 369 — CLI Tools (répétition dernier mot, LABEL Dockerfile, prepend dotenv, 2e colonne CSV, VOLUME Dockerfile)
- [x] text-repeat-last-word : répète le dernier mot de chaque ligne N fois (prepend/append, --sep, --skip-first, --check CI, JSON) ✓ 2026-08-02
- [x] dockerfile-extract-label : extrait les LABEL d'un Dockerfile par stage (multi-pairs, quotes, legacy form, --names-only/--values-for, --require-label/--forbid-label/--require-value CI, JSON) ✓ 2026-08-02
- [x] env-prepend-line : prépend des KEY=VALUE en tête d'un dotenv (ordre préservé, --after-comments, --no-clobber/--update, --export, --stdin-pairs, --dry-run, --require-present CI, JSON) ✓ 2026-08-02
- [x] csv-second-column : imprime les valeurs de la 2e colonne CSV (header-aware, délimiteur custom, short rows skippés, --unique/--count, --require-min/max/value CI, JSON) ✓ 2026-08-02
- [x] dockerfile-extract-volume : extrait les VOLUME d'un Dockerfile par stage (formes JSON/shell, continuations, --paths-only, --final-only, --require-volume/--forbid-volume CI, JSON) ✓ 2026-08-02

## Vague 368 — CLI Tools (répétition premier mot, MAINTAINER Dockerfile, append dotenv, 1re/dernière colonne CSV, tags FROM)
- [x] text-repeat-first-word : répète le premier mot de chaque ligne N fois (prepend/append, --sep, --skip-first, --check CI, JSON) ✓ 2026-08-02
- [x] dockerfile-extract-maintainer : extrait les MAINTAINER et LABEL maintainer= par stage (--values-only, --include-labels, --require/--forbid-maintainer CI, JSON) ✓ 2026-08-02
- [x] env-append-line : ajoute des KEY=VALUE à un dotenv (idempotent --no-clobber/--update, --export, --stdin-pairs, --dry-run, --require-present CI, JSON) ✓ 2026-08-02
- [x] csv-last-column : imprime les valeurs de la dernière colonne CSV (header-aware, délimiteur custom, --unique/--count, --require-min/max/value CI, JSON) ✓ 2026-08-02
- [x] csv-first-column : imprime les valeurs de la première colonne CSV (header-aware, délimiteur custom, --unique/--count, --require-min/max/value CI, JSON) ✓ 2026-08-02
- [x] dockerfile-extract-from-tag : extrait name/tag/digest des FROM par stage (latest implicite, registry:port, ARG détecté, --require-pin/--require-digest/--forbid-latest CI, JSON) ✓ 2026-08-02

## Vague 367 — CLI Tools (ADD Dockerfile, STOPSIGNAL Dockerfile, ONBUILD Dockerfile, inversion JSON, snake→kebab, valeurs dotenv)
- [x] dockerfile-extract-add : liste les ADD d'un Dockerfile par stage (flags --checksum/--chown, exec/shell form, --flags-only, --forbid-add/--require-add/--forbid-remote/--forbid-checksum-missing CI, JSON) ✓ 2026-08-02
- [x] dockerfile-extract-stopsignal : extrait le STOPSIGNAL effectif par stage (last-wins, défaut SIGTERM, noms/numéros, --final-only, --require-signal/--forbid-default/--forbid-invalid CI, JSON) ✓ 2026-08-02
- [x] dockerfile-extract-onbuild : liste les triggers ONBUILD par stage (détection triggers invalides FROM/ONBUILD/MAINTAINER, --keywords-only, --forbid-chaining CI, JSON) ✓ 2026-08-02
- [x] json-invert-key-values : inverse un objet JSON plat (clés↔valeurs, politiques de collision error/first/last, JSONL, --require-colision-free CI, compact/indent) ✓ 2026-08-02
- [x] text-snake-to-kebab : convertit snake_case en kebab-case par ligne (mode token/whole-line, --lower, filtres regex, --in-place, --check CI, JSON) ✓ 2026-08-02
- [x] env-list-values : liste les valeurs d'un dotenv (export/quotes, --pairs, --mask/--mask-full, --sort/--unique, --require-value/--forbid-empty/min/max CI, JSON) ✓ 2026-08-02

## Vague 366 — CLI Tools (moyenne colonne CSV, delete valeur JSON, swap case lignes, HEALTHCHECK Dockerfile, listing clés dotenv)
- [x] csv-mean-column : moyenne arithmétique d'une colonne numérique CSV (nom/index, délimiteur custom/échappements, --all-rows, --decimals, --require-min/max CI, JSON) ✓ 2026-08-02
- [x] json-delete-value : supprime la valeur à un dot-path JSON/JSONL (indices négatifs, auto-detect JSONL, --compact/--in-place, --require-present CI, JSON) ✓ 2026-08-02
- [x] text-swap-case-lines : inverse la casse de chaque lettre par ligne (Unicode swapcase, --only-letters ASCII, --skip-first, --in-place, --check CI, JSON) ✓ 2026-08-02
- [x] dockerfile-extract-healthcheck : extrait les HEALTHCHECK d'un Dockerfile (flags, exec/shell form, NONE, continuations, --last-only, --require-healthcheck/--require-timeout/--forbid-none CI, JSON) ✓ 2026-08-02
- [x] env-list-keys : liste les clés d'un dotenv (export/bare keys, --with-markers, --unique/--sort/--count, --require-key/min/max CI, JSON) ✓ 2026-08-02 (repo pré-existant, contenu synchronisé)

## Vague 365 — CLI Tools (numérotation CSV, get valeur JSON, uppercase lignes, comptage instructions Dockerfile, title case)
- [x] csv-add-line-number : ajoute une colonne de numérotation à un CSV (--start/--step, --end, --no-header, délimiteur custom) ✓ 2026-08-02
- [x] json-get-value : lit une valeur à un dot-path JSON/JSONL (indices négatifs, --raw-string, --default, manquant = exit 2 CI) ✓ 2026-08-02
- [x] text-uppercase-lines : met chaque ligne en majuscules (Unicode, --turkic i→İ, --skip-first/--only-matching, --check CI, JSON) ✓ 2026-08-02
- [x] dockerfile-count-instructions : compte les instructions d'un Dockerfile par type (-i KW, --require KW=N/--require-min CI, JSON) ✓ 2026-08-02
- [x] text-title-case-lines : convertit chaque ligne en title case (petits mots anglais préservés, --all-words, --check CI, JSON) ✓ 2026-08-02

## Vague 364 — CLI Tools (min/max colonne CSV, set valeur JSON, mot de passe URL, comptage clés dotenv)
- [x] csv-max-column : affiche le max d'une colonne numérique CSV (par nom/index, --all-rows, --require-min/max CI, JSON) ✓ 2026-08-02
- [x] json-set-value : écrit une valeur à un dot-path dans un JSON/JSONL (création de segments, --jsonl, --check CI, compact/indent) ✓ 2026-08-02
- [x] url-extract-password : extrait le mot de passe (userinfo) des URLs (--scan, --mask, --unique, --require-password/--forbid-password CI, JSON) ✓ 2026-08-02
- [x] env-count-keys : compte les clés d'un dotenv (export/comments, doublons, --names, --require-min/max/--forbid-duplicates CI, JSON) ✓ 2026-08-02
- [x] csv-min-column : affiche le min d'une colonne numérique CSV (par nom/index, --all-rows, --require-min/max CI, JSON) ✓ 2026-08-02

## Vague 363 — CLI Tools (comptage colonnes CSV, rapport types JSON, usernames URL, dequotage dotenv, lowercase lignes)
- [x] csv-count-columns : compte les colonnes d'un CSV et détecte les lignes ragged (header/first row, --delimiter, --require-count/min/max, --check CI, JSON) ✓ 2026-08-02
- [x] json-type-report : rapporte le type JSON du document ou de noeuds par dot-path (object/array/string/number/boolean/null, JSONL, --require-type/--forbid-type CI, JSON) ✓ 2026-08-02
- [x] url-extract-username : extrait le username (userinfo) des URLs (--scan texte libre, --with-password, --unique comptes, --require-username/--forbid-password CI, JSON) ✓ 2026-08-02
- [x] env-strip-quotes : retire les quotes entourant les valeurs dotenv (commentaires/export préservés, quotes non matchés conservés, --in-place, --check CI, JSON) ✓ 2026-08-02
- [x] text-lowercase-lines : met chaque ligne en minuscules (Unicode lower/--fold casefold, --skip-first/--skip-matching/--only-matching, --check CI, JSON) ✓ 2026-08-02

## Vague 362 — CLI Tools (somme colonne CSV, extraction valeurs JSON, collapse lignes vides, hosts URL, unset dotenv)
- [x] csv-sum-column : somme et agrégats d'une colonne numérique CSV (count/min/max/mean, --no-header, --require-min/max CI, JSON) ✓ 2026-08-02
- [x] json-extract-values : extrait les valeurs d'une clé dans des objets JSON (JSON/JSONL, --path avec *, --unique/--count, --require-min/max CI, JSON) ✓ 2026-08-02
- [x] text-collapse-blank-lines : réduit les suites de lignes vides à N max (--keep, --strip-blanks, --in-place, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-host : extrait le hostname des URLs (bare hosts, --scan texte libre, --unique fréquences, --check/--forbid-host/--allow CI, JSON) ✓ 2026-08-02
- [x] env-unset-keys : retire des clés d'un dotenv (batch positionnel/--key/--keys-file, --in-place, --check, --require-absent CI, JSON) ✓ 2026-08-02

## Vague 361 — CLI Tools (ARG Dockerfile, suppression colonnes CSV, comptage clés JSON, squeeze espaces, schemes URL)
- [x] dockerfile-extract-arg : liste les ARG d'un Dockerfile (scope global + par stage, défauts, --names-only, --require-arg/--forbid-arg CI, JSON) ✓ 2026-08-02
- [x] csv-delete-columns : supprime des colonnes d'un CSV (noms/indices, --no-header, --require-absent CI, JSON) ✓ 2026-08-02
- [x] json-count-keys : compte les clés d'objets JSON récursivement (JSONL, --path avec *, --per-document, --require-min/max CI, JSON) ✓ 2026-08-02
- [x] text-squeeze-spaces : réduit les répétitions d'espaces/tabs à un seul espace (indent préservée, --no-keep-indent, --strip-eol, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-scheme : extrait le scheme des URLs (lowercase, --scan texte libre, --unique/--count, --require-scheme/--forbid-scheme CI, JSON) ✓ 2026-08-02

## Vague 360 — CLI Tools (EXPOSE Dockerfile, garder colonnes CSV, clés objets JSON, reformatage lignes, ports URL)
- [x] dockerfile-extract-expose : liste les ports EXPOSE d'un Dockerfile par stage (protocole tcp/udp, --forbid-udp/--require-port CI, JSON) ✓ 2026-08-02
- [x] csv-keep-columns : ne garde que certaines colonnes d'un CSV (noms/indices, --drop inverse, --require-column CI, JSON) ✓ 2026-08-02
- [x] json-object-keys-list : liste les clés des objets JSON (document/JSONL, --path avec *, --require-key/--forbid-key CI, JSON) ✓ 2026-08-02
- [x] text-wrap-lines : reformate les lignes longues à une largeur cible (indentation préservée, --join, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-port : extrait le port effectif des URLs (défauts 30+ schemes, --require-range/--forbid-port CI, JSON) ✓ 2026-08-02

## Vague 359 — CLI Tools (COPY Dockerfile, swap colonnes CSV, comptage tableaux JSON, indentation, lectures dotenv)
- [x] dockerfile-extract-copy : liste les instructions COPY d'un Dockerfile par stage (flags --from/--chown, --forbid-from CI, JSON) ✓ 2026-08-02
- [x] csv-swap-columns : permute deux colonnes d'un CSV (header inclus, noms ou indices, --no-header) ✓ 2026-08-02
- [x] json-array-count-items : compte les éléments de tableaux JSON (document/JSONL, --path, --require-min/max CI) ✓ 2026-08-02
- [x] text-indent-lines : indente chaque ligne (espaces/tabs/prefixe, --skip-first, --check CI, JSON) ✓ 2026-08-02
- [x] env-get-value : lit des valeurs d'un fichier dotenv (batch ordonné, --default, --require-nonempty CI, --export, JSON) ✓ 2026-08-02

## Vague 358 — CLI Tools (WORKDIR Dockerfile, réordonnancement colonnes CSV, dédup tableaux JSON, suffixe commun, sections INI)
- [x] dockerfile-extract-workdir : extrait le WORKDIR effectif par stage d'un Dockerfile (last-wins, --final-only, --require-workdir CI, JSON) ✓ 2026-08-02
- [x] csv-reorder-columns : réordonne les colonnes d'un CSV (ordre voulu, --drop-rest, --check CI, JSON) ✓ 2026-08-02
- [x] json-array-unique : déduplique les tableaux JSON en gardant les 1res occurrences (JSONL, --path, --sort, --check CI) ✓ 2026-08-02
- [x] text-strip-suffix : retire un suffixe fixe ou auto-détecté commun de chaque ligne (--ignore-case, --repeat, --check CI, JSON) ✓ 2026-08-02
- [x] ini-section-list : liste les sections d'un INI (compteurs de clés, --require/--forbid CI, JSON) ✓ 2026-08-02

## Vague 357 — CLI Tools (bases de stages Dockerfile, valeurs colonne CSV, merge tableaux JSON, trim gauche, valeurs query URL)
- [x] dockerfile-extract-stage-base : extrait l'image de base de chaque stage FROM d'un Dockerfile (ARG expansés, --check-image/--require-stage CI, JSON) ✓ 2026-08-02
- [x] csv-extract-column-values : imprime les valeurs d'une colonne CSV (nom ou index, --unique/--count, --require-min/max CI, JSON) ✓ 2026-08-02
- [x] json-merge-arrays : concatène les tableaux racine de plusieurs documents JSON (mode JSONL, --path, --unique, --require-min CI) ✓ 2026-08-02
- [x] text-trim-left : retire les espaces/caractères en tête de chaque ligne (--chars, --max, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-query-value : extrait la valeur d'un paramètre query d'URLs en batch (--all, --default, --require-match CI, JSON) ✓ 2026-08-02

## Vague 356 — CLI Tools (USER Dockerfile, comptage regex CSV, slice tableaux JSON, authority URL, lectures INI)
- [x] dockerfile-extract-user : extrait le USER effectif par stage d'un Dockerfile (root implicite détecté, --check non-root CI, JSON) ✓ 2026-08-02
- [x] csv-count-rows-matching : compte les lignes CSV dont les cellules matchent un regex (colonne ciblée, --invert, --print, --require CI, JSON) ✓ 2026-08-02
- [x] json-array-slice : tranche les tableaux de docs JSON/JSONL (start/stop/step, indices négatifs, --path dot-path, --check CI, JSON) ✓ 2026-08-02
- [x] url-parse-authority : éclate des URLs en user/host/port effectif (défauts par scheme, --scan, --fields, --json, --check CI) ✓ 2026-08-02
- [x] ini-get-value : lit des valeurs INI par section.key (batch stdin, --default, --fallback-section, --strict CI, JSON) ✓ 2026-08-02

## Vague 355 — CLI Tools (colonne CSV, minify JSON stream, padding lignes, ENTRYPOINT Dockerfile, remotes git)
- [x] csv-append-column : ajoute une colonne à un CSV (valeur constante ou template {field}, --overwrite, --check CI, JSON) ✓ 2026-08-02
- [x] json-minify-stream : minifie des documents JSON en une ligne chacun (mode JSONL, --sort-keys, --ascii, --check CI, JSON) ✓ 2026-08-02
- [x] text-pad-right : pousse chaque ligne à une largeur cible (right/left/center, --fill, --truncate, --check CI, JSON) ✓ 2026-08-02
- [x] dockerfile-extract-entrypoint : extrait l'ENTRYPOINT effectif par stage d'un Dockerfile (formes exec/shell, continuations, --check CI, JSON) ✓ 2026-08-02
- [x] git-remote-normalize : normalise les URLs de remotes git entre formes canoniques SSH/HTTPS (stdin/args/--repo, --check CI, JSON) ✓ 2026-08-02

## Vague 354 — CLI Tools (premières lignes CSV, nombres JSON, chomp lignes vides, casse path URL, CMD Dockerfile)
- [x] csv-remove-first-rows : retire les N premières lignes de données d'un CSV (header conservé, --no-header, --check CI, JSON) ✓ 2026-08-02
- [x] json-normalize-numbers : normalise les floats d'un JSON/JSONL (round --decimals, --strip-trailing 2.0→2, auto-detect JSONL, --check CI, JSON) ✓ 2026-08-02
- [x] text-chomp-lines : retire les lignes vides finales d'un texte (--keep N, --no-newline, --in-place, --check CI, JSON) ✓ 2026-08-02
- [x] url-normalize-path-case : met le path des URLs en minuscules (query/fragment préservés, --scan texte libre, --check CI, JSON) ✓ 2026-08-02
- [x] dockerfile-extract-cmd : extrait le CMD effectif par stage d'un Dockerfile (formes exec/shell, last-CMD-wins, continuations, --check CI, JSON) ✓ 2026-08-02

## Vague 353 — CLI Tools (trim cellules CSV, champ fixe JSON, lignes répétées, slash final URL, ENV Dockerfile)
- [x] csv-trim-cells : retire les espaces début/fin des cellules d'un CSV (colonnes ciblées, --report, --check CI, JSON) ✓ 2026-08-02 (repo pré-existant, contenu synchronisé)
- [x] json-append-field : ajoute une clé fixe k=v à chaque objet JSON d'un flux JSONL (types auto, --overwrite, --check CI, JSON) ✓ 2026-08-02
- [x] text-repeat-line-report : rapport des lignes identiques répétées consécutivement (--ignore-case, --min-repeats, --check CI, JSON) ✓ 2026-08-02
- [x] url-strip-trailing-slash : retire le slash final des paths d'URLs (racine conservée, --strip-root, --check CI, JSON) — Go ✓ 2026-08-02
- [x] dockerfile-extract-env : extrait les variables ENV définies dans un Dockerfile (valeurs, stage, redéfinitions, --check CI, JSON) — Go ✓ 2026-08-02

## Vague 352 — CLI Tools (lignes vides CSV, champ index JSON, compte mots, ancre URL, stages Dockerfile)
- [x] csv-count-empty-rows : rapport lignes vides et cellules vides d'un CSV (par colonne, --check CI, JSON) ✓ 2026-08-02
- [x] json-add-index-field : ajoute un champ index auto-incrémenté à chaque objet d'un array JSON ou flux JSONL (--start/--step, fallback JSONL, --check CI, JSON) ✓ 2026-08-02
- [x] text-word-count-report : compte les mots par ligne et au total (stats min/max/moyenne, histogramme, --top, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-anchor : extrait le fragment (ancre) des URLs d'un flux ou texte libre (--scan, --unique, stats top ancres, --check CI, JSON) ✓ 2026-08-02
- [x] dockerfile-parse-stages : parse les stages d'un Dockerfile multi-stage (base, COPY --from, EXPOSE, ENTRYPOINT, --check-stages CI, JSON) — Go ✓ 2026-08-02

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

## Vague 284 — CLI Tools (médiane CSV, chemins JSON en table, déwrap texte, domaine de base URL, lignes par extension)
- [x] csv-median-column : calcule la médiane d'une colonne numérique CSV (interpolation paire, --check bornes CI, JSON) ✓ 2026-08-01
- [x] json-paths-as-table : rend tous les chemins feuilles d'un JSON en tableau colonnes séparées (orienté diff, --check CI, JSON) ✓ 2026-08-01
- [x] text-strip-word-wrap : défait un word-wrap et rejoint les lignes en paragraphes (heuristique, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-base-domain : extrait le domaine enregistré approximatif (eTLD+1 simple) de chaque URL (--check CI, JSON) ✓ 2026-08-01
- [x] file-count-lines-per-ext : compte les lignes des fichiers texte par extension dans une arborescence (binaires exclus, --check CI, JSON) ✓ 2026-08-01

## Vague 285 — CLI Tools (filtrage par numéro de ligne CSV, diff tableaux JSONL, tri mots, scheme par défaut URL, min/max mtime par extension)
- [x] csv-row-number-filter : garde les lignes dont le numéro matche une expression (ranges 1-5,8, mod N, --check CI, JSON) ✓ 2026-08-01
- [x] json-compare-arrays : compare deux tableaux JSONL élément par élément et rapporte les écarts (index, --check CI, JSON) ✓ 2026-08-01
- [x] text-reorder-words : trie les mots de chaque ligne alphabétiquement (ignore-case, --check CI, JSON) ✓ 2026-08-01
- [x] url-add-default-scheme : ajoute https:// aux URLs sans scheme (exceptions bare-IP, --check CI, JSON) ✓ 2026-08-01
- [x] file-newest-oldest-per-ext : plus récent/plus ancien fichier par extension (mtime, --check CI, JSON) ✓ 2026-08-01

## Vague 288 — CLI Tools (suppression CSV, tri clés JSON, mots uniques, strip query, fichiers récents)
- [x] csv-remove-column : supprime une ou plusieurs colonnes d'un CSV (par nom ou index, --require CI, JSON) ✓ 2026-08-01
- [x] json-sort-keys-deep : trie les clés de tous les objets d'un JSON/JSONL récursivement (arrays préservés, --check CI, JSON) ✓ 2026-08-01
- [x] text-unique-words : extrait les mots uniques d'un texte (normalisation, tri, --check CI, JSON) ✓ 2026-08-01
- [x] url-remove-query-param : retire des paramètres query ciblés des URLs (blacklist, keep-blank, --check CI, JSON) ✓ 2026-08-01
- [x] file-recent-hours : liste les fichiers modifiés dans les N dernières heures (--hours, --check CI, JSON) ✓ 2026-08-01

## Vague 289 — CLI Tools (à faire)
- [x] csv-trim-cells : retire les espaces de début/fin de chaque cellule CSV (--check CI, JSON) ✓ 2026-08-01 (déjà publié)
- [x] json-renumber-array-field : réécrit un champ entier séquentiel dans les objets d'un JSONL (start/step, --check CI, JSON) ✓ 2026-08-01
- [x] text-squeeze-repeats : réduit les répétitions consécutives d'un caractère à N occurrences (--char, --check CI, JSON) ✓ 2026-08-01
- [x] url-swap-host-path-prefix : déplace le premier segment du path vers le host (a.b.com/p1 -> p1.a.b.com, --check CI, JSON) ✓ 2026-08-01
- [x] file-count-by-hour-of-day : histogramme des fichiers par heure de mtime (0-23, --check CI, JSON) ✓ 2026-08-01

## Vague 290 — CLI Tools (commentaires CSV, arrondi JSON, blanks finaux, valeur query, cibles symlinks)
- [x] csv-drop-comment-rows : retire les lignes commençant par un préfixe commentaire (#, //, --prefix, --check CI, JSON) ✓ 2026-08-01
- [x] json-round-numbers : arrondit tous les nombres flottants d'un JSON/JSONL à N décimales (int préservés, --check CI, JSON) ✓ 2026-08-01
- [x] text-rstrip-blank-tail : supprime les lignes vides en fin de fichier (--head aussi, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-query-value : extrait la valeur d'un paramètre query donné par URL (--all occurrences, --check CI, JSON) ✓ 2026-08-01
- [x] file-symlink-target-report : rapport des cibles de symlinks (absolue/relative, existante/cassée, --check CI, JSON) ✓ 2026-08-01

## Vague 291 — CLI Tools (guillemets CSV, unflatten JSON, whitespace visible, port défaut URL, sous-arbres dupliqués)
- [x] csv-quoted-cells-report : rapport des cellules CSV qui nécessitent des guillemets (--check CI, JSON) ✓ 2026-08-01
- [x] json-unflatten-keys : reconstruit un JSON depuis des paires chemin=valeur (dot-paths, arrays indexés, --check CI, JSON) ✓ 2026-08-01
- [x] text-visualize-whitespace : rend visibles espaces/tabs/fins de ligne (·, →, $, --check CI, JSON) ✓ 2026-08-01
- [x] url-normalize-default-port : retire le port par défaut explicite des URLs (:80, :443, --check CI, JSON) ✓ 2026-08-01
- [x] file-duplicate-dir-trees : détecte des sous-arbres de dossiers au contenu identique (hash manifeste, --check CI, JSON) ✓ 2026-08-01

## Vague 292 — CLI Tools (header CSV, inversion JSON, strip numéros, période URL, jour semaine mtime)
- [x] csv-first-row-as-header : promeut la première ligne comme header si absent (synthèse col_N, --check CI, JSON) ✓ 2026-08-01
- [x] json-invert-mapping : inverse un objet clé->valeur en valeur->clés (arrays pour collisions, --check CI, JSON) ✓ 2026-08-01
- [x] text-strip-line-numbers : retire une numérotation de début de ligne (N:, N., [N], --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-year-month : extrait /YYYY/ ou /YYYY/MM/ du path des URLs (stats par période, --check CI, JSON) ✓ 2026-08-01
- [x] file-group-by-weekday : histogramme des fichiers par jour de semaine de mtime (--require DAYS CI, JSON) ✓ 2026-08-01

## Vague 293 — CLI Tools (dates CSV, moyenne JSON, guillemets typo, filename URL, longueur chemins)
- [x] csv-detect-date-column : identifie les colonnes contenant des dates (formats variés, score, --check CI, JSON) ✓ 2026-08-01
- [x] json-average-path : moyenne arithmétique des valeurs numériques à un dot-path à travers JSONL (--check CI, JSON) ✓ 2026-08-01
- [x] text-replace-smart-quotes : convertit guillemets/apostrophes typographiques en ASCII (“”‘’ -> ", --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-filename : extrait le segment filename final des URLs (avec/sans extension, --check CI, JSON) ✓ 2026-08-01
- [x] file-path-length-report : rapport des chemins dépassant une longueur seuil (PATH_MAX, --max CI, JSON) ✓ 2026-08-01

## Vague 294 — CLI Tools (types colonnes CSV, échappement JSON, tabs, JWT URL, mtime par nom)
- [x] csv-column-type-report : rapport du type dominant par colonne CSV (int/float/date/string/bool, score, --check CI, JSON) ✓ 2026-08-01
- [x] json-escape-strings : échappe les caractères spéciaux des valeurs string d'un JSONL (\n, \t, \uXXXX, --unescape, --check CI, JSON) ✓ 2026-08-01
- [x] text-expand-tabs : convertit les tabulations en espaces (tabstop custom, --in-place, --check CI, JSON) ✓ 2026-08-01
- [x] url-parse-jwt-segment : décode les segments JWT (header.payload) présents dans les URLs (base64url, --check CI, JSON) ✓ 2026-08-01
- [x] file-mtime-set-from-name : applique le mtime depuis un timestamp présent dans le nom de fichier (regex, --dry-run, --check CI, JSON) ✓ 2026-08-01

## Vague 295 — CLI Tools (unicité CSV, types JSON, préfixe commun, schème auth URL, octets NUL)
- [x] csv-column-unique-check : vérifie l'unicité des valeurs d'une ou plusieurs colonnes CSV (rapport doublons, --check CI, JSON) ✓ 2026-08-01
- [x] json-count-by-type : compte les valeurs JSON par type (null/bool/number/string/array/object) à travers JSONL (--check CI, JSON) ✓ 2026-08-01
- [x] text-prefix-common-strip : retire le préfixe commun à toutes les lignes (auto-detect, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-auth-scheme : classe les URLs par schème d'auth (basic userinfo, bearer param, api_key query, --check CI, JSON) ✓ 2026-08-01
- [x] file-content-null-bytes : détecte les fichiers contenant des octets NUL dans une arborescence (--check CI, JSON) ✓ 2026-08-01

## Vague 296 — CLI Tools (ordre colonnes CSV, timestamps JSON, IP littérales URL, caractères contrôle)
- [x] csv-column-order : réordonne les colonnes d'un CSV selon une liste (colonnes manquantes en fin, --check CI, JSON) ✓ 2026-08-01
- [x] json-path-exists : vérifie qu'un dot-path existe dans chaque document JSONL (--require PATHS, --check CI, JSON) ✓ 2026-08-01 (déjà publié Vague 273)
- [x] text-column-align : aligne les champs d'un texte en colonnes à largeur fixe (séparateur custom, --check CI, JSON) ✓ 2026-08-01 (déjà publié)
- [x] url-extract-ip : liste les URLs utilisant une IP littérale au lieu d'un hostname (v4/v6, --check CI, JSON) ✓ 2026-08-01
- [x] file-permission-audit : rapport des fichiers avec permissions inhabituelles (world-writable, setuid, --check CI, JSON) ✓ 2026-08-01 (déjà publié Vague 260)
- [x] csv-count-empty-cells : rapport des cellules vides par colonne CSV (pourcentages, seuils --max-empty CI, JSON) ✓ 2026-08-01
- [x] json-detect-timestamps : détecte et classe les valeurs timestamp d'un JSON/JSONL (iso8601/epoch s/ms/us, --require-format CI, JSON) ✓ 2026-08-01
- [x] text-replace-non-printable : remplace/retire/échappe les caractères de contrôle non imprimables (modes replace/remove/escape, --check CI, JSON) ✓ 2026-08-01

## Vague 297 — CLI Tools (hash lignes CSV, join JSONL, phrases, ports URL, rebase chemins)
- [x] csv-row-hash : ajoute une colonne hash par ligne CSV pour détecter les doublons de lignes entières (--check CI, JSON) ✓ 2026-08-01
- [x] json-merge-lines : fusionne deux fichiers JSONL clé par clé (join inner/left/right/full, --check CI, JSON) ✓ 2026-08-01
- [x] text-split-sentences : découpe un texte en phrases (ponctuation, abréviations en/fr, --check CI, JSON) ✓ 2026-08-01
- [x] url-normalize-port-default : retire les ports par défaut explicites (80/443/21..., --add-default inverse, --check CI, JSON) ✓ 2026-08-01
- [x] file-chdir-relative-paths : rend les chemins relatifs d'un manifeste par rapport à un nouveau dossier de base (--check CI, JSON) ✓ 2026-08-01

## Vague 298 — CLI Tools (remplissage CSV, nulls JSON, ellipse milieu, fuites credentials URL, noms de fichiers)
- [x] csv-column-fill-forward : propage la dernière valeur non vide vers le bas dans des colonnes CSV (groupes par clé, --check CI, JSON) ✓ 2026-08-01
- [x] json-count-null-fields : compte les champs null par chemin dans un JSONL (top-N, --threshold CI, JSON) ✓ 2026-08-01
- [x] text-truncate-middle : tronque les lignes en conservant début et fin avec ellipse (largeur max, --check CI, JSON) ✓ 2026-08-01
- [x] url-detect-credential-leak : repère les URLs embarquant des secrets (userinfo user:pass@, token=/api_key=/auth= en query, masquage, --check CI, JSON) ✓ 2026-08-01
- [x] file-duplicate-basename : liste les noms de fichiers en double dans une arborescence regardless of dir (match exact ou normalisé, --check CI, JSON) ✓ 2026-08-01

## Vague 299 — CLI Tools (colonnes CSV, longueurs tableaux JSON, lignes longues, fragments URL, âge fichiers)
- [x] csv-remove-columns : retire des colonnes d'un CSV par nom ou index (blacklist, inverse de keep, --check CI, JSON) ✓ 2026-08-01 (déjà publié)
- [x] json-count-array-lengths : statistiques sur les longueurs de tableaux par dot-path (min/max/mean, --threshold CI, JSON) ✓ 2026-08-01 (déjà publié)
- [x] text-longest-line-report : rapport des lignes les plus longues d'un texte (top-N, largeur moyenne, --max-width CI, JSON) ✓ 2026-08-01
- [x] url-extract-fragment-keys : éclate les fragments de type #k=v&k2=v2 en paires clé=valeur (SPA routes, --check CI, JSON) ✓ 2026-08-01
- [x] file-age-buckets-report : répartit les fichiers d'une arborescence par tranche d'âge (<1h, <1j, <1sem, <1mois, plus, --check CI, JSON) ✓ 2026-08-01

## Vague 300 — CLI Tools (jointure CSV, Unicode JSON, paragraphes, encodage query, modes octaux)
- [x] csv-join-two-files : jointure de deux CSV sur une colonne clé (inner/left/right/full, --require-matches CI, JSON) ✓ 2026-08-01
- [x] json-normalize-unicode : applique une normalisation Unicode NFC/NFKC aux clés et valeurs string d'un JSON/JSONL (--check conformité CI, JSON) ✓ 2026-08-01
- [x] text-paragraph-split : découpe un texte en paragraphes (blocs séparés par lignes vides, indices, --require RANGE CI, JSON) ✓ 2026-08-01
- [x] url-canonicalize-query-encoding : ré-encode uniformément les valeurs query des URLs (percent-encoding minimal, --check CI, JSON) ✓ 2026-08-01
- [x] file-permission-octal-report : liste les fichiers par mode octal exact d'une arborescence (histogramme, --allow CI, JSON) ✓ 2026-08-01

## Vague 301 — CLI Tools (longueurs cellules CSV, inversion index JSON, similarité, tri segments URL, chaînes symlinks)
- [x] csv-value-length-report : statistiques de longueur des cellules par colonne CSV (min/max/mean, --max-len CI, JSON) ✓ 2026-08-01
- [x] json-invert-array-index : transforme {clé: [items]} en {item: [clés]} (inversion d'index, --require-items CI, JSON) ✓ 2026-08-01
- [x] text-line-diff-ratio : ratio de similarité ligne-à-ligne entre deux fichiers (SequenceMatcher, --min-ratio CI, JSON) ✓ 2026-08-01
- [x] url-sort-path-segments : trie les segments du path d'URLs alphabétiquement (premier ancré, --check CI, JSON) ✓ 2026-08-01
- [x] file-symlink-chain-report : résout les chaînes de symlinks d'une arborescence (longueur, boucles, cibles cassées, --max-depth CI, JSON) ✓ 2026-08-01

## Vague 302 — CLI Tools (bfill CSV, wildcard JSON, runs caractères, casse segments URL, contenu stdin)
- [x] csv-column-fill-backward : propage la première valeur non vide vers le haut dans des colonnes CSV (groupes par clé, --check CI, JSON) ✓ 2026-08-01
- [x] json-extract-path-values : extrait toutes les valeurs sous un dot-path incluant tableaux (wildcard *, --require CI, JSON) ✓ 2026-08-01
- [x] text-count-char-runs : compte les runs d'un caractère donné dans chaque ligne (--char, --min-run CI, JSON) ✓ 2026-08-01
- [x] url-normalize-segment-case : met les segments path d'URLs en minuscules sauf segments protégés (regex, --check CI, JSON) ✓ 2026-08-01
- [x] file-same-content-as-stdin : liste les fichiers d'un dossier dont le contenu matche exactement stdin (sha256, --require CI, JSON) ✓ 2026-08-01

## Vague 303 — CLI Tools (merge non-vide CSV, canonical JSON, purge contexte, localhost URL, collisions noms)
- [x] csv-last-non-empty-row : fusionne les lignes CSV partageant une clé, dernière valeur non vide par colonne (--check CI, JSON) ✓ 2026-08-01
- [x] json-round-trip-check : vérifie qu'un JSON re-sérialisé canoniquement est identique octet-pour-octet (sorted keys, --write, --check CI, JSON) ✓ 2026-08-01
- [x] text-strip-matching-lines : retire les lignes matchant un pattern ET leurs N lignes de contexte (--before/--after, --check CI, JSON) ✓ 2026-08-01
- [x] url-detect-localhost : détecte les URLs pointant vers localhost/loopback/privé/.local (report, --check CI, JSON) ✓ 2026-08-01
- [x] file-name-collision-report : détecte les fichiers dont les noms entrent en collision après normalisation (lower, accents, --dirs, --check CI, JSON) ✓ 2026-08-01

## Vague 304 — CLI Tools (ordre colonnes CSV, chemins floats JSON, inversion casse, eTLD+1 URL, blocs disque)
- [x] csv-detect-out-of-order : détecte les lignes dont une colonne clé n'est pas monotonique (asc/desc, --check CI, JSON) ✓ 2026-08-01
- [x] json-extract-float-paths : liste les chemins contenant des floats non entiers avec statistiques (--check CI, JSON) ✓ 2026-08-01
- [x] text-invert-case : inverse la casse des lettres ligne par ligne (upper<->lower, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-registered-domain-rdap-hint : classe les domaines d'URLs par eTLD+1 approximatif + TLD connu inconnu (--check CI, JSON) ✓ 2026-08-01
- [x] file-block-count-report : rapport du nombre de blocs disque (st_blocks) par fichier (--threshold CI, JSON) ✓ 2026-08-01

## Vague 305 — CLI Tools (valeurs manquantes CSV, longueurs strings JSONL, tabs→spaces, auth HTTP URL, hardlinks)
- [x] csv-missing-value-report : rapport des valeurs manquantes par colonne CSV (null/NA/vide, --max-missing CI, JSON) ✓ 2026-08-01
- [x] json-extract-string-lengths : distribution des longueurs de chaînes par chemin dans un JSONL (--max-len CI, JSON) ✓ 2026-08-01
- [x] text-tabs-to-spaces : convertit les tabs en espaces avec largeur configurable (--tab-width, --check CI, JSON) ✓ 2026-08-01
- [x] url-detect-http-auth : détecte les URLs nécessitant une auth (basic-auth, /admin, query tokens, --check CI, JSON) ✓ 2026-08-01
- [x] file-hardlink-tree : reconstruit l'arbre des liens durs par inode avec root commun (--check CI, JSON) ✓ 2026-08-01

## Vague 306 — CLI Tools (histogramme CSV, fréquence clés JSON, strip non-ASCII, eTLD+1 URL, classification contenu)
- [x] csv-column-histogram : histogramme visuel ASCII par colonne numérique CSV (buckets, --width, --check CI, JSON) ✓ 2026-08-01
- [x] json-key-frequency : fréquence d'apparition de chaque clé à travers les docs JSONL (--top, --check CI, JSON) ✓ 2026-08-01
- [x] text-strip-non-ascii : retire les caractères non-ASCII de chaque ligne (--replace X, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-domains-tld-aware : extrait unique domaines depuis flux d'URLs en respectant l'eTLD+1 (--check CI, JSON) ✓ 2026-08-01
- [x] file-content-classify : classe fichiers par contenu (code/data/media/docs) via magic+extension (--check CI, JSON) ✓ 2026-08-01

## Vague 307 — CLI Tools (jour semaine CSV, couverture types JSON, alignement commentaires, segment URL, mois mtime)
- [x] csv-add-weekday-column : ajoute une colonne jour de semaine dérivé d'une colonne date CSV (locale en/fr, --check CI, JSON) ✓ 2026-08-01
- [x] json-type-coverage : couverture des types JSON par chemin à travers JSONL (type union, --check CI, JSON) ✓ 2026-08-01
- [x] text-align-comments : aligne les commentaires en fin de ligne sur une colonne fixe (#, //, --check CI, JSON) ✓ 2026-08-01
- [x] url-replace-path-segment : remplace le N-ième segment du path d'URLs en flux (1-based, négatifs, --check CI, JSON) ✓ 2026-08-01
- [x] file-count-by-month : histogramme des fichiers par mois de mtime (YYYY-MM, --require PATTERN CI, JSON) ✓ 2026-08-01

## Vague 308 — CLI Tools (âge jour CSV, coalescence chemins JSON, runs consonnes, tri valeurs query URL, buckets taille)
- [x] csv-add-age-days : ajoute une colonne âge en jours depuis une colonne date CSV (référence --as-of, --check CI, JSON) ✓ 2026-08-01
- [x] json-coalesce-paths : fusionne plusieurs dot-paths alias en un seul canonique (--prefer first/last, --check CI, JSON) ✓ 2026-08-01
- [x] text-collapse-consonant-runs : réduit les runs de consonnes répétées (>2) dans chaque mot (--keep N, --check CI, JSON) ✓ 2026-08-01
- [x] url-query-sort-values : trie les valeurs multi-occurrences de chaque clé query (dedup optionnel, --check CI, JSON) ✓ 2026-08-01
- [x] file-group-by-size-prefix : regroupe les fichiers par préfixe de taille lisible (K/M/G buckets, --check CI, JSON) ✓ 2026-08-01

## Vague 309 — CSV/JSON/text/URL/file mix
- [x] csv-numeric-format : reformate les colonnes numériques (décimales, séparateur milliers, --check CI, JSON) ✓ 2026-08-02
- [x] json-flatten-to-csv : aplatit chaque objet JSONL en une ligne CSV (union des clés, --check CI, JSON) ✓ 2026-08-02
- [x] text-average-line-length : rapport des longueurs de lignes (min/max/mean, --max-len CI, JSON) ✓ 2026-08-02
- [x] url-drop-query-keys : retire des clés query spécifiques des URLs (keep-list inverse, --check CI, JSON) ✓ 2026-08-02
- [x] file-snapshot-manifest : génère un manifeste chemin+taille+sha256 d'une arborescence (--check manifeste CI, JSON) ✓ 2026-08-02

## Vague 310 — CSV/JSON/text/URL/file mix
- [x] csv-swap-columns : échange deux colonnes d'un CSV (par nom ou index, --check CI, JSON) ✓ 2026-08-02
- [x] json-unwrap-single-key : extrait la valeur d'un champ racine unique par ligne JSONL (--key, --check CI, JSON) ✓ 2026-08-02
- [x] text-collapse-blank-runs : réduit les runs de lignes vides à N max (--max N, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-id-path : extrait le dernier segment numérique/slug du path d'URLs (--pattern, --check CI, JSON) ✓ 2026-08-02
- [x] file-extension-report : rapport par extension (count, taille totale/moyenne, --top N CI, JSON) ✓ 2026-08-02

## Vague 311 — CSV/JSON/text/URL/file mix
- [x] csv-row-filter-regex : garde les lignes CSV dont une colonne matche un regex (--invert, --check CI, JSON) ✓ 2026-08-02 (déjà publié Vague 206)
- [x] json-sort-keys-recursive : trie récursivement les clés d'objets JSONL (ordre alpha, --check CI, JSON) ✓ 2026-08-02 (déjà publié Vague 237)
- [x] text-wrap-hard-crlf : normalise les fins de lignes CRLF/LF/CR vers LF (--check CI, JSON) ✓ 2026-08-02
- [x] url-compose-from-parts : reconstruit des URLs depuis des champs tabulés scheme/host/path/query (--check CI, JSON) ✓ 2026-08-02
- [x] file-newest-per-dir : trouve le fichier le plus récent par dossier d'une arborescence (--check CI, JSON) ✓ 2026-08-02 (déjà publié Vague 240)

## Vague 312 — CSV/JSON/text/URL/file mix
- [x] csv-trim-cells : retire les espaces de début/fin des cellules d'un CSV (--check CI, JSON) ✓ 2026-08-02 (déjà publié)
- [x] json-pretty-print : reformate un JSON/JSONL avec indentation et ordre de clés optionnels (--indent N, --check CI, JSON) ✓ 2026-08-02
- [x] text-count-paragraphs : statistiques sur les paragraphes d'un texte (count, taille min/max/mean, --check CI, JSON) ✓ 2026-08-02
- [x] url-strip-userinfo : retire les informations user:pass@ des URLs (masquage, --check CI, JSON) ✓ 2026-08-02
- [x] file-count-by-size-range : histogramme des fichiers par tranche de taille (--ranges, --check CI, JSON) ✓ 2026-08-02

## Vague 313 — CSV/JSON/text/URL/file mix
- [x] csv-column-stats-summary : résumé statistique global par colonne CSV (count/empty/unique/top/mean, --check CI, JSON) ✓ 2026-08-02
- [x] json-extract-top-level-keys : liste les clés racine d'un document JSON avec leurs types et cardinalités (--check CI, JSON) ✓ 2026-08-02
- [x] text-suffix-common-strip : retire le suffixe commun à toutes les lignes (auto-detect, --check CI, JSON) ✓ 2026-08-02
- [x] url-detect-static-assets : classe les URLs par type de ressource statique (css/js/img/font/doc, --check CI, JSON) ✓ 2026-08-02
- [x] file-group-by-month-total : regroupe les fichiers par mois de mtime avec taille totale (--check CI, JSON) ✓ 2026-08-02

## Vague 314 — CSV/JSON/text/URL/file mix
- [x] csv-header-rename-sanitize : normalise les noms de colonnes CSV (lower_snake, strip accents, --check CI, JSON) ✓ 2026-08-02
- [x] json-dedupe-array-by-key : déduplique les tableaux d'objets JSON par clé (--keep first/last, --check CI, JSON) ✓ 2026-08-02
- [x] text-extract-urls-cli : extrait les URLs brutes d'un texte (http/https/ftp, dedup, --check CI, JSON) ✓ 2026-08-02
- [x] url-sort-query-params-stable : trie les clés query en préservant les doublons (--check CI, JSON) ✓ 2026-08-02
- [x] file-extension-case-normalize : uniformise la casse des extensions en minuscules (--apply, --check CI, JSON) ✓ 2026-08-02

## Vague 315 — CSV/JSON/text/URL/file mix
- [x] csv-row-shuffle : mélange aléatoirement les lignes d'un CSV (seed reproductible, header préservé, --check CI, JSON) ✓ 2026-08-02
- [x] json-invert-mapping : inverse un mapping JSON clé->valeur en valeur->clés (collisions en array, --check CI, JSON) ✓ 2026-08-02
- [x] text-extract-ips : extrait les adresses IPv4/IPv6 d'un texte (validation stricte, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-anchor-text : associe chaque URL à un texte d'ancre markdown/html adjacent ([text](url) ou <a>, --check CI, JSON) ✓ 2026-08-02
- [x] file-broken-symlink-detect : liste les liens symboliques morts d'une arborescence (--prune option, --check CI, JSON) ✓ 2026-08-02

## Vague 316 — CSV/JSON/text/URL/file mix
- [x] csv-swap-column-order : réordonne les colonnes d'un CSV via nom1,nom2,... (--check CI, JSON) ✓ 2026-08-02
- [x] json-pluck-values : extrait toutes les valeurs scalaires d'un JSON/JSONL (path, --check CI, JSON) ✓ 2026-08-02
- [x] text-strip-control-chars : retire les caractères de contrôle non-imprimables sauf tab/NL (--check CI, JSON) ✓ 2026-08-02
- [x] url-punycode-decode : décode les hosts IDN punycode (xn--) en Unicode (--encode inverse, --check CI, JSON) ✓ 2026-08-02
- [x] file-same-content-pairs-rolling : détecte les paires de fichiers au contenu identique via rolling hash préfiltre (fast, --check CI, JSON) ✓ 2026-08-02

## Vague 317 — CSV/JSON/text/URL/file mix
- [x] csv-number-format-check : vérifie qu'une colonne CSV ne contient que des nombres bien formatés (thousands, decimales, --check CI, JSON) ✓ 2026-08-02
- [x] json-merge-prefer-non-null : fusionne deux JSONL en préférant les valeurs non-null par clé (--check CI, JSON) ✓ 2026-08-02
- [x] text-extract-quoted-lines : garde les lignes contenant une chaîne quotée (--quote char, --invert, --check CI, JSON) ✓ 2026-08-02
- [x] url-path-depth-limit : filtre les URLs dont la profondeur de path dépasse N (--check CI, JSON) ✓ 2026-08-02
- [x] file-group-by-first-byte : regroupe les fichiers d'une arborescence par leur premier octet magic (--check CI, JSON) ✓ 2026-08-02

## Vague 318 — CSV/JSON/text/URL/file mix
- [x] csv-row-position-shift : décale les lignes d'un CSV de N positions (wrap-around, --check CI, JSON) ✓ 2026-08-02
- [x] json-array-of-arrays-to-csv : rend un tableau de tableaux JSON en CSV (--header, --check CI, JSON) ✓ 2026-08-02
- [x] text-markup-strip : retire le markup léger (bold/italique/html simple) d'un texte (--check CI, JSON) ✓ 2026-08-02
- [x] url-hostname-rotate : remplace le host d'URLs par rotation sur une liste (--list, --seed, --check CI, JSON) ✓ 2026-08-02
- [x] file-mtime-range-filter : ne garde que les fichiers dont le mtime est dans [min,max] (--since/--until, --check CI, JSON) ✓ 2026-08-02

## Vague 319 — CSV/JSON/text/URL/file mix
- [x] csv-shuffle-rows : mélange aléatoirement les lignes d'un CSV (seed, header conservé, --check CI, JSON) ✓ 2026-08-02
- [x] json-values-unique : compte les valeurs uniques par chemin dans un JSONL (top-N, --check CI, JSON) ✓ 2026-08-02
- [x] text-swap-case-lines : inverse la casse des lignes alternées ou selon un motif (--every N, --check CI, JSON) ✓ 2026-08-02
- [x] url-default-port-strip : retire le port par défaut des URLs (80/443 selon scheme, --check CI, JSON) ✓ 2026-08-02
- [x] file-empty-dirs-prune : liste et supprime optionnellement les dossiers vides d'une arborescence (--apply, --check CI, JSON) ✓ 2026-08-02

## Vague 320 — CSV/JSON/text/URL/file mix
- [x] csv-drop-columns : supprime des colonnes d'un CSV par nom ou index (keep-list inverse, --check CI, JSON) ✓ 2026-08-02
- [x] json-booleans-report : rapport des champs booléens par chemin dans un JSONL (true/false counts, --check CI, JSON) ✓ 2026-08-02
- [x] text-indent-convert : convertit indentation tabs<->espaces d'un texte (--width N, --check CI, JSON) ✓ 2026-08-02
- [x] url-add-trailing-slash : normalise les paths d'URLs avec slash final (exclude extensions, --check CI, JSON) ✓ 2026-08-02
- [x] file-newest-per-dir : affiche le fichier le plus récent de chaque dossier d'une arborescence (--check CI, JSON) ✓ 2026-08-02

## Vague 321 — CSV/JSON/text/URL/file mix
- [x] csv-row-filter-regex : filtre les lignes CSV où une colonne matche un regex (--invert, --check CI, JSON) ✓ 2026-08-02
- [x] json-null-fields-report : rapport des champs null par chemin dans un JSONL (% null, --check CI, JSON) ✓ 2026-08-02
- [x] text-wrap-diff : compare le wrapping de deux textes ligne à ligne (--width, --check CI, JSON) ✓ 2026-08-02
- [x] url-dedupe-canonical : déduplique des URLs après normalisation canonique (lowercase host, sort query, --check CI, JSON) ✓ 2026-08-02
- [x] file-content-grep-count : compte les occurrences d'un motif par fichier d'une arborescence (totaux, --threshold CI, JSON) ✓ 2026-08-02

## Vague 322 — CSV/JSON/text/URL/file mix
- [x] csv-add-line-hash : ajoute une colonne avec le hash de chaque ligne (md5/sha256, --check CI, JSON) ✓ 2026-08-02
- [x] json-extract-strings : liste toutes les chaînes d'un document JSON avec leur chemin (min-length, --check CI, JSON) ✓ 2026-08-02
- [x] text-first-words : garde les N premiers mots de chaque ligne (--ellipsis, --check CI, JSON) ✓ 2026-08-02
- [x] url-scheme-downgrade-check : détecte les URLs https ayant un équivalent http dans la liste (--check CI, JSON) ✓ 2026-08-02
- [x] file-duplicate-names : liste les noms de fichiers apparaissant en plusieurs endroits d'une arborescence (--check CI, JSON) ✓ 2026-08-02

## Vague 323 — CSV/JSON/text/URL/file mix
- [x] csv-sort-by-column : trie les lignes d'un CSV par une colonne (numérique, naturelle, alpha, --reverse, --check CI, JSON) ✓ 2026-08-02
- [x] json-path-value-set : assigne une valeur à un chemin dot dans un JSONL (create missing, arrays indexés, --check CI, JSON) ✓ 2026-08-02
- [x] text-column-extract : extrait une colonne N d'un texte whitespace-séparé (1-based, indices négatifs, --check CI, JSON) ✓ 2026-08-02
- [x] url-query-remove-params : retire des paramètres nommés de la query des URLs (--keep-only inverse, --check CI, JSON) ✓ 2026-08-02
- [x] file-age-buckets : regroupe les fichiers par tranche d'âge (boundary custom, top N oldest, --check CI, JSON) ✓ 2026-08-02

## Vague 324 — CSV/JSON/text/URL/file mix
- [x] csv-normalize-whitespace : nettoie les espaces dans les cellules CSV (trim, collapse, colonnes ciblées, --check CI, JSON) ✓ 2026-08-02
- [x] json-array-shuffle : mélange déterministe des lignes JSONL (--seed, --reverse undo, --check CI, JSON) ✓ 2026-08-02
- [x] text-line-diff-count : compte lignes ajoutées/supprimées/communes entre deux fichiers (multiset, seuils CI, JSON) ✓ 2026-08-02
- [x] url-extract-fragment : extrait le fragment (#...) des URLs (stats, top-N, --check CI, JSON) ✓ 2026-08-02
- [x] file-permission-check : vérifie les permissions d'une arborescence (--expect/--max octal, --apply, CI, JSON) ✓ 2026-08-02

## Vague 325 — CSV/JSON/text/URL/file mix
- [x] csv-median-column : calcule la médiane d'une colonne CSV (-c, percentiles, --check CI, JSON) ✓ 2026-08-02 (déjà publié)
- [x] json-key-frequency : fréquence d'occurrence de chaque clé sur un JSONL (top-N, --check CI, JSON) ✓ 2026-08-02 (déjà publié Vague 306)
- [x] text-longest-line : affiche la plus longue ligne et son numéro (--top, --check CI, JSON) ✓ 2026-08-02
- [x] url-swap-scheme : permute http/https (ws/wss, ftp/sftp) des URLs (--to, --check CI, JSON) ✓ 2026-08-02
- [x] file-empty-dirs : liste les dossiers vides d'une arborescence (--prune, --check CI, JSON) ✓ 2026-08-02 (déjà publié)

## Vague 326 — CSV/JSON/text/URL/file mix
- [x] csv-row-swap : échange deux lignes de données d'un CSV par position (1-based, --check CI, JSON) ✓ 2026-08-02
- [x] json-extract-booleans : liste toutes les valeurs booléennes d'un JSON/JSONL avec leur chemin (--check CI, JSON) ✓ 2026-08-02
- [x] text-strip-trailing-punct : retire la ponctuation terminale de chaque ligne (. , ; : ! ?, --chars, --check CI, JSON) ✓ 2026-08-02
- [x] url-add-path-prefix : préfixe le path des URLs d'un segment (/api/v1, trailing slash gérée, --check CI, JSON) ✓ 2026-08-02
- [x] file-permission-setuid-detect : liste les fichiers setuid/setgid d'une arborescence (rapport, --check CI, JSON) ✓ 2026-08-02

## Vague 327 — CSV/JSON/text/URL/file mix
- [x] csv-column-last-non-empty : remplace chaque cellule vide par la dernière valeur non vide au-dessus dans la colonne (--cols, --check CI, JSON) ✓ 2026-08-02
- [x] json-count-array-items : compte total d'éléments de tous les tableaux d'un JSON/JSONL (par chemin, --check CI, JSON) ✓ 2026-08-02
- [x] text-duplicate-word-detect : détecte les mots répétés consécutivement dans chaque ligne ("the the", --ignore-case, --check CI, JSON) ✓ 2026-08-02
- [x] url-strip-path-suffix : retire un suffixe de path des URLs (/index.html, --suffix, --check CI, JSON) ✓ 2026-08-02
- [x] file-extension-lowercase-check : vérifie que toutes les extensions sont en minuscules (--apply rename, --check CI, JSON) ✓ 2026-08-02

## Vague 328 — CSV/JSON/text/URL/file mix
- [x] csv-collapse-blank-rows : retire les lignes vides (ou ne contenant que des virgules/espaces) d'un CSV (--check CI, JSON) ✓ 2026-08-02
- [x] json-extract-top-level-arrays : extrait les tableaux racine d'un JSON/JSONL en lignes JSONL séparées (--min-len, --check CI, JSON) ✓ 2026-08-02
- [x] text-swap-word-pairs : échange deux mots donnés dans chaque ligne (--a/--b, simultané --check CI, JSON) ✓ 2026-08-02
- [x] url-host-replace-mapping : remplace des hosts d'URLs via mapping fichier old=new (port préservé, --check CI, JSON) ✓ 2026-08-02
- [x] file-oldest-per-build-dir : trouve le fichier le plus ancien de chaque sous-dossier (bascule de build stale, --check CI, JSON) ✓ 2026-08-02

## Vague 329 — CSV/JSON/text/URL/file mix
- [x] csv-swap-columns : échange deux colonnes d'un CSV par nom ou index (--check CI, JSON) ✓ 2026-08-02
- [x] json-pick-random-key : sélectionne une clé aléatoire (seed) par objet JSONL (--check CI, JSON) ✓ 2026-08-02
- [x] text-wrap-prefix-indent : re-wrap les lignes longues en indentant la continuation (--width, --check CI, JSON) ✓ 2026-08-02
- [x] url-remove-default-port : retire le port par défaut du scheme (:80 http, :443 https) des URLs (--check CI, JSON) ✓ 2026-08-02
- [x] file-empty-dirs-list : liste les dossiers vides d'une arborescence (--delete, --check CI, JSON) ✓ 2026-08-02

## Vague 330 — CSV/JSON/text/URL/file mix
- [x] csv-column-quarter : extrait le quartile demandé (q1/q2/q3) d'une colonne numérique CSV (--check CI, JSON) ✓ 2026-08-02
- [x] json-set-difference : calcule les clés présentes dans un JSON mais pas dans l'autre (--symmetric, --check CI, JSON) ✓ 2026-08-02
- [x] text-strip-trailing-punct : retire la ponctuation terminale des lignes (configurable, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-auth-userinfo : extrait user/password éventuels des URLs et les rapporte séparément (--redact, --check CI, JSON) ✓ 2026-08-02
- [x] file-broken-pipe-links : détecte les symlinks circulaires ou trop profonds d'une arborescence (--max-depth, --check CI, JSON) ✓ 2026-08-02

## Vague 331 — CSV/JSON/text/URL/file mix
- [x] csv-add-row-hash : ajoute une colonne hash de chaque ligne CSV (md5/sha256, --check CI, JSON) ✓ 2026-08-02
- [x] json-extract-unique-paths : liste les chemins uniques d'un JSON/JSONL avec types (--check CI, JSON) ✓ 2026-08-02
- [x] text-strip-leading-punct : retire la ponctuation en début de ligne (configurable, --check CI, JSON) ✓ 2026-08-02
- [x] url-normalize-case-path : met les segments path d'URLs en kebab-case (--check CI, JSON) ✓ 2026-08-02
- [x] file-deep-empty-dirs : détecte les dossiers vides récursivement (bottom-up, --prune, --check CI, JSON) ✓ 2026-08-02

## Vague 332 — CSV/JSON/text/URL/file mix
- [x] csv-row-select-pattern : sélectionne les lignes CSV dont une colonne matche un pattern glob (--check CI, JSON) ✓ 2026-08-02
- [x] json-count-by-depth : compte les noeuds JSON par profondeur (histogramme, --check CI, JSON) ✓ 2026-08-02
- [x] text-extract-between-markers : extrait le texte entre deux marqueurs début/fin (multi-bloc, --check CI, JSON) ✓ 2026-08-02
- [x] url-replace-host-prefix : ajoute/remplace un préfixe de host dans les URLs (www., api., --check CI, JSON) ✓ 2026-08-02
- [x] file-duplicate-names-report : rapport des noms de fichiers en double avec chemins complets (--check CI, JSON) ✓ 2026-08-02

## Vague 333 — CSV/JSON/text/URL/file mix
- [x] csv-add-column-index : ajoute une colonne d'index 1-based à un CSV (--start, --check CI, JSON) ✓ 2026-08-02
- [x] json-array-length-check : vérifie que tous les tableaux d'un chemin JSON ont la même longueur (--length, --check CI, JSON) ✓ 2026-08-02
- [x] text-strip-inline-whitespace : réduit les whitespace internes multiples à un seul espace (--check CI, JSON) ✓ 2026-08-02
- [x] url-extract-port-list : extrait et liste les ports uniques des URLs (--check CI, JSON) ✓ 2026-08-02
- [x] file-large-files-report : rapport des fichiers dépassant une taille seuil (--min-size, --check CI, JSON) ✓ 2026-08-02

## Vague 334 — CSV/JSON/text/URL/file mix
- [x] csv-filter-empty-rows : retire les lignes entièrement vides d'un CSV (--check CI, JSON) ✓ 2026-08-02
- [x] json-key-case-report : rapport des styles de casse des clés JSON (camel/snake/kebab, --check CI, JSON) ✓ 2026-08-02
- [x] text-collapse-multiple-blank : réduit les lignes vides multiples à une seule (--max, --check CI, JSON) ✓ 2026-08-02
- [x] url-query-add-param : ajoute un paramètre query à des URLs (--check CI, JSON) ✓ 2026-08-02
- [x] file-permission-world-writable : liste les fichiers world-writable d'une arborescence (--check CI, JSON) ✓ 2026-08-02

## Vague 335 — CSV/JSON/text/URL/file mix
- [x] csv-strip-empty-columns : retire les colonnes entièrement vides d'un CSV (--check CI, JSON) ✓ 2026-08-02
- [x] json-path-tail : extrait le dernier segment des chemins d'un JSON (dot-path leaves, --check CI, JSON) ✓ 2026-08-02
- [x] text-first-n-lines-of-blocks : garde les N premières lignes de chaque bloc séparé par des lignes vides (--check CI, JSON) ✓ 2026-08-02
- [x] url-extract-host : extrait le hostname de chaque URL (dedup, stats, --check CI, JSON) ✓ 2026-08-02
- [x] file-executable-report : liste les fichiers exécutables d'une arborescence (shebang/ELF/bit x, --check CI, JSON) ✓ 2026-08-02

## Vague 336 — CSV/JSON/text/URL/file mix
- [x] csv-drop-trailing-commas : retire les virgules terminales en excès des lignes CSV (--check CI, JSON) ✓ 2026-08-02
- [x] json-type-coerce-strings : convertit les strings JSON "123"/"true"/"null" en types natifs (récursif, --check CI, JSON) ✓ 2026-08-02
- [x] text-reverse-words-line : inverse l'ordre des mots de chaque ligne (--min-len, --check CI, JSON) ✓ 2026-08-02
- [x] url-ensure-scheme : garantit un scheme par défaut (https) aux URLs sans scheme (--check CI, JSON) ✓ 2026-08-02
- [x] file-hardlink-count-report : rapport du nombre de liens durs par fichier (--min-links, --check CI, JSON) ✓ 2026-08-02

## Vague 341 — CSV/JSON/text/URL/file mix
- [x] csv-detect-header : détecte si la première ligne d'un CSV est un header réel (--check CI, JSON) ✓ 2026-08-02
- [x] json-count-numeric-paths : compte les chemins numériques dans un JSONL (--min-value CI, JSON) ✓ 2026-08-02
- [x] text-extract-ipv4 : extrait les adresses IPv4 d'un texte (validation stricte, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-scheme-report : rapport des schemes d'URLs (comptage, --require http/https CI, JSON) ✓ 2026-08-02
- [x] file-sticky-bit-detect : liste les dossiers avec sticky bit inhabituel (hors /tmp, --check CI, JSON) ✓ 2026-08-02

## Vague 342 — CSV/JSON/text/URL/file mix (à définir)
- [x] csv-max-column-length : rapport de la longueur max par colonne CSV (--threshold CI, JSON) ✓ 2026-08-02
- [x] json-detect-timestamp-formats : détecte les formats de dates/timestamps en valeurs JSON (iso/unix/ms, --check CI, JSON) ✓ 2026-08-02
- [x] text-split-sentences : découpe un texte en phrases (abbr-aware, --min-words, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-fragment-report : rapport des fragments #section d'URLs (dedup, counts, --check CI, JSON) ✓ 2026-08-02
- [x] file-mtime-future-detect : liste les fichiers dont le mtime est dans le futur (--tolerance CI, JSON) ✓ 2026-08-02

## Vague 343 — CSV/JSON/text/URL/file mix (à définir)
- [x] csv-quote-all-fields : re-quote toutes les cellules d'un CSV (guillemets uniformes, --minimal, --check CI, JSON) ✓ 2026-08-02
- [x] json-count-by-key-length : rapport des longueurs de clés JSON par profondeur (histogramme, --max CI, JSON) ✓ 2026-08-02
- [x] text-wrap-indent-preserve : wrap de texte qui conserve l'indentation initiale (--width, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-path-extension-report : rapport par extension du path des URLs (counts, mime guess, --check CI, JSON) ✓ 2026-08-02
- [x] file-same-content-links : détecte les fichiers identiques pouvant être remplacés par des hardlinks (--apply, --check CI, JSON) ✓ 2026-08-02

## Vague 350 — CSV/JSON/text/URL/file mix
- [x] csv-remove-empty-rows : retire les lignes CSV 100% vides/whitespace (--strip, --keep-header, --check CI, JSON) ✓ 2026-08-02
- [x] json-add-uuid-field : ajoute un UUID déterministe par doc basé sur SHA-256 canonique (key custom, --overwrite, --check CI, JSON) ✓ 2026-08-02
- [x] text-trailing-newline-report : détecte/normalise newline final (ok/no_newline/multi_newline/empty, --fix, --check CI, JSON) ✓ 2026-08-02
- [x] url-strip-index-file : retire les trailing index.html/default.aspx (--extra custom, --check CI, JSON) ✓ 2026-08-02
- [x] file-charset-detect-report : heuristique charset per fichier (utf-8/16/32/ascii/latin-1/binary, --require-charset CI, JSON) ✓ 2026-08-02

## Vague 349 — CSV/JSON/text/URL/file mix
- [x] csv-column-unique-values : distribution des valeurs uniques par colonne CSV (count, percent, --require-unique/--require-distinct, JSON) ✓ 2026-08-02
- [x] json-values-length-report : stats de longueur des valeurs string par dot-path (min/max/mean/total, --min-len/--max-len/--require, JSON) ✓ 2026-08-02
- [x] text-repeat-word-count : détecte les mots consécutifs répétés par ligne (--ignore-case/--min-length/whitelist, --check CI, JSON) ✓ 2026-08-02
- [x] url-add-query-prefix : préfixe toutes les clés query d'URLs (--prefix/--separator/--exclude, --check CI, JSON) ✓ 2026-08-02
- [x] file-empty-lines-report : compte les lignes vides/whitespace d'une arborescence (runs, --max-empty-per-file/--max-files, JSON) ✓ 2026-08-02

## Vague 348 — CSV/JSON/text/URL/file mix
- [x] csv-value-replace-map : remplacement des valeurs CSV via mapping old=new (colonne cible, --strict/--check CI, JSON) ✓ 2026-08-02
- [x] json-compact-spaces : normalise l'espacement JSON (compact/indent/séparateurs, --strip, --check CI, JSON) ✓ 2026-08-02
- [x] text-delete-blank-prefix : retire/compte/enforce l'absence de lignes vides en tête (--max-keep, --in-place, --check CI, JSON) ✓ 2026-08-02
- [x] url-strip-www-host : normalise les hosts URL en retirant/ajoutant www. (port/creds/IPv6 préservés, --check CI, JSON) ✓ 2026-08-02
- [x] text-sort-by-number : tri numérique de lignes via préfixe/field/regex (desc, unique, --check CI, JSON) ✓ 2026-08-02
- [x] file-owner-group-report : histogramme owner:group d'une arborescence (--only/--forbid/--require, --check CI, JSON) ✓ 2026-08-02

## Vague 347 — CSV/JSON/text/URL/file mix
- [x] csv-row-window-shift : décale une colonne CSV de N lignes lag/lead (--fill, --require CI, JSON) ✓ 2026-08-02
- [x] json-count-null-paths : compte les null par dot-path dans JSONL (taux global, --max-null-rate CI, JSON) ✓ 2026-08-02
- [x] text-sentence-length-report : mots par phrase + distribution (abbr-aware, --max-words lint CI, JSON) ✓ 2026-08-02
- [x] url-query-value-length : stats de longueur des valeurs query par paramètre (--max-length CI, JSON) ✓ 2026-08-02
- [x] file-atime-older-than-mtime : fichiers écrits jamais relus (atime < mtime, --tolerance, --check CI, JSON) ✓ 2026-08-02

## Vague 346 — CSV/JSON/text/URL/file mix
- [x] csv-row-ordinal-rank : colonne de rang ordinal dense/compétition (--desc, --dense, --require CI, JSON) ✓ 2026-08-02
- [x] json-merge-prefer-newer : fusionne deux JSONL en gardant le record au timestamp le plus récent (-k, -t, --require-match CI, JSON) ✓ 2026-08-02
- [x] text-word-length-histogram : histogramme longueurs de mots + stats + barres ASCII (--max-length lint CI, JSON) ✓ 2026-08-02
- [x] url-domains-common-prefix : préfixe commun / suffixe DNS commun des hostnames (--suffix, --require CI, JSON) ✓ 2026-08-02
- [x] file-atime-stale-report : fichiers dont l'atime dépasse un seuil de jours (--days, --check CI, JSON) ✓ 2026-08-02

## Vague 345 — CSV/JSON/text/URL/file mix
- [x] csv-keep-unique-rows : garde uniquement les lignes dont la clé est unique (--require CI, JSON) ✓ 2026-08-02
- [x] json-insert-index-field : insère un champ index monotone dans JSONL/tableau (--check, --overwrite CI, JSON) ✓ 2026-08-02
- [x] text-syllable-estimate : estime les syllabes par mot/ligne + Flesch (en/fr, --max lint CI, JSON) ✓ 2026-08-02
- [x] url-fragment-append : ajoute/remplace le fragment #section des URLs (--keep-existing, --check CI, JSON) ✓ 2026-08-02
- [x] file-disk-usage-report : rapport blocs alloués vs taille logique, détection sparse (--check CI, JSON) ✓ 2026-08-02

## Vague 344 — CSV/JSON/text/URL/file mix (à définir)
- [x] csv-trim-cells : retire les espaces débuts/fins des cellules CSV (déjà publié) ✓ 2026-08-02
- [x] csv-non-numeric-report : rapport des valeurs non numériques par colonne CSV (--threshold CI, JSON) ✓ 2026-08-02
- [x] json-array-unique-report : rapport des tableaux JSONL avec doublons d'éléments (--count CI, JSON) ✓ 2026-08-02
- [x] text-hyphenate-count : rapport des mots avec tirets vs sans par ligne (histogramme, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-default-port-clean : retire le port par défaut des URLs (80/443 normalisés, --check CI, JSON) ✓ 2026-08-02
- [x] file-shebang-detect : liste les scripts exécutables avec shebang + type (python/bash/node, --check CI, JSON) ✓ 2026-08-02

## Vague 340 — CSV/JSON/text/URL/file mix
- [x] csv-add-hash-column : ajoute une colonne de hash (md5/sha256) calculée sur les colonnes choisies de chaque ligne (--check CI, JSON) ✓ 2026-08-02
- [x] json-extract-max-value : trouve la valeur maximale à un dot-path numérique à travers JSONL (--check CI, JSON) ✓ 2026-08-02
- [x] text-blank-line-guard : vérifie qu'un fichier ne contient aucune ligne vide (--allow-trailing, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-query-count : rapport du nombre de paramètres query par URL (stats, --threshold CI, JSON) ✓ 2026-08-02
- [x] file-world-readable-report : rapport des fichiers lisibles par tout le monde (o+r) d'une arborescence (--check CI, JSON) ✓ 2026-08-02

## Vague 339 — CSV/JSON/text/URL/file mix
- [x] csv-fill-random : remplit les cellules vides d'une colonne CSV par des valeurs aléatoires de la colonne (seed reproductible, --check CI, JSON) ✓ 2026-08-02
- [x] json-merge-prefer-longer : fusionne deux JSONL en préférant la valeur la plus longue par clé (strings/arrays, --check CI, JSON) ✓ 2026-08-02
- [x] text-strip-matching-prefix-lines : retire les lignes commençant par un préfixe donné (--invert, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-ipv6 : repère les URLs utilisant une adresse IPv6 littérale ([::1], --check CI, JSON) ✓ 2026-08-02
- [x] file-group-by-uid-gid : regroupe les fichiers d'une arborescence par couple uid:gid (counts, tailles, --check CI, JSON) ✓ 2026-08-02

## Vague 338 — CSV/JSON/text/URL/file mix
- [x] csv-escape-formula-prefix : préfixe les cellules CSV commençant par =, +, -, @ d'une apostrophe (anti formula-injection, --check CI, JSON) ✓ 2026-08-02
- [x] json-unwrap-single-element : déplie les objets JSON à une seule clé en leur valeur directe (récursif, --check CI, JSON) ✓ 2026-08-02
- [x] text-collapse-blank-runs : écrase les runs de lignes vides consécutives en une seule (--max N, --check CI, JSON) ✓ 2026-08-02
- [x] url-extract-tld-domain : extrait le registered domain (tld+1) d'URLs (public suffix embarqué, --check CI, JSON) ✓ 2026-08-02
- [x] file-utf8-validity-check : vérifie que chaque fichier texte d'une arborescence est UTF-8 valide (rapport, --check CI, JSON) ✓ 2026-08-02

## Vague 337 — CSV/JSON/text/URL/file mix
- [x] csv-quote-all-cells : force le quoting de toutes les cellules CSV (QUOTE_ALL, --minimal inverse, --check CI, JSON) ✓ 2026-08-02
- [x] json-boolean-report : rapport des valeurs booléennes par dot-path dans un JSONL (ratio true/false, --check CI, JSON) ✓ 2026-08-02
- [x] text-squeeze-spaces : écrase les espaces multiples en un seul (hors indentation, --keep-indent, --check CI, JSON) ✓ 2026-08-02
- [x] url-lowercase-host-path-split : passe host en minuscules mais préserve la casse du path (--check CI, JSON) ✓ 2026-08-02
- [x] file-size-dupes-same-name : détecte les fichiers de même nom et même taille dans des dossiers différents (--check CI, JSON) ✓ 2026-08-02

## Vague 287 — CLI Tools (rename CSV, flatten JSON, tri par longueur, query k=v, plus anciens fichiers)
- [x] csv-column-rename : renomme des colonnes CSV via mapping nom=nouveau (--require CI, JSON) ✓ 2026-08-01
- [x] json-flatten-keys : aplatit un JSON en paires chemin=valeur (séparateur dot, arrays indexés, --check CI, JSON) ✓ 2026-08-01
- [x] text-sort-by-length : trie les lignes par longueur (asc/desc, tie-break alpha, --check CI, JSON) ✓ 2026-08-01
- [x] url-extract-query-keys-values : éclate chaque URL en lignes key=value pour la query (dedup, --check CI, JSON) ✓ 2026-08-01
- [x] file-oldest-first-list : liste les fichiers d'une arborescence triés par mtime croissant (--top, --check CI, JSON) ✓ 2026-08-01

## Vague 286 — CLI Tools (whitelist CSV, unescape JSON, table→kv, strip www, taille par owner)
- [x] csv-column-keep : garde seulement certaines colonnes d'un CSV (whitelist par nom ou index, --require CI, JSON) ✓ 2026-08-01
- [x] json-unescape-strings : décode les séquences \n \t \uXXXX dans les valeurs string d'un JSONL (--check CI, JSON) ✓ 2026-08-01
- [x] text-columns-to-lines : convertit un tableau whitespace en lignes clé=valeur (--kv, --check CI, JSON) ✓ 2026-08-01
- [x] url-strip-www : retire le préfixe www. des hosts d'URLs (--check CI, JSON) ✓ 2026-08-01
- [x] file-group-by-owner-size : regroupe les fichiers d'une arborescence par owner avec total de taille (--threshold CI, JSON) ✓ 2026-08-01


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
