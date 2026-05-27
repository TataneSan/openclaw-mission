# OpenClaw — File d'attente des outils

## RÈGLE
Chaque outil = son propre repo Git sur github.com/TataneSan.
JAMAIS mentionner IA/agent dans le code ou les commits.
Push automatique après chaque outil.

## Terminés
(mettre à jour ici chaque outil fini)

## En cours
(aucun)

## File d'attente — CLI Tools
- [x] json-dive : CLI pour explorer/query du JSON avec des chemins type jq mais plus intuitif ✓ 2026-05-20
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
- [ ] hex-convert : convertisseur entre hex, décimal, binaire, octal
- [ ] markdown-preview : prévisualise du markdown en terminal (coloré)
- [ ] port-check : vérifie si un port est ouvert sur un hôte distant
- [ ] regex-extract : extrait des données avec regex depuis des fichiers
- [ ] string-manip : utilitaires de manipulation de chaînes (upper, lower, slug, etc.)

## Notes
- Dernière mise à jour: 2026-05-27
- Priorité: CLI tools d'abord (rapides à coder)
- Chaque outil: max 1-2h de dev
