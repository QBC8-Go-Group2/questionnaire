package app

import "github.com/QBC8-Go-Group2/questionnaire/config"

type App interface {
	Config() config.Config
}
