package rtusite

type LessonBell struct {
	Start string
	End string
}

type LessonForm struct {
	Form string
	Links []string
}

type Meta struct {
	Name string
	Places []string
	Bells []LessonBell
	Forms []LessonForm
}