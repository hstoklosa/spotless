package main

import "github.com/hstoklosa/spotless/cmd"

/** Commands represent actions, Args are things and Flags are modifiers for those actions.
 *  APPNAME VERB NOUN --ADJECTIVE or APPNAME COMMAND ARG --FLAG.
 */

func main() {
	cmd.Execute()
}
