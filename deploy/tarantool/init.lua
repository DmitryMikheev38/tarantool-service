box.cfg({listen="3301"})
box.schema.user.create('root', {password='passw0rd', if_not_exists=true})
box.schema.user.grant('root', 'super', nil, nil, {if_not_exists=true})
require('msgpack').cfg{encode_invalid_as_nil = true}

books = box.schema.create_space('books', { if_not_exists = true })
books:format({
    { name = 'book_id', type = 'string'},
    { name = 'title', type = 'string'},
    { name = 'description', type = 'string'},
})
books:create_index('pk', {if_not_exists = true, parts = {{ field = 'book_id', type = 'string'}}})

authors = box.schema.create_space('authors', { if_not_exists = true })
authors:format({
    { name = 'author_id', type = 'string'},
    { name = 'name', type = 'string'},
    { name = 'books_count', type = 'integer', is_nullable = true},
})
authors:create_index('pk', {if_not_exists = true, parts = { { field = 'author_id', type = 'string'}}})
authors:create_index('books_count', {if_not_exists = true, parts = { { field = 'books_count', type = 'integer'}}})

books_authors =  box.schema.create_space('books_authors', { if_not_exists = true })
books_authors:format({{ name = 'author_id', type = 'string'}, { name = 'book_id', type = 'string'}})
books_authors:create_index('pk', {if_not_exists = true, parts = {
    { field = 'author_id', type = 'string'},
    { field = 'book_id', type = 'string'}
}})