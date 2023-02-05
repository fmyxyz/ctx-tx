package test

import (
	"context"
	"errors"
	"log"

	tx "github.com/fmyxyz/ctx-tx"
)

/*
	 create table item
	 (

		id  int auto_increment primary key,

		qty int not null

	 );

	 insert into item (id, qty) values (88, 89);

	 insert into item (id, qty) values (99, 100);
*/
type Item struct {
	ID  int
	Qty int
}

// 测试时对此函数覆盖
var Update = func(ctx context.Context, id, num int) error {
	panic("Update:此函数未覆盖")
}

// 测试时对此函数覆盖
var Update88 = func(ctx context.Context) error {
	panic("Update88:此函数未覆盖")
}

// 测试时对此函数覆盖
var Update99 = func(ctx context.Context) error {
	panic("Update99:此函数未覆盖")
}

func rr(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}

		err = tx.WithTx(ctx, Update99)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println(err)
	}

	return err
}

func rrl(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update99(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	return err
}
func rrle(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}

		if err != nil {
			log.Println(err)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update99(ctx)
		if err != nil {
			return err
		}

		return errors.New("err99")
	})

	if err != nil {
		log.Println(err)
	}

	return err
}

func rre(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}

		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		})

		if err != nil {
			log.Println(err)
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	return err
}
func rse(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		}, tx.PropagationSupported())

		if err != nil {
			log.Println(err)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	return err
}
func rNestE(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		}, tx.PropagationNested())

		if err != nil {
			log.Println(err)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}

	return err
}
func rNewE(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		}, tx.PropagationNew())

		if err != nil {
			log.Println(err)
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	return err
}
func sre(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		})

		if err != nil {
			log.Println(err)
		}
		return nil
	}, tx.PropagationSupported())

	if err != nil {
		log.Println(err)
	}

	return err
}
func sse(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		}, tx.PropagationSupported())

		if err != nil {
			log.Println(err)
		}
		return nil
	}, tx.PropagationSupported())

	if err != nil {
		log.Println(err)
	}

	return err
}
func sNestE(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		}, tx.PropagationNested())

		if err != nil {
			log.Println(err)
		}
		return nil
	}, tx.PropagationSupported())

	if err != nil {
		log.Println(err)
	}

	return err
}
func sNewE(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		}, tx.PropagationNew())

		if err != nil {
			log.Println(err)
		}
		return nil
	}, tx.PropagationSupported())

	if err != nil {
		log.Println(err)
	}

	return err
}
func nestRE(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		})

		if err != nil {
			log.Println(err)
		}
		return nil
	}, tx.PropagationNested())

	if err != nil {
		log.Println(err)
	}

	return err
}
func nestSE(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		}, tx.PropagationSupported())

		if err != nil {
			log.Println(err)
		}
		return nil
	}, tx.PropagationNested())

	if err != nil {
		log.Println(err)
	}

	return err
}
func nestNestE(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		}, tx.PropagationNested())

		if err != nil {
			log.Println(err)
		}
		return nil
	}, tx.PropagationNested())

	if err != nil {
		log.Println(err)
	}

	return err
}
func nestNewE(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		}, tx.PropagationNew())

		if err != nil {
			log.Println(err)
		}
		return nil
	}, tx.PropagationNested())

	if err != nil {
		log.Println(err)
	}

	return err
}
func newRE(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		})

		if err != nil {
			log.Println(err)
		}
		return nil
	}, tx.PropagationNew())

	if err != nil {
		log.Println(err)
	}

	return err
}
func newSE(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		}, tx.PropagationSupported())

		if err != nil {
			log.Println(err)
		}
		return nil
	}, tx.PropagationNew())

	if err != nil {
		log.Println(err)
	}

	return err
}
func newNestE(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		}, tx.PropagationNested())

		if err != nil {
			log.Println(err)
		}
		return nil
	}, tx.PropagationNew())

	if err != nil {
		log.Println(err)
	}

	return err
}
func newNewE(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return errors.New("error99")

		}, tx.PropagationNew())

		if err != nil {
			log.Println(err)
		}
		return nil
	}, tx.PropagationNew())

	if err != nil {
		log.Println(err)
	}

	return err
}

func rer(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}

		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}
			return nil

		})

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")

	})
	if err != nil {
		log.Println(err)
	}

	return err
}
func res(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		}, tx.PropagationSupported())

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	})

	if err != nil {
		log.Println(err)
	}

	return err
}
func rENest(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		}, tx.PropagationNested())

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	})

	if err != nil {
		log.Println(err)
	}

	return err
}
func rENew(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		}, tx.PropagationNew())

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	})

	if err != nil {
		log.Println(err)
	}

	return err
}
func ser(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		})

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	}, tx.PropagationSupported())

	if err != nil {
		log.Println(err)
	}

	return err
}
func ses(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		}, tx.PropagationSupported())

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	}, tx.PropagationSupported())

	if err != nil {
		log.Println(err)
	}

	return err
}
func sENest(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		}, tx.PropagationNested())

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	}, tx.PropagationSupported())

	if err != nil {
		log.Println(err)
	}

	return err
}
func sENew(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		}, tx.PropagationNew())

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	}, tx.PropagationSupported())

	if err != nil {
		log.Println(err)
	}

	return err
}
func nestER(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		})

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	}, tx.PropagationNested())

	if err != nil {
		log.Println(err)
	}

	return err
}
func nestES(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		}, tx.PropagationSupported())

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	}, tx.PropagationNested())

	if err != nil {
		log.Println(err)
	}

	return err
}
func nestENest(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		}, tx.PropagationNested())

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	}, tx.PropagationNested())

	if err != nil {
		log.Println(err)
	}

	return err
}
func nestENew(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		}, tx.PropagationNew())

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	}, tx.PropagationNested())

	if err != nil {
		log.Println(err)
	}

	return err
}
func newER(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		})

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	}, tx.PropagationNew())

	if err != nil {
		log.Println(err)
	}

	return err
}
func newES(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		}, tx.PropagationSupported())

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	}, tx.PropagationNew())

	if err != nil {
		log.Println(err)
	}

	return err
}
func newENest(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		}, tx.PropagationNested())

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	}, tx.PropagationNew())

	if err != nil {
		log.Println(err)
	}

	return err
}
func newENew(ctx context.Context) (err error) {

	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update88(ctx)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {

			err := Update99(ctx)
			if err != nil {
				return err
			}

			return nil

		}, tx.PropagationNew())

		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	}, tx.PropagationNew())

	if err != nil {
		log.Println(err)
	}

	return err
}

func comprehensive(ctx context.Context) (err error) {
	err = tx.WithTx(ctx, func(ctx context.Context) error {
		err := Update(ctx, 88, 1)
		if err != nil {
			return err
		}
		err = tx.WithTx(ctx, func(ctx context.Context) error {
			err := Update(ctx, 88, 2)
			if err != nil {
				return err
			}
			err = tx.WithTx(ctx, func(ctx context.Context) error {
				err := Update(ctx, 88, 2)
				if err != nil {
					return err
				}
				err = tx.WithTx(ctx, func(ctx context.Context) error {
					err := Update(ctx, 88, 4)
					if err != nil {
						return err
					}
					err = tx.WithTx(ctx, func(ctx context.Context) error {
						err := Update(ctx, 88, 5)
						if err != nil {
							return err
						}
						err = tx.WithTx(ctx, func(ctx context.Context) error {
							err := Update(ctx, 88, 6)
							if err != nil {
								return err
							}
							return nil
						}, tx.PropagationSupported())
						return nil
					})
					return nil
				}, tx.PropagationNested())
				return nil
			})
			return nil
		}, tx.PropagationSupported())
		if err != nil {
			log.Println(err)
		}
		return errors.New("error88")
	}, tx.PropagationNew())

	if err != nil {
		log.Println(err)
	}

	return err
}

var Tests = []struct {
	Name string

	Fc func(ctx context.Context) (err error)

	Eq88    bool
	Eq99    bool
	WantErr bool
}{
	{
		Name:    "Required-Required",
		Fc:      rr,
		WantErr: false,
		Eq88:    true,
		Eq99:    true,
	},
	{
		Name:    "Required-Required-平级",
		Fc:      rrl,
		WantErr: false,
		Eq88:    true,
		Eq99:    true,
	},
	{
		Name:    "Required-Required-平级-err",
		Fc:      rrle,
		WantErr: true,
		Eq88:    true,
		Eq99:    false,
	},

	{
		Name:    "Required-Required-err",
		Fc:      rre,
		WantErr: false,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "Required-Supported-err",
		Fc:      rse,
		WantErr: false,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "Required-Nested-err",
		Fc:      rNestE,
		WantErr: false,
		Eq88:    true,
		Eq99:    false,
	},
	{
		Name:    "Required-New-err",
		Fc:      rNewE,
		WantErr: false,
		Eq88:    true,
		Eq99:    false,
	},
	{
		Name:    "Supported-Required-err",
		Fc:      sre,
		WantErr: false,
		Eq88:    true,
		Eq99:    false,
	},
	{
		Name:    "Supported-Supported-err",
		Fc:      sse,
		WantErr: false,
		Eq88:    true,
		Eq99:    true,
	},
	{
		Name:    "Supported-Nested-err",
		Fc:      sNestE,
		WantErr: false,
		Eq88:    true,
		Eq99:    false,
	},
	{
		Name:    "Supported-New-err",
		Fc:      sNewE,
		WantErr: false,
		Eq88:    true,
		Eq99:    false,
	},
	{
		Name:    "Nested-Required-err",
		Fc:      nestRE,
		WantErr: false,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "Nested-Supported-err",
		Fc:      nestSE,
		WantErr: false,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "Nested-Nested-err",
		Fc:      nestNestE,
		WantErr: false,
		Eq88:    true,
		Eq99:    false,
	},
	{
		Name:    "Nested-New-err",
		Fc:      nestNewE,
		WantErr: false,
		Eq88:    true,
		Eq99:    false,
	},
	{
		Name:    "New-Required-err",
		Fc:      newRE,
		WantErr: false,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "New-Supported-err",
		Fc:      newSE,
		WantErr: false,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "New-Nested-err",
		Fc:      newNestE,
		WantErr: false,
		Eq88:    true,
		Eq99:    false,
	},
	{
		Name:    "New-New-err",
		Fc:      newNewE,
		WantErr: false,
		Eq88:    true,
		Eq99:    false,
	},

	{
		Name:    "Required-err-Required",
		Fc:      rer,
		WantErr: true,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "Required-err-Supported",
		Fc:      res,
		WantErr: true,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "Required-err-Nested",
		Fc:      rENest,
		WantErr: true,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "Required-err-New",
		Fc:      rENew,
		WantErr: true,
		Eq88:    false,
		Eq99:    true,
	},
	{
		Name:    "Supported-err-Required",
		Fc:      ser,
		WantErr: true,
		Eq88:    true,
		Eq99:    true,
	},
	{
		Name:    "Supported-err-Supported",
		Fc:      ses,
		WantErr: true,
		Eq88:    true,
		Eq99:    true,
	},
	{
		Name:    "Supported-err-Nested",
		Fc:      sENest,
		WantErr: true,
		Eq88:    true,
		Eq99:    true,
	},
	{
		Name:    "Supported-err-New",
		Fc:      sENew,
		WantErr: true,
		Eq88:    true,
		Eq99:    true,
	},
	{
		Name:    "Nested-err-Required",
		Fc:      nestER,
		WantErr: true,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "Nested-err-Supported",
		Fc:      nestES,
		WantErr: true,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "Nested-err-Nested",
		Fc:      nestENest,
		WantErr: true,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "Nested-err-New",
		Fc:      nestENew,
		WantErr: true,
		Eq88:    false,
		Eq99:    true,
	},
	{
		Name:    "New-err-Required",
		Fc:      newER,
		WantErr: true,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "New-err-Supported",
		Fc:      newES,
		WantErr: true,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "New-err-Nested",
		Fc:      newENest,
		WantErr: true,
		Eq88:    false,
		Eq99:    false,
	},
	{
		Name:    "New-err-New",
		Fc:      newENew,
		WantErr: true,
		Eq88:    false,
		Eq99:    true,
	},
	{
		Name:    "comprehensive",
		Fc:      comprehensive,
		WantErr: true,
		Eq88:    false,
		Eq99:    false,
	},
}
