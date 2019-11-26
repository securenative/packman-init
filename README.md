# Packman Seed Template

This is the most basic project packman can generate, and it will simply template this readme file with your input data.  
Feel free to modify this seed to satisfy your needs.

Flags given to packman:
{{{ range $k, $v :=.Flags }}}
- {{{ $k }}} => {{{ $v }}}
{{{ end }}}

This template was generated at: {{{ .Timestamp }}}