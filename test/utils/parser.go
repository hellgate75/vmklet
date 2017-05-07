package utils

type Person struct {
	Name		string	`json:"Name" xml:"Name" yaml:"Name"`
	Surname	string	`json:"Surname" xml:"Surname" yaml:"Surname"`
	Age			int			`json:"Age" xml:"Age" yaml:"Age"`
}
