package ru

type Image struct {
	URL              string `json:"url" xml:",innerxml"`
	ImageTypeID      `json:"imageTypeId" xml:"ImageTypeID,attr"`
	ImageReferenceID `json:"imageReferenceId" xml:"ImageReferenceID,attr"`
}

type ImageTypeID uint64
type ImageReferenceID uint64

type allImageTypes struct {
	MainImage    ImageType
	PropertyPlan ImageType
	Interior     ImageType
	Exterior     ImageType
}

type ImageType struct {
	ID          ImageTypeID
	Description string
}

func getAllImageTypes() allImageTypes {
	return allImageTypes{
		MainImage: ImageType{
			ID:          1,
			Description: "MainImage",
		},
		PropertyPlan: ImageType{
			ID:          2,
			Description: "PropertyPlan",
		},
		Interior: ImageType{
			ID:          3,
			Description: "Interior",
		},
		Exterior: ImageType{
			ID:          4,
			Description: "Exterior",
		},
	}
}
