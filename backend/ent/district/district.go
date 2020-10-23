// Code generated by entc, DO NOT EDIT.

package district

const (
	// Label holds the string label denoting the district type in the database.
	Label = "district"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDistrict holds the string denoting the district field in the database.
	FieldDistrict = "district"

	// EdgeProvince holds the string denoting the province edge name in mutations.
	EdgeProvince = "province"

	// Table holds the table name of the district in the database.
	Table = "districts"
	// ProvinceTable is the table the holds the province relation/edge.
	ProvinceTable = "provinces"
	// ProvinceInverseTable is the table name for the Province entity.
	// It exists in this package in order to avoid circular dependency with the "province" package.
	ProvinceInverseTable = "provinces"
	// ProvinceColumn is the table column denoting the province relation/edge.
	ProvinceColumn = "district"
)

// Columns holds all SQL columns for district fields.
var Columns = []string{
	FieldID,
	FieldDistrict,
}

var (
	// DistrictValidator is a validator for the "district" field. It is called by the builders before save.
	DistrictValidator func(string) error
)