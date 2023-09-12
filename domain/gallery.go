package domain

// PhotoService : represent the gallery's services
type GalleryService interface {
	FindPhotoByID(uint) (Photo, error)
	FindAllPhoto() ([]Photo, error)
	FindNextAndPrevPhotoID(string) (string, string, error)
}
