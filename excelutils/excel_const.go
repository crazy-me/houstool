package excelutils

const (
	MAX_LINE = 1000

	TABLE_HEADER = `{
		"alignment":
		{
			"horizontal": "center",
			"ident": 1,
			"justify_last_line": true
		},
		"border": [
		{
			"type": "left",
			"color": "000000",
			"style": 1
		},
		{
			"type": "top",
			"color": "000000",
			"style": 1
		},
		{
			"type": "bottom",
			"color": "000000",
			"style": 1
		},
		{
			"type": "right",
			"color": "000000",
			"style": 1
		}]
	}`

	TABLE_DATA = `{
		"alignment":
		{
			"horizontal": "center",
			"ident": 1,
			"justify_last_line": true
		}
	}`
)
