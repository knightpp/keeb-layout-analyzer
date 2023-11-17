package layout

var FerrisSweepColemakDH = Flatten(
	Cluster(LeftSide, PinkyFinger,
		Column(
			Key{Char: 'q'},
			Key{
				Char:   'a',
				IsHome: true,
			},
			Key{Char: 'x'},
		),
		ActivationGroup([]string{"rshift"},
			Column(
				Key{Char: 'Q'},
				Key{Char: 'A'},
				Key{Char: 'X'},
			)...,
		),
		ActivationGroup([]string{"layer 1"},
			Key{Char: '1'},
		),
	),
	Cluster(LeftSide, RingFinger,
		Column(
			Key{Char: 'w'},
			Key{
				ID:     "lshift",
				Char:   'r',
				IsHome: true,
			},
			Key{Char: 'c'},
		),
		ActivationGroup([]string{"rshift"},
			Column(
				Key{Char: 'W'},
				Key{Char: 'R'},
				Key{Char: 'C'},
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
			Key{Char: 'f'},
			Key{
				Char:   's',
				IsHome: true,
			},
			Key{Char: 'd'},
		),
		ActivationGroup([]string{"rshift"},
			Column(
				Key{Char: 'F'},
				Key{Char: 'S'},
				Key{Char: 'D'},
			)...,
		),
		ActivationGroup([]string{"layer 1"},
			Key{Char: '3'},
		),
	),
	Cluster(LeftSide, IndexFinger,
		Column(
			Key{Char: 'p'},
			Key{
				Char:   't',
				IsHome: true,
			},
			Key{Char: 'v'},
		),
		Column(
			Key{
				Char: 'b',
				Pos: Vec2{
					X: int(U1),
				},
			},
			Key{
				Char: 'g',
				Pos: Vec2{
					X: int(U1),
				},
			},
			Key{
				Char: 'z',
				Pos: Vec2{
					X: int(U1),
				},
			},
		),
		ActivationGroup([]string{"rshift"},
			Flatten(
				Column(
					Key{Char: 'P'},
					Key{Char: 'T'},
					Key{Char: 'V'},
				),
				Column(
					Key{
						Char: 'B',
						Pos: Vec2{
							X: int(U1),
						},
					},
					Key{
						Char: 'G',
						Pos: Vec2{
							X: int(U1),
						},
					},
					Key{
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
			Key{Char: ';'},
			Key{
				Char:   'o',
				IsHome: true,
			},
			Key{Char: '/'},
		),
		ActivationGroup([]string{"lshift"},
			Column(
				Key{Char: ':'},
				Key{Char: 'O'},
				Key{Char: '?'},
			)...,
		),
		ActivationGroup([]string{"layer 1"},
			Key{Char: '0'},
			Key{
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
			Key{Char: 'y'},
			Key{
				ID:     "rshift",
				Char:   'i',
				IsHome: true,
			},
			Key{Char: '.'},
		),
		ActivationGroup([]string{"lshift"},
			Column(
				Key{Char: 'Y'},
				Key{Char: 'I'},
				Key{Char: '>'},
			)...,
		),
		ActivationGroup([]string{"layer 1"},
			Key{Char: '9'},
		),
	),
	Cluster(RightSide, MiddleFinger,
		Column(
			Key{Char: 'u'},
			Key{
				Char:   'e',
				IsHome: true,
			},
			Key{Char: ','},
		),
		ActivationGroup([]string{"lshift"},
			Column(
				Key{Char: 'U'},
				Key{Char: 'E'},
				Key{Char: '<'},
			)...,
		),
		ActivationGroup([]string{"layer 1"},
			Key{Char: '8'},
		),
	),
	Cluster(RightSide, IndexFinger,
		Column(
			Key{Char: 'l'},
			Key{
				Char:   'n',
				IsHome: true,
			},
			Key{Char: 'h'},
		),
		Column(
			Key{
				Char: 'j',
				Pos: Vec2{
					X: int(U1),
				},
			},
			Key{
				Char: 'm',
				Pos: Vec2{
					X: int(U1),
				},
			},
			Key{
				Char: 'k',
				Pos: Vec2{
					X: int(U1),
				},
			},
		),
		ActivationGroup([]string{"lshift"},
			Flatten(
				Column(
					Key{Char: 'L'},
					Key{Char: 'N'},
					Key{Char: 'H'},
				),
				Column(
					Key{
						Char: 'J',
						Pos: Vec2{
							X: int(U1),
						},
					},
					Key{
						Char: 'M',
						Pos: Vec2{
							X: int(U1),
						},
					},
					Key{
						Char: 'K',
						Pos: Vec2{
							X: int(U1),
						},
					},
				),
			)...,
		),
		ActivationGroup([]string{"layer 1"},
			Key{Char: '7'},
			Key{Char: '6'},
		),
	),

	Cluster(LeftSide, ThumbFinger,
		[]Key{
			{
				ID:     "left space",
				Char:   ' ',
				Finger: ThumbFinger,
				IsHome: true,
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
				IsHome: true,
			},
		},
	),
)
