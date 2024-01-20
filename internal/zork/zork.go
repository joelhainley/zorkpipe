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

	rooms = append(rooms, "East West Passage\n\nThis is a narrow east-west passageway. There is a narrow stairway leading down at the north end of the room.")
	rooms = append(rooms, "Round Room\n\nThis is a circular stone room with passages in all directions. Several of them have unfortunately been blocked by cave-ins.")
	// add deep canyon
	rooms = append(rooms, "Damp Cave\n\nThis cave has exits to the west and east, and narrows to a crack toward the south. The earth is particularly damp here.")
	// add loud room
	rooms = append(rooms, "North-South Passage\n\nThis is a high north-south passage, which forks to the northeast.")
	rooms = append(rooms, "Chasm Room\n\nA chasm runs southwest to northeast and the path follows it. You are on the south side of the chasm, where a crack opens into a passage.")

	rooms = append(rooms, "Entrance To Hades\n\n")
	rooms = append(rooms, "Land of the Dead\n\nYou have entered the Land of the Living Dead. Thousands of lost souls can be heard weeping and moaning. In the corner are stacked the remains of dozens of previous adventurers less fortunate than yourself. A passage exits to the north.")
	rooms = append(rooms, "Engravings Cave\n\nYou have entered a low cave with passages leading northwest and east.")
	rooms = append(rooms, "Egyptian Room\n\nThis is a room which looks like an Egyptian tomb. There is an ascending staircase to the west.")
	rooms = append(rooms, "Dome Room\n\n")
	rooms = append(rooms, "Torch Room\n\n")
	rooms = append(rooms, "North Temple\n\nThis is the north end of a large temple. On the east wall is an ancient inscription, probably a prayer in a long-forgotten language. Below the prayer is a staircase leading down. The west wall is solid granite. The exit to the north end of the room is through huge marble pillars.")
	rooms = append(rooms, "South Temple\n\nThis is the south end of a large temple. In front of you is what appears to be an altar. In one corner is a small hole in the floor which leads into darkness. You probably could not get back up it.")
	rooms = append(rooms, "Dam\n\n")
	rooms = append(rooms, "Dam Lobby\n\nThis room appears to have been the waiting room for groups touring the dam.There are open doorways here to the north and east marked\"Private\", and there is a path leading south over the top of the dam.")
	rooms = append(rooms, "Maintenance Room\n\nThis is what appears to have been the maintenance room for Flood Control Dam #3. Apparently, this room has been ransacked recently, for most of the valuable equipment is gone. On the wall in front of you is a group of buttons colored blue, yellow, brown, and red. There are doorways to the west and south.")
	rooms = append(rooms, "Dam Base\n\nYou are at the base of Flood Control Dam #3, which looms above you and to the north. The Frigid River is flowing by here. Along the river are the White Cliffs which seem to form giant walls stretching from north to south along the shores of the river as it winds its way downstream.")
	rooms = append(rooms, "Frigid River\n\nYou are on the Frigid River in the vicinity of the Dam. The river flows quietly here. There is a landing on the west shore.")
	rooms = append(rooms, "Frigid River\n\nThe river turns a corner here making it impossible to see the Dam. The White Cliffs loom on the east bank and large rocks prevent landing on the west.")
	rooms = append(rooms, "Frigid River\n\nThe river descends here into a valley. There is a narrow beach on the west shore below the cliffs. In the distance a faint rumbling can be heard.\"")
	rooms = append(rooms, "White Cliffs Beach\n\nYou are on a narrow strip of beach which runs along the base of the White Cliffs. There is a narrow path heading south along the Cliffs and a tight passage leading west into the cliffs themselves.")
	rooms = append(rooms, "White Cliffs Beach\n\nYou are on a rocky, narrow strip of beach beside the Cliffs. A narrow path leads north along the shore.")
	rooms = append(rooms, "Frigid River\n\nThe river is running faster here and the sound ahead appears to be that of rushing water. On the east shore is a sandy beach. A small area of beach can also be seen below the cliffs on the west shore.")
	rooms = append(rooms, "Frigid River\n\nThe river is running faster here and the sound ahead appears to be that of rushing water. On the east shore is a sandy beach. A small area of beach can also be seen below the cliffs on the west shore.")
	rooms = append(rooms, "Frigid River\n\nThe sound of rushing water is nearly unbearable here. On the east shore is a large landing area.")
	rooms = append(rooms, "Shore\n\nYou are on the east shore of the river. The water here seems somewhat treacherous. A path travels from north to south here, the south end quickly turning around a sharp corner.")
	rooms = append(rooms, "Sandy Beach\n\nYou are on a large sandy beach on the east shore of the river, which is flowing quickly by. A path runs beside the river to the south here, and a passage is partially buried in sand to the northeast.")
	rooms = append(rooms, "Sandy Cave\n\nThis is a sand-filled cave whose exit is to the southwest.")
	rooms = append(rooms, "Aragain Falls\n\n")
	rooms = append(rooms, "On the Rainbow\n\nYou are on top of a rainbow (I bet you never thought you would walk on a rainbow), with a magnificent view of the Falls. The rainbow travels east-west here.")
	rooms = append(rooms, "End of Rainbow\n\nYou are on a small, rocky beach on the continuation of the Frigid River past the Falls. The beach is narrow due to the presence of the White Cliffs. The river canyon opens here and sunlight shines in from above. A rainbow crosses over the falls to the east and a narrow path continues to the southwest.")
	rooms = append(rooms, "Canyon Bottom\n\nYou are beneath the walls of the river canyon which may be climbable here. The lesser part of the runoff of Aragain Falls flows by below. To the north is a narrow path.")
	rooms = append(rooms, "Rocky Ledge\n\nYou are on a ledge about halfway up the wall of the river canyon. You can see from here that the main flow from Aragain Falls twists along a passage which it is impossible for you to enter. Below you is the canyon bottom. Above you is more cliff, which appears climbable.")
	rooms = append(rooms, "Canyon View\n\nYou are at the top of the Great Canyon on its west wall. From here there is a marvelous view of the canyon and parts of the Frigid River upstream. Across the canyon, the walls of the White Cliffs join the mighty ramparts of the Flathead Mountains to the east. Following the Canyon upstream to the north, Aragain Falls may be seen, complete with rainbow. The mighty Frigid River flows out from a great dark cavern. To the west and south can be seen an immense forest, stretching for miles around. A path leads northwest. It is possible to climb down into the canyon from here.")

	rooms = append(rooms, "Mine Entrance\n\nYou are standing at the entrance of what might have been a coal mine. The shaft enters the west wall, and there is another exit on the south end of the room.")
	rooms = append(rooms, "Squeaky Room\n\nYou are in a small room. Strange squeaky sounds may be heard coming from the passage at the north end. You may also escape to the east.")
	rooms = append(rooms, "Bat Room\n\n")
	rooms = append(rooms, "Shaft Room\n\nThis is a large room, in the middle of which is a small shaft descending through the floor into darkness below. To the west and the north are exits from this room. Constructed over the top of the shaft is a metal framework to which a heavy iron chain is attached.")
	rooms = append(rooms, "Smelly Room\n\nThis is a small nondescript room. However, from the direction of a small descending staircase a foul odor can be detected. To the south is a narrow tunnel.")
	rooms = append(rooms, "Gas Room\n\nThis is a small room which smells strongly of coal gas. There is a short climb up some stairs and a narrow tunnel leading east.")
	rooms = append(rooms, "Ladder Top\n\nThis is a very small room. In the corner is a rickety wooden ladder, leading downward. It might be safe to descend. There is also a staircase leading upward.")
	rooms = append(rooms, "Ladder Bottom\n\nThis is a rather wide room. On one side is the bottom of a narrow wooden ladder. To the west and the south are passages leaving the room.")
	rooms = append(rooms, "Dead End\n\nYou have come to a dead end in the mine.")
	rooms = append(rooms, "Timber Room\n\nThis is a long and narrow passage, which is cluttered with broken timbers. A wide passage comes from the east and turns at the west end of the room into a very narrow passageway. From the west comes a strong draft.")
	rooms = append(rooms, "Drafty Room\n\nThis is a small drafty room in which is the bottom of a long shaft. To the south is a passageway and to the east a very narrow passage. In the shaft can be seen a heavy iron chain.")
	rooms = append(rooms, "Drafty Room\n\nThis is a small drafty room in which is the bottom of a long shaft. To the south is a passageway and to the east a very narrow passage. In the shaft can be seen a heavy iron chain.")
	rooms = append(rooms, "Machine Room\n\n")
	rooms = append(rooms, "Mine\n\nThis is a nondescript part of a coal mine.")
	rooms = append(rooms, "Slide Room\n\nThis is a small chamber, which appears to have been part of a coal mine. On the south wall of the chamber the letters \"Granite Wall\" are etched in the rock. To the east is a long passage, and there is a steep metal slide twisting downward. To the north is a small opening.")
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
