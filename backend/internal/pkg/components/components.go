package components

import "decert/internal/pkg/db/mysql"

func Init() error {
	if err := mysql.Init(); err != nil {
		return err
	}

	return nil
}
