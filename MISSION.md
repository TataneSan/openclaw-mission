# OpenClaw — File d'attente des outils

## RÈGLE
Chaque outil = son propre repo Git sur github.com/TataneSan.
JAMAIS mentionner IA/agent dans le code ou les commits.
Push automatique après chaque outil.

## Terminés
(mettre à jour ici chaque outil fini)

## En cours
- env-export (Vague 9)

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
- [ ] vm-stats : collecte et affiche les stats système (CPU, RAM, disk, net)
- [ ] docker-logs-cli : agrège et filtre les logs de containers Docker
- [ ] release-notes : génère des release notes depuis git tags

## Vague 9 — Productivité
- [ ] journal-cli : journal quotidien en CLI avec tags et recherche
- [ ] goal-tracker : suivi d'objectifs avec sous-tâches et progression
- [ ] flashcards-cli : système de répétition espacée en terminal
- [ ] budget-cli : gestion budgétaire mensuelle en CLI
- [ ] meeting-notes : template et gestion de notes de réunion

## Notes
- Dernière mise à jour: 2026-05-28
- Priorité: CLI tools d'abord (rapides à coder)
- Chaque outil: max 1-2h de dev
