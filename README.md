# TataneSan Service Workspace

Workspace de déploiement et de maintenance des services hébergés sur le serveur.

## Services

| Service | Port | Unité systemd | État attendu |
|---|---:|---|---|
| Proxy HTTP | 8088 | `openclaw-proxy.service` | selon configuration fournisseur |
| Scraper web | 8089 | `openclaw-scraper.service` | selon besoin |
| Convertisseur PDF | 8090 | `openclaw-converter.service` | actif |
| API image | 8091 | `openclaw-image.service` | actif |
| Suivi SEO | — | `openclaw-seo.service` | actif |
| Suivi crypto | — | `openclaw-crypto.service` | selon configuration |

Vérifier les services :

```bash
cd /root/openclaw
./manage.sh status
systemctl --type=service --all | grep openclaw
```

## Organisation

```text
/root/openclaw/
├── MISSION.md          # source de vérité de la file des outils
├── tools/              # dépôts Git indépendants, un outil par dépôt
├── crypto-monitor/     # service de suivi crypto
├── image-api/          # API de traitement d'images
├── llm-proxy/          # proxy HTTP et sa configuration locale
├── pdf-converter/      # service de conversion PDF
├── seo-toolkit/        # outils SEO et API
├── telegram-bot/       # bot optionnel
├── web-scraper/        # scraper et API web
├── runtime/            # logs d'exécution, ignorés par Git
├── archive/            # anciennes sorties et tracker historique
└── wallets/            # données locales sensibles
```

Les répertoires `node_modules/`, `venv/`, les bases SQLite, les logs, les PID et les fichiers `.env` sont locaux et ne sont pas versionnés. Les dépôts situés sous `tools/` possèdent leur propre historique Git et leur propre remote.

## Runner de mission

Le runner est géré par systemd :

```bash
systemctl status openclaw.service
systemctl restart openclaw.service
```

Son journal local est écrit dans :

```text
/root/openclaw/runtime/logs/runner.log
```

## Sécurité

- Ne jamais commiter un fichier `.env`, une clé d'administration ou une base de production.
- Charger les secrets depuis l'environnement ou un gestionnaire de secrets.
- Les valeurs d'exemple dans la documentation sont des placeholders.
- Les données du dossier `wallets/` doivent rester accessibles uniquement à l'utilisateur de service.

## Dépôts

Le dépôt de suivi est :

```text
git@github.com:TataneSan/openclaw-mission.git
```

Chaque outil publié est un dépôt séparé sous `github.com/TataneSan/`.
