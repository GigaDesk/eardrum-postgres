package mockproduct

import "time"

var MultipleCategoryNodes = []MockCategory{
	{
		Id:                      3,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		Name:                    "Stationery and School Supplies",
		Description: "This category includes essential items that students use daily. Products include notebooks, pens, pencils, erasers, rulers, calculators, and textbooks. These are fundamental for academic activities",
	},
	{
		Id:                      4,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		Name:                    "Snacks and Beverages",
		Description: "This category covers food and drinks that students can consume during breaks or lunch. Items such as packaged snacks, juices, bottled water, and healthy snacks like fruits can be included",
	},
	{
		Id:                      5,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		Name:                    "School Uniforms and Attire",
		Description: "Many schools in Kenya have specific uniform requirements. This category includes the full school uniform, such as shirts, trousers, skirts, dresses, sweaters, and sports attire",
	},
	{
		Id:                      6,
		CreatedAt:               time.Now().UTC(),
		UpdatedAt:               time.Now().UTC(),
		Name:                    "Personal Care Products",
		Description: "These are items that help students maintain personal hygiene and grooming. Products can include soap, sanitizers, tissue paper, and for boarding schools, it could expand to include items like toothpaste, toothbrushes, and hair products",
	},
}


var UpdatedCategory = MockCategory {
	Id:                      4,
	CreatedAt:               time.Now().UTC(),
	UpdatedAt:               time.Now().UTC(),
	Name:                    "Drinks and Packaged Food",
	Description: "This category covers food and drinks that students can consume during breaks or lunch. Items such as packaged food products, canned drinks, bottled water, and healthy snacks like fruits can be included",
}