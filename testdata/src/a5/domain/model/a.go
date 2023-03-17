package model

import (
	"a5/usecase" // want "error: found bug in dependency import"
)

type A struct{
	uc *usecase.A
}
