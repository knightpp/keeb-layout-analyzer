package layout

var ColemakDH34Keys = Flatten(
	Cluster(LeftSide, PinkyFinger,
		AddUppercase([]string{"rshift"}, Column(
			Key{Char: 'q'},
			Key{
				Char:   'a',
				IsHome: true,
			},
			Key{Char: 'x'},
		)),
		ActivationGroup([]string{"layer 1"},
			Key{Char: '1'},
		),
	),
	Cluster(LeftSide, RingFinger,
		AddUppercase([]string{"rshift"}, Column(
			Key{Char: 'w'},
			Key{
				ID:     "lshift",
				Char:   'r',
				IsHome: true,
			},
			Key{Char: 'c'},
		)),
		ActivationGroup([]string{"layer 1"},
			Key{
				ID:   "2",
				Char: '2',
			},
		),
	),
	Cluster(LeftSide, MiddleFinger,
		AddUppercase([]string{"rshift"}, Column(
			Key{Char: 'f'},
			Key{
				Char:   's',
				IsHome: true,
			},
			Key{Char: 'd'},
		)),
		ActivationGroup([]string{"layer 1"},
			Key{Char: '3'},
		),
	),
	Cluster(LeftSide, IndexFinger,
		AddUppercase([]string{"rshift"}, Column(
			Key{Char: 'p'},
			Key{
				Char:   't',
				IsHome: true,
			},
			Key{Char: 'v'},
		)),
		AddUppercase([]string{"rshift"}, Column(
			Key{
				Char: 'b',
				Pos:  Vec2{X: int(U1)},
			},
			Key{
				Char: 'g',
				Pos:  Vec2{X: int(U1)},
			},
			Key{
				Char: 'z',
				Pos:  Vec2{X: int(U1)},
			},
		)),
		ActivationGroup([]string{"layer 1"},
			Key{Char: '4'},
			Key{
				Char: '5',
				Pos:  Vec2{X: int(U1)},
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
		AddUppercase([]string{"lshift"}, Column(
			Key{
				Char: 'l',
				Pos: Vec2{
					X: int(U1),
				},
			},
			Key{
				Char:   'n',
				IsHome: true,
				Pos: Vec2{
					X: int(U1),
				},
			},
			Key{
				Char: 'h',
				Pos: Vec2{
					X: int(U1),
				},
			},
		)),
		AddUppercase([]string{"lshift"}, Column(
			Key{Char: 'j'},
			Key{Char: 'm'},
			Key{Char: 'k'},
		)),
		ActivationGroup([]string{"layer 1"},
			Key{
				Char: '7',
				Pos: Vec2{
					X: int(U1),
				},
			},
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
					X: int(U1),
					Y: int(U1) / 2,
				},
			},
		},
	),
	Cluster(RightSide, ThumbFinger,
		[]Key{
			{
				ID:     "right esc",
				IsHome: true,
			},
			{
				ID: "right tab",
				Pos: Vec2{
					X: int(U1),
					Y: int(U1) / 2,
				},
			},
		},
	),
)
