package layout

var Ukrainian34 = Flatten(
	Cluster(LeftSide, PinkyFinger,
		AddUppercase([]string{"rshift"}, Column(
			Key{Char: 'й'},
			Key{
				Char:   'ф',
				IsHome: true,
			},
			Key{Char: 'я'},
		)),
		ActivationGroup([]string{"ukr"},
			Column(
				Key{},
				Key{},
				Key{},
			)...,
		),
		ActivationGroup([]string{"l1"},
			Key{Char: '1'},
		),
		ActivationGroup([]string{"l1", "rb"},
			Key{Char: '!'},
		),
	),
	Cluster(LeftSide, RingFinger,
		AddUppercase([]string{"rshift"}, Column(
			Key{Char: 'ц'},
			Key{
				ID:     "lshift",
				Char:   'і',
				IsHome: true,
			},
			Key{Char: 'ч'},
		)),
		ActivationGroup([]string{"ukr"},
			Column(
				Key{},
				Key{Char: 'ї'},
				Key{Char: 'Ї'},
			)...,
		),
		ActivationGroup([]string{"l1"},
			Key{Char: '2'},
		),
		ActivationGroup([]string{"l1", "rb"},
			Key{Char: '"'},
		),
	),
	Cluster(LeftSide, MiddleFinger,
		AddUppercase([]string{"rshift"}, Column(
			Key{Char: 'у'},
			Key{
				Char:   'в',
				IsHome: true,
			},
			Key{Char: 'с'},
		)),
		ActivationGroup([]string{"ukr"},
			Column(
				Key{},
				Key{},
				Key{},
			)...,
		),
		ActivationGroup([]string{"l1"},
			Key{Char: '3'},
		),
		ActivationGroup([]string{"l1", "rb"},
			Key{Char: '№'},
		),
	),
	Cluster(LeftSide, IndexFinger,
		AddUppercase([]string{"rshift"}, Column(
			Key{Char: 'к'},
			Key{
				Char:   'а',
				IsHome: true,
			},
			Key{Char: 'м'},
		)),
		AddUppercase([]string{"rshift"}, Column(
			Key{
				Char: 'е',
				Pos:  Vec2{X: int(U1)},
			},
			Key{
				Char: 'п',
				Pos:  Vec2{X: int(U1)},
			},
			Key{
				Char: 'и',
				Pos:  Vec2{X: int(U1)},
			},
		)),
		ActivationGroup([]string{"ukr"},
			Column(
				Key{Char: 'х'},
				Key{Char: 'Х'},
				Key{},
			)...,
		),
		ActivationGroup([]string{"ukr"},
			Column(
				Key{
					Char: 'є',
					Pos:  Vec2{X: int(U1)},
				},
				Key{
					Char: 'Є',
					Pos:  Vec2{X: int(U1)},
				},
				Key{
					Pos: Vec2{X: int(U1)},
				},
			)...,
		),
		ActivationGroup([]string{"l1"},
			Key{Char: '4'},
			Key{
				Char: '5',
				Pos:  Vec2{X: int(U1)},
			},
		),
		ActivationGroup([]string{"l1", "rb"},
			Key{Char: ';'},
			Key{
				Char: '%',
				Pos:  Vec2{X: int(U1)},
			},
		),
	),
	///////////// right hand
	Cluster(RightSide, PinkyFinger,
		AddUppercase([]string{"lshift"}, Column(
			Key{Char: 'з'},
			Key{
				Char:   'ж',
				IsHome: true,
			},
			Key{Char: '.'},
		)),
		ActivationGroup([]string{"lshift"},
			Key{
				Char: ',',
				Pos: Vec2{
					Y: 3 * int(U1),
				},
			},
		),
		ActivationGroup([]string{"ukr"},
			Column(
				Key{},
				Key{Char: '\''},
				Key{},
			)...,
		),
		ActivationGroup([]string{"l1"},
			Key{Char: '0'},
			Key{
				Char: '\'',
				Pos: Vec2{
					X: 0,
					Y: int(U1),
				},
			},
		),
		ActivationGroup([]string{"l1", "rb"},
			Key{Char: ')'},
		),
	),
	Cluster(RightSide, RingFinger,
		AddUppercase([]string{"lshift"}, Column(
			Key{Char: 'щ'},
			Key{
				ID:     "rshift",
				Char:   'д',
				IsHome: true,
			},
			Key{Char: 'ю'},
		)),
		ActivationGroup([]string{"ukr"},
			Column(
				Key{},
				Key{},
				Key{},
			)...,
		),
		ActivationGroup([]string{"l1"},
			Key{Char: '9'},
		),
		ActivationGroup([]string{"l1", "rb"},
			Key{Char: '('},
		),
	),
	Cluster(RightSide, MiddleFinger,
		AddUppercase([]string{"lshift"}, Column(
			Key{Char: 'ш'},
			Key{
				Char:   'л',
				IsHome: true,
			},
			Key{Char: 'б'},
		)),
		ActivationGroup([]string{"ukr"},
			Column(
				Key{},
				Key{},
				Key{},
			)...,
		),
		ActivationGroup([]string{"l1"},
			Key{Char: '8'},
		),
		ActivationGroup([]string{"l1", "rb"},
			Key{Char: '*'},
		),
	),
	Cluster(RightSide, IndexFinger,
		AddUppercase([]string{"lshift"}, Column(
			Key{
				Char: 'г',
				Pos: Vec2{
					X: int(U1),
				},
			},
			Key{
				Char:   'о',
				IsHome: true,
				Pos: Vec2{
					X: int(U1),
				},
			},
			Key{
				Char: 'ь',
				Pos: Vec2{
					X: int(U1),
				},
			},
		)),
		AddUppercase([]string{"lshift"}, Column(
			Key{Char: 'н'},
			Key{Char: 'р'},
			Key{Char: 'т'},
		)),
		ActivationGroup([]string{"rb"},
			Key{
				Char: '-',
				Pos: Vec2{
					Y: int(U1),
				},
			},
		),
		ActivationGroup([]string{"ukr"},
			Column(
				Key{},
				Key{},
				Key{},
			)...,
		),
		ActivationGroup([]string{"ukr"},
			Column(
				Key{
					Char: 'ґ',
					Pos:  Vec2{X: int(U1)},
				},
				Key{
					Char: 'Ґ',
					Pos:  Vec2{X: int(U1)},
				},
				Key{
					Pos: Vec2{X: int(U1)},
				},
			)...,
		),
		ActivationGroup([]string{"l1"},
			Key{
				Char: '7',
				Pos: Vec2{
					X: int(U1),
				},
			},
			Key{Char: '6'},
		),
		ActivationGroup([]string{"l1", "rb"},
			Key{
				Char: '?',
				Pos: Vec2{
					X: int(U1),
				},
			},
			Key{Char: ':'},
		),
	),

	Cluster(LeftSide, ThumbFinger,
		[]Key{
			{
				ID:     "lb",
				Char:   ' ',
				Finger: ThumbFinger,
				IsHome: true,
			},
			{
				ID: "l1",
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
				ID:     "rb",
				IsHome: true,
			},
			{
				ID: "ukr",
			},
			{
				ID: "rs",
				Pos: Vec2{
					X: int(U1),
					Y: int(U1) / 2,
				},
			},
		},
	),
)
