db.auth('mongo', 'mongo');

nomeDB = 'aluguel'

let testes = [
  '',
  'TestConectarDB',
];

testes.forEach(teste => {
  if (teste == '')  {
    db = db.getSiblingDB(nomeDB);
  } else {
    db = db.getSiblingDB(nomeDB + "-" + teste);
  }

  db.createUser({
    user: 'aluguel',
    pwd: 'aluguel',
    roles: [
      {
        role: 'root',
        db: 'admin',
      },
    ],
  });

});