datatype MultisetContainer = EmptySet | BooleanMultiset(containerSet: multiset<bool>)

method Main() {
    // Start with a multiset containing one `false`
    var initialMultiset := multiset{false};
    
    for iteration := 0 to 98 {
        // Create a new instance with the current multiset
        var multisetInstance := BooleanMultiset(initialMultiset); 
        // Double the contents of the multiset
        initialMultiset := initialMultiset + multisetInstance.containerSet;  
    }

    // Compare the multiset with itself, which will always be false
    print initialMultiset > initialMultiset;
}
