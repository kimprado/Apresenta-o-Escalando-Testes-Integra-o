db.auth('mongo', 'mongo');

let nomeDB = 'aluguel'

let testes = [
    '',
    'TestConectarDB',
    'TestConsultarMartelo',
    'TestConsultarAtivos',
];

let hoje = new Date();
let inicio = new Date(hoje);

testes.forEach(teste => {
    if (teste == '')  {
      db = db.getSiblingDB(nomeDB);
    } else {
      db = db.getSiblingDB(nomeDB + "-" + teste);
    }

    let data = hoje.setDate(new Date(hoje).getDate() + 1);
    data = new Date(data);

    let Martelo = {
        "_id" : ObjectId("555b6e830850536438063761"),
        "codigo": "e-001",
        "descricao": "Martelo",
        "criacao": data,
        "ativo": true
    };

    let Furadeira = {
        "_id" : ObjectId("555b6e830850536438063762"),
        "codigo": "e-002",
        "descricao": "Furadeira",
        "rpm": "1500",
        "criacao": data,
        "ativo": true
    };

    let res = [
        db.Equipamento.drop(),
        db.Equipamento.insert(Martelo),
        db.Equipamento.insert(Furadeira),
        db.Equipamento.createIndex({ ativo: 1 }),
    ];

    printjson(res)

});