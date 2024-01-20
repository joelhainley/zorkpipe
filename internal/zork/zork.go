package zork

import (
	"fmt"
	"math/rand"
)

var rooms []string
var weapons []string
var artifacts []string
var villains []string

func initializeRooms() {

	rooms = append(rooms, "Stone Barrow\n\nYou are standing in front of a massive barrow of stone. In the east face is a huge stone door which is open. You cannot see into the dark of the tomb.")
	rooms = append(rooms, "North Of House\n\nYou are facing the north side of a white house. There is no door here, and all the windows are boarded up. To the north a narrow path winds through the trees.")
	rooms = append(rooms, "South of House\n\nYou are facing the south side of a white house. There is no door here, and all the windows are boarded.")
	rooms = append(rooms, "Forest\n\nThis is a forest, with trees in all directions. To the east, there appears to be sunlight.")
	rooms = append(rooms, "Forest\n\nThis is a dimly lit forest, with large trees all around.")
	rooms = append(rooms, "Forest\n\nThe forest thins out, revealing impassable mountains.")
	rooms = append(rooms, "Forest Path\n\nThis is a path winding through a dimly lit forest. The path heads north-south here. One particularly large tree with some low branches stands at the edge of the path.")
	rooms = append(rooms, "Clearing\n\nYou are in a small clearing in a well marked forest path that\nextends to the east and west.")
	rooms = append(rooms, "Attic\n\nThis is the Attic. The only exit is stairs that lead down.")
	rooms = append(rooms, "Troll Room\n\nThis is a small room with passages to the east and south and a forbidding hole leading west. Bloodstains and deep scratches (perhaps made by an axe) mar the walls.")
	rooms = append(rooms, "East of Chasm\n\nYou are on the east edge of a chasm, the bottom of which cannot be\nseen. A narrow passage goes north, and the path you are on continues\nto the east.")
	rooms = append(rooms, "Gallery\n\nThis is an art gallery. Most of the paintings have been stolen by vandals with exceptional taste. The vandals left through either the north or west exits.")
	rooms = append(rooms, "Studio\n\nThis appears to have been an artist's studio. The walls and floors are splattered with paints of 69 different colors. Strangely enough, nothing of value is hanging here. At the south end of the room is an open door (also covered with paint). A dark and narrow chimney leads up from a\nfireplace; although you might be able to get up it, it seems unlikely\nyou could get back down.")
	rooms = append(rooms, "Maze\n\nThis is part of a maze of twisty little passages, all alike.")
	rooms = append(rooms, "Dead End\n\nYou have come to a dead end in the maze.")
	rooms = append(rooms, "Maze\n\nThis is part of a maze of twisty little passages, all alike. A skeleton, probably the remains of a luckless adventurer, lies here.")
	rooms = append(rooms, "Dead End\n\nYou have come to a dead end in the maze.")
	rooms = append(rooms, "Grating Room\n\nThe grating is closed.")
	rooms = append(rooms, "Cyclops Room\n\nA cyclops, who looks prepared to eat horses (much less mere adventurers), blocks the staircase. From his state of health, and the bloodstains on the walls, you gather that he is not very friendly, though he likes people.")
	rooms = append(rooms, "Strange Passage\n\nThis is a long passage. To the west is one entrance. On the\neast there is an old wooden door, with a large opening in it (about\ncyclops sized).")
	rooms = append(rooms, "Treasure Room\n\nThis is a large room, whose east wall is solid granite. A number\nof discarded bags, which crumble at your touch, are scattered about\non the floor. There is an exit down a staircase.")
	// add reservoir-south
	// add reservoir
	// add reservoir-north
	rooms = append(rooms, "Stream View\n\nYou are standing on a path beside a gently flowing stream. The path\nfollows the stream, which flows from west to east.")
	rooms = append(rooms, "In Stream\n\nYou are on the gently flowing stream. The upstream route is too narrow to navigate, and the downstream route is invisible due to twisting walls. There is a narrow beach to land on.")
	// add mirror room 1 / 2
	rooms = append(rooms, "Small Cave\n\nThis is a tiny cave with entrances west and north, and a staircase leading down.")
	rooms = append(rooms, "Tiny Cave\n\nThis is a tiny cave with entrances west and north, and a dark, forbidding staircase leading down.")
	rooms = append(rooms, "Cold Passage\n\nThis is a cold and damp corridor where a long east-west passageway turns into a southward path.")
	rooms = append(rooms, "Narrow Passage\n\nThis is a long and narrow corridor where a long north-south passageway briefly narrows even further.")
	rooms = append(rooms, "Winding Passage\n\nThis is a winding passage. It seems that there are only exits on the east and north.")
	rooms = append(rooms, "Twisting Passage\n\nThis is a winding passage. It seems that there are only exits on the east and north.")
	rooms = append(rooms, "Atlantis Room\n\nThis is an ancient room, long under water. There is an exit to the south and a staircase leading up.")

}

func GetDescription() string {
	initializeRooms()

	roomIndex := rand.Intn(len(rooms) - 1)
	return fmt.Sprintf("%s\n", rooms[roomIndex])
}

//
//func initializeOther() {
//	weapons = append(weapons, "On a table is a nasty-looking knife.")
//	weapons = append(weapons, "On hooks above the mantelpiece hangs an elvish sword of great antiquity.")
//
//	artifacts = append(artifacts, "Lying in one corner of the room is a beautifully carved crystal skull.\nIt appears to be grinning at you rather nastily.")
//	artifacts = append(artifacts, "An ornamented sceptre, tapering to a sharp point, is here.")
//	artifacts = append(artifacts, "A sceptre, possibly that of ancient Egypt itself, is in the coffin. The sceptre is ornamented with colored enamel, and tapers to a sharp point.")
//
//	villains = append(villains, "A cyclops, who looks prepared to eat horses (much less mere adventurers), blocks the staircase. From his state of health, and\nthe bloodstains on the walls, you gather that he is not very\nfriendly, though he likes people.")
//	villains = append(villains, "A \"lean and hungry\" gentleman just wandered through, carrying a large bag. Finding nothing of value, he left disgruntled.")
//	villains = append(villains, "There is a suspicious-looking individual, holding a large bag, leaning\nagainst one wall. He is armed with a deadly stiletto.")
//}
