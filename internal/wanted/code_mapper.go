package wanted

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

	// HTML classes
	jobList          = "List_List_container__JnQMS"
	skillAddButton   = "FilterInWdlist_skillsFilterContainer__UZGLH"
	skillInput       = "SkillsSearch_SkillsSearch__At_ER"
	skillApplyButton = "explore__filter__update"
	workFromHomeTag  = "Work From Home"
	flexibleTag      = "Flexible Hours"
	flatTag          = "Company Culture"
	snackTag         = "Free Meals"
	militaryTag      = "Military exempt"
	nextButton       = "NavButton_NavButton__j73pc NavButton_isNext__m3fGk"
	card             = "Card_className__u5rsb"
	company          = "job-card-company-name"
	position         = "job-card-position"
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

func mapBenefitsToTags(benefits []string) []string {
	var tags []string
	for _, benefit := range benefits {
		switch benefit {
		case "재택 근무":
			tags = append(tags, workFromHomeTag)
		case "유연 근무":
			tags = append(tags, flexibleTag)
		case "수평적 조직":
			tags = append(tags, flatTag)
		case "간식":
			tags = append(tags, snackTag)
		}
	}
	return tags
}
