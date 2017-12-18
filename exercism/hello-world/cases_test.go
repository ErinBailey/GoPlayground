package greeting

// Say hello to someone based on their name test cases

var addCases = []struct {
	description string
	in string
	expected string 
}{
	{
		"first test",
		"Barb",
		"Hello, Barb!",
	},
	{
		"second test",
		"Dan",
		"Hello, Dan!",
	},
	{
		"third test",
		"Steve",
		"Hello, Steve!",
	},
	{
		"fourth test",
		"",
		"Hello, World!",
	},
	{
		"fourth test",
		"Sir Charles IV",
		"Hello, Sir Charles IV!",
	},
}
