package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/cale-i/building-modern-web-applications-with-go-bookings-project/internal/models"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name   string
	url    string
	method string
	// params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", http.StatusOK},
	{"about", "/about", "GET", http.StatusOK},
	{"gq", "/generals-quarters", "GET", http.StatusOK},
	{"ms", "/majors-suite", "GET", http.StatusOK},
	{"sa", "/search-availability", "GET", http.StatusOK},
	{"contact", "/contact", "GET", http.StatusOK},
	{"non-existent", "/hoge/hoge/hoge", "GET", http.StatusNotFound},

	// new routes
	{"login", "/user/login", "GET", http.StatusOK},
	{"logout", "/user/logout", "GET", http.StatusOK},
	{"dashboard", "/admin/dashboard", "GET", http.StatusOK},
	{"new_res", "/admin/reservations-new", "GET", http.StatusOK},
	{"all_res", "/admin/reservations-all", "GET", http.StatusOK},
	{"show_res", "/admin/reservations/new/1/show", "GET", http.StatusOK},

	// {"post-search-avail", "/search-availability", "POST", []postData{
	// 	{key: "start-date", value: "2020-01-01"},
	// 	{key: "end-date", value: "2020-01-02"},
	// }, http.StatusOK},
	// {"post-search-avail-json", "/search-availability-json", "POST", []postData{
	// 	{key: "start-date", value: "2020-01-01"},
	// 	{key: "end-date", value: "2020-01-02"},
	// }, http.StatusOK},

	// {"make_reservation_post", "/make-reservation", "POST", []postData{
	// 	{key: "first_name", value: "John"},
	// 	{key: "last_name", value: "Smith"},
	// 	{key: "email", value: "me@here.com"},
	// 	{key: "phone", value: "055-555-5555"},
	// }, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)

	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
			// } else {
			// 	values := url.Values{}
			// 	for _, x := range e.params {
			// 		values.Add(x.key, x.value)
			// 	}
			// 	resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			// 	if err != nil {
			// 		t.Log(err)
			// 		t.Fatal(err)
			// 	}

			// 	if resp.StatusCode != e.expectedStatusCode {
			// 		t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			// 	}
		}
	}
}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()

	session.Put(ctx, "reservation", reservation)
	handler := http.HandlerFunc(Repo.Reservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
	}

	// test case where reservation is not in session (reset everything)
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test with non-existent room
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()

	reservation.RoomID = 1000000

	session.Put(ctx, "reservation", reservation)
	handler = http.HandlerFunc(Repo.Reservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func TestRepository_PostReservation(t *testing.T) {

	reservation := models.Reservation{
		FirstName: "John",
		LastName:  "Smith",
		Email:     "john@smith.com",
		Phone:     "12345789",
		StartDate: time.Now(),
		EndDate:   time.Now(),
		RoomID:    1,
	}

	postedData := getDefaultPostedData(reservation)

	req, _ := http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Repo.PostReservation)

	// Added own test
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for missing session: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// original test code
	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr = httptest.NewRecorder()

	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// test for missing post body
	req, _ = http.NewRequest("POST", "/make-reservation", nil) // make a post request
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr = httptest.NewRecorder()

	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for missing post body: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for invalid start date
	postedData = getDefaultPostedData(reservation)
	postedData.Set("start_date", "invalid")
	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode())) // make a post request
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr = httptest.NewRecorder()

	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for invalid start date: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
	// test for invalid end date
	postedData = getDefaultPostedData(reservation)
	postedData.Set("end_date", "invalid")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode())) // make a post request
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr = httptest.NewRecorder()

	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for invalid end date: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for invalid room id -errors when room is not a number
	postedData = getDefaultPostedData(reservation)
	postedData.Set("room_id", "invalid")
	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode())) // make a post request
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr = httptest.NewRecorder()

	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler returned wrong response code for invalid room id: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for invalid data
	postedData = getDefaultPostedData(reservation)
	postedData.Set("first_name", "J")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode())) // make a post request
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr = httptest.NewRecorder()

	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		// テンプレートにFormとDataを渡すため、リダイレクトからrender.Templateに変更
		t.Errorf("PostReservation handler returned wrong response code for invalid data: got %d, wanted %d", rr.Code, http.StatusOK)
	}

	// test for failure to insert reservation into database
	reservation.RoomID = 2
	postedData = getDefaultPostedData(reservation)

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode())) // make a post request
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr = httptest.NewRecorder()

	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler failed when trying to fail inserting reservation: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for failure to insert restriction into database
	reservation.RoomID = 1000
	postedData = getDefaultPostedData(reservation)

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode())) // make a post request
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr = httptest.NewRecorder()

	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler failed when trying to fail inserting restriction: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func TestRepository_AvailabilityJson(t *testing.T) {
	// first case - rooms are not available
	reqBody := "start_date=2020-01-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end_date=2020-01-02")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=2")

	// create request
	req, _ := http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))

	// get context
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// make handler handlerfunc
	handler := http.HandlerFunc(Repo.AvailabilityJSON)

	// get response recoder
	rr := httptest.NewRecorder()
	// make request
	handler.ServeHTTP(rr, req)

	// session.Put(ctx, "reservation", reservation)

	var j jsonResponse
	err := json.Unmarshal([]byte(rr.Body.String()), &j)
	if err != nil {
		t.Error("failed to parse json")
	}

	// if rr.Code != http.StatusTemporaryRedirect {
	// 	t.Errorf("PostReservation handler failed when trying to fail inserting restriction: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	// }
}

var loginTests = []struct {
	name               string
	email              string
	expectedStatusCode int
	expectedHTML       string
	expectedLocation   string
}{
	{"valid-credentials", "valid@valid.com", http.StatusSeeOther, "", "/"},
	{"invalid-credentials", "invalid@invalid.com", http.StatusSeeOther, "", "/user/login"},
	{"invalid-data", "invalid", http.StatusOK, `action="/user/login"`, ""},
}

func TestLogin(t *testing.T) {
	// range through all tests

	for _, e := range loginTests {
		postedData := url.Values{}
		postedData.Add("email", e.email)
		postedData.Add("password", "password")

		// create request
		req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(postedData.Encode()))
		ctx := getCtx(req)
		req = req.WithContext(ctx)

		// set the header
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		// call the handler
		handler := http.HandlerFunc(Repo.PostShowLogin)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatusCode {
			t.Errorf("failed %s: expected code %d, but got %d", e.name, e.expectedStatusCode, rr.Code)
		}

		if e.expectedLocation != "" {
			// get the URL from test
			actualLoc, _ := rr.Result().Location()
			if actualLoc.String() != e.expectedLocation {
				t.Errorf("failed %s: expected location %s, but got location %s", e.name, e.expectedLocation, actualLoc)
			}
		}

		// checking for expected values in HTML
		if e.expectedHTML != "" {
			// read the response body into a string
			html := rr.Body.String()
			if !strings.Contains(html, e.expectedHTML) {
				t.Errorf("failed %s: expected to find %s, but did not", e.name, e.expectedHTML)
			}
		}
	}
}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}

func getDefaultPostedData(reservation models.Reservation) url.Values {
	layout := "2006-01-02"
	postedData := url.Values{}
	postedData.Add("start_date", reservation.StartDate.Format(layout))
	postedData.Add("end_date", reservation.EndDate.Format(layout))
	postedData.Add("first_name", reservation.FirstName)
	postedData.Add("last_name", reservation.LastName)
	postedData.Add("email", reservation.Email)
	postedData.Add("phone", reservation.Phone)
	postedData.Add("room_id", strconv.Itoa(reservation.RoomID))

	return postedData
}
