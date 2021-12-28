box.cfg({listen="127.0.0.1:3301"})
require('msgpack').cfg{encode_invalid_as_nil = true}
users = box.schema.create_space('users', { if_not_exists = true })
users:format({{ name = 'user_id', type = 'number'}, { name = 'fullname', type = 'string'}})

users:create_index('pk', {if_not_exists = true, parts = { { field = 'user_id', type = 'number'}}})