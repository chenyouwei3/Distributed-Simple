package grades

var Stus []Student

func Init() {
	Stus = []Student{
		{
			Id:   1,
			Name: "Alice",
			Grades: []Grade{
				{
					Title: "Math Exam",
					Type:  GradeExam,
					Score: 85.0,
				},
				{
					Title: "Physics Exam",
					Type:  GradeExam,
					Score: 78.5,
				},
				{
					Title: "English Test",
					Type:  GradeTest,
					Score: 92.3,
				},
			},
		},
		{
			Id:   2,
			Name: "Bob",
			Grades: []Grade{
				{
					Title: "Math Exam",
					Type:  GradeExam,
					Score: 90.5,
				},
				{
					Title: "Physics Exam",
					Type:  GradeExam,
					Score: 82.0,
				},
				{
					Title: "English Test",
					Type:  GradeTest,
					Score: 88.7,
				},
			},
		},
	}
}
