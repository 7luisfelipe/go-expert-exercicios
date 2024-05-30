CREATE DATABASE IF NOT EXISTS ordersdb;

USE ordersdb;


-- Table de pedidos realizados
CREATE TABLE IF NOT EXISTS pedidos(
    id INT AUTO_INCREMENT PRIMARY KEY,
    numero_pedido INT NOT NULL,
    nome_produto VARCHAR(100) NOT NULL,
    quantidade INT NOT NULL,
    preco_unitario DECIMAL(10,2) NOT NULL
);

-- Associações aleatórias entre produtos e pedidos
INSERT INTO pedidos (numero_pedido, nome_produto, quantidade, preco_unitario) VALUES
(700300, 'Notebook acer nitro 5', 1,6499.99),
(700400, 'MacBook Air 13 chip M2', 2,7249.99);