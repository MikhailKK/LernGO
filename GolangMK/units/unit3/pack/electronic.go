package lerngo

// интерфейсы

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
}

type SmartPhone interface {
	OS() string
}

// структуры

type ApplePhone struct {
	Brand string
	Model string
	Type  string
	OS    string
}

type AndroidPhone struct {
	Brand string
	Model string
	Type  string
	OS    string
}

type RadioPhone struct {
	Brand        string
	Model        string
	Type         string
	ButtonsCount int
}

// функции конструкторы для iOS
func (apple *ApplePhone) GetBrand() string {
	return apple.Brand
}

func (apple *ApplePhone) GetModel() string {
	return apple.Model
}

func (apple *ApplePhone) GetType() string {
	apple.Type = "smartPhone"
	return apple.Type
}

func (apple *ApplePhone) GetOS() string {
	return apple.OS
}

// функции конструкторы для Android
func (android *AndroidPhone) GetBrand() string {
	return android.Brand
}

func (android *AndroidPhone) GetModel() string {
	return android.Model
}

func (android *AndroidPhone) GetType() string {
	android.Type = "smartPhone"
	return android.Type
}

func (android *AndroidPhone) GetOS() string {
	return android.OS
}

// функции конструкторы для стационарного телефона
func (radio *RadioPhone) GetBrand() string {
	return radio.Brand
}

func (radio *RadioPhone) GetModel() string {
	return radio.Model
}

func (radio *RadioPhone) GetType() string {
	radio.Type = "station"
	return radio.Type
}

func (radio *RadioPhone) GetCountButtons() int {
	return radio.ButtonsCount
}
