# zork pipe

I thought it would be fun/nostalgic to occasionally output random descriptions from Zork into a terminal, perhaps occurring if the user were to list the content of a directory. (much like the 'l' command in Zork)  The simplist approach seemed to me to support piping through a utility that would handle putting the "zorkish bits" at the end of a stream. Hence. `zorkpipe`

In the first iteration, I've simply put the room descriptions in place, though I thought at some point it might be fun to put some of the villain, artifact, and weapon descriptions in as well. Simply randomly throwing things in.

There's not much value to this, just a bit of a saturday morning goof.

-- 

I've been pulling the strings from this repo: https://github.com/historicalsource/zork1

## open items
* a few of the rooms have empty descriptions, I suspect they have so logic to generate the descriptions so sort that out and add them in.
* colorize the output from zorkpipe
* add villains, weapons, and artifacts
* add support for a config file at ~/.zorkpipe for customizing the output
