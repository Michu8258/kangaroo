package commands

import (
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services"
)

type CommandContext struct {
	Settings          *models.Settings
	ServiceCollection *services.ServiceCollection
}
