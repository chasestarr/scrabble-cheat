var Tree = function(value) {
  var newTree = Object.create(treeMethods);
  newTree.v = value;
  newTree.c = [];
  return newTree;
};

var treeMethods = {};

treeMethods.addChild = function(value) {
  var newChild = Tree(value);
  this.c.push(newChild);
  return newChild;
};

treeMethods.contains = function(target) {
  var found;
  if (arguments[1] === undefined) {
    found = false;
  } else {
    found = arguments[1];
  }
  // Look over the children
  this.c.forEach(function(child) {
    // With each child
      // If value matches
    if (child.v === target) {
      found = true;
    } else {
      // If doesn't; recurse
      found = child.contains(target, found);
    }
  });
  // Otherwise return false
  return found;
};

module.exports = Tree;