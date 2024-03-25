CREATE TABLE Products (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE Shelves (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE ProductsShelves (
    product_id INTEGER,
    shelf_id INTEGER,
    is_main BOOLEAN,
    FOREIGN KEY (product_id) REFERENCES Products(id),
    FOREIGN KEY (shelf_id) REFERENCES Shelves(id)
);

CREATE TABLE Orders (
    id INTEGER PRIMARY KEY
);

CREATE TABLE ProductsOrders (
    product_id INTEGER,
    order_id INTEGER,
    quantity INTEGER,
    FOREIGN KEY (product_id) REFERENCES Products(id),
    FOREIGN KEY (order_id) REFERENCES Orders(id)
);
