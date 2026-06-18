# ☀️ Soulèu — Forum communautaire provençale sous le soleil

**Soulèu** (*Soleil* en provençal) est une plateforme de discussion et d'entraide communautaire dédiée aux personnes qui veulent soutenir la culture et les loisirs du Sud : la pétanque, la cuisine provençale et la nature. Inspiré de Reddit et Discord, ce projet inclu de nombreuses fonctionnalités.

Ce projet s'inscrit dans le cadre de la validation de la première année de Bachelor Informatique (2026).

---

## Membres de l'équipe
* **FUENTES Quentin**
* **PAYAN Lisa**

---

## Technologies utilisées

* **Backend, Frontend & Logique :** Golang
* **Base de données :** SQL
* **Templates :** HTML5
* **Esthétique :** CSS et JavaScript
* **Sécurité :** Tokens JWT (JSON Web Tokens) pour les sessions et Hashage SHA-512 pour les mots de passe

---

## Spécifications Fonctionnelles et Règles de Gestion

L'ensemble des fonctionnalités obligatoires décrites ci-dessous a été intégralement implémenté dans l'application :

### Gestion des Comptes & Sécurité

#### FT-1 – Système d’inscription
* **Objectif :** Permettre à un nouvel utilisateur de créer un compte afin de rejoindre la plateforme communautaire. Constitue le point d’entrée principal pour accéder aux fonctionnalités réservées aux membres authentifiés.
* **Règles de gestion :**
  * Un utilisateur non authentifié doit pouvoir créer un compte.
  * Le nom d’utilisateur et le mail doivent être uniques sur la plateforme.
  * Le mot de passe doit contenir au minimum **12 caractères**, dont au moins **une majuscule** et **un caractère spécial**.
  * Le mot de passe ne doit jamais être stocké en clair dans la base de données.
  * Le mot de passe doit être hashé avec l’algorithme **SHA-512** au minimum.

#### FT-2 – Système de connexion
* **Objectif :** Permettre à un utilisateur inscrit d’accéder à son compte et aux fonctionnalités nécessitant une authentification en vérifiant son identité à partir des informations saisies.
* **Règles de gestion :**
  * Un utilisateur non authentifié doit pouvoir se connecter.
  * L’utilisateur doit pouvoir se connecter avec son **nom d’utilisateur** ou son **adresse email**.
  * En cas d’informations incorrectes, la connexion doit être refusée.
  * Une fois l’authentification validée, un token **JWT** doit être généré.
  * Le token JWT doit permettre d’identifier l’utilisateur lors des requêtes nécessitant une authentification.

---

### Fils de Discussion & Messages

#### FT-3 – Création d’un fil de discussion
* **Objectif :** Permettre à un utilisateur authentifié de lancer un nouveau sujet d’échange sur la plateforme afin de structurer les conversations autour de thèmes précis.
* **Règles de gestion :**
  * Seul un utilisateur authentifié peut créer un fil de discussion.
  * Un fil de discussion doit être associé à son créateur.
  * Un fil de discussion doit pouvoir être associé à un ou plusieurs tags ou catégories.
  * Un fil de discussion doit posséder un état parmi les valeurs suivantes : **ouvert**, **fermé** ou **archivé**.
  * Un fil **ouvert** permet aux utilisateurs authentifiés de publier des messages.
  * Un fil **fermé** reste visible et consultable, mais ne permet plus de publier de nouveaux messages.
  * Un fil **archivé** ne doit plus être affiché ni accessible aux utilisateurs.

#### FT-4 – Consulter un fil de discussion
* **Objectif :** Permettre aux utilisateurs (visiteurs anonymes et membres authentifiés) d'accéder aux échanges déjà publiés, d'en parcourir le contenu et de lire les messages associés.
* **Règles de gestion :**
  * Un utilisateur authentifié ou non doit pouvoir consulter un fil de discussion.
  * Les fils **ouverts** peuvent être consultés.
  * Les fils **fermés** peuvent être consultés.
  * Les fils **archivés** ne doivent pas être visibles.
  * Les messages associés au fil doivent être affichés selon les règles de tri définies.

#### FT-5 – Publication d’un message dans un fil de discussion
* **Objectif :** Permettre à un utilisateur authentifié de participer activement à une discussion et de transformer le fil en véritable espace d’échange.
* **Règles de gestion :**
  * Seul un utilisateur authentifié peut publier un message.
  * Un message doit être publié dans un fil de discussion existant.
  * Un message ne peut être publié que dans un fil **ouvert**.
  * Un fil **fermé** ne doit pas accepter de nouveaux messages.
  * Un fil **archivé** ne doit pas accepter de nouveaux messages.
  * Chaque message doit être associé à son auteur.
  * Chaque message doit être associé au fil de discussion concerné.
  * La date d’envoi du message doit pouvoir être connue.

---

### Interactions, Tris & Pagination

#### FT-6 – Réagir à un message d’un fil de discussion
* **Objectif :** Mettre en place un système de like/dislike permettant aux utilisateurs authentifiés de réagir afin d'introduire une interaction sociale et de faire ressortir les messages pertinents.
* **Règles de gestion :**
  * Seul un utilisateur authentifié peut réagir à un message.
  * Un utilisateur ne peut pas liker et disliker le même message en même temps.
  * Un utilisateur ne peut avoir qu’**une seule réaction active** sur un même message.
  * Un **like** ajoute `+1` au score de popularité du message.
  * Un **dislike** retire `-1` au score de popularité du message.
  * Le score de popularité doit pouvoir être utilisé pour trier les messages.

#### FT-8 – Tri des messages d’un fil de discussion
* **Objectif :** Améliorer la lisibilité d’un fil de discussion en proposant plusieurs modes d’affichage selon le besoin de l'utilisateur.
* **Règles de gestion :**
  * Les messages doivent être affichés par défaut **du plus récent au plus ancien**.
  * L’utilisateur doit pouvoir afficher les messages par **ordre chronologique**.
  * L’utilisateur doit pouvoir afficher les messages selon leur **score de popularité**.
  * Le score de popularité est calculé à partir des likes et des dislikes (`+1` par like, `-1` par dislike).
  * Le tri par popularité doit tenir compte du score calculé pour chaque message.

#### FT-9 – Pagination des fils de discussion et des messages
* **Objectif :** Limiter le nombre d’éléments affichés à l’écran afin d’améliorer la lisibilité et les performances de la plateforme.
* **Règles de gestion :**
  * Les listes de fils de discussion doivent pouvoir être paginées.
  * Les listes de messages doivent pouvoir être paginées.
  * L’utilisateur doit pouvoir choisir un affichage par lots de **10**, **20** ou **30** éléments.
  * L’utilisateur doit pouvoir choisir d’**afficher la totalité** des éléments.
  * Par défaut, les éléments doivent être affichés par lots de **10**.

---

### Recherche & Navigation

#### FT-10 – Afficher les fils de discussion par tag/catégorie
* **Objectif :** Permettre aux utilisateurs de retrouver plus facilement les discussions liées à un thème précis afin de rendre la navigation plus intuitive.
* **Règles de gestion :**
  * Les fils de discussion doivent pouvoir être associés à un ou plusieurs tags ou catégories.
  * Un utilisateur doit pouvoir afficher les fils liés à un tag ou catégorie précis.
  * Seuls les fils visibles doivent apparaître dans les résultats.
  * Les fils archivés ne doivent pas apparaître dans les résultats.
  * Le filtrage doit conserver un affichage cohérent avec les règles de pagination.

#### FT-11 – Rechercher un fil de discussion
* **Objectif :** Permettre à un utilisateur authentifié de retrouver rapidement un fil de discussion sans parcourir toute la plateforme.
* **Règles de gestion :**
  * Seul un utilisateur authentifié peut effectuer une recherche.
  * La recherche doit pouvoir porter sur le **titre** d’un fil de discussion.
  * La recherche doit pouvoir porter sur un **tag/catégorie**.
  * Le système doit déterminer automatiquement si la recherche correspond à un titre, un tag ou une catégorie.
  * Les fils archivés ne doivent pas apparaître dans les résultats.
  * Les résultats doivent respecter les règles de visibilité des fils de discussion.

---

### Droits, Modération & Administration

#### FT-7 – Gestion des messages et des fils de discussion
* **Objectif :** Permettre aux utilisateurs de contrôler les contenus dont ils sont propriétaires et aux administrateurs de modérer la plateforme.
* **Règles de gestion :**
  * Un utilisateur authentifié peut modifier et supprimer un fil de discussion dont il est propriétaire.
  * Un utilisateur authentifié peut modifier et supprimer un message dont il est l’auteur.
  * Un administrateur peut supprimer n’importe quel fil de discussion ou message sans avoir besoin d'en être le propriétaire.
  * La suppression d’un fil de discussion doit entraîner la suppression des données qui lui sont liées.
  * La suppression d’un message doit entraîner la suppression des données qui lui sont liées, notamment les réactions associées.

#### FT-12 – Gestion et administration de la plateforme
* **Objectif :** Donner accès à un tableau de bord regroupant les actions importantes de gestion, assurant une séparation claire entre les droits d’un utilisateur classique et d'un administrateur.
* **Règles de gestion :**
  * Seuls les utilisateurs possédant le rôle **administrateur** peuvent accéder au tableau de bord d’administration.
  * Un administrateur peut modifier l’état d’un fil de discussion.
  * Un administrateur peut supprimer un fil de discussion ou un message.
  * Un administrateur peut bannir un compte utilisateur.
  * Un administrateur peut agir sur un fil ou un message même s’il n’en est pas le propriétaire/auteur.
  * Un utilisateur banni ne doit plus pouvoir accéder aux fonctionnalités nécessitant une authentification.

  ---

## Architecture du Projet

Le projet est structuré en deux grands sous-systèmes indépendants (Client et Serveur), chacun disposant de son propre cycle de vie et d'une architecture modulaire stricte :

```text
PROJET_FORUM/
│
├── client/                  # PARTIE CLIENT (Vues, Styles & Logique Frontend)
│   ├── api/                 # Appels et intégration des endpoints distants
│   ├── app/                 # Logique et scripts de l'application cliente
│   ├── assets/              # Fichiers statiques (css, img, js)
│   ├── config/              # Configuration spécifique au client
│   ├── controllers/         # Gestionnaires de rendu et flux côté client
│   ├── dto/                 # Objets de transfert de données d'affichage
│   ├── middleware/          # Intercepteurs (ex: gestion locale des sessions)
│   ├── routers/             # Définition des routes de l'interface graphique
│   ├── services/            # Traitements spécifiques à l'expérience utilisateur
│   ├── templates/           # Vues HTML / Templates Go distribués au navigateur
│   ├── .env                 # Variables d'environnement client
│   ├── main.go              # Point d'entrée du serveur web frontal (Client)
│   └── go.mod               # Dépendances du module client
│
└── server/                  # PARTIE SERVEUR (API REST, Métier & Données)
    ├── auth/                # Logique d'authentification et génération JWT
    ├── config/              # Connexion base de données et clés secrètes
    ├── controllers/         # Contrôleurs HTTP de l'API (FT-1 à FT-12)
    ├── dto/                 # Data Transfer Objects pour valider les requêtes de l'API
    ├── helper/              # Fonctions utilitaires partagées
    ├── middleware/          # Sécurité, validation JWT, gestion CORS et logs
    ├── migration/           # Scripts d'initialisation et d'évolution de la base SQL
    ├── models/              # Définition des entités de la base de données (User, Fil, Message)
    ├── repositories/        # Couche d'accès aux données (Requêtes SQL brutes ou ORM)
    ├── routers/             # Endpoints et routes de l'API REST
    ├── services/            # Couche Métier (Règles de gestion, chiffrement SHA-512)
    ├── .env                 # Variables d'environnement serveur (clés privées, credentials DB)
    ├── main.go              # Point d'entrée de l'API de données (Serveur)
    └── go.mod               # Dépendances du module serveur
```
    
    
## Prérequis à l'installation

Avant de lancer le projet, assurez-vous d'avoir installé sur votre machine :
* **Visual Studio Code** (Environnement de développement préconisé).
* **Go (Golang)** (v1.22+) configuré dans votre variable d'environnement PATH.
* **WampServer** (Serveur MySQL local).
* **DBeaver** (Gestionnaire de base de données universel).

---

## Procédure d'installation et de Lancement (De A à Z)

### Étape 1 : Ouverture du projet dans VS Code
1. Clonez ce dépôt Git sur votre machine locale ou téléchargez le dossier.
2. Ouvrez l'application **Visual Studio Code**.
3. Cliquez sur **Fichier** > **Ouvrir le dossier...** (ou *Open Folder...*) et sélectionnez le répertoire racine du projet contenant les deux sous-dossiers (`client` et `server`).
4. *Conseil :* Si VS Code l'indique, acceptez l'installation des dépendances recommandées de l'extension Go (by Google).

### Étape 2 : Initialisation de la Base de Données (Wamp & DBeaver)
1. Démarrez **WampServer** sur votre ordinateur et attendez que l'icône de la barre des tâches devienne **verte** (signifiant que MySQL et Apache tournent correctement).
2. Ouvrez **DBeaver**.
3. Créez une nouvelle connexion vers votre base locale : sélectionnez **MySQL**, configurez l'hôte sur `localhost` ou `127.0.0.1`, utilisez le nom d'utilisateur `root` et laissez le champ mot de passe complètement vide (configuration classique locale).
4. Faites un clic droit sur vos connexions et créez une nouvelle base de données nommée **`souleu_db`**.
5. Allez dans le dossier `server/migration/` depuis l'explorateur de VS Code, copiez l'intégralité du contenu des scripts SQL présents (`.sql`) qui contiennent les définitions de tables et d'insertions.
6. De retour sur DBeaver, ouvrez un nouvel **Éditeur SQL** rattaché à votre base `souleu_db`, collez le script complet et cliquez sur le bouton d'exécution générale (**Execute SQL Script** / `Alt + X`). Vos tables (*users*, *threads*, *messages*, *reactions*, *tags*, etc.) sont prêtes !

### Étape 3 : Fichiers de Configuration (.env)
1. Dans le dossier `server/`, créez un fichier nommé `.env` à la racine de ce sous-répertoire et collez-y les accès suivants :
   ```env
   DB_HOST=127.0.0.1
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=
   DB_NAME=souleu_db
   JWT_SECRET=votre_super_cle_secrete_jwt_sha512_pour_souleu
   ```

---
### Étape 4 : Téléchargement, installation des paquets et Lancement simultané
Pour lancer l'ensemble, ouvrez le terminal intégré de VS Code (`Ctrl + \`` ou via le menu **Terminal** > **Nouveau Terminal**).

#### Action 1 : Installer l'intégralité des paquets et lancer le Serveur (API)
Dans le premier onglet du terminal dédié au backend, exécutez les commandes suivantes :
```bash
cd server

# 1. Nettoyage et synchronisation du module Go
go mod tidy

# 2. Installation manuelle de chaque paquet requis côté Serveur
go get [github.com/gorilla/mux](https://github.com/gorilla/mux)                  # Le routeur de requêtes (supermux)
go get [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)          # Le driver de connexion MySQL pour Wamp
go get [github.com/joho/godotenv](https://github.com/joho/godotenv)                # Le gestionnaire de fichier de configuration .env
go get [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt/v5)            # Le système de génération de jetons de session (JWT)
go get filippo.io/edwards25519                 # Dépendance cryptographique requise par MySQL
go get golang.org/x/text                       # Gestion des encodages et des textes sécurisés

# 3. Lancement du fichier principal
go run main.go
```

#### Action 2 : Installer l'intégralité des paquets et lancer le Client (Interface Front)
Cliquez sur le bouton **`+`** en haut à droite de l'espace terminal de VS Code pour ouvrir un second onglet indépendant, puis exécutez les commandes suivantes :

```bash
cd client

# 1. Nettoyage et synchronisation du module Go
go mod tidy

# 2. Installation manuelle de chaque paquet requis côté Client
go get [github.com/gorilla/mux](https://github.com/gorilla/mux)                  # Le routeur de requêtes (supermux)
go get [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)          # Le driver MySQL (requis pour les structures communes)
go get [github.com/joho/godotenv](https://github.com/joho/godotenv)                # Le gestionnaire de fichier de configuration .env
go get filippo.io/edwards25519                 # Dépendance de sécurité et cryptographie
go get golang.org/x/text                       # Gestion des encodages de texte pour les templates HTML

# 3. Lancement du fichier principal frontal
go run main.go
```

> 🖥️ **Accès à l'Interface Utilisateur :** Une fois les deux terminaux lancés, ouvrez votre navigateur internet et rendez-vous sur **`http://localhost:3000`** pour charger l'interface visuelle de **Soulèu** et naviguer sur la plateforme.

---

## Cartographie complète des Routes (Localhost:3000)

L'application distribue ses interfaces utilisateur et ses formulaires sur le port `3000` via le routeur **Gorilla Mux**. Voici la liste complète des adresses disponibles pour naviguer sur la plateforme :

### Vues Publiques (Accessibles par tous)
* **`GET http://localhost:3000/`** : Page d'accueil principale avec affichage de la liste des fils.
* **`GET http://localhost:3000/all`** : Flux global affichant l'intégralité des fils avec gestion de la pagination générale.
* **`GET http://localhost:3000/fils/petanque`** : Espace thématique dédié à la pétanque avec pagination.
* **`GET http://localhost:3000/fils/cuisine`** : Espace thématique culinaire provençal avec pagination.
* **`GET http://localhost:3000/fils/nature`** : Espace thématique nature et paysages avec pagination.
* **`GET http://localhost:3000/fil/{id}/messages`** : Consultation détaillée d'un fil (remplacez `{id}` par le numéro du fil) affichant ses messages paginés.

---

### Vues & Actions Sécurisées (Requiert d'être connecté — `AuthMiddleware`)

#### Recherche
* **`GET http://localhost:3000/search`** : Permet aux membres d'effectuer une recherche par titre, pseudo ou par catégorie.

#### Gestion des Fils de Discussion
* **`GET http://localhost:3000/fils/create`** : Accès au formulaire de création d'un nouveau fil.
* **`POST http://localhost:3000/fils`** : Soumission et enregistrement du nouveau fil en base de données.
* **`GET http://localhost:3000/fil/{id}/edit`** : Accès au formulaire de modification d'un fil existant.
* **`POST http://localhost:3000/fil/{id}/update`** : Envoi des modifications apportées au fil.
* **`POST http://localhost:3000/fil/{id}/delete`** : Suppression définitive d'un fil et cascades associées.

#### Gestion des Messages & Interactions
* **`POST http://localhost:3000/fil/{id}/messages`** : Envoi et publication d'un nouveau message dans un fil ouvert.
* **`POST http://localhost:3000/message/reaction`** : Ajout, modification ou suppression d'une réaction Like/Dislike sur un message.
* **`GET http://localhost:3000/message/edit`** : Accès au formulaire d'édition d'un message spécifique.
* **`POST http://localhost:3000/message/update`** : Validation et modification du texte d'un message.
* **`POST http://localhost:3000/message/delete`** : Suppression d'un message et de ses réactions rattachées.