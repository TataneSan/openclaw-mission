# Independent tools

Chaque sous-répertoire de `tools/` est un dépôt Git autonome publié séparément sous `github.com/TataneSan/`.

Le dépôt parent ne versionne pas les worktrees enfants. Pour travailler sur un outil :

```bash
cd /root/openclaw/tools/<nom-outil>
git status
git log --oneline -5
```

Pour publier un nouvel outil, suivre la structure et les contrôles documentés dans `/root/openclaw/MISSION.md`.
