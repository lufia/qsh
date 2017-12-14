PATH=(/bin /usr/bin)

echo a.b | cut -d . -f 2 | tr b c
# Output: c
