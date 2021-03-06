# Building Modern Web Applications with Go (Golang)

https://www.udemy.com/course/building-modern-web-applications-with-go/

## Bookings and Reservations

This is the repository for my booking and reservations project

- Built in Go version 1.16
- [alex edwards SCS](https://github.com/alexedwards/scs) session management - v2.4.0
- [chi router](https://github.com/go-chi/chi) - v1.5.4
- [nosurf](https://github.com/justinas/nosurf) - v1.1.1
- [govalidator](https://github.com/asaskevich/govalidator) - v11.0.0
- [soda](https://gobuffalo.io/en/docs/db/getting-started) - v5.3.1
- [pgx](https://github.com/jackc/pgx/) - v4.11.0
- [Go Simple Mail](https://github.com/xhit/go-simple-mail) - v2.8.1

Frontend

- [Bootstrap](https://getbootstrap.com/) - v5.0
- [vanillajs-datepicker](https://github.com/mymth/vanillajs-datepicker) - v1.1.4
- [notie](https://github.com/jaredreich/notie)
- [sweetalert2](https://github.com/sweetalert2/sweetalert2) -v10
- [RoyalUI-Free-Bootstrap-Admin-Template](https://github.com/BootstrapDash/RoyalUI-Free-Bootstrap-Admin-Template) admin dashboard
- [Simple-DataTables](https://github.com/fiduswriter/Simple-DataTables)

Misc

- [MailHog](https://github.com/mailhog/MailHog) local mail server
- [Foundation](https://get.foundation/emails.html)

### Reminder

#### How to Install

[pgx]

```
go get github.com/jackc/pgx/v4
```

(driver.go)

```
package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

```
