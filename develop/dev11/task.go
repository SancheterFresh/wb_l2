package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

type Event struct {
	Id   int       `json:"Id"`
	User int       `json:"User"`
	Text string    `json:"Text"`
	Date time.Time `json:"Date"`
}

type Storage struct {
	mu     *sync.Mutex
	events map[int][]Event
}

func (s *Storage) Create(e *Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if events, ok := s.events[e.User]; ok {
		for _, event := range events {
			if event.Id == e.Id {
				return fmt.Errorf("событие с таким id уже существует для данного пользователя")
			}
		}
	}
	s.events[e.User] = append(s.events[e.User], *e)

	return nil
}

func (s *Storage) Update(e *Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	found := false

	var events []Event
	ok := false

	if events, ok = s.events[e.User]; !ok {
		return fmt.Errorf("пользователь не найден")
	}

	for i, event := range events {
		if event.Id == e.Id {
			s.events[e.User][i] = *e
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("событие с таким id не найдено для данного пользователя")
	}

	return nil
}

func (s *Storage) Delete(e *Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	found := false

	var events []Event
	ok := false

	if events, ok = s.events[e.User]; !ok {
		return fmt.Errorf("пользователь не найден")
	}

	for i, event := range events {
		if event.Id == e.Id {
			found = true
			l := len(s.events[e.User])
			s.events[e.User][i] = s.events[e.User][l-1]
			s.events[e.User] = s.events[e.User][:l-1]
			break
		}
	}

	if !found {
		return fmt.Errorf("событие с таким id не найдено для данного пользователя")
	}

	return nil
}

func (s *Storage) getEventsForDay(user int, date time.Time) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var res []Event
	var events []Event
	ok := false

	if events, ok = s.events[user]; !ok {
		return nil, fmt.Errorf("пользователь не найден")
	}

	for _, e := range events {

		if e.Date.Year() == date.Year() && e.Date.Month() == date.Month() && e.Date.Day() == date.Day() {
			res = append(res, e)
		}
	}

	return res, nil
}

func (s *Storage) getEventsForWeek(user int, date time.Time) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var res []Event

	var events []Event
	ok := false

	if events, ok = s.events[user]; !ok {
		return nil, fmt.Errorf("пользователь не найден")
	}

	for _, e := range events {
		wyear, week := e.Date.ISOWeek()
		swyear, sweek := date.ISOWeek()
		if wyear == swyear && week == sweek {
			res = append(res, e)
		}
	}

	return res, nil
}

func (s *Storage) getEventsForMonth(user int, date time.Time) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var res []Event

	var events []Event
	ok := false

	if events, ok = s.events[user]; !ok {
		return nil, fmt.Errorf("пользователь не найден")
	}

	for _, e := range events {
		if e.Date.Year() == date.Year() && e.Date.Month() == date.Month() {
			res = append(res, e)
		}
	}

	return res, nil
}

type dataStore struct {
	data *Storage
}

func newDataStore() *dataStore {
	data := &Storage{}
	data.mu = &sync.Mutex{}
	data.events = make(map[int][]Event)
	return &dataStore{data: data}
}

func getResponse(w http.ResponseWriter, r string, ev []Event, status int) {
	resp := struct {
		Result string
		Events []Event
	}{Result: r, Events: ev}

	marshaled, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshaled)
}

func getErrResponse(w http.ResponseWriter, e string, status int) {
	errResp := struct {
		Error string
	}{Error: e}

	marshaled, err := json.Marshal(errResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshaled)
}

func getPostData(w http.ResponseWriter, r *http.Request) (Event, error) {

	w.Header().Set("Content-Type", "application/json")

	var newEvent Event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		getErrResponse(w, " error with reading json", http.StatusNotAcceptable)
		return newEvent, errors.New("400: bad int")
	}

	json.Unmarshal(reqBody, &newEvent)

	return newEvent, nil
}

func (d *dataStore) createEvent(w http.ResponseWriter, r *http.Request) {
	var e Event
	var err error
	e, err = getPostData(w, r)

	if err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := d.data.Create(&e); err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	getResponse(w, "Событие добавлено", []Event{e}, http.StatusCreated)

}

func (d *dataStore) updateEvent(w http.ResponseWriter, r *http.Request) {
	var e Event
	var err error
	e, err = getPostData(w, r)
	if err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := d.data.Update(&e); err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	getResponse(w, "Событие обновлено", []Event{e}, http.StatusOK)

}

func (d *dataStore) deleteEvent(w http.ResponseWriter, r *http.Request) {
	var e Event
	var err error
	e, err = getPostData(w, r)
	if err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = d.data.Delete(&e); err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
	}

	getResponse(w, "Событие удалено", []Event{e}, http.StatusOK)
}

func (d *dataStore) eventsForDay(w http.ResponseWriter, r *http.Request) {
	var ev []Event

	user, err := strconv.Atoi(r.URL.Query().Get("user"))
	if err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	if ev, err = d.data.getEventsForDay(user, date); err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	getResponse(w, "Выполнено", ev, http.StatusOK)
}
func (d *dataStore) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	var ev []Event

	user, err := strconv.Atoi(r.URL.Query().Get("user"))
	if err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ev, err = d.data.getEventsForWeek(user, date); err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	getResponse(w, "Выполнено", ev, http.StatusOK)
}

func (d *dataStore) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	var ev []Event

	user, err := strconv.Atoi(r.URL.Query().Get("user"))
	if err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ev, err = d.data.getEventsForMonth(user, date); err != nil {
		getErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	getResponse(w, "Выполнено", ev, http.StatusOK)
}

func main() {
	mux := http.NewServeMux()

	ds := newDataStore()
	//POST
	mux.HandleFunc("/create_event", ds.createEvent)
	mux.HandleFunc("/update_event", ds.updateEvent)
	mux.HandleFunc("/delete_event", ds.deleteEvent)
	//GET
	mux.HandleFunc("/events_for_day", ds.eventsForDay)
	mux.HandleFunc("/events_for_week", ds.eventsForWeek)
	mux.HandleFunc("/events_for_month", ds.eventsForMonth)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))

}
