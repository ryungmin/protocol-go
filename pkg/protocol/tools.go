package protocol

import (
	"fmt"
	"strconv"
	"strings"
)

// The center() method will center align the string, using a specified
// character as the fill character.
func center(s string, n int, character rune) string {
	m := len(s)
	if m >= n {
		return s
	}

	d := (n - m)

	left := d / 2
	right := d - left

	fill := string(character)
	return strings.Repeat(fill, left) + s + strings.Repeat(fill, right)
}

// @return a string representing the bit units and bit tens on top of the
// protocol header. Note that a proper string is only returned if one or
// both self.do_print_top_tens and self.do_print_top_units is True.
// The returned string is not \n terminated, but it may contain a newline
// character in the middle.
func _get_top_numbers(opts ...option) string {
	cfgs := &configEntity{
		bits_per_line:      32,
		do_print_top_tens:  false,
		do_print_top_units: false,
	}

	for _, opt := range opts {
		opt(cfgs)
	}

	lines := []string{"", ""}

	if cfgs.do_print_top_tens {
		for i := 0; i < cfgs.bits_per_line; i++ {
			unit := i % 10
			n := i / 10

			if unit == 0 {
				lines[0] += fmt.Sprintf(" %d", n)
			} else {
				lines[0] += "  "
			}
		}
		lines[0] += "\n"
	}

	if cfgs.do_print_top_units {
		for i := 0; i < cfgs.bits_per_line; i++ {
			unit := i % 10
			lines[1] += fmt.Sprintf(" %d", unit)
		}
	}
	result := strings.Join(lines, "")

	return result
}

// Parses a textual protocol spec and stores the relevant internal state
// so such spec can be later converted to a nice ASCII diagram.
// @return the list of protocol fields, as a dictionary containing
// keys 'len' and 'text'. The list is returned for completeness but no
// caller is expected to store or use such list.
// @raise ProtocolException in case the supplied spec is not valid
func parseSpec(spec string) ([]string, []string) {
	var items []string
	options := []string{}

	if strings.Contains(spec, "?") {
		// split '?'
		parts := strings.Split(spec, "?")

		items = strings.Split(parts[0], ",")
		options = strings.Split(parts[1], ",")
	} else {
		items = strings.Split(spec, ",")
	}

	return items, options
}

func parseFieldList(items []string) ([]fieldEntity, error) {
	fieldList := []fieldEntity{}

	for _, item := range items {
		columns := strings.Split(item, ":")
		if len(columns) < 2 {
			return nil, fmt.Errorf("FATAL: Invalid field_list specification (%s)", item)
		}

		text := strings.TrimSpace(columns[0])
		bits, err := strconv.Atoi(strings.TrimSpace(columns[1]))

		if err != nil {
			return nil, err
		}

		if bits <= 0 {
			return nil, fmt.Errorf("FATAL: Fields must be at least one bit long (%s)", item)
		}

		fieldList = append(fieldList, fieldEntity{
			text,
			bits,
			false,
		})
	}

	return fieldList, nil
}
