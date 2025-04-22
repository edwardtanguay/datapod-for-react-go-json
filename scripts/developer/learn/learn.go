package learn

import (
	"datapod-for-react-go-json/qtools/qcli"
	"datapod-for-react-go-json/qtools/qdev"
)

func Ex001() {
	qcli.Message("Working with map[string][string]", "star")

	frameworks := map[string]string{
		"svelte":  "Svelte",
		"angular": "Angular",
		"solid":   "SolidJS",
	}
	frameworks["react"] = "React"
	frameworks["nextjs"] = "Next.js"
	frameworks["vue"] = "Vue.js"

	qdev.DisplayStringStringMap(frameworks)

	delete(frameworks, "solid")

	qdev.DisplayStringStringMap(frameworks)
}

func Ex002() {
	qcli.Message("This is exercise 002.", "success")
}
