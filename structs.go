package main

import (
	"strings"
)

type XMLResponseString struct {
	Value string `xml:",chardata"`
}

type Coursework struct {
	Courses  []Course  `json:"CurrentCourses"`
	Grades   []Grade   `json:"CourseWork"`
	Midterms []Midterm `json:"Midterm"`
}

type Course struct {
	Id   string `json:"sm_crs_id"`
	Name string `json:"course_short_name"`
}

type Grade struct {
	CourseId   string `json:"sm_crs_id"`
	ModuleName string `json:"eval_method_name"`
	Point      string `json:"grade"`
	MaxPoint   string `json:"max_point"`
}

type Midterm struct {
	CourseName string `json:"course_full_name"`
	Percentage string `json:"total_perc"`
}

type Absence struct {
	AbsenceReports []AbsenceReport `json:"AbsenceReport"`
}

type AbsenceReport struct {
	CourseName   string `json:"Name"`
	AbsenceLevel string `json:"AbsenceLevel"`
}

type ResponseAPI struct {
	Error interface{} `json:"error"`
	Data  interface{} `json:"data"`
}

type AuthorizedAPI struct {
	IsAuthorized bool `json:"authorized"`
}

type CourseworkAPI struct {
	Id     string     `json:"-"`
	Code   string     `json:"code"`
	Name   string     `json:"name"`
	Grades []GradeAPI `json:"grades"`
}

type GradeAPI struct {
	Module   string `json:"module"`
	Point    string `json:"point"`
	MaxPoint string `json:"maxPoint"`
}

type MidtermAPI struct {
	Name       string `json:"name"`
	Percentage string `json:"percentage"`
}

type AbsenceReportAPI struct {
	CourseName string `json:"name"`
	Level      string `json:"level"`
}

func NewAuthorizedAPI(authorized string) AuthorizedAPI {
	authorizedAPI := AuthorizedAPI{}

	if strings.Compare(authorized, "True") == 0 {
		authorizedAPI.IsAuthorized = true
	} else {
		authorizedAPI.IsAuthorized = false
	}

	return authorizedAPI
}

func NewCourseworkAPI(course Course) CourseworkAPI {
	courseAPI := CourseworkAPI{}

	courseAPI.Id = course.Id
	courseAPI.Grades = []GradeAPI{}

	courseNameSplit := strings.Split(course.Name, "(")
	courseAPI.Name = strings.TrimSpace(courseNameSplit[0])
	courseAPI.Code = courseNameSplit[1][0 : len(courseNameSplit[1])-1]

	return courseAPI
}

func NewGradeAPI(grade Grade) GradeAPI {
	gradeAPI := GradeAPI{}

	gradeAPI.Module = grade.ModuleName
	gradeAPI.Point = grade.Point
	gradeAPI.MaxPoint = grade.MaxPoint

	return gradeAPI
}

func NewMidtermAPI(midterm Midterm) MidtermAPI {
	midtermAPI := MidtermAPI{}

	midtermAPI.Name = midterm.CourseName
	midtermAPI.Percentage = midterm.Percentage

	return midtermAPI
}

func NewAbsenceReportAPI(absenceReport AbsenceReport) AbsenceReportAPI {
	absenceReportAPI := AbsenceReportAPI{}

	absenceReportAPI.CourseName = absenceReport.CourseName
	absenceReportAPI.Level = absenceReport.AbsenceLevel

	return absenceReportAPI
}