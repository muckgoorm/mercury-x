package wanted

import "strings"

const (
	// Roles
	dev        = "518"
	be         = "/872"
	fe         = "/669"
	sw         = "/10110"
	android    = "/677"
	ios        = "/678"
	xplatform  = "/10111"
	devOps     = "/674"
	de         = "/655"
	ds         = "/1024"
	ml         = "/1634"
	dba        = "/10231"
	infra      = "/665"
	qa         = "/676"
	support    = "/1026"
	security   = "/671"
	blockchain = "/1027"
	hw         = "/672"

	// Stacks
	java     = "1540"
	kotlin   = "1544"
	js       = "1541"
	ts       = "1564"
	go1      = "1702"
	go2      = "4335"
	python   = "1554"
	spring   = "10169"
	jpa      = "3451"
	node1    = "1547"
	node2    = "2080"
	nest     = "10288"
	react1   = "1469"
	react2   = "9551"
	next     = "10286"
	vue1     = "1478"
	vue2     = "10342"
	angular1 = "1439"
	angular2 = "10352"
	rn       = "10168"
	flutter  = "10176"
	swift    = "1563"
	aws      = "1698"
	gcp      = "3468"
	git      = "1411"
	github   = "1412"
	linux    = "1459"
	docker   = "2217"
	kube     = "10268"
	sql      = "1562"
	mysql    = "1464"
	psql     = "2683"
	oracle   = "1465"
	mongo    = "1462"
	redis    = "1470"
	c        = "1663"
	cpp      = "1786"
	csharp   = "1533"
	dotnet   = "1445"
	rust     = "1557"

	// HTML classes
	card     = "Card_className__u5rsb"
	company  = "job-card-company-name"
	position = "job-card-position"
)

func mapRole(role string) string {
	switch role {
	case "백엔드":
		return be
	case "프론트엔드":
		return fe
	case "SW 엔지니어":
		return sw
	case "안드로이드":
		return android
	case "iOS":
		return ios
	case "크로스플랫폼 앱":
		return xplatform
	case "DevOps":
		return devOps
	case "데이터 엔지니어":
		return de
	case "데이터 사이언티스트":
		return ds
	case "머신러닝":
		return ml
	case "DBA":
		return dba
	case "시스템/네트워크":
		return infra
	case "QA/테스트":
		return qa
	case "기술지원":
		return support
	case "보안":
		return security
	case "블록체인":
		return blockchain
	case "HW/임베디드":
		return hw
	default:
		return ""
	}
}

func mapExperience(exp string) string {
	switch exp {
	case "신입":
		return "0"
	case "1년차":
		return "1"
	case "2년차":
		return "2"
	case "3년차":
		return "3"
	case "4년차":
		return "4"
	case "5년차":
		return "5"
	case "6년차":
		return "6"
	case "7년차":
		return "7"
	case "8년차":
		return "8"
	case "9년차":
		return "9"
	case "10년 이상":
		return "10"
	default:
		return "-1"
	}
}

func mapStacks(stacks []string) string {
	var qb strings.Builder
	for _, stack := range stacks {
		switch stack {
		case "Java":
			qb.WriteString("&skill_tags=" + java)
		case "Kotlin":
			qb.WriteString("&skill_tags=" + kotlin)
		case "JavaScript":
			qb.WriteString("&skill_tags=" + js)
		case "TypeScript":
			qb.WriteString("&skill_tags=" + ts)
		case "Go":
			qb.WriteString("&skill_tags=" + go1)
			qb.WriteString("&skill_tags=" + go2)
		case "Python":
			qb.WriteString("&skill_tags=" + python)
		case "Spring Boot":
			qb.WriteString("&skill_tags=" + spring)
		case "JPA":
			qb.WriteString("&skill_tags=" + jpa)
		case "Node.js":
			qb.WriteString("&skill_tags=" + node1)
			qb.WriteString("&skill_tags=" + node2)
		case "NestJS":
			qb.WriteString("&skill_tags=" + nest)
		case "React":
			qb.WriteString("&skill_tags=" + react1)
			qb.WriteString("&skill_tags=" + react2)
		case "Next.js":
			qb.WriteString("&skill_tags=" + next)
		case "Vue.js":
			qb.WriteString("&skill_tags=" + vue1)
			qb.WriteString("&skill_tags=" + vue2)
		case "Angular":
			qb.WriteString("&skill_tags=" + angular1)
			qb.WriteString("&skill_tags=" + angular2)
		case "React Native":
			qb.WriteString("&skill_tags=" + rn)
		case "Flutter":
			qb.WriteString("&skill_tags=" + flutter)
		case "Swift":
			qb.WriteString("&skill_tags=" + swift)
		case "AWS":
			qb.WriteString("&skill_tags=" + aws)
		case "GCP":
			qb.WriteString("&skill_tags=" + gcp)
		case "Git":
			qb.WriteString("&skill_tags=" + git)
		case "Github":
			qb.WriteString("&skill_tags=" + github)
		case "Linux":
			qb.WriteString("&skill_tags=" + linux)
		case "Docker":
			qb.WriteString("&skill_tags=" + docker)
		case "Kubernetes":
			qb.WriteString("&skill_tags=" + kube)
		case "SQL":
			qb.WriteString("&skill_tags=" + sql)
		case "MySQL":
			qb.WriteString("&skill_tags=" + mysql)
		case "PostgreSQL":
			qb.WriteString("&skill_tags=" + psql)
		case "Oracle":
			qb.WriteString("&skill_tags=" + oracle)
		case "MongoDB":
			qb.WriteString("&skill_tags=" + mongo)
		case "Redis":
			qb.WriteString("&skill_tags=" + redis)
		case "C":
			qb.WriteString("&skill_tags=" + c)
		case "C++":
			qb.WriteString("&skill_tags=" + cpp)
		case "C#":
			qb.WriteString("&skill_tags=" + csharp)
		case ".NET":
			qb.WriteString("&skill_tags=" + dotnet)
		case "Rust":
			qb.WriteString("&skill_tags=" + rust)
		}
	}

	return qb.String()
}
