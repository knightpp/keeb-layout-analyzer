package layout

var FerrisSweepColemakDH = Flatten(
	Cluster(LeftSide, PinkyFinger,
		Column(
			Key{
				ID:   "q",
				Char: 'q',
			},
			Key{
				ID:   "a",
				Char: 'a',
			},
			Key{
				ID:   "lshift|x",
				Char: 'x',
			},
		),
		ActivationGroup([]string{"rshift|/"},
			Column(
				Key{
					ID:   "Q",
					Char: 'Q',
				},
				Key{
					ID:   "A",
					Char: 'A',
				},
				Key{
					ID:   "X",
					Char: 'X',
				},
			)...,
		),
		ActivationGroup([]string{"layer 1"},
			Key{
				ID:   "1",
				Char: '1',
			},
		),
	),
	Cluster(LeftSide, RingFinger,
		Column(
			Key{
				ID:   "w",
				Char: 'w',
			},
			Key{
				ID:   "r",
				Char: 'r',
			},
			Key{
				ID:   "c",
				Char: 'c',
			},
		),
		ActivationGroup([]string{"rshift|/"},
			Column(
				Key{
					ID:   "W",
					Char: 'W',
				},
				Key{
					ID:   "R",
					Char: 'R',
				},
				Key{
					ID:   "C",
					Char: 'C',
				},
			)...,
		),
		ActivationGroup([]string{"layer 1"},
			Key{
				ID:   "2",
				Char: '2',
			},
		),
	),
	Cluster(LeftSide, MiddleFinger,
		Column(
			Key{
				ID:   "f",
				Char: 'f',
			},
			Key{
				ID:   "s",
				Char: 's',
			},
			Key{
				ID:   "d",
				Char: 'd',
			},
		),
		ActivationGroup([]string{"rshift|/"},
			Column(
				Key{
					ID:   "F",
					Char: 'F',
				},
				Key{
					ID:   "S",
					Char: 'S',
				},
				Key{
					ID:   "D",
					Char: 'D',
				},
			)...,
		),
		ActivationGroup([]string{"layer 1"},
			Key{
				ID:   "3",
				Char: '3',
			},
		),
	),
	Cluster(LeftSide, IndexFinger,
		Column(
			Key{
				ID:   "p",
				Char: 'p',
			},
			Key{
				ID:   "t",
				Char: 't',
			},
			Key{
				ID:   "v",
				Char: 'v',
			},
		),
		Column(
			Key{
				ID:   "b",
				Char: 'b',
				Pos: Vec2{
					X: int(U1),
				},
			},
			Key{
				ID:   "g",
				Char: 'g',
				Pos: Vec2{
					X: int(U1),
				},
			},
			Key{
				ID:   "z",
				Char: 'z',
				Pos: Vec2{
					X: int(U1),
				},
			},
		),
		ActivationGroup([]string{"rshift|/"},
			Flatten(
				Column(
					Key{
						ID:   "P",
						Char: 'P',
					},
					Key{
						ID:   "T",
						Char: 'T',
					},
					Key{
						ID:   "V",
						Char: 'V',
					},
				),
				Column(
					Key{
						ID:   "B",
						Char: 'B',
						Pos: Vec2{
							X: int(U1),
						},
					},
					Key{
						ID:   "G",
						Char: 'G',
						Pos: Vec2{
							X: int(U1),
						},
					},
					Key{
						ID:   "Z",
						Char: 'Z',
						Pos: Vec2{
							X: int(U1),
						},
					},
				),
			)...,
		),
		ActivationGroup([]string{"layer 1"},
			Key{
				ID:   "4",
				Char: '4',
			},
			Key{
				ID:   "5",
				Char: '5',
			},
		),
	),
	///////////// right hand
	Cluster(RightSide, PinkyFinger,
		Column(
			Key{
				ID:   ";",
				Char: ';',
			},
			Key{
				ID:   "o",
				Char: 'o',
			},
			Key{
				ID:   "rshift|/",
				Char: '/',
			},
		),
		ActivationGroup([]string{"lshift|x"},
			Column(
				Key{
					ID:   ":",
					Char: ':',
				},
				Key{
					ID:   "O",
					Char: 'O',
				},
				Key{
					ID:   "?",
					Char: '?',
				},
			)...,
		),
		ActivationGroup([]string{"layer 1"},
			Key{
				ID:   "0",
				Char: '0',
			},
			Key{
				ID:   "'",
				Char: '\'',
				Pos: Vec2{
					X: 0,
					Y: int(U1),
				},
			},
		),
	),
	Cluster(RightSide, RingFinger,
		Column(
			Key{
				ID:   "y",
				Char: 'y',
			},
			Key{
				ID:   "i",
				Char: 'i',
			},
			Key{
				ID:   ".",
				Char: '.',
			},
		),
		ActivationGroup([]string{"lshift|x"},
			Column(
				Key{
					ID:   "Y",
					Char: 'Y',
				},
				Key{
					ID:   "I",
					Char: 'I',
				},
				Key{
					ID:   ">",
					Char: '>',
				},
			)...,
		),
		ActivationGroup([]string{"layer 1"},
			Key{
				ID:   "9",
				Char: '9',
			},
		),
	),
	Cluster(RightSide, MiddleFinger,
		Column(
			Key{
				ID:   "u",
				Char: 'u',
			},
			Key{
				ID:   "e",
				Char: 'e',
			},
			Key{
				ID:   ",",
				Char: ',',
			},
		),
		ActivationGroup([]string{"lshift|x"},
			Column(
				Key{
					ID:   "U",
					Char: 'U',
				},
				Key{
					ID:   "E",
					Char: 'E',
				},
				Key{
					ID:   "<",
					Char: '<',
				},
			)...,
		),
		ActivationGroup([]string{"layer 1"},
			Key{
				ID:   "8",
				Char: '8',
			},
		),
	),
	Cluster(RightSide, IndexFinger,
		Column(
			Key{
				ID:   "l",
				Char: 'l',
			},
			Key{
				ID:   "n",
				Char: 'n',
			},
			Key{
				ID:   "h",
				Char: 'h',
			},
		),
		Column(
			Key{
				ID:   "j",
				Char: 'j',
				Pos: Vec2{
					X: int(U1),
				},
			},
			Key{
				ID:   "m",
				Char: 'm',
				Pos: Vec2{
					X: int(U1),
				},
			},
			Key{
				ID:   "k",
				Char: 'k',
				Pos: Vec2{
					X: int(U1),
				},
			},
		),
		ActivationGroup([]string{"lshift|x"},
			Flatten(
				Column(
					Key{
						ID:   "L",
						Char: 'L',
					},
					Key{
						ID:   "N",
						Char: 'N',
					},
					Key{
						ID:   "H",
						Char: 'H',
					},
				),
				Column(
					Key{
						ID:   "J",
						Char: 'J',
						Pos: Vec2{
							X: int(U1),
						},
					},
					Key{
						ID:   "M",
						Char: 'M',
						Pos: Vec2{
							X: int(U1),
						},
					},
					Key{
						ID:   "K",
						Char: 'K',
						Pos: Vec2{
							X: int(U1),
						},
					},
				),
			)...,
		),
		ActivationGroup([]string{"layer 1"},
			Key{
				ID:   "7",
				Char: '7',
			},
			Key{
				ID:   "6",
				Char: '6',
			},
		),
	),

	Cluster(LeftSide, ThumbFinger,
		[]Key{
			{
				ID:     "left space",
				Char:   ' ',
				Finger: ThumbFinger,
			},
			{
				ID: "layer 1",
				Pos: Vec2{
					X: 0,
					Y: 0,
				},
			},
		},
	),
	Cluster(RightSide, ThumbFinger,
		[]Key{
			{
				ID: "right space",
				Pos: Vec2{
					X: 0,
					Y: 0,
				},
			},
		},
	),
)
