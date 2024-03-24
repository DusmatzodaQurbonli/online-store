
INSERT INTO Products VALUES (1, 'Ноутбук'), (2, 'Телевизор'), (3, 'Телефон'), (4, 'Системный блок'), (5, 'Часы'), (6, 'Микрофон');
INSERT INTO Shelves VALUES (1, 'A'), (2, 'B'), (3, 'Ж'), (4, 'З'), (5, 'В');
INSERT INTO ProductsShelves VALUES (1, 1, TRUE), (2, 1, TRUE), (3, 2, TRUE), (4, 3, TRUE), (5, 3, TRUE), (6, 3, TRUE), (3, 4, FALSE), (3, 5, FALSE), (5, 1, FALSE);
INSERT INTO Orders VALUES (10), (11), (14), (15);
INSERT INTO ProductsOrders VALUES (1, 10, 2), (3, 10, 1), (6, 10, 1), (2, 11, 3), (1, 14, 3), (4, 14, 4), (5, 15, 1);
