# OpenClaw — File d'attente des outils

## RÈGLE
Chaque outil = son propre repo Git sur github.com/TataneSan.
JAMAIS mentionner IA/agent dans le code ou les commits.
Push automatique après chaque outil.

## Terminés
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

## Vague 43 — CLI Tools
- [x] go-struct-gen : génère des structs Go depuis du JSON ✓ 2026-05-28
- [x] json-to-go : génère des structs Go depuis du JSON (avec type hints) ✓ 2026-05-28
- [x] yaml-to-go : génère des structs Go depuis du YAML ✓ 2026-05-28
- [x] sql-migrate : système de migrations SQL simple (up/down) ✓ 2026-05-28
- [x] go-mod-tidy-check : vérifie si go.mod est à jour ✓ 2026-05-28

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
