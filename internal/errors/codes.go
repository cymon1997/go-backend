package errors

const (
	//2xx Success
	Success Code = 200

	//4xx Client Error
	InvalidRequest Code = 400
	Forbidden      Code = 403
	NoDataFound    Code = 404

	//5xx Server Error
	InternalServer     Code = 500
	ServiceUnavailable Code = 504
)

type Code int

func (c *Code) Status() string {
	switch *c {
	case Success:
		return "ok"
	case InvalidRequest:
		return "invalid request"
	case Forbidden:
		return "invalid access"
	case NoDataFound:
		return "no data found"
	case InternalServer:
		return "internal server error"
	case ServiceUnavailable:
		return "service unavailable"
	default:
		return "unknown error"
	}
}

func (c *Code) ShouldRetry() bool {
	return *c >= InternalServer
}
