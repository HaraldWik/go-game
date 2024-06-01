package err_

const (
	COLOR_RESET  = "\033[0m"
	COLOR_RED    = "\033[31m"
	COLOR_GREEN  = "\033[32m"
	COLOR_BLUE   = "\033[34m"
	COLOR_YELLOW = "\033[33m"
	COLOR_PURPLE = "\033[35m"
	COLOR_CYAN   = "\033[36m"
	COLOR_GRAY   = "\033[37m"
)

const (
	ERROR   string = COLOR_RED + "[ERROR] " + COLOR_RESET
	WARNING string = COLOR_YELLOW + "[INFO] " + COLOR_RESET
)

const (
	FAILED_TO_INIT   string = ERROR + "Failed to init:"
	NOT_DESIRED_TYPE string = WARNING + "Not a desired type"
)
