package contracts
// RabbitMQ is a message broker that implements AMQP (Advanced Message Queuing Protocol)
// Think of it as a post office for microservices:
// - A service publishes a message (Trip created), another service subscribes and consumes that message(Payment service gets notified)
// AmqpMessage is the message structure for AMQP.
type AmqpMessage struct {
	OwnerID string `json:"ownerId"`
	Data    []byte `json:"data"`
}

// Routing keys - using consistent event/command patterns
const (
	// Trip events (trip.event.*)
	// publishes an event that trip event was created.
	TripEventCreated             = "trip.event.created"
	TripEventDriverAssigned      = "trip.event.driver_assigned"
	TripEventNoDriversFound      = "trip.event.no_drivers_found"
	TripEventDriverNotInterested = "trip.event.driver_not_interested"

	// Driver commands (driver.cmd.*)
	DriverCmdTripRequest = "driver.cmd.trip_request"
	DriverCmdTripAccept  = "driver.cmd.trip_accept"
	DriverCmdTripDecline = "driver.cmd.trip_decline"
	DriverCmdLocation    = "driver.cmd.location"
	DriverCmdRegister    = "driver.cmd.register"

	// Payment events (payment.event.*)
	PaymentEventSessionCreated = "payment.event.session_created"
	PaymentEventSuccess        = "payment.event.success"
	PaymentEventFailed         = "payment.event.failed"
	PaymentEventCancelled      = "payment.event.cancelled"

	// Payment commands (payment.cmd.*)
	PaymentCmdCreateSession = "payment.cmd.create_session"
)
