package glob

import "strings"

const glob = "*"

// Matcher creates a function that return true for subjects that match the glob
func Matcher(pattern string) func(subject string) bool {
	if pattern == glob {
		return func(string) bool {
			return true
		}
	}

	parts := strings.Split(pattern, glob)
	switch len(parts) {
	case 0:
		// should actually never happen
		fallthrough
	case 1:
		return func(subject string) bool {
			return pattern == subject
		}
	case 2:
		prefix := parts[0]
		suffix := parts[1]

		if prefix == "" {
			return func(subject string) bool {
				return strings.HasSuffix(subject, suffix)
			}
		}

		if suffix == "" {
			return func(subject string) bool {
				return strings.HasPrefix(subject, prefix)
			}
		}

		return func(subject string) bool {
			return strings.HasPrefix(subject, prefix) && strings.HasSuffix(subject, suffix)
		}
	default:
		prefix := parts[0]
		suffix := parts[len(parts)-1]

		var infixes []string
		for _, p := range parts[1 : len(parts)-1] {
			if p != "" {
				infixes = append(infixes, p)
			}
		}

		return func(subject string) bool {
			if prefix != "" && !strings.HasPrefix(subject, prefix) {
				return false
			}
			if suffix != "" && !strings.HasSuffix(subject, suffix) {
				return false
			}
			i := strings.LastIndex(subject, suffix)
			for _, in := range infixes {
				i = strings.LastIndex(subject[0:i], in)
				if i < len(prefix) {
					return false
				}
			}
			return true
		}
	}
}
