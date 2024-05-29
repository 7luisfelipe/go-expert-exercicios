CREATE DATABASE IF NOT EXISTS ordersdb;

USE ordersdb;

-- Table de produtos
CREATE TABLE IF NOT EXISTS produtos(
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    preco DECIMAL(10,2) NOT NULL
);

-- Table de pedidos realizados
CREATE TABLE IF NOT EXISTS pedidos(
    id INT AUTO_INCREMENT PRIMARY KEY,
    numero_pedido INT NOT NULL
);

-- Table com os itens do pedido
CREATE TABLE IF NOT EXISTS itens_pedido(
    id INT AUTO_INCREMENT PRIMARY KEY,
    produto_id INT NOT NULL,
    pedido_id INT NOT NULL,
    quantidade INT NOT NULL,
    preco DECIMAL(10,2) NOT NULL,
    FOREIGN KEY (produto_id) REFERENCES produtos(id),
    FOREIGN KEY (pedido_id) REFERENCES pedidos(id)
);

-- Inserts para tabela de produtos
INSERT INTO produtos (nome, preco) VALUES
('Curso de GO', 49.99),
('Curso de Python', 59.99),
('Curso de Java', 69.99),
('Curso de C++', 79.99),
('Curso de Ruby', 94.99);

-- Inserts para tabela de pedidos
INSERT INTO pedidos (numero_pedido) VALUES
(1001),
(1002),
(1003),
(1004),
(1005);

-- Inserts para tabela de itens de pedido
-- Associações aleatórias entre produtos e pedidos
INSERT INTO itens_pedido (produto_id, pedido_id, quantidade, preco) VALUES
(1, 1, 1, 49.99),  -- Curso de GO para Pedido 1
(2, 2, 2, 59.99),  -- Curso de JavaScript para Pedido 2
(3, 3, 3, 69.99),  -- Curso de Java para Pedido 3
(4, 4, 4, 79.99),  -- Curso de C++ para Pedido 4
(5, 5, 5, 94.99),  -- Curso de Ruby para Pedido 5
(1, 1, 2, 49.99),  -- Curso de GO para Pedido 1
(2, 2, 3, 59.99),  -- Curso de JavaScript para Pedido 2
(3, 3, 4, 69.99),  -- Curso de Java para Pedido 3
(4, 4, 5, 79.99),  -- Curso de C++ para Pedido 4
(5, 5, 1, 94.99);  -- Curso de Ruby para Pedido 1