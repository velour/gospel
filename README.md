We have all kinds of tools that check for errors in code (compiler,
linter, vetter, model checker, etc.), but only clumsy humans ever
check the comments—more than half of them don't care if what they're
reading is correct. Enter gospel: it checks spelling (and possibly
grammar, if we get super ambitious) in your Go doc comments. It is
aware of identifiers, so it won't flag things like "fmt" as errors.
