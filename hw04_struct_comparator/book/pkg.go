package book

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func (b Book) GetID() int {
	return b.id
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b Book) GetTitle() string {
	return b.title
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b Book) GetAuthor() string {
	return b.author
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b Book) GetYear() int {
	return b.year
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b Book) GetSize() int {
	return b.size
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b Book) GetRate() float32 {
	return b.rate
}

func (b *Book) SeRate(rate float32) {
	b.rate = rate
}
