USE forum_souleu;

INSERT INTO users (id, pseudo, email, passwd) VALUES
(1, 'Quentin', 'quentin@email.com', 'hash_secure_1'),
(2, 'Michel', 'michel@email.com', 'hash_secure_2'),
(3, 'Thomas', 'thomas@email.com', 'hash_secure_3'),
(4, 'Gaëtan', 'gaetan@email.com', 'hash_secure_4'),
(5, 'Mamie Minou', 'minou@email.com', 'hash_secure_5'),
(6, 'Papi Li', 'papili@email.com', 'hash_secure_6');

INSERT INTO tags (id, name, description) VALUES
(1, 'Pétanque', 'Règles officielles et astuces du carreau.'),
(2, 'Cuisine du Sud', 'Les secrets des recettes provençales.');

INSERT INTO fils (id, titre, statut, score, fk_user) VALUES
(1, 'Règle officielle : Pousser une boule adverse au milieu ?', 0, 12, 1),
(2, 'La vraie recette de la Ratatouille niçoise', 0, 25, 4);

INSERT INTO fils_tags (fk_fil, fk_tag) VALUES
(1, 1),
(2, 2);

INSERT INTO messages (contenu, is_published, fk_user, fk_fil) VALUES

('Salut l''équipe, j''ai un doute sur une règle. Est-ce qu''on a le droit de shooter la boule de l''adversaire pour la sortir du cadre ? Merci !', 1, 1, 1),
('Ah fada, bien sûr que tu as le droit ! C''est ce qu''on appelle faire un carreau si tu prends sa place, ou un palet !', 1, 2, 1),
('Je confirme Michel, par contre attention à ne pas sortir le cochonnet du terrain sinon la mène est annulée si les deux équipes ont encore des boules.', 1, 3, 1),

('Bonjour à tous, je cherche la véritable recette de la ratatouille du Sud. Faut-il cuire les légumes ensemble ou séparément ?', 1, 4, 2),
('Mon p''tit Gaëtan, le secret c''est de faire revenir chaque légume séparément dans de l''huile d''olive avant de tout mijoter ensemble. Bon appétit !', 1, 5, 2),
('Tout à fait d''accord avec Mamie Minou ! Et n''oublie surtout pas les herbes de Provence fraîchement cueillies !', 1, 6, 2);