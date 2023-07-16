package internal

const (
	ExitCodeNoData       = 1
	ExitCodeInvalidUsage = 2
)

var (
	Platforms = []string{
		"wanted",
		"rallit",
	}

	Roles = []string{
		"백엔드",
		"프론트엔드",
		"SW 엔지니어",
		"안드로이드",
		"iOS",
		"크로스플랫폼 앱",
		"데이터 엔지니어",
		"데이터 사이언티스트",
		"머신러닝",
		"DBA",
		"DevOps",
		"시스템/네트워크",
		"QA/테스트",
		"기술지원",
		"보안",
		"블록체인",
		"HW/임베디드",
	}

	Experiences = []string{
		"신입",
		"1년차",
		"2년차",
		"3년차",
		"4년차",
		"5년차",
		"6년차",
		"7년차",
		"8년차",
		"9년차",
		"10년차 이상",
	}

	Stacks = []string{
		"Java",
		"Kotlin",
		"JavaScript",
		"TypeScript",
		"Go",
		"Python",
		"Spring Boot",
		"JPA",
		"Node.js",
		"Nest.js",
		"React.js",
		"Next.js",
		"Vue.js",
		"Angular",
		"React Native",
		"Flutter",
		"Swift",
		"AWS",
		"GCP",
		"Git",
		"GitHub",
		"Linux",
		"Docker",
		"Kubernetes",
		"Jenkins",
		"SQL",
		"MySQL",
		"PostgreSQL",
		"Oracle",
		"MongoDB",
		"Redis",
		"C/C++",
		"C#",
		".NET",
		"Rust",
	}

	// TODO: 원티드에서 더 다양한 조건으로 검색할 수 있을지 확인 필요 (랠릿이나 점핏에는 더욱 다양한 옵션들이 있음)
	Benefits = []string{
		"간식",
		"수평적 조직",
		"유연 근무",
		"재택 근무",
	}
)
