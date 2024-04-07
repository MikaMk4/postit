# Anleitung zur lokalen Ausführung:

(Docker Engine (Linux-basiert) mit Compose Plugin wird benötigt)  

Im Terminal im root-Ordner des Projektes folgenden Befehl ausführen:
```bash
docker compose up --build
```

Das wird ein wenig dauern. Sobald 'App running at:' zu sehen ist, kann über einen beliebigen Browser http://localhost:8080 aufgerufen werden.

Allerdings müssen noch Daten in die Datenbank eingespeist werden. In einem weiteren Terminal folgenden Befehl ausführen:
```bash
docker exec -it postit-db sh -c "mysql -proot postit < /dump/data.sql"
```
Danach sollte alles klappen.