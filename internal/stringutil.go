package internal

func FindNextAndPrevPhotoID(slice []string, target string) (prevPhotoID, nextPhotoID string) {
	prevPhotoID, nextPhotoID = "", ""

	for i, v := range slice {
		if v == target {
			if i > 0 {
				prevPhotoID = slice[i-1]
			}
			if i < len(slice)-1 {
				nextPhotoID = slice[i+1]
			}
			return
		}
	}
	return
}
