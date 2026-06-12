USE forum_souleu;

SET FOREIGN_KEY_CHECKS = 0;
TRUNCATE TABLE reactions;
TRUNCATE TABLE messages;
TRUNCATE TABLE fils;
TRUNCATE TABLE tags;
TRUNCATE TABLE users;
SET FOREIGN_KEY_CHECKS = 1;

INSERT INTO users (id, pseudo, email, passwd, is_admin, is_ban) VALUES
(1, 'Quentin', 'quentin@email.com', 'hash_secure_1', FALSE, FALSE),
(2, 'Michel', 'michel@email.com', 'hash_secure_2', FALSE, FALSE),
(3, 'Thomas', 'thomas@email.com', 'hash_secure_3', FALSE, FALSE),
(4, 'Gaëtan', 'gaetan@email.com', 'hash_secure_4', FALSE, FALSE),
(5, 'Mamie Minou', 'minou@email.com', 'hash_secure_5', FALSE, FALSE),
(6, 'Papi Li', 'papili@email.com', 'hash_secure_6', FALSE, FALSE),
(7, 'Admin_Marius', 'marius@souleu.fr', 'hash_admin_7', TRUE, FALSE); 

INSERT INTO tags (id, nom, description) VALUES
(1, 'Pétanque', 'Règles officielles et astuces du carreau.'),
(2, 'Cuisine du Sud', 'Les secrets des recettes provençales.'),
(3, 'Coins secrets', 'Les plus belles calanques et balades caches de Provence.');

INSERT INTO fils (id, titre, statut, score, fk_user, fk_tag) VALUES
(1, 'Règle officielle : Pousser une boule adverse au milieu ?', 'ouvert', 12, 1, 1),      
(2, 'La vraie recette de la Ratatouille niçoise', 'ouvert', 25, 4, 2),                  
(3, 'Le meilleur spot pour voir le coucher de soleil à Cassis', 'ouvert', 45, 3, 3),    
(4, '[Concours Clos] Grand tournoi de pétanque de Cabriès 2026', 'fermé', 8, 7, 1),     
(5, 'Ancien fil de test à ne pas afficher', 'archivé', -5, 2, 1);                       

INSERT INTO messages (id, contenu, fk_user, fk_fil) VALUES
(1, 'Salut l''équipe, j''ai un doute sur une règle. Est-ce qu''on a le droit de shooter la boule de l''adversaire pour la sortir du cadre ? Merci !', 1, 1),
(2, 'Ah fada, bien sûr que tu as le droit ! C''est ce qu''on appelle faire un carreau si tu prends sa place, ou un palet !', 2, 1),
(3, 'Je confirme Michel, par contre attention à ne pas sortir le cochonnet du terrain sinon la mène est annulée si les deux équipes ont encore des boules.', 3, 1),

(4, 'Bonjour à tous, je cherche la véritable recette de la ratatouille du Sud. Faut-il cuire les légumes ensemble ou séparément ?', 4, 2),
(5, 'Mon p''tit Gaëtan, le secret c''est de faire revenir chaque légume séparément dans de l''huile d''olive avant de tout mijoter ensemble. Bon appétit !', 5, 2),
(6, 'Tout à fait d''accord avec Mamie Minou ! Et n''oublie surtout pas les herbes de Provence fraîchement cueillies !', 6, 2),

(7, 'Je cherche un endroit un peu caché à Cassis pour poser les serviettes et regarder le coucher de soleil sans la foule. Des idées ?', 3, 3),
(8, 'Va du côté des roches plates après la plage du Corton. C''est un peu escarpé mais la vue sur le Cap Canaille est incroyable !', 2, 3),

(9, 'Les inscriptions pour le tournoi annuel de Cabriès sont désormais ouvertes. Équipes de 3 uniquement !', 7, 4),
(10, 'Dommage que ce soit déjà fini, l''ambiance avait l''air incroyable cette année.', 4, 4);