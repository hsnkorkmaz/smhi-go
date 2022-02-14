package utilities

import (
	"github.com/hsnkorkmaz/smhi-go/config"
	"fmt"
)

func GetEndpoint(partial string) string {
	return fmt.Sprintf("%s%s.%s", config.SMHI_ENDPOINT, partial, config.SMHI_RESPONSE_TYPE)
}
