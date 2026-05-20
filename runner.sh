#!/bin/bash
# OpenClaw Autonomous Tool Builder
# Creates and publishes tools to GitHub continuously

export PATH="/root/.local/bin:$PATH"
export HOME="/root"

HERMES="/root/.local/bin/hermes"
LOG="/root/openclaw/runner.log"
MISSION="/root/openclaw/MISSION.md"

log() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $*" >> "$LOG"
}

BATCH=1
while true; do
    log "=== Batch $BATCH starting ==="

    PROMPT="Tu es un développeur open source prolifique. Tu codes des outils utiles et tu les publies sur GitHub.

RÈGLES STRICTES:
1. UN OUTIL = UN REPO GIT. Toujours.
2. JAMAIS mentionner IA, AI, LLM, ChatGPT, OpenClaw, agent dans les commits, README, ou le code.
3. Commits signés normalement, style humain.
4. Code propre, fonctionnel, bien documenté.
5. Push sur github.com/TataneSan après chaque outil.

WORKFLOW:
1. Lis /root/openclaw/MISSION.md
2. Prends le premier outil non-fait dans la file d'attente
3. Crée /root/openclaw/tools/NOM-OUTIL
4. Code l'outil complet (CLI, API, bot, ce que c'est)
5. README.md pro avec description, install, usage, exemples
6. git init, git add, git commit -m 'feat: initial commit'
7. git remote add origin git@github.com:TataneSan/NOM-OUTIL.git
8. git push -u origin main
9. Coche l'outil dans MISSION.md comme terminé
10. Passe au suivant IMMÉDIATEMENT

Tu as: Node.js, Python, Go, Docker, git, gh CLI.
Ne demande JAMAIS de permission. Code et push.
Commence MAINTENANT."

    timeout 600 $HERMES chat -q "$PROMPT" >> "$LOG" 2>&1

    EXIT_CODE=$?
    log "Batch $BATCH finished (exit: $EXIT_CODE)"

    BATCH=$((BATCH + 1))
    log "Pausing 30s before next batch..."
    sleep 30
done
