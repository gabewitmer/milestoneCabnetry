package main

import (
	"net/http"

	"github.com/cagnosolutions/web"
)

const (
	USERNAME = "admin"
	PASSWORD = "admin"
)

var mux = web.NewMux()
var tmpl = web.NewTmplCache()

func main() {
	mux.AddRoutes(home, bath, build, contact, decks, design, flooring, garage, kitchen, more, otherrooms, roof, windowsDoors)
	mux.AddSecureRoutes(ADMIN)
	http.ListenAndServe(":8080", mux)
}

var home = web.Route{"GET", "/", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "index.tmpl", web.Model{
		"page": "home",
	})
	return
}}

var bath = web.Route{"GET", "/bath", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "bath.tmpl", web.Model{
		"page": "bath",
	})
	return
}}

var build = web.Route{"GET", "/build", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "build.tmpl", web.Model{
		"page": "more",
	})
	return
}}

var contact = web.Route{"GET", "/contact", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "contact.tmpl", web.Model{
		"page": "",
	})
	return
}}

var decks = web.Route{"GET", "/decks", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "decks.tmpl", web.Model{
		"page": "more",
	})
	return
}}

var design = web.Route{"GET", "/design", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "design.tmpl", web.Model{
		"page": "more",
	})
	return
}}

var flooring = web.Route{"GET", "/flooring", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "flooring.tmpl", web.Model{
		"page": "more",
	})
	return
}}

var garage = web.Route{"GET", "/garage", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "garage.tmpl", web.Model{
		"page": "garage",
	})
	return
}}

var kitchen = web.Route{"GET", "/kitchen", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "kitchen.tmpl", web.Model{
		"page": "kitchen",
	})
	return
}}

var more = web.Route{"GET", "/more", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "more.tmpl", web.Model{
		"page": "more",
	})
	return
}}

var otherrooms = web.Route{"GET", "/otherrooms", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "otherrooms.tmpl", web.Model{
		"page": "otherrooms",
	})
	return
}}

var roof = web.Route{"GET", "/roof", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "roof.tmpl", web.Model{
		"page": "roof",
	})
	return
}}

var windowsDoors = web.Route{"GET", "/windowsDoors", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "windowsDoors.tmpl", web.Model{
		"page": "more",
	})
	return
}}
