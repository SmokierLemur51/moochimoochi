package models 

// data for rendering regular pages
type PageData struct {
	Page 	string
	Title	string
	CoilSelections []GutterCoil
}

type Product struct {

}

type GutterCoil struct {
	ID	int
	Size  string
	
}



type AdministrativeFeed struct {
	Page string
	Title string
}


