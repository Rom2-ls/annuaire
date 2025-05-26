const express = require('express');
const chokidar = require('chokidar');
const WebSocket = require('ws');
const fs = require('fs');

const app = express();
const PORT = process.env.PORT || 3000;

// Configuration du serveur WebSocket
const server = require('http').createServer(app);
const wss = new WebSocket.Server({ server });

// Middleware pour parser JSON
app.use(express.json());

// Route pour obtenir les contacts depuis le fichier JSON Go
app.get('/api/contacts', (req, res) => {
  try {
    if (fs.existsSync('annuaire.json')) {
      const data = fs.readFileSync('annuaire.json', 'utf8');
      const contacts = JSON.parse(data);
      const contactArray = Object.values(contacts);
      res.json(contactArray);
    } else {
      res.json([]);
    }
  } catch (error) {
    console.error('Erreur lors de la lecture du fichier:', error);
    res.status(500).json({ error: 'Erreur serveur' });
  }
});

// Route principale avec interface HTML
app.get('/', (req, res) => {
  res.send(`
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Annuaire - Auto-Refresh</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 1000px; margin: 0 auto; padding: 20px; background: #f0f2f5; }
        .container { background: white; padding: 30px; border-radius: 10px; box-shadow: 0 4px 20px rgba(0,0,0,0.1); }
        h1 { text-align: center; color: #333; margin-bottom: 30px; }
        .status { padding: 15px; border-radius: 8px; margin: 20px 0; text-align: center; font-weight: bold; }
        .status.connected { background: #d4edda; color: #155724; border: 1px solid #c3e6cb; }
        .status.disconnected { background: #f8d7da; color: #721c24; border: 1px solid #f5c6cb; }
        .stats { background: #e3f2fd; padding: 15px; border-radius: 8px; margin: 20px 0; text-align: center; color: #1976d2; }
        .contact-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(250px, 1fr)); gap: 15px; margin: 20px 0; }
        .contact-card { background: #f8f9fa; border: 1px solid #dee2e6; padding: 15px; border-radius: 8px; }
        .contact-name { font-weight: bold; font-size: 1.1em; margin-bottom: 5px; }
        .contact-phone { color: #666; }
        .commands { background: #f8f9fa; padding: 20px; border-radius: 8px; margin: 20px 0; }
        .command { background: #e9ecef; padding: 8px; border-radius: 4px; font-family: monospace; margin: 5px 0; display: block; }
    </style>
</head>
<body>
    <div class="container">
        <h1>📞 Annuaire - Interface Auto-Refresh</h1>
        
        <div id="status" class="status disconnected">🔌 Connexion en cours...</div>
        
        <div class="stats" id="stats">📊 Chargement...</div>

        <div id="contacts">⏳ Chargement des contacts...</div>

        <div class="commands">
            <h3>💡 Commandes CLI</h3>
            <p>Testez l'auto-refresh avec ces commandes :</p>
            <code class="command">go run cmd/annuaire/main.go --action ajouter --nom "Test" --prenom "Auto" --tel "0123456789"</code>
            <code class="command">go run cmd/annuaire/main.go --action lister</code>
        </div>
    </div>

    <script>
        let ws = null;
        let isConnected = false;

        function connectWebSocket() {
            try {
                ws = new WebSocket('ws://localhost:' + window.location.port || '3000');
                
                ws.onopen = function() {
                    isConnected = true;
                    updateStatus('connected', '✅ Auto-refresh activé');
                    loadContacts();
                };
                
                ws.onmessage = function(event) {
                    const data = JSON.parse(event.data);
                    if (data.type === 'refresh') {
                        console.log('Changement détecté, rechargement...');
                        loadContacts();
                    }
                };
                
                ws.onclose = function() {
                    isConnected = false;
                    updateStatus('disconnected', '❌ Connexion fermée - Reconnexion...');
                    setTimeout(connectWebSocket, 3000);
                };
                
                ws.onerror = function() {
                    isConnected = false;
                    updateStatus('disconnected', '⚠️ Erreur de connexion');
                };
            } catch (error) {
                updateStatus('disconnected', '❌ WebSocket indisponible');
            }
        }

        function updateStatus(type, message) {
            const statusEl = document.getElementById('status');
            statusEl.className = 'status ' + type;
            statusEl.textContent = message;
        }

        async function loadContacts() {
            try {
                const response = await fetch('/api/contacts');
                const contacts = await response.json();
                displayContacts(contacts);
                updateStats(contacts.length);
            } catch (error) {
                document.getElementById('contacts').innerHTML = '<div style="text-align:center;color:#999;">❌ Erreur de chargement</div>';
            }
        }

        function displayContacts(contacts) {
            const contactsEl = document.getElementById('contacts');
            
            if (contacts.length === 0) {
                contactsEl.innerHTML = '<div style="text-align:center;color:#999;padding:40px;">📭 Aucun contact</div>';
                return;
            }

            const contactsHtml = contacts.map(contact => 
                '<div class="contact-card">' +
                '<div class="contact-name">' + contact.prenom + ' ' + contact.nom + '</div>' +
                '<div class="contact-phone">📞 ' + contact.telephone + '</div>' +
                '</div>'
            ).join('');

            contactsEl.innerHTML = '<div class="contact-grid">' + contactsHtml + '</div>';
        }

        function updateStats(count) {
            const lastUpdate = new Date().toLocaleTimeString('fr-FR');
            document.getElementById('stats').innerHTML = 
                '📊 Total: <strong>' + count + '</strong> | 🕒 MAJ: <strong>' + lastUpdate + '</strong>';
        }

        // Initialisation
        connectWebSocket();
        setInterval(() => { if (!isConnected) loadContacts(); }, 10000);
    </script>
</body>
</html>`);
});

// Surveillance du fichier annuaire.json
const watcher = chokidar.watch('annuaire.json', {
  persistent: true,
  ignoreInitial: true,
});

watcher.on('change', () => {
  console.log('📝 Changement détecté dans annuaire.json');
  wss.clients.forEach((client) => {
    if (client.readyState === WebSocket.OPEN) {
      client.send(JSON.stringify({ type: 'refresh' }));
    }
  });
});

// Gestion des connexions WebSocket
wss.on('connection', (ws) => {
  console.log('🔌 Nouvelle connexion WebSocket');
  ws.on('close', () => console.log('🔌 Connexion fermée'));
});

// Démarrage du serveur
server.listen(PORT, () => {
  console.log(`🚀 Serveur Node.js sur http://localhost:${PORT}`);
  console.log(`📁 Surveillance: annuaire.json`);
  console.log(`🔄 Auto-refresh: WebSocket activé`);
});

process.on('SIGINT', () => {
  console.log('\n🛑 Arrêt du serveur...');
  watcher.close();
  server.close(() => process.exit(0));
});
