package constString

import "strings"

func StripSpecialCharacters(input string) string {
	output := strings.ReplaceAll(input, "-", "")
	output = strings.ReplaceAll(output, "_", "")
	return output
}

func GetLiveStoragePath(appID, roomID, envType string) string {
	appIDProcess := StripSpecialCharacters(appID)
	roomIDProcess := StripSpecialCharacters(roomID)
	return "live/v5live/" + envType + "/" + appIDProcess + "/" + roomIDProcess + "/"
}
