# Event Booking REST API

This is a simple API that enable users to create events and book events.

## Dependencies

- `github.com/gin-gonic/gin` : web framework
- `sql` + `github.com/mattn/go-sqlite3` for database
- `golang.org/x/crypto/bcrypt`: encrypt user password
- `golang-jwt/jwt/v5`: generate jwt token