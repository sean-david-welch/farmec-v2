package lib

type StatsItem struct {
	Link        string
	Title       string
	Icon        string
	Description string
}

type SpecialsItem struct {
	Link        string
	Title       string
	Icon        string
	Description string
}

func GetDisplayData() ([]StatsItem, []SpecialsItem) {
	specialsItems := []SpecialsItem{
		{
			Title:       "Quality Stock",
			Description: "Farmec is committed to the importation and distribution of only quality brands of unique farm machinery. We guarantee that all our suppliers are committed to providing farmers with durable and superior stock",
			Icon:        "tractor.svg",
			Link:        "/suppliers",
		},
		{
			Title:       "Assemably",
			Description: "Farmec have a team of qualified and experienced staff that ensure abundant care is taken during the assembly process; we make sure that a quality supply chain is maintained from manufacturer to customer",
			Icon:        "toolbox.svg",
			Link:        "/suppliers",
		},
		{
			Title:       "Spare Parts",
			Description: "Farmec offers a diverse and complete range of spare parts for all its machinery. Quality stock control and industry expertise ensures parts finds their way to you efficiently",
			Icon:        "gears.svg",
			Link:        "/spareparts",
		},
		{
			Title:       "Customer Service",
			Description: "Farmec is a family run company, we make sure we extend the ethos of a small community to our customers. We build established relationships with our dealers that provide them and the farmers with extensive guidance",
			Icon:        "user-plus.svg",
			Link:        "/contact",
		},
	}

	statsItems := []StatsItem{
		{
			Title:       "Large Network",
			Description: "50+ Dealers Nationwide",
			Icon:        "people.svg",
			Link:        "/suppliers",
		},
		{
			Title:       "Experience",
			Description: "25+ Years in Business",
			Icon:        "business-time.svg",
			Link:        "/about",
		},
		{
			Title:       "Diverse Range",
			Description: "10+ Quality Suppliers",
			Icon:        "handshake.svg",
			Link:        "/suppliers",
		},
		{
			Title:       "Committment",
			Description: "Warranty Guarentee",
			Icon:        "wrench.svg",
			Link:        "/spareparts",
		},
	}

	return statsItems, specialsItems
}
