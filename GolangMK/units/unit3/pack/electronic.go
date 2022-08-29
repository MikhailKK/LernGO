package electronic

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

type applePhone struct {
	Brand string
	Model string
	Type  string
	OS    string
}

type androidPhone struct {
	Brand string
	Model string
	Type  string
	OS    string
}

type radioPhone struct {
	Brand        string
	Model        string
	Type         string
	ButtonsCount int
}

// функции конструкторы для iOS
func (apple *applePhone) GetBrand() string {
	return apple.Brand
}

func (apple *applePhone) GetModel() string {
	return apple.Model
}

func (apple *applePhone) GetType() string {
	apple.Type = "smartPhone"
	return apple.Type
}

func (apple *applePhone) GetOS() string {
	return apple.OS
}

// функции конструкторы для Android
func (android *androidPhone) GetBrand() string {
	return android.Brand
}

func (android *androidPhone) GetModel() string {
	return android.Model
}

func (android *androidPhone) GetType() string {
	android.Type = "smartPhone"
	return android.Type
}

func (android *androidPhone) GetOS() string {
	return android.OS
}

// функции конструкторы для стационарного телефона
func (radio *radioPhone) GetBrand() string {
	return radio.Brand
}

func (radio *radioPhone) GetModel() string {
	return radio.Model
}

func (radio *radioPhone) GetType() string {
	radio.Type = "station"
	return radio.Type
}

func (radio *radioPhone) GetCountButtons() int {
	return radio.ButtonsCount
}
