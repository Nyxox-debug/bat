This is where something called the environment comes into play. The environment is what we
use to keep track of value by associating them with a name. The name “environment” is a
classic one, used in a lot of other interpreters, especially Lispy ones. But even though the name
may sound sophisticated, at its heart the environment is a hash map that associates strings
with objects. And that’s exactly what we’re going to use for our implementation.
