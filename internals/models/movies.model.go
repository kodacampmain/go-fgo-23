package models

type Movie struct {
	title    string
	cast     []string
	genre    []string
	duration uint
	price    uint
	image    string
}

type Movies struct {
	Data      []Movie
	IsSuccess bool
	IsFailed  bool
}

// movie constructor
func NewMovie(title, image string, cast, genre []string, duration, price uint) Movie {
	return Movie{
		title:    title,
		image:    image,
		cast:     cast,
		genre:    genre,
		duration: duration,
		price:    price,
	}
}

// Method Getter untuk Movie
func (m *Movie) GetMovieTitle() string {
	return m.title
}
func (m *Movie) GetMoviePrice() uint {
	return m.price
}

// Method Setter untuk Movie
func (m *Movie) ChangePrice(newPrice uint) {
	m.price = newPrice
}
