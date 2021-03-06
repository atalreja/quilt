package db

// An Image row represents a Docker image that should be built by the Quilt
// masters.
type Image struct {
	ID int

	// The desired name for the image.
	Name string

	// The Dockerfile with which to build the image.
	Dockerfile string

	// The ID of the built image.
	DockerID string
}

// InsertImage creates a new image row and inserts it into the database.
func (db Database) InsertImage() Image {
	result := Image{ID: db.nextID()}
	db.insert(result)
	return result
}

// SelectFromImage gets all images in the database that satisfy 'check'.
func (db Database) SelectFromImage(check func(Image) bool) []Image {
	imageTable := db.accessTable(ImageTable)
	result := []Image{}
	for _, row := range imageTable.rows {
		if check == nil || check(row.(Image)) {
			result = append(result, row.(Image))
		}
	}
	return result
}

func (image Image) getID() int {
	return image.ID
}

func (image Image) tt() TableType {
	return ImageTable
}

func (image Image) String() string {
	return defaultString(image)
}

func (image Image) less(r row) bool {
	return image.ID < r.(Image).ID
}

// ImageSlice is an alias for []Image to allow for joins
type ImageSlice []Image

// Get returns the value contained at the given index
func (slc ImageSlice) Get(ii int) interface{} {
	return slc[ii]
}

// Len returns the number of items in the slice.
func (slc ImageSlice) Len() int {
	return len(slc)
}
