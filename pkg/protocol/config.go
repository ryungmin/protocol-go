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

type option func(c *configEntity)

func with_bits_per_line(bits_per_line int) option {
	return func(c *configEntity) {
		c.bits_per_line = bits_per_line
	}
}

func with_do_print_top_tens(do_print_top_tens bool) option {
	return func(c *configEntity) {
		c.do_print_top_tens = do_print_top_tens
	}
}

func with_do_print_top_units(do_print_top_units bool) option {
	return func(c *configEntity) {
		c.do_print_top_units = do_print_top_units
	}
}

func (h *configEntity) fromKeyAndValue(key, value string) error {
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
	return nil
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

		err := h.fromKeyAndValue(key, value)

		if err != nil {
			return err
		}
	}
	return nil
}
