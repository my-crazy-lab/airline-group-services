Imagine: 
- I'm a founder of the airlines
- I have 5 planes (P1, ...P5), 1 airport (X) with 3 ports (X1, X2, X3)
- My planes can go to A, B, C (round trip)

### Airport Management Service
- Manage planes, ports

### Flight Booking Management Service
- Booking processes. (8)
- Reservation processes. (7)
- Seat allocation.
- Manages flight schedules, availability and routes
- Payments for bookings and ticket purchases.

### Check-in Service
- Passenger check-in, baggage handling, and boarding passes.

### Baggage Management Service
- Tracks the status and location of passengers' baggage. (5)
- Tracks baggage handling, baggage allowances, and baggage claims.

### Flight Management Service
- Provides real-time information on flight statuses (3)
- Handles delays, cancellations, and diversions (4)

### Crew Management Service
- Manages the scheduling and assignment of flight crews (1)
- Tracks crew availability (2)

### Invertory Management Service
- Manages inventory for items such as meals, entertainment, and amenities onboard flights.

## Sequence
- Note: P is passenger, AE is airport employee, C is crew, SE is system employee, S is system(automate)
- Main flow
  - SE: create available flight, prices, available seats, schedules (depart time, duration) (9)
  - P: search for available flights
  - P: (8) choose depart time + destination -> choose trip -> choose number of passenger (adult, baby, child) -> chose type of ticket (vip / normal) -> choose type of trip (normal / round trip) -> fill information -> choose seat -> choose luggage options -> pay 
  - P: go to airport
  - P: show code or qr code
  - AE: verify in system (6)
  - P: ship luggage
  - AE: check luggage's weight
  - S: tracking passenger's luggage realtime (5)
  - P: go to waiting hall, see flight's status (3)
  - S: announce passenger go to port and plane
  - P: go to port and plane
  - S: check crew availability (2)
  - S: assign and announce the crew (prepare to start the flight) (1)
  - C: approve/reject assign
  - C: go to port and plane
- When a flight cancel/delay (4)
  - SE: announce why cancel/delay
  - S: update status flight at waiting hall
- System manage reservation process (7)
