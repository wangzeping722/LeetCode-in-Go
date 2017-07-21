package main

const (
	// USER 是你在leetcode.com的用户名
	// 登录后，在页面右上角可以看到
	USER = "aQuaYi"
)

func main() {
	categories := []string{
		"algorithms",
		// "database",
		"draft",
		// "operating-system",
		// "shell",
		// "system-design",
	}

	solvedProblems := Solveds{}
	cs := Categories{}
	total := &Category{
		Name: "Total",
	}

	for i := range categories {
		c := getCategory(categories[i])

		solvedProblems = append(solvedProblems, c.run()...)

		total.update(c)
		cs = append(cs, c)
	}
	cs = append(cs, total)

	makeREADME(cs, solvedProblems)
}
