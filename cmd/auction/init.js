// init.js
db = db.getSiblingDB('auctions');
db.users.insertMany([
    { 
        _id:"a6deb003-ba12-4dae-801f-e2f75a56de47", 
        name: "Luis Felipe"
    },
]);

/*
db.auctions.insertMany([
    {
        _id: "064f9f8a-42f6-4a12-9b03-9620d001aef5",
        product_name: "IMAC",
        category: "Eletrônicos",
        description: "Produto em excelente estado, com pouco uso e garantia.",
        condition: 1,
        status: 0,
        timestamp: Long('1730158227')
    },
]);
*/