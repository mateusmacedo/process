/* eslint-disable no-undef */
db = db.getSiblingDB('admin')
db.createUser({
  user: 'process_user',
  pwd: 'Secret*123',
  roles: [
    {
      role: 'dbOwner',
      db: 'process_db'
    }
  ]
})
