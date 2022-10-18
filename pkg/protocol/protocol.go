package protocol

import (
	"fmt"
	"strconv"
	"strings"
)

type configEntity struct {
	hdr_char_start     rune
	hdr_char_end       rune
	hdr_char_fill_odd  rune
	hdr_char_fill_even rune
	hdr_char_sep       rune
	bits_per_line      int
	do_print_top_tens  bool
	do_print_top_units bool
}

type Protocol interface {
	String() string
}

type protocolImpl struct {
	config     configEntity
	field_list []fieldEntity
}

type fieldEntity struct {
	text           string
	len            int
	more_fragments bool
}

// Converts the protocol specification stored in the object to a nice
// ASCII diagram like the ones that appear in RFCs. Conversion supports
// fields of any length, and supports field that span more than one
// line in the diagram.
// @return a string containing the ASCII representation of the protocol
// header.
func (pt protocolImpl) String() string {
	// First of all, process our field list. This does some magic to make
	// the algorithm work for fields that span more than one line
	source := make([]fieldEntity, len(pt.field_list))
	copy(source, pt.field_list)

	proto_fields := pt._process_field_list()
	lines := []string{}
	numbers := _get_top_numbers(
		pt.config.bits_per_line,
		pt.config.do_print_top_tens,
		pt.config.do_print_top_units)

	if len(numbers) > 0 {
		lines = append(lines, numbers)
	}
	lines = append(lines, pt._get_horizontal())

	// Print all protocol fields
	bits_in_line := 0
	current_line := ""
	fields_done := 0

	for p := -1; p < len(proto_fields)-1; {
		p += 1

		// Extract all the info we need about the field
		field := proto_fields[p]
		field_text := field.text
		field_len := field.len
		field_mf := field.more_fragments // Field has more fragments

		// If the field text is too long, we truncate it, and add a dot
		// at the end.
		if len(field_text) > (field_len*2 - 1) {
			field_text = field_text[0:(field_len*2 - 1)]
			if len(field_text) > 1 {
				// python: field_text = field_text[0:-1] + "."
				field_text = field_text[0:len(field_text)-2] + "."
			}
		}

		// If we have space for the whole field in the current line, go
		// ahead and add it
		if pt.config.bits_per_line-bits_in_line >= field_len {
			// If this is the first thing we print on a line, add the
			// starting character
			if bits_in_line == 0 {
				current_line += pt._get_separtor("")
			}

			// Add the whole field
			current_line += center(field_text, field_len*2-1, ' ')

			// Update counters
			bits_in_line += field_len
			fields_done += 1

			// If this is the last character in the line, store the line
			if bits_in_line == pt.config.bits_per_line {
				current_line += pt._get_separtor("")
				lines = append(lines, current_line)
				current_line = ""
				bits_in_line = 0
				// When we have a fragmented field, we may need to suppress
				// the floor of the field, so the current line connects
				// with the one that follows. E.g.:
				// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
				// |            field16            |                               |
				// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+                               +
				// |                             field                             |
				// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
				if field_mf {
					if proto_fields[p+1].len > pt.config.bits_per_line-field_len {
						// Print some +-+-+ to cover the previous field
						line_left := pt._get_horizontal_width(pt.config.bits_per_line - field_len)

						if len(line_left) == 0 {
							line_left = string(pt.config.hdr_char_start)
						}

						// Now print some empty space to cover the part that
						// we can join with the field below.
						// Case 1: If the next field reaches the end of its
						// line, then we need to print whitespace until the
						// end our line
						line_center := ""
						line_right := ""
						if proto_fields[p+1].len >= pt.config.bits_per_line {
							line_center = strings.Repeat(" ", 2*field_len-1)
							line_right = string(pt.config.hdr_char_end)
						} else {
							// Case 2: the field in the next row is not big enough
							// to cover all the space we'd like to join, so we
							// just print whitespace to cover as much as we can
							repeat_count := ((2 * (proto_fields[p+1].len - (pt.config.bits_per_line - field_len))) - 1)
							line_center = strings.Repeat(" ", repeat_count)
							line_right = pt._get_horizontal_width(pt.config.bits_per_line - proto_fields[p+1].len)
						}

						lines = append(lines, line_left+line_center+line_right)
					} else {
						lines = append(lines, pt._get_horizontal())
					}
				} else {
					lines = append(lines, pt._get_horizontal())
				}
			} else if fields_done == len(proto_fields) {
				// If this is not the last character of the line but we have no
				// more fields to print, wrap up
				current_line += pt._get_separtor("")
				lines = append(lines, current_line)
				lines = append(lines, pt._get_horizontal_width(bits_in_line))
			} else {
				// Add the separator character
				current_line += string(pt.config.hdr_char_sep)
			}
		} else {
			// We don't have enough space for the field on this line.
			if bits_in_line == 0 {
				// Case 1: We are at the beginning of a new line and we need to
				// span more than one line
				if (field_len % pt.config.bits_per_line) == 0 {
					// Case 1a: We have a multiple of the number of bits per line
					// Compute how many lines in total we need to print for this
					// big field.
					lines_to_print := int(((field_len / pt.config.bits_per_line) * 2) - 1)
					// We print the field text in the central line
					central_line := int(lines_to_print / 2)
					// Print all those lines
					for i := 0; i < lines_to_print; i++ {
						// Let's figure out which character we need to use
						// to start and end the current line
						start_line := string(pt.config.hdr_char_sep)
						end_line := string(pt.config.hdr_char_sep)
						if i%2 == 1 {
							start_line = string(pt.config.hdr_char_start)
							end_line = string(pt.config.hdr_char_end)
						}

						// This is the line where we need to print the field
						// text.
						if i == central_line {
							line := start_line + center(field_text, pt.config.bits_per_line*2-1, ' ') + end_line
							lines = append(lines, line)
						} else {
							// This is a line we need to leave blank
							line := start_line + strings.Repeat(" ", pt.config.bits_per_line*2-1) + end_line
							lines = append(lines, line)
						}

						if i == lines_to_print-1 {
							// If we just added the last line, add a horizontal separator
							lines = append(lines, pt._get_horizontal())
						}
					}

				}
			} else {
				// Case 2: We are not at the beginning of the line and we need
				// to print something that does not fit in the current line

				// This should never happen, since our _process_field_list()
				// divides fields in chunks so we never have the case of
				// something spanning lines in a weird manner
				panic(fmt.Errorf("case 2: We are not at the beginning of the line and we need to print something that does not fit in the current line"))
			}
		}
	}

	lines = append(lines, "")

	{ // print description
		// * {{proto_fields[0].text}} ({{proto_fields[0].len}} bytes)
		// * {{proto_fields[1].text}} ({{proto_fields[1].len}} bytes)
		// ...
		// * {{proto_fields[n].text}} ({{proto_fields[n].len}} bytes)
		// total {{SUM(proto_fields[0:n].len)}} bytes

		total := 0
		for _, field := range source {
			var line string
			total += field.len
			if field.len < 2 {
				line = fmt.Sprintf("* %s (%d byte)", field.text, field.len)
			} else {
				line = fmt.Sprintf("* %s (%d bytes)", field.text, field.len)
			}

			lines = append(lines, line)
		}

		var line string
		if total < 2 {
			line = fmt.Sprintf("total %d byte", total)
		} else {
			line = fmt.Sprintf("total %d bytes", total)
		}
		lines = append(lines, line)
	}

	result := strings.Join(lines, "\n")

	return result
}

// Parse field spec
func (p *protocolImpl) setFields(items []string) error {
	fields, err := parseFieldList(items)
	if err != nil {
		return err
	}

	p.field_list = append(p.field_list, fields...)
	return nil
}

// Parse options
func (p *protocolImpl) setOptions(options []string) error {
	return p.config.parseOptionList(options)
}

func (h *configEntity) parseOptionList(options []string) error {
	if len(options) == 0 {
		return nil
	}

	for _, option := range options {
		columns := strings.Split(option, "=")

		if len(columns) < 2 {
			continue
		}

		key := strings.ToLower(strings.TrimSpace(columns[0]))
		value := strings.ToLower(strings.TrimSpace(columns[1]))

		switch key {
		case "bits":
			n, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			h.bits_per_line = n
			if n <= 0 {
				return fmt.Errorf("FATAL: Invalid value for 'bits' option (%s)", value)
			}
		case "numbers":
			switch value {
			case "0", "n", "no", "none", "false":
				h.do_print_top_tens = false
				h.do_print_top_units = false
			case "1", "y", "yes", "true":
				h.do_print_top_tens = true
				h.do_print_top_units = true
			default:
				return fmt.Errorf("FATAL: Invalid value for 'numbers' option (%s)", value)
			}

		case "oddchar", "evenchar", "startchar", "endchar", "sepchar":
			if len(value) > 1 || len(value) <= 0 {
				return fmt.Errorf("FATAL: Invalid value for '%s' option (%s)", key, value)
			}

			switch key {
			case "oddchar":
				h.hdr_char_fill_odd = rune(value[0])
			case "evenchar":
				h.hdr_char_fill_even = rune(value[0])
			case "startchar":
				h.hdr_char_start = rune(value[0])
			case "endchar":
				h.hdr_char_end = rune(value[0])
			case "sepchar":
				h.hdr_char_sep = rune(value[0])
			}
		}
	}
	return nil
}

func (p protocolImpl) _get_horizontal() string {
	return p._get_horizontal_width(p.config.bits_per_line)
}

// @return the horizontal border line that separates field rows.
// @param width controls how many field bits the line should cover. By
// default, if no width is supplied, the line covers the hole length of
// the header.
func (p protocolImpl) _get_horizontal_width(width int) string {
	if width <= 0 {
		return ""
	}

	a := fmt.Sprintf("%c", p.config.hdr_char_start)
	b := strings.Repeat(fmt.Sprintf("%c%c", p.config.hdr_char_fill_even, p.config.hdr_char_fill_odd), width-1)
	c := fmt.Sprintf("%c%c", p.config.hdr_char_fill_even, p.config.hdr_char_end)

	return a + b + c
}

func (p protocolImpl) _get_separtor(line_end string) string {
	return string(p.config.hdr_char_sep)
}

// Processes the list of protocol fields that we got from the spec and turns
// it into something that we can print easily (useful for cases when we have
// protocol fields that span more than one line). This is just a helper
// function to make __str__()'s life easier.
func (p *protocolImpl) _process_field_list() []fieldEntity {
	new_fields := []fieldEntity{}
	bits_in_line := 0

	for i := 0; i < len(p.field_list); {
		// Extract all the info we need about the field
		field := p.field_list[i]
		field_text := field.text
		field_len := field.len

		field.more_fragments = false

		available_in_line := p.config.bits_per_line - bits_in_line

		if available_in_line >= field_len {
			new_fields = append(new_fields, field)
			bits_in_line += field_len

			i += 1

			if bits_in_line == p.config.bits_per_line {
				bits_in_line = 0
			}
		} else {
			if bits_in_line == 0 && (field_len%p.config.bits_per_line) == 0 {
				// Case 1: We have a field that is perfectly aligned and it
				// has a length that is multiple of our line length
				new_fields = append(new_fields, field)
				i += 1
				bits_in_line = 0
			} else {
				// Case 2: We weren't that lucky and the field is either not
				// aligned or we can't print it using an exact number of full
				// lines
				if available_in_line >= (field_len - available_in_line) {
					// If we have more space in the current line than in the next,
					// then put the field text in this one
					new_field := fieldEntity{
						field_text,
						available_in_line,
						true,
					}
					new_fields = append(new_fields, new_field)

					p.field_list[i].text = ""
					p.field_list[i].len = field_len - available_in_line
					p.field_list[i].more_fragments = false
				} else {
					new_field := fieldEntity{
						"",
						available_in_line,
						true,
					}
					new_fields = append(new_fields, new_field)

					p.field_list[i].text = field_text
					p.field_list[i].len = field_len - available_in_line
					p.field_list[i].more_fragments = false
				}
				bits_in_line = 0
				continue
			}
		}
	}

	return new_fields
}

func NewProtocol(spec string) (Protocol, error) {
	p := &protocolImpl{
		config: configEntity{
			hdr_char_start:     '+',
			hdr_char_end:       '+',
			hdr_char_fill_odd:  '+',
			hdr_char_fill_even: '-',
			hdr_char_sep:       '|',
			bits_per_line:      32,
			do_print_top_tens:  true,
			do_print_top_units: true,
		},
		field_list: []fieldEntity{},
	} // default

	// parse fields, options
	items, options := parseSpec(spec)

	// set fields
	err := p.setFields(items)

	if err != nil {
		return nil, err
	}

	// set options
	err = p.setOptions(options)

	if err != nil {
		return nil, err
	}

	return p, err
}
