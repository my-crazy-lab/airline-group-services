package models

/*
ID: Unique identifier for the airport.
Name: Name of the airport.
Location: Geographic coordinates (latitude and longitude) of the airport.
Country: Country where the airport is located.
Capacity: Maximum capacity of the airport, indicating the number of passengers it can handle.
Runways: Number of runways available for aircraft operations.
Gates: Number of gates available for boarding and deboarding passengers.
Planes: List of planes currently stationed at the airport.
Flights: List of scheduled flights departing from or arriving at the airport.
*/
/*
ID: "airport123"
Name: "John F. Kennedy International Airport"
Location: {latitude: 40.6413, longitude: -73.7781}
Country: "United States"
Capacity: 75,000 passengers per day
Runways: 4
Gates: 128
Planes: ["plane456", "plane789"]
Flights: ["flight001", "flight002", "flight003"]
*/
type Airport struct {
	ID       string `bson:"_id,omitempty"`
	Name     string `bson:"name"`
	Location struct {
		Latitude  float64 `bson:"latitude"`
		Longitude float64 `bson:"longitude"`
	} `bson:"location"`
	Country  string   `bson:"country"`
	Capacity int      `bson:"capacity"`
	Runways  int      `bson:"runways"`
	Gates    int      `bson:"gates"`
	Planes   []string `bson:"planes"`
	Flights  []string `bson:"flights"`
}
