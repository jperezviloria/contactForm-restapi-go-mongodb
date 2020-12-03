package repository

import "gopkg.in/mgo.v2"

type Context struct {
	Session *mgo.Session
}

func (context *Context) Close() {
	context.Session.Close()
}

func (context *Context) DBCollection(name string) *mgo.Collection {
	return context.Session.DB(DBName).C(CName)
}

func NewContext() *Context {
	session := getSession()
	context := &Context{
		Session: session,
	}
	return context
}
